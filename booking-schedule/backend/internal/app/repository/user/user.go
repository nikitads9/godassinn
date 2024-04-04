package user

import (
	"context"
	"errors"
	"log/slog"

	"github.com/nikitads9/godassinn/booking-schedule/backend/internal/app/model"
	"github.com/nikitads9/godassinn/booking-schedule/backend/internal/pkg/db"

	"github.com/jackc/pgx/v5/pgconn"
	"go.opentelemetry.io/otel/trace"
)

type Repository interface {
	CreateUser(ctx context.Context, user *model.User) (int64, error)
	GetUser(ctx context.Context, userID int64) (*model.User, error)
	GetUserByNickname(ctx context.Context, nickName string) (*model.User, error)
	EditUser(ctx context.Context, user *model.UpdateUserInfo) error
	DeleteUser(ctx context.Context, userID int64) error
}

var (
	ErrAlreadyExists = errors.New("this user already exists")
	ErrDuplicate     = "ERROR: duplicate key value violates unique constraint \"users_login_key\" (SQLSTATE 23505)"

	ErrNotFound       = errors.New("no user with this id")
	ErrNoRowsAffected = errors.New("no database entries affected by this operation")

	ErrQuery        = errors.New("failed to execute query")
	ErrQueryBuild   = errors.New("failed to build query")
	ErrPgxScan      = errors.New("failed to read database response")
	ErrNoConnection = errors.New("could not connect to database")
	pgNoConnection  = new(*pgconn.ConnectError)
)

type repository struct {
	client db.Client
	log    *slog.Logger
	tracer trace.Tracer
}

func NewUserRepository(client db.Client, log *slog.Logger, tracer trace.Tracer) Repository {
	return &repository{
		client: client,
		log:    log,
		tracer: tracer,
	}
}
