package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	httpserver "github.com/Uikola/task-manager/internal/adapters/transport/http"
	"github.com/Uikola/task-manager/internal/config"
	"github.com/Uikola/task-manager/pkg/closer"
)

type App struct {
	serviceProvider *serviceProvider
	httpServer      *http.Server
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, fmt.Errorf("new app: %w", err)
	}

	return a, nil
}

func (a *App) Run(ctx context.Context) error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	logWriter := a.serviceProvider.AsyncLogWriter()

	logWriter.Start()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		a.serviceProvider.Logger().Info("starting HTTP server")
		if err := a.runHTTPServer(); err != nil {
			a.serviceProvider.Logger().Error("start HTTP server error", err)
		}
	}()

	<-stop

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	a.serviceProvider.Logger().Info("shutting down http server")
	if err := a.shutdownHTTPServer(ctx); err != nil {
		a.serviceProvider.Logger().Error("shutdown error", err)
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	if err := config.Load(); err != nil {
		return err
	}

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initHTTPServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return fmt.Errorf("init deps: %w", err)
		}
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initHTTPServer(_ context.Context) error {
	srv := httpserver.NewServer(a.serviceProvider.TaskHandler())

	httpServer := &http.Server{
		Addr:    a.serviceProvider.HTTPConfig().Address(),
		Handler: srv,
	}
	a.httpServer = httpServer

	return nil
}

func (a *App) runHTTPServer() error {
	if err := a.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("failed to start http server: %w", err)
	}

	return nil
}

func (a *App) shutdownHTTPServer(ctx context.Context) error {
	if err := a.httpServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to shutdown http server: %w", err)
	}

	return nil
}
