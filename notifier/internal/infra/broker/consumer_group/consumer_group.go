package consumer_group

import (
	"context"
	"fmt"
	"route256/notifier/internal/pkg/logger"

	"github.com/IBM/sarama"
)

type ConsumerGroup struct {
	sarama.ConsumerGroup
	Handlers *ConsumerGroupHandler
	Topics   []string
}

type Config interface {
	GetBrokers() []string
	GetGroupID() string
	GetTopics() []string
}

func New(ctx context.Context, configMain Config) (*ConsumerGroup, error) {
	handlers := NewHandler()
	config := sarama.NewConfig()

	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Offsets.AutoCommit.Enable = false

	sCG, err := sarama.NewConsumerGroup(
		configMain.GetBrokers(),
		configMain.GetGroupID(),
		config,
	)
	if err != nil {
		return nil, fmt.Errorf("sarama.NewConsumerGroup %w", err)
	}

	ConsumerGroup := &ConsumerGroup{
		sCG,
		handlers,
		configMain.GetTopics(),
	}

	go ConsumerGroup.Run(ctx)

	return ConsumerGroup, nil
}

func (c *ConsumerGroup) Run(ctx context.Context) {
	for {
		if err := c.ConsumerGroup.Consume(ctx, c.Topics, c.Handlers); err != nil {
			logger.Errorw(ctx, err.Error())
		}

		if ctx.Err() != nil {
			logger.Infow(ctx, "ConsumerGroup.Run ctx closed")
			return
		}
	}
}
