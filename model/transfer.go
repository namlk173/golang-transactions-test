package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type TransferResponse struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	From      Account            `json:"from" bson:"from" binding:"required"`
	To        Account            `json:"to" bson:"to" binding:"required"`
	Amount    int                `json:"amount" bson:"amount" binding:"required"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}

type TransferRequest struct {
	From   primitive.ObjectID `json:"from" bson:"from" binding:"required"`
	To     primitive.ObjectID `json:"to" bson:"to" binding:"required"`
	Amount uint               `json:"amount" bson:"amount" binding:"required"`
}

type TransferWriter struct {
	From      Account   `json:"from" bson:"from" binding:"required"`
	To        Account   `json:"to" bson:"to" binding:"required"`
	Amount    int       `json:"amount" bson:"amount" binding:"required"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}

type (
	TransferRepository interface {
		ListTransferByAccount(ctx context.Context, id primitive.ObjectID) ([]TransferResponse, error)
		GetAccountByID(ctx context.Context, id primitive.ObjectID) (*Account, error)
		GetTransferByID(ctx context.Context, id primitive.ObjectID) (*TransferResponse, error)
		Transfer(ctx context.Context, transfer *TransferRequest) (interface{}, error)
	}
	TransferUseCase interface {
		ListTransferByAccount(ctx context.Context, id primitive.ObjectID) ([]TransferResponse, error)
		Transfer(ctx context.Context, transfer *TransferRequest) (interface{}, error)
	}
)
