package transport

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"onelab/internal/config"
	"onelab/internal/jsonlog"
)

func NewServer(cfg *config.Config, logger *jsonlog.Logger, handler http.Handler) error {
	srv := &http.Server{
		Addr:         cfg.HTTP.HttpPort,
		Handler:      handler,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	shotdownError := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <-quit

		logger.PrintInfo("shutting down server signal:", map[string]string{
			"signal": s.String(),
		})

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		shotdownError <- srv.Shutdown(ctx)
	}()

	logger.PrintInfo("starting server", map[string]string{
		"addr": srv.Addr,
	})

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shotdownError
	if err != nil {
		return err
	}

	logger.PrintInfo("stopped server", map[string]string{
		"addr": srv.Addr,
	})

	return nil
}
