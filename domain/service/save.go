package service

import (
	"context"
	"trading_journal/domain/entity"
)

func Save(ctx context.Context, trade *entity.Trade) error {
	return getProvider().Repo.Save(ctx, trade)
}
