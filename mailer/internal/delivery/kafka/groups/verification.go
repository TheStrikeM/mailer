package groups

import (
	"fmt"
	"github.com/IBM/sarama"
	"log/slog"
	"mailer/internal/delivery"
	"mailer/pkg/e"
)

type VerificationConsumerGroupHandler struct {
	log         *slog.Logger
	authUsecase delivery.AuthUsecase
}

func NewVerificationConsumerGroupHandler(log *slog.Logger, authUsecase delivery.AuthUsecase) *VerificationConsumerGroupHandler {
	return &VerificationConsumerGroupHandler{
		log:         log,
		authUsecase: authUsecase,
	}
}

func (VerificationConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (VerificationConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (h VerificationConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	const op = "VerificationConsumerGroupHandler.ConsumeClaim()"
	for msg := range claim.Messages() {
		h.log.Debug(fmt.Sprintf("%s, partition - %d, offset - %d", msg.Value, int(msg.Partition), int(msg.Offset)))
		err := h.authUsecase.SendVerification(msg.Value)
		if err != nil {
			return e.WrapIfErr(op, err)
		}
		sess.MarkMessage(msg, "")
	}
	return nil
}
