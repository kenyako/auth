package app

import (
	"context"
	"log"

	"github.com/kenyako/auth/internal/client/db"
	"github.com/kenyako/auth/internal/client/db/pg"
	"github.com/kenyako/auth/internal/client/db/transaction"
	"github.com/kenyako/auth/internal/closer"
	"github.com/kenyako/auth/internal/config"
	"github.com/kenyako/auth/internal/config/env"
	"github.com/kenyako/auth/internal/repository"
	"github.com/kenyako/auth/internal/service"

	userAPI "github.com/kenyako/auth/internal/api/user"
	userRepo "github.com/kenyako/auth/internal/repository/user"
	userService "github.com/kenyako/auth/internal/service/user"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	dbClient       db.Client
	txManager      db.TxManager
	authRepository repository.UserRepository

	userService service.UserService

	authImpl *userAPI.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {

	if s.pgConfig == nil {
		cfg, err := env.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {

	if s.grpcConfig == nil {
		cfg, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) DBCClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %v", err)
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBCClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) AuthRepository(ctx context.Context) repository.UserRepository {
	if s.authRepository != nil {
		s.authRepository = userRepo.NewRepository(s.DBCClient(ctx))
	}

	return s.authRepository
}

func (s *serviceProvider) AuthService(ctx context.Context) service.UserService {
	if s.userService == nil {
		s.userService = userService.NewService(
			s.AuthRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.userService
}

func (s *serviceProvider) AuthImpl(ctx context.Context) *userAPI.Implementation {
	if s.authImpl == nil {
		s.authImpl = userAPI.NewImplementation(s.AuthService(ctx))
	}

	return s.authImpl
}
