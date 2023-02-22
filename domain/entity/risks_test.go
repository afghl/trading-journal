package entity

import (
	"reflect"
	"testing"
)

func TestCalRisk(t *testing.T) {
	type args struct {
		entryPrice    float64
		targetPrice   float64
		stopLossPrice float64
		amount        float64
	}
	tests := []struct {
		name string
		args args
		want Risk
	}{
		{name: "ok", args: args{
			entryPrice:    35.05,
			targetPrice:   41,
			stopLossPrice: 32,
			amount:        65,
		}, want: Risk{
			TargetPrice:     41,
			StopLossPrice:   32,
			RiskRewardRatio: 1.95,
			Risk:            198.25,
			RiskPercent:     0.087,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalRisk(tt.args.entryPrice, tt.args.targetPrice, tt.args.stopLossPrice, tt.args.amount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalRisk() = %v, want %v", got, tt.want)
			}
		})
	}
}
