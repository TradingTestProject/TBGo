package main

import (
	investapi "TBGo/proto/generated"
	"github.com/shopspring/decimal"
)

func MapToCandleInterval(in CandleInterval) investapi.CandleInterval {
	switch in {
	case CandleInterval1Min:
		return investapi.CandleInterval_CANDLE_INTERVAL_1_MIN
	case CandleInterval5Min:
		return investapi.CandleInterval_CANDLE_INTERVAL_5_MIN
	case CandleInterval15Min:
		return investapi.CandleInterval_CANDLE_INTERVAL_15_MIN
	case CandleInterval1Hour:
		return investapi.CandleInterval_CANDLE_INTERVAL_HOUR
	case CandleInterval1Day:
		return investapi.CandleInterval_CANDLE_INTERVAL_DAY
	}

	return investapi.CandleInterval_CANDLE_INTERVAL_UNSPECIFIED
}

func MapToQuotation(in decimal.Decimal) *investapi.Quotation {
	v := in.Sub(decimal.NewFromInt(in.IntPart())).CoefficientInt64()
	return &investapi.Quotation{
		Units: in.IntPart(),
		Nano:  int32(calcFactor(v, -9-in.Exponent())),
	}
}

func calcFactor(v int64, exp int32) int64 {
	if exp == 0 {
		return v
	}
	return calcFactor(v*10, exp+1)
}
