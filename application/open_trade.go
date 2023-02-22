package application

import (
	"context"
	"trading_journal/domain/service"
)

// OpenTrade 新的trade
func OpenTrade(ctx context.Context, cmd service.OpenTradeCommand) error {
	trade, err := service.OpenTrade(ctx, cmd)
	if err != nil {
		return err
	}
	if err := service.ValidateTrade(ctx, trade); err != nil {
		return err
	}
	if err := service.Save(ctx, trade); err != nil {
		return err
	}
	return nil
}
