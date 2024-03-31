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
	"github.com/georgysavva/scany/pgxscan"
	"github.com/go-chi/chi/middleware"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func (r *repository) GetVacantOffers(ctx context.Context, startDate time.Time, endDate time.Time) ([]*model.Offer, error) {
	const op = "repository.booking.GetVacantOffers"

	requestID := middleware.GetReqID(ctx)

	log := r.log.With(
		slog.String("op", op),
		slog.String("request_id", requestID),
	)

	ctx, span := r.tracer.Start(ctx, op, trace.WithAttributes(attribute.String("request_id", requestID)))
	defer span.End()

	builder := sq.Select("*").
		Distinct().
		From(t.OfferTable).
		PlaceholderFormat(sq.Dollar)
	subQuery, subQueryArgs, err := sq.Select("1").
		From(t.BookingTable + " AS e").
		Where(sq.And{
			sq.ConcatExpr("e."+t.OfferID+"=", t.OfferTable+".id"),
			sq.Or{sq.And{
				sq.Lt{"e." + t.StartDate: startDate},
				sq.Gt{"e." + t.EndDate: endDate},
			},
				sq.And{
					sq.Lt{"e." + t.StartDate: endDate},
					sq.Gt{"e." + t.EndDate: startDate},
				}},
		},
		).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error("failed to build subquery", sl.Err(err))
		return nil, ErrQueryBuild
	}

	builder = builder.Where("NOT EXISTS ("+subQuery+") OR NOT EXISTS (SELECT DISTINCT "+t.OfferID+" FROM "+t.BookingTable+")", subQueryArgs...)

	query, args, err := builder.ToSql()
	if err != nil {
		log.Error("failed to build a query", sl.Err(err))
		return nil, ErrQueryBuild
	}

	span.AddEvent("query built")

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	var res []*model.Offer
	err = r.client.DB().SelectContext(ctx, &res, q, args...)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		if errors.As(err, pgNoConnection) {
			log.Error("no connection to database host", sl.Err(err))
			return nil, ErrNoConnection
		}
		if pgxscan.NotFound(err) {
			log.Error("no vacant offers within this period", sl.Err(err))
			return nil, ErrNotFound
		}
		log.Error("query execution error", sl.Err(err))
		return nil, ErrQuery
	}

	span.AddEvent("query successfully executed")

	return res, nil
}
