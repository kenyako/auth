package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/brianvoe/gofakeit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
	desc "github.com/kenyako/auth/pkg/auth_v1"
)

const (
	grpcPort = 50051
	dbDSN    = "host=localhost port=54321 dbname=auth user=auth-user password=auth-password sslmode=disable"
)

type server struct {
	desc.UnimplementedUserAPIServer

	db *pgxpool.Pool
	qb sq.StatementBuilderType
}

func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {

	userName := req.GetName()
	userEmail := req.GetEmail()
	userPassword := req.GetPassword()
	userPasConfirm := req.GetPasswordConfirm()
	userRole := req.GetRole().String()

	builderInsert := s.qb.Insert("users").
		PlaceholderFormat(sq.Dollar).
		Columns("name", "email", "password", "password_confirm", "role").
		Values(userName, userEmail, userPassword, userPasConfirm, userRole).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
	}

	var userID int64

	err = s.db.QueryRow(ctx, query, args...).Scan(&userID)
	if err != nil {
		log.Fatalf("failed to insert user: %v", err)
	}

	log.Printf("inserted user with ID: %d", userID)

	return &desc.CreateResponse{
		Id: userID,
	}, nil
}

func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("Note id: %d", req.GetId())

	return &desc.GetResponse{
		Id:        req.GetId(),
		Name:      gofakeit.Name(),
		Email:     gofakeit.Email(),
		Role:      desc.UserRole_USER,
		CreatedAt: timestamppb.New(gofakeit.Date()),
		UpdatedAt: timestamppb.New(gofakeit.Date()),
	}, nil
}

func (s *server) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {

	newName := gofakeit.Name()
	newEmail := gofakeit.Email()
	newRole := desc.UserRole(gofakeit.Number(1, 0))

	log.Printf("Updating user with ID: %d", req.GetId())
	log.Printf("New name: %s\nNew Email: %s\nNew role: %v", newName, newEmail, newRole)

	return &emptypb.Empty{}, nil
}

func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {

	err := deleteUser(req.GetId())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func deleteUser(userID int64) error {
	return nil
}

func main() {
	ctx := context.Background()

	pool, err := pgxpool.Connect(ctx, dbDSN)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	desc.RegisterUserAPIServer(s, &server{
		db: pool,
		qb: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
