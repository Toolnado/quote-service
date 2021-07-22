package onederx

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	onederxWsUrl         = "wss://api.onederx.com/v1/ws"
	defaultRetryInterval = time.Second
	defaultReadTimeout   = time.Second * 5
	defaultWriteTimeout  = time.Second * 5
)

type Source struct {
}

func NewSource() *Source {
	return &Source{}
}

func (s *Source) Start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Printf("source stopped: context cancelled")
			return
		default:
			if err := s.receiveData(ctx); err != nil {
				log.Printf("sleep for %v", defaultRetryInterval)
				time.Sleep(defaultRetryInterval)
			}
		}

	}
}

func (s *Source) receiveData(ctx context.Context) error {
	dialer := websocket.Dialer{}
	conn, _, err := dialer.DialContext(ctx, onederxWsUrl, nil)
	if err != nil {
		return err
	}

	if err := conn.SetWriteDeadline(time.Now().Add(defaultWriteTimeout)); err != nil {
		return err
	}

	if err := conn.WriteJSON(GetWsL2SubscribeRequest()); err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			mt, data, err := conn.ReadMessage()
			if err != nil {
				return err
			}

			if mt != websocket.TextMessage {
				return fmt.Errorf("unexpected message type")
			}

			header := struct{ Type string }{}

			if err := json.Unmarshal(data, &header); err != nil {
				return err
			}

			switch header.Type {
			case "snapshot":
				s.onSnapshot(data)
			case "update":
				s.onUpdate(data)
			default:
			}
		}

	}

}

func (s *Source) onSnapshot(data []byte) error {
	return nil
}

func (s *Source) onUpdate(data []byte) error {
	return nil
}
