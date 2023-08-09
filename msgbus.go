package msgbuzz

type MessageBus interface {
	Publish(topicName string, msg []byte, options ...func(*PublishOption)) error
	On(topicName string, consumerName string, handlerFunc MessageHandler) error
}

type MessageHandler func(MessageConfirm, []byte) error

type MessageConfirm interface {
	Ack() error
	Nack() error
	Retry(delay int64, maxRetry int) error
}

type PublishOption struct {
	RoutingKey   string
	ExchangeType string
}

func WithRoutingKey(routingKey string) func(*PublishOption) {
	return func(p *PublishOption) {
		p.RoutingKey = routingKey
	}
}

func WithExchangeType(exchangeType string) func(*PublishOption) {
	return func(p *PublishOption) {
		p.ExchangeType = exchangeType
	}
}

func (p *PublishOption) GetExchangeType() string {
	if p.ExchangeType == "" {
		p.ExchangeType = "fanout"
	}
	return p.ExchangeType
}
