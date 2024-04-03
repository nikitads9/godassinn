package user

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/nikitads9/godassinn/booking-schedule/backend/internal/app/model"
	t "github.com/nikitads9/godassinn/booking-schedule/backend/internal/app/repository/table"
	"github.com/nikitads9/godassinn/booking-schedule/backend/internal/logger/sl"
	"github.com/nikitads9/godassinn/booking-schedule/backend/internal/pkg/db"

	sq "github.com/Masterminds/squirrel"
	"github.com/go-chi/chi/middleware"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func (r *repository) EditUser(ctx context.Context, user *model.UpdateUserInfo) error {
	const op = "bookings.repository.EditUser"

	requestID := middleware.GetReqID(ctx)

	log := r.log.With(
		slog.String("op", op),
		slog.String("request_id", requestID),
	)

	ctx, span := r.tracer.Start(ctx, op, trace.WithAttributes(attribute.String("request_id", requestID)))
	defer span.End()

	builder := sq.Update(t.UserTable).
		Set(t.UpdatedAt, time.Now()).
		Where(sq.Eq{t.ID: user.ID})

	if user.Name.Valid {
		builder = builder.Set(t.Name, user.Name.String)
	}

	if user.Login.Valid {
		builder = builder.Set(t.Login, user.Login.String)
	}

	if user.PhoneNumber.Valid {
		builder = builder.Set(t.PhoneNumber, user.PhoneNumber.String)
	}

	if user.Password.Valid {
		builder = builder.Set(t.Password, user.Password.String)
	}

	query, args, err := builder.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error("failed to build a query", sl.Err(err))
		return ErrQueryBuild
	}

	span.AddEvent("query built")

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	result, err := r.client.DB().ExecContext(ctx, q, args...)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		if errors.As(err, pgNoConnection) {
			log.Error("no connection to database host", sl.Err(err))
			return ErrNoConnection
		}
		if errors.As(err, &ErrDuplicate) {
			log.Error("this user already exists", sl.Err(err))
			return ErrAlreadyExists
		}
		log.Error("query execution error", sl.Err(err))
		return ErrQuery
	}

	if result.RowsAffected() == 0 {
		span.RecordError(ErrNoRowsAffected)
		span.SetStatus(codes.Error, ErrNoRowsAffected.Error())
		log.Error("unsuccessful update", sl.Err(ErrNoRowsAffected))
		return ErrNotFound
	}

	span.AddEvent("query successfully executed")

	return nil
}
