package entity

import "time"

type TradeState int
type Style int
type Market int

const (
	Open TradeState = iota + 1
	Closed

	BreakOut Style = iota + 1
	Reversal
	Macro
	GutFeel

	AShare Market = iota + 1
	US
	HK
)

type Result struct {
	//盈亏
	//%
	//结果review
	//Mistake
	//Lesson
	Gain        float64  // profit or loss
	GainPercent float64  // % of this trade
	Mistakes    []string //
	Lessons     []string // review的可优化点
}

type DecisionProcess struct {
	Economy     []string // 大市
	Sector      []string // 板块
	Fundamental []string // 基础分析
	Technical   []string // TA
	Thought     []string // 想法
}

// Trade a Trade
type Trade struct {
	TradeID int64
	Buys    []*Action
	Sells   []*Action

	//交易类型
	//市场
	//代号
	TradeState TradeState // 这个交易是完结还是进行中
	Style      Style
	Market     Market
	Ticker     string
	OpenDate   time.Time

	DecisionProcess DecisionProcess
	Risk            Risk
	Result          Result
}
