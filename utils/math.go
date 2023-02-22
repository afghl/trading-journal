package utils

import (
	"math"
)

func Round2Decimal(x float64) float64 {
	return math.Round(x*100) / 100
}

func Round3Decimal(x float64) float64 {
	return math.Round(x*1000) / 1000
}
