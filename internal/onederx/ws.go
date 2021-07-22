package onederx

func GetWsL2SubscribeRequest() interface{} {
	subscribeReq := struct {
		Type    string
		Payload struct {
			Subscriptions []struct {
				Channel string
				Params  struct {
					Symbol string
				}
			}
		}
	}{}
	subscribeReq.Type = "subscribe"
	subscribeReq.Payload.Subscriptions = make([]struct {
		Channel string
		Params  struct{ Symbol string }
	}, 1)

	subscribeReq.Payload.Subscriptions[0].Channel = "l2"
	subscribeReq.Payload.Subscriptions[0].Params.Symbol = "BTCUSD_P"

	return subscribeReq
}
