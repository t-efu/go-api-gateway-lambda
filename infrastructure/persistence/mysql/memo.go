package mysql

import (
	"context"

	"github.com/pkg/errors"
	"github.com/t-efu/go-api-gateway-lambda/domain/entity"
	"github.com/t-efu/go-api-gateway-lambda/domain/repository"

	"gorm.io/gorm"
)

type memoRepository struct {
	conn *gorm.DB
}

func NewMemoRepository(conn *gorm.DB) repository.MemoRepository {
	return &memoRepository{
		conn: conn,
	}
}

func (r *memoRepository) Find(ctx context.Context) ([]entity.Memo, error) {
	var memos []entity.Memo
	err := r.conn.Find(&memos).Error
	if err != nil {
		return nil, errors.Wrap(err, "failed find memos")
	}
	return nil, nil
}

func (r *memoRepository) Create(ctx context.Context, memo *entity.Memo) (*entity.Memo, error) {
	err := r.conn.Create(memo).Error
	if err != nil {
		return nil, errors.Wrap(err, "failed create memo")
	}
	return memo, nil
}
