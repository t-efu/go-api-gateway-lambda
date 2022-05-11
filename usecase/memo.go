package usecase

import (
	"context"

	"github.com/pkg/errors"
	"github.com/t-efu/go-api-gateway-lambda/domain/entity"
	"github.com/t-efu/go-api-gateway-lambda/domain/repository"
)

type MemoUsecase interface {
	Find(ctx context.Context) ([]entity.Memo, error)
	Create(ctx context.Context, memo *entity.Memo) error
}

type memoUsecase struct {
	memoRepository repository.MemoRepository
}

func NewMemoUsecase(memoRepository repository.MemoRepository) MemoUsecase {
	return &memoUsecase{
		memoRepository: memoRepository,
	}
}

func (u *memoUsecase) Find(ctx context.Context) ([]entity.Memo, error) {
	memos, err := u.memoRepository.Find(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed find memos")
	}
	return memos, nil
}

func (u *memoUsecase) Create(ctx context.Context, memo *entity.Memo) error {
	_, err := u.memoRepository.Create(ctx, memo)
	if err != nil {
		return errors.Wrap(err, "failed create memo")
	}
	return nil
}
