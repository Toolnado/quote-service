package types

type Source interface {
	GetL2OrderBook(symbol string, size int) (*L2OrderBook, error)
}
