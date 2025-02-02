package main

import (
	"context"
	"github.com/hussainmuzamil/students-api/internal/config"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := config.MustLoad()

	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello World"))
		if err != nil {
			log.Fatal(err)
		}
	})

	//setup server
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}
	slog.Info("Starting server...", slog.String("addr", cfg.Addr))
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	<-done
	slog.Info("Shutting down server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("Failed to shutdown server gracefully", slog.String("error", err.Error()))
	}
	slog.Info("Server gracefully stopped")
}
