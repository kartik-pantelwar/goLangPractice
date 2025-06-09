package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kartik-pantelwar/new-api.git/internal/config"
	"github.com/kartik-pantelwar/new-api.git/internal/http/handlers/student"
)

func main() {
	cnf := config.MustLoad()
	router := http.NewServeMux()
	router.HandleFunc("POST /api/students", student.New())
	server := http.Server{
		Addr:    cnf.Addr,
		Handler: router,
	}

	slog.Info("Server started", slog.String("at- ", cnf.Addr))

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("Unable to start server: %s", err)
		}

	}()
	<-done
	slog.Info("shhutting down server!")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("failed to shutdown", slog.String("error", err.Error()))
	}
	slog.Info("Server shutdown successfully")
}
