package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
	conn, err := pgx.Connect(context.Background(), "user=pqgotest dbname=pqgotest sslmode=verify-full")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}

	l := slog.New(slog.NewJSONHandler(os.Stdout, opts))
	l.Info("test message", "arg1", "arg2")
}
