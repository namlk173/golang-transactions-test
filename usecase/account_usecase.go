package usecase

import (
	"banking/model"
	"context"
	"time"
)

type accountUseCase struct {
	model.AccountRepository
	Timeout time.Duration
}

func NewAccountUseCase(accountRepository model.AccountRepository, timeout time.Duration) model.AccountUseCase {
	return &accountUseCase{
		accountRepository,
		timeout,
	}
}

func (acc *accountUseCase) ListAccount(ctx context.Context) ([]model.Account, error) {
	c, cancel := context.WithTimeout(ctx, acc.Timeout)
	defer cancel()
	return acc.AccountRepository.ListAccount(c)
}

func (acc *accountUseCase) InsertAccount(ctx context.Context, account *model.Account) (interface{}, error) {
	c, cancel := context.WithTimeout(ctx, acc.Timeout)
	defer cancel()
	return acc.AccountRepository.InsertAccount(c, account)
}
