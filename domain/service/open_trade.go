package service

import (
	"context"
	"time"
	"trading_journal/domain/entity"
)

// OpenTradeCommand 开仓必须的参数
type OpenTradeCommand struct {

	// 交易类型	市场	代号	买入时间	买入价	股数
	// 目标	止损
	// 大市分析	基本分析	技术分析	想法
	Style  entity.Style
	Market entity.Market
	Ticker string

	BuyTime *time.Time
	Price   float64
	Amount  float64
	Target  float64
	SL      float64

	Economy     []string // 大市
	Sector      []string // 板块
	Fundamental []string // 基础分析
	Technical   []string // TA
	Thought     []string // 想法
}

// OpenTrade 开仓
func OpenTrade(ctx context.Context, cmd OpenTradeCommand) (*entity.Trade, error) {
	var buyTime time.Time
	if cmd.BuyTime != nil {
		buyTime = *cmd.BuyTime
	} else {
		buyTime = time.Now()
	}
	trade := &entity.Trade{
		Buys: []*entity.Action{{
			ActType: entity.Buy,
			Date:    buyTime,
			Price:   cmd.Price,
			Amount:  cmd.Amount,
		}},
		TradeState: entity.Open,
		Style:      cmd.Style,
		Market:     cmd.Market,
		Ticker:     cmd.Ticker,
		OpenDate:   buyTime,
		DecisionProcess: entity.DecisionProcess{
			Economy:     cmd.Economy,
			Sector:      cmd.Sector,
			Fundamental: cmd.Fundamental,
			Technical:   cmd.Technical,
			Thought:     cmd.Thought,
		},
		Risk: entity.CalRisk(cmd.Price, cmd.Target, cmd.SL, cmd.Amount),
	}
	return trade, nil
}
