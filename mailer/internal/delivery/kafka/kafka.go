package kafka

import (
	"context"
	"github.com/IBM/sarama"
	"log/slog"
	"mailer/internal/config"
	"mailer/internal/delivery"
	"mailer/internal/delivery/kafka/groups"
	"mailer/pkg/e"
)

type KafkaConsumer struct {
	config      *config.Config
	log         *slog.Logger
	authUsecase delivery.AuthUsecase
}

func NewKafkaConsumer(cfg *config.Config, log *slog.Logger, authUsecase delivery.AuthUsecase) *KafkaConsumer {
	return &KafkaConsumer{
		config:      cfg,
		log:         log,
		authUsecase: authUsecase,
	}
}

func (kc *KafkaConsumer) StartVerificationReading() (err error) {
	const op = "KafkaConsumer.StartVerificationReading()"
	defer func() { err = e.WrapIfErr(op, err) }()

	consumerGroup, err := sarama.NewConsumerGroup(
		[]string{kc.config.Kafka.Addr},
		kc.config.Kafka.VerificationConsumerGroup,
		nil,
	)
	if err != nil {
		return err
	}
	ctx := context.Background()
	handler := groups.NewVerificationConsumerGroupHandler(kc.log, kc.authUsecase)
	for {
		err := consumerGroup.Consume(ctx, []string{kc.config.Kafka.VerificationTopic}, *handler)
		if err != nil {
			return err
		}
	}
}
