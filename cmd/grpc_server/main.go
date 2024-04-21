package main

import (
	"context"
	"flag"
	"log"

	"github.com/kenyako/auth/internal/app"
	"github.com/kenyako/auth/internal/config"
)

func main() {
	config.InitFlags()

	flag.Parse()
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init app: %v", err)
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("failed to run app: %v", err)
	}
}
