package repository

import (
	"banking/model"
	"banking/mongoImplement"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
)

type accountRepository struct {
	db         mongoImplement.Database
	collection string
}

func NewAccountRepository(db mongoImplement.Database, collection string) model.AccountRepository {
	return &accountRepository{
		db:         db,
		collection: collection,
	}
}

func (acc *accountRepository) ListAccount(ctx context.Context) ([]model.Account, error) {
	var accounts []model.Account
	collection := acc.db.Collection(acc.collection)
	filter := bson.D{}
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return accounts, err
	}

	if err := cur.All(ctx, &accounts); err != nil {
		return accounts, err
	}

	if accounts == nil {
		return accounts, errors.New("no accounts found")
	}

	return accounts, nil
}

func (acc *accountRepository) InsertAccount(ctx context.Context, account *model.Account) (interface{}, error) {
	collection := acc.db.Collection(acc.collection)
	return collection.InsertOne(ctx, &account)
}
