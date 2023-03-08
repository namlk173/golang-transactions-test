package repository

import (
	"banking/model"
	"banking/mongoImplement"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"time"
)

type transferRepository struct {
	db           mongoImplement.Database
	TransferColl string
	AccountColl  string
}

func NewTransferRepository(db mongoImplement.Database, transferColl, accountColl string) model.TransferRepository {
	return &transferRepository{
		db:           db,
		TransferColl: transferColl,
		AccountColl:  accountColl,
	}
}

func (tran *transferRepository) ListTransferByAccount(ctx context.Context, id primitive.ObjectID) ([]model.TransferResponse, error) {
	var transfers []model.TransferResponse
	collection := tran.db.Collection(tran.TransferColl)
	filter := bson.D{
		{"$or", []interface{}{
			bson.D{{"from._id", id}},
			bson.D{{"to._id", id}},
		}},
	}
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return transfers, err
	}

	err = cur.All(ctx, &transfers)
	if err != nil {
		return []model.TransferResponse{}, err
	}

	if transfers == nil {
		return transfers, errors.New("no transfer for this account")
	}

	return transfers, err
}

func (tran *transferRepository) GetAccountByID(ctx context.Context, id primitive.ObjectID) (*model.Account, error) {
	collection := tran.db.Collection(tran.AccountColl)
	var account model.Account
	err := collection.FindOne(ctx, bson.D{{"_id", id}}).Decode(&account)
	if err != nil {
		return &model.Account{}, errors.New("not found account")
	}

	return &account, nil
}

func (tran *transferRepository) GetTransferByID(ctx context.Context, id primitive.ObjectID) (*model.TransferResponse, error) {
	var transfer model.TransferResponse
	collection := tran.db.Collection(tran.TransferColl)
	err := collection.FindOne(ctx, bson.D{{"_id", id}}).Decode(&transfer)
	if err != nil {
		return &model.TransferResponse{}, err
	}

	return &transfer, nil
}

func (tran *transferRepository) Transfer(ctx context.Context, transfer *model.TransferRequest) (interface{}, error) {
	wcMajority := writeconcern.New(writeconcern.WMajority(), writeconcern.WTimeout(time.Second*1))
	wcMajorityCollectionOpts := options.Collection().SetWriteConcern(wcMajority)

	accountCollection := tran.db.Collection(tran.AccountColl, wcMajorityCollectionOpts)
	transferCollection := tran.db.Collection(tran.TransferColl, wcMajorityCollectionOpts)
	var idTransfer interface{}

	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {

		// Increase value of received account more (Amount transfer)
		toAccountUpdateQuery := bson.D{{"$inc", bson.D{{"value", int(transfer.Amount)}}}}
		_, err := accountCollection.UpdateOne(sessCtx, bson.D{{"_id", transfer.To}}, toAccountUpdateQuery)
		if err != nil {
			return nil, err
		}

		// Decrease value of send account about (Amount transfer)
		fromAccountUpdateQuery := bson.D{{"$inc", bson.D{{"value", -int(transfer.Amount)}}}}
		_, err = accountCollection.UpdateOne(sessCtx, bson.D{{"_id", transfer.From}}, fromAccountUpdateQuery)
		if err != nil {
			return nil, err
		}

		// Get information of sending account
		//	IF ERROR (NOT FOUND THIS ACCOUNT) => ROLLBACK
		// 	ELSE IF VALUE < 0 => ROLLBACK
		accountFrom, err := tran.GetAccountByID(sessCtx, transfer.From)
		if err != nil {
			return nil, err
		}

		if accountFrom.Value < 0 {
			return nil, errors.New("not enough money to transfer")
		}

		// IF accountTo not exist, Need ROLLBACK too
		accountTo, err := tran.GetAccountByID(sessCtx, transfer.To)
		if err != nil {
			return nil, err
		}

		transferWriter := model.TransferWriter{
			From:      *accountFrom,
			To:        *accountTo,
			Amount:    int(transfer.Amount),
			CreatedAt: time.Now(),
		}

		idTransfer, err = transferCollection.InsertOne(sessCtx, transferWriter)
		if err != nil {
			return nil, err
		}
		return idTransfer, nil
	}
	session, err := tran.db.Client().StartSession()
	if err != nil {
		return nil, err
	}
	defer session.EndSession(ctx)

	result, err := session.WithTransaction(ctx, callback)
	if err != nil {
		return nil, err
	}

	return result, nil
}
