package entity

import "trading_journal/utils"

// Risk 风险
type Risk struct {
	//目标
	//止损
	//R/R
	//风险
	//%
	TargetPrice     float64
	StopLossPrice   float64
	RiskRewardRatio float64 // TODO: 怎么算
	Risk            float64 // TODO: risk 是否在整个trade里记录？
	RiskPercent     float64 // TODO: 如果多次买，怎么算
}

// CalRisk 计算风险
func CalRisk(entryPrice, targetPrice, stopLossPrice, amount float64) Risk {
	return Risk{
		TargetPrice:     targetPrice,
		StopLossPrice:   stopLossPrice,
		RiskRewardRatio: utils.Round2Decimal((targetPrice - entryPrice) / (entryPrice - stopLossPrice)),
		Risk:            utils.Round2Decimal((entryPrice - stopLossPrice) * amount),
		RiskPercent:     utils.Round3Decimal(1 - (stopLossPrice / entryPrice)),
	}
}
