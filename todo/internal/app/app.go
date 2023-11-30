package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"todo/config" 
	loggerPackage "todo/pkg/logger"
	"todo/internal/database/postgres"
	"todo/internal/transport/rest/router"
	"todo/internal/repository"
	"todo/internal/service"

)

func Run() {
	logger := loggerPackage.Get()

	config, err := config.Load()
	if err != nil {
		logger.Errorf("Error loading config variables: %s", err.Error())
		return
	}

	db, err := postgres.NewPostgresClient(config.Postgres)
	if err != nil {
		logger.Fatal("Error loading database: %s", err.Error())
		return
	}
	defer db.Close()

	mainRouter := router.Init(config.App)

	repository := repository.NewTaskRepository(db)
	service := service.NewTaskService(repository)
	handler := NewTaskHandler(service)

	handler.InitializeExternalRoutes(mainRouter, handler)

	ctx,cancel := context.WithCancel(context.Background())

	host := fmt.Sprintf("%s:%d", config.Http.Host, config.Http.Port)

	docs.SwaggerInfo.Host = host

	srv := &http.Server{
		Addr:    host,
		Handler: mainRouter,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("Error starting HTTP server: %s\n", err.Error())
		}
	}()

	logger.Info("Server started on port: ", config.Http.Port)

	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Graceful shutdown server ....")

	ctx.cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Error server shutdown:", err.Error())
	}
	select {
	case <-ctx.Done():
		logger.Info("Server shutdown timed out")
	}
	logger.Info("Exiting")
}
