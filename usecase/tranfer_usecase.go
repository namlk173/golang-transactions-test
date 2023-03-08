package usecase

import (
	"banking/model"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type transferUseCase struct {
	model.TransferRepository
	Timeout time.Duration
}

func NewTransferUseCase(transferRepository model.TransferRepository, timeout time.Duration) model.TransferUseCase {
	return &transferUseCase{
		transferRepository,
		timeout,
	}
}

func (tran *transferUseCase) ListTransferByAccount(ctx context.Context, id primitive.ObjectID) ([]model.TransferResponse, error) {
	c, cancel := context.WithTimeout(ctx, tran.Timeout)
	defer cancel()

	return tran.TransferRepository.ListTransferByAccount(c, id)
}

func (tran *transferUseCase) Transfer(ctx context.Context, transfer *model.TransferRequest) (interface{}, error) {
	c, cancel := context.WithTimeout(ctx, tran.Timeout)
	defer cancel()

	return tran.TransferRepository.Transfer(c, transfer)
}
