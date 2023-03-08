package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Account struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"name" bson:"name" binding:"required"`
	Value int                `json:"value" bson:"value" binding:"required"`
}

type (
	AccountRepository interface {
		ListAccount(ctx context.Context) ([]Account, error)
		InsertAccount(ctx context.Context, account *Account) (interface{}, error)
	}
	AccountUseCase interface {
		ListAccount(ctx context.Context) ([]Account, error)
		InsertAccount(ctx context.Context, account *Account) (interface{}, error)
	}
)
