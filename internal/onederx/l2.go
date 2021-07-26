package onederx

import (
	"time"

	"github.com/HuKeping/rbtree"
	"github.com/Toolnado/quote-service.git/internal/types"
	"github.com/shopspring/decimal"
)

type L2OrderBookItem types.L2OrderBookItem

func (item L2OrderBookItem) Less(than rbtree.Item) bool {
	return item.Price.LessThan(than.(*L2OrderBookItem).Price)
}

type L2OrderBook struct {
	bid, ask *rbtree.Rbtree
}

func NewL2OrderBook() *L2OrderBook {
	return &L2OrderBook{
		bid: rbtree.New(),
		ask: rbtree.New(),
	}
}

func (ob *L2OrderBook) Apply(price decimal.Decimal, side types.SideType, volume uint64, tm time.Time) {
	obs := ob.bid
	if side == types.SideAsk {
		obs = ob.ask
	}

	item := obs.InsertOrGet(&L2OrderBookItem{Price: price})

	if volume == 0 {
		obs.Delete(item)
		return
	}
	item.(*L2OrderBookItem).Volume = volume
	item.(*L2OrderBookItem).Time = tm
}

func (ob *L2OrderBook) GetBid(size int) []*types.L2OrderBookItem {
	ret := make([]*types.L2OrderBookItem, 0, size)
	ob.bid.Descend(ob.bid.Max(), func(item rbtree.Item) bool {
		itemCopy := types.L2OrderBookItem(*item.(*L2OrderBookItem))
		ret = append(ret, &itemCopy)
		size--
		return size != 0
	})
	return ret
}

func (ob *L2OrderBook) GetAsk(size int) []*types.L2OrderBookItem {
	ret := make([]*types.L2OrderBookItem, 0, size)
	ob.bid.Descend(ob.bid.Min(), func(item rbtree.Item) bool {
		itemCopy := types.L2OrderBookItem(*item.(*L2OrderBookItem))
		ret = append(ret, &itemCopy)
		size--
		return size != 0
	})
	return ret
}
