package main

import (
	"auth/adapters/controllers/routes"
	routes_v1 "auth/adapters/controllers/routes/v1"
	domain_config "auth/domains/config"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	setUpDataSource()
}

func setUpDataSource() {
	time.Local, _ = time.LoadLocation("America/Sao_Paulo")
	h := echo.New()

	envs := domain_config.LoadEnvVars()

	routes.SetUpHealthRoute(h)
	routes_v1.SetUpAuthRoute(h)

	err := start(h, envs)
	if err != nil {
		fmt.Println("error up api")
	}
}

func start(handler http.Handler, envs *domain_config.Environments) error {
	srv := &http.Server{
		ReadTimeout:  envs.SelfTimeOut,
		WriteTimeout: envs.SelfTimeOut,
		Addr:         ":" + envs.SelfPort,
		Handler:      handler,
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	go func() {
		<-ctx.Done()
		fmt.Println("stopping api")
		err := srv.Shutdown(context.Background())
		if err != nil {
			panic(err)
		}
	}()

	fmt.Println("api running on port ", envs.SelfPort)
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}

	return nil
}
