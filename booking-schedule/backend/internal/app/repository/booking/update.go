package booking

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

func (r *repository) UpdateBooking(ctx context.Context, mod *model.BookingInfo) error {
	const op = "repository.booking.UpdateBooking"

	requestID := middleware.GetReqID(ctx)

	log := r.log.With(
		slog.String("op", op),
		slog.String("request_id", requestID),
	)

	ctx, span := r.tracer.Start(ctx, op, trace.WithAttributes(attribute.String("request_id", requestID)))
	defer span.End()

	builder := sq.Update(t.BookingTable).
		Set(t.UpdatedAt, time.Now()).
		Set("start_date", mod.StartDate).
		Set("end_date", mod.EndDate).
		Set("suite_id", mod.OfferID).
		Where(sq.And{
			sq.Eq{t.ID: mod.ID},
			sq.Eq{t.UserID: mod.UserID},
		}).
		PlaceholderFormat(sq.Dollar)

	if mod.NotifyAt != 0 {
		builder = builder.Set(t.NotifyAt, mod.NotifyAt)
	}

	query, args, err := builder.ToSql()
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
		log.Error("query execution error", sl.Err(err))
		return ErrQuery
	}

	if result.RowsAffected() == 0 {
		span.RecordError(ErrNoRowsAffected)
		span.SetStatus(codes.Error, ErrNoRowsAffected.Error())
		log.Error("update unsuccessful", sl.Err(ErrNoRowsAffected))
		return ErrNotFound
	}

	span.AddEvent("query successfully executed")

	return nil
}
