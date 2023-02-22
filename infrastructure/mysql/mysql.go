package mysql

import (
	"context"
	"errors"
	"trading_journal/domain/entity"
	"trading_journal/domain/repository"
)

type repoImpl struct {
}

func GetRepo() repository.TradeRepository {
	return &repoImpl{}
}

func (r repoImpl) Fetch(ctx context.Context, ID uint64) (*entity.Trade, error) {
	//TODO implement me
	panic("implement me")
}

func (r repoImpl) Save(ctx context.Context, t *entity.Trade) error {
	if t.TradeID == 0 {
		return insert(ctx, t)
	} else {
		return errors.New("not implement")
	}
}

func insert(ctx context.Context, t *entity.Trade) error {
	return errors.New("not implements" +
		"")
}
