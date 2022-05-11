package repository

import (
	"context"

	"github.com/t-efu/go-api-gateway-lambda/domain/entity"
)

type MemoRepository interface {
	Find(ctx context.Context) ([]entity.Memo, error)
	Create(ctx context.Context, memo *entity.Memo) (*entity.Memo, error)
}
