package types

import (
	"time"

	"github.com/shopspring/decimal"
)

type SideType int8

const (
	SideBid SideType = iota
	SideAsk
)

func SideFromString(s string) SideType {
	if s == "buy" {
		return SideBid
	}
	return SideAsk
}

type L2OrderBookItem struct {
	Price  decimal.Decimal
	Volume uint64
	Time   time.Time
}

type L2OrderBook struct {
	Bid, Ask []*L2OrderBookItem
}
