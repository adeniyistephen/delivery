package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/adeniyistephen/delivery/database"
	"github.com/adeniyistephen/delivery/delivery"
	"github.com/ardanlabs/conf/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var build = "Development"

func New(service string) (*zap.SugaredLogger, error) {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout"}
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.DisableStacktrace = true
	config.InitialFields = map[string]interface{}{
		"service": service,
	}

	log, err := config.Build()
	if err != nil {
		return nil, err
	}

	return log.Sugar(), nil
}

func main() {

	log, err := New("Delivery")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer log.Sync()

	if err := run(log); err != nil {
		log.Errorw("main: error:", err)
		os.Exit(1)
	}
}

func run(log *zap.SugaredLogger) error {
	cfg := struct {
		conf.Version
		DB struct {
			User         string `conf:"default:mydb"`
			Password     string `conf:"default:pass"`
			Host         string `conf:"default:db"`
			Name         string `conf:"default:delivery"`
			Port         int    `conf:"default:3306"`
			MaxIdleConns int    `conf:"default:0"`
			MaxOpenConns int    `conf:"default:0"`
			DisableTLS   bool   `conf:"default:true"`
		}
	}{
		Version: conf.Version{
			Build: build,
		},
	}
	const prefix = "Delivery"
	help, err := conf.Parse(prefix, &cfg)
	if err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			fmt.Println(help)
			return nil
		}
		return fmt.Errorf("parsing config: %w", err)
	}

	//====================
	// Database Support
	db, err := database.Open(database.Config{
		User:         cfg.DB.User,
		Password:     cfg.DB.Password,
		Host:         cfg.DB.Host,
		Port:         cfg.DB.Port,
		Name:         cfg.DB.Name,
		MaxIdleConns: cfg.DB.MaxIdleConns,
		MaxOpenConns: cfg.DB.MaxOpenConns,
		DisableTLS:   cfg.DB.DisableTLS,
	})
	if err != nil {
		return fmt.Errorf("connecting to db: %w", err)
	}
	defer func() {
		log.Infof("shutdown", "status", "stopping database support", "host", cfg.DB.Host)
		db.Close()
	}()

	//====================
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	api := http.Server{
		Addr:    ":8080",
		Handler: delivery.Test(log, db),
		IdleTimeout: 5 * time.Second,
	}

	serverErrors := make(chan error, 1)

	go func() {
		log.Infow("main: API listening on %s", api.Addr)
		serverErrors <- api.ListenAndServe()
	}()

	// Shutdown
	// Blocking main and waiting for shutdown.
	select {
	case err := <-serverErrors:
		return errors.New(err.Error())

	case sig := <-shutdown:
		log.Infow("main: %v: Start shutdown", sig)

		// Give outstanding requests a deadline for completion.
		ctx, cancel := context.WithTimeout(context.Background(), api.IdleTimeout)
		defer cancel()

		// Asking listener to shutdown and shed load.
		if err := api.Shutdown(ctx); err != nil {
			api.Close()
			return errors.New(err.Error())
		}
	}
	return nil
}
