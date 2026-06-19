package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"marginalia/internal/server"
	"marginalia/internal/storage"
	"marginalia/internal/store"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	dbPath := os.Getenv("DATABASE_PATH")
	if dbPath == "" {
		dbPath = store.DefaultPath
	}
	db, err := store.Open(ctx, dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := store.Migrate(ctx, db); err != nil {
		log.Fatal(err)
	}
	log.Printf("store: sqlite ready at %s", dbPath)

	storagePath := os.Getenv("STORAGE_PATH")
	if storagePath == "" {
		storagePath = storage.DefaultPath
	}
	log.Printf("storage: library path %s", storagePath)

	cfg, err := server.ConfigFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	srv := server.New(cfg, db, storagePath)
	if err := srv.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
