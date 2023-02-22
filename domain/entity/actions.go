package entity

import "time"

type ActType int

const (
	Buy ActType = iota + 1
	Sell
)

// Action buy or sell
type Action struct {
	ActType ActType
	Date    time.Time
	Price   float64
	Amount  float64
}
