package rpc

import (
	"log"
	"time"

	"github.com/Toolnado/quote-service.git/api"
	"github.com/Toolnado/quote-service.git/internal/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	sources []types.Source
	api.UnimplementedQuotesServer
}

func NewService() *Service {
	return &Service{
		sources: make([]types.Source, 0),
	}
}

func (s *Service) AddSource(source types.Source) {
	s.sources = append(s.sources, source)
}

func (s *Service) GetL2OrderBook(req *api.L2OrderBookRequest, stream api.Quotes_GetL2OrderBookServer) error {
	log.Printf("client connected")

	if req.Size <= 0 {
		return status.Error(codes.InvalidArgument, "invalid size")
	}

	if req.Interval <= 0 {
		return status.Error(codes.InvalidArgument, "invalid interval")
	}
	var (
		stop bool
		l2   *types.L2OrderBook
		err  error
	)

	for !stop {
		select {
		case <-stream.Context().Done():
			stop = true
			break

		case <-time.After(time.Duration(req.Interval) * time.Microsecond):
			for _, source := range s.sources {
				l2, err = source.GetL2OrderBook(req.Symbol, int(req.Size))
				if err != nil {
					stop = true
					break
				}

				if err = stream.Send(ConvertToProtoL2(req.Symbol, *l2)); err != nil {
					stop = true
					break
				}
			}
		}
	}

	log.Printf("client disconnected")
	return err
}
