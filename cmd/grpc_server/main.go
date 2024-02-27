package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/brianvoe/gofakeit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	desc "github.com/kenyako/auth/pkg/auth_v1"
)

const (
	grpcPort = 50051
	dbDSN    = "host=localhost port=54321 dbname=auth user=auth-user password=auth-password sslmode=disable"
)

type User struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	PassConf  string    `db:"password_confirm"`
	Role      string    `db:"role"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

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
	userID := req.GetId()

	builderSelect := s.qb.Select("id", "name", "email", "role", "created_at", "updated_at").
		From("users").
		Where(sq.Eq{
			"id": userID,
		})

	query, args, err := builderSelect.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
	}

	row, err := s.db.Query(ctx, query, args...)
	if err != nil {
		log.Fatalf("failed to get user from query: %v", err)
	}

	user, err := pgx.CollectOneRow(row, pgx.RowToAddrOfStructByNameLax[User])
	if err != nil {
		log.Fatalf("failed to collect user from db: %v", err)
	}

	roleNum := desc.UserRole_value[user.Role]

	userData := desc.GetResponse{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      desc.UserRole(roleNum),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}

	return &userData, nil
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

	return &emptypb.Empty{}, nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// конфиг с ДСН
	pgxCfg, err := pgxpool.ParseConfig(dbDSN)
	if err != nil {
		log.Fatalf("failed to patde config: %v", err)
	}

	// создание нового соединения с бд через конфиг
	pool, err := pgxpool.NewWithConfig(ctx, pgxCfg)
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}

	// проверка на то, что база откликается и работает
	err = pool.Ping(ctx)
	if err != nil {
		log.Fatalf("ping to postgres failed: %v", err)
	}

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
