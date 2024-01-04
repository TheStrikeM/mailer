package app

import (
	"fmt"
	"log/slog"
	"mailer/internal/config"
	"mailer/internal/delivery/kafka"
	"mailer/internal/repository"
	"mailer/internal/usecase"
	"mailer/pkg/logger/handlers/slogpretty"
	"os"
	"sync"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func Run(cfg *config.Config) {
	fmt.Println("Starting loading, opening config...")

	fmt.Println("Start loading logger...")
	logger := setupLogger(cfg.Env)
	logger.Debug("Yey! Logger, Config enabled!")

	var wg sync.WaitGroup
	wg.Add(1)
	RunKafkaVerification(logger, cfg, &wg)
	wg.Wait()
}

func RunKafkaVerification(log *slog.Logger, cfg *config.Config, wg *sync.WaitGroup) {

	authRepo := repository.NewAuthRepository(log, cfg)
	authUsecase := usecase.NewAuthUsecase(log, authRepo)
	kafkaDelivery := kafka.NewKafkaConsumer(cfg, log, authUsecase)
	go func() {
		defer wg.Done()
		err := kafkaDelivery.StartVerificationReading()
		if err != nil {
			panic(err)
		}
	}()
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = setupPrettySlog()
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
