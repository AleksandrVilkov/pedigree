package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"pedigree/internal/storage/postgreSQL/storage"
	endpoints "pedigree/internal/transport/http/auth"
	"pedigree/internal/usecase"
	"time"
)

type App struct {
	httpServer *http.Server

	AuthUseCase *usecase.AuthUsecase
}

func NewApp() *App {
	userStorage := storage.UserStorage{}
	return &App{
		AuthUseCase: usecase.NewAuthUsecase(
			&userStorage,
			viper.GetString("auth.hash_salt"),
			[]byte(viper.GetString("auth.signing_key")),
			viper.GetDuration("auth.token.ttl"),
		)}
}

func (a *App) Run(port string) error {
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	//Регистрируем эндпоинты
	endpoints.RegisterHTTPEndpoints(router, *a.AuthUseCase)

	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatal("Failed listen and serve on "+a.httpServer.Addr+" port: %+v", err)
		}

	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)
	<-quit
	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()
	return a.httpServer.Shutdown(ctx)
}
