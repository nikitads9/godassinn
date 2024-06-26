package booking

import (
	"context"
	"errors"
	"log/slog"
	"strings"
	"time"

	"github.com/nikitads9/godassinn/booking-schedule/backend/internal/app/model"
	t "github.com/nikitads9/godassinn/booking-schedule/backend/internal/app/repository/table"
	"github.com/nikitads9/godassinn/booking-schedule/backend/internal/logger/sl"
	"github.com/nikitads9/godassinn/booking-schedule/backend/internal/pkg/db"

	"github.com/go-chi/chi/middleware"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	sq "github.com/Masterminds/squirrel"
	"github.com/gofrs/uuid"
)

func (r *repository) AddBooking(ctx context.Context, mod *model.BookingInfo) (uuid.UUID, error) {
	const op = "repository.booking.AddBooking"

	requestID := middleware.GetReqID(ctx)

	log := r.log.With(
		slog.String("op", op),
		slog.String("request_id", requestID),
	)

	ctx, span := r.tracer.Start(ctx, op, trace.WithAttributes(attribute.String("request_id", requestID)))
	defer span.End()

	var builder sq.InsertBuilder

	newID, err := uuid.NewV4()
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error("failed to generate uuid", sl.Err(err))
		return uuid.Nil, ErrUuid
	}

	span.AddEvent("uuid generated")

	if mod.NotifyAt != 0 {
		builder = sq.Insert(t.BookingTable).
			Columns(t.ID, t.UserID, t.OfferID, t.StartDate, t.EndDate, t.CreatedAt, t.NotifyAt).
			Values(newID, mod.UserID, mod.OfferID, mod.StartDate, mod.EndDate, time.Now(), mod.NotifyAt)
	} else {
		builder = sq.Insert(t.BookingTable).
			Columns(t.ID, t.UserID, t.OfferID, t.StartDate, t.EndDate, t.CreatedAt).
			Values(newID, mod.UserID, mod.OfferID, mod.StartDate, mod.EndDate, time.Now())
	}

	query, args, err := builder.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error("failed to build a query", sl.Err(err))
		return uuid.Nil, ErrQueryBuild
	}

	span.AddEvent("query built")

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	_, err = r.client.DB().ExecContext(ctx, q, args...)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		if errors.As(err, pgNoConnection) {
			log.Error("no connection to database host", sl.Err(err))
			return uuid.Nil, ErrNoConnection
		}
		if strings.EqualFold(err.Error(), ErrNoSuchUser) {
			return uuid.Nil, ErrUnauthorized
		}
		log.Error("query execution error", sl.Err(err))
		return uuid.Nil, ErrQuery
	}

	span.AddEvent("query successfully executed")

	return newID, nil
}
