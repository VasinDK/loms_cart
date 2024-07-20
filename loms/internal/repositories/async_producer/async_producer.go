package async_producer

import (
	"context"
	"fmt"
	"route256/loms/internal/model"

	"github.com/IBM/sarama"
)

type Producer struct {
	AsyncProducer sarama.AsyncProducer
	Partitions    []int32
}

type Config interface {
	GetBrokers() *[]string
}

func (p *Producer) Message(msg *model.ProducerMessage) *sarama.ProducerMessage {
	message := &sarama.ProducerMessage{}

	if msg.Topic != "" {
		message.Topic = msg.Topic
	}

	if msg.Key != "" {
		message.Key = sarama.StringEncoder(msg.Key)
	}

	if msg.Value != "" {
		message.Value = sarama.StringEncoder(msg.Value)
	}

	if !msg.Timestamp.IsZero() {
		message.Timestamp = msg.Timestamp
	}

	message.Partition = msg.Partition

	if len(msg.Headers) > 0 {
		var Headers []sarama.RecordHeader
		for i := range msg.Headers {
			Headers = append(Headers, msg.Headers[i])
		}
		message.Headers = Headers
	}

	return message
}

func (p *Producer) Push(msg *sarama.ProducerMessage) {
	p.AsyncProducer.Input() <- msg
}

func (p *Producer) MessagePush(msg *model.ProducerMessage) {
	p.Push(p.Message(msg))
}

func (p *Producer) GetPartition(k int32) int32 {
	return k % int32(len(p.Partitions))
}

func NewAsyncProducer(ctx context.Context, config Config) (*Producer, error) {
	prodConf := sarama.NewConfig()
	prodConf.Producer.Partitioner = sarama.NewManualPartitioner
	prodConf.Producer.RequiredAcks = sarama.NoResponse
	prodConf.Producer.Idempotent = false

	asyncProd, err := sarama.NewAsyncProducer(*config.GetBrokers(), prodConf)
	if err != nil {
		return nil, fmt.Errorf("sarama.NewAsyncProducer %w", err)
	}

	client, err := sarama.NewClient(*config.GetBrokers(), prodConf)
	if err != nil {
		return nil, fmt.Errorf("sarama.NewClient %w", err)
	}
	defer client.Close()

	partitions, err := client.Partitions(model.TopicLomsOrderEvents)
	if err != nil {
		return nil, fmt.Errorf("client.Partitions %w", err)
	}

	return &Producer{
		asyncProd,
		partitions,
	}, nil
}
