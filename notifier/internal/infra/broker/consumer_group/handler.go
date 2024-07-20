package consumer_group

import (
	"context"
	"route256/notifier/internal/pkg/logger"

	"github.com/IBM/sarama"
)

var _ sarama.ConsumerGroupHandler = (*ConsumerGroupHandler)(nil)

type ConsumerGroupHandler struct {
	ready chan bool
}

func NewHandler() *ConsumerGroupHandler {
	return &ConsumerGroupHandler{}
}

func (c *ConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (c *ConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (c *ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	ctx := context.Background()

	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				return nil
			}

			msg := convertMsg(message)
			logger.Infow(ctx, msg.Payload)

			session.MarkMessage(message, "")
			session.Commit()

		case <-session.Context().Done():
			return nil
		}
	}
}

type Msg struct {
	Topic     string `json:"topic"`
	Partition int32  `json:"partition"`
	Offset    int64  `json:"offset"`
	Key       string `json:"key"`
	Payload   string `json:"payload"`
}

func convertMsg(in *sarama.ConsumerMessage) *Msg {
	return &Msg{
		Topic:     in.Topic,
		Partition: in.Partition,
		Offset:    in.Offset,
		Key:       string(in.Key),
		Payload:   string(in.Value),
	}
}
