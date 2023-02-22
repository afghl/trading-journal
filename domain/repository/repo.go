package repository

import (
	"context"
	"trading_journal/domain/entity"
)

type TradeRepository interface {
	Fetch(ctx context.Context, ID uint64) (*entity.Trade, error)

	Save(ctx context.Context, t *entity.Trade) error
}
