package app

import (
	"context"
	"os/signal"
	"route256/notifier/internal/infra/broker/consumer_group"
	"route256/notifier/internal/pkg/config"
	"route256/notifier/internal/pkg/logger"
	"syscall"
)

func Run(config *config.Config) {
	logger.New()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	consumerGroup, err := consumer_group.New(ctx, config)
	if err != nil {
		logger.Panicw(ctx, err.Error())
	}
	defer consumerGroup.Close()

	<-ctx.Done()
	// gracefulShutdown
	consumerGroup.Close()
	logger.Infow(ctx, "Graceful Shutdown")
}
