package main

import (
	"context"
	"flag"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kenyako/auth/internal/config"
	"github.com/kenyako/auth/internal/config/env"

	authAPI "github.com/kenyako/auth/internal/api/auth"
	authRepo "github.com/kenyako/auth/internal/repository/auth"
	authService "github.com/kenyako/auth/internal/service/auth"

	desc "github.com/kenyako/auth/pkg/auth_v1"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

func main() {
	flag.Parse()
	ctx := context.Background()

	err := config.Load(configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	grpcConfig, err := env.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to get grpc config: %v", err)
	}

	pgConfig, err := env.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to get pg config: %v", err)
	}

	lis, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	pool, err := pgxpool.New(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatalf("failed to create pool: %v", err)
	}
	defer pool.Close()

	authRepo := authRepo.NewRepository(pool)
	authSrv := authService.NewService(authRepo)

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserAPIServer(s, authAPI.NewImplementation(authSrv))

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
