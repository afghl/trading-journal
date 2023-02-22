package mysql

import "time"

type TTrade struct {
	TradeID         int64
	TradeState      int // 这个交易是完结还是进行中
	Style           int
	Market          int
	Ticker          string
	Economy         []string // 大市
	Sector          []string // 板块
	Fundamental     []string // 基础分析
	Technical       []string // TA
	Thought         []string // 想法
	OpenDate        time.Time
	TargetPrice     float64
	StopLossPrice   float64
	RiskRewardRatio float64
	Risk            float64
	RiskPercent     float64
	Gain            float64
	GainPercent     float64
	Mistakes        []string
	Lessons         []string
}
