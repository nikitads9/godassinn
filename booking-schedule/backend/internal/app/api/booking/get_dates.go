package booking

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/nikitads9/godassinn/booking-schedule/backend/internal/app/api"
	"github.com/nikitads9/godassinn/booking-schedule/backend/internal/logger/sl"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

// GetVacantDates godoc
//
//	@Summary		Get vacant intervals
//	@Description	Responds with list of vacant intervals within month for selected offer.
//	@ID				getDatesByOfferID
//	@Tags			bookings
//	@Produce		json
//	@Param			offer_id path	int	true	"offer_id"	Format(int64) default(1)
//	@Success		200	{object}	api.GetVacantDatesResponse
//	@Failure		400	{object}	api.errResponse
//	@Failure		404	{object}	api.errResponse
//	@Failure		503	{object}	api.errResponse
//	@Router			/{offer_id}/get-vacant-dates [get]
func (i *Implementation) GetVacantDates(logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "api.booking.GetVacantDates"

		ctx := r.Context()
		requestID := middleware.GetReqID(ctx)

		log := logger.With(
			slog.String("op", op),
			slog.String("request_id", requestID),
		)
		ctx, span := i.tracer.Start(ctx, op, trace.WithAttributes(attribute.String("request_id", requestID)))
		defer span.End()

		offerID := chi.URLParam(r, "suite_id")
		if offerID == "" {
			span.RecordError(errNoOfferID)
			span.SetStatus(codes.Error, errNoOfferID.Error())
			log.Error("invalid request", sl.Err(errNoOfferID))
			api.WriteWithError(w, http.StatusBadRequest, errNoOfferID.Error())
			return
		}

		span.AddEvent("offerID extracted from path", trace.WithAttributes(attribute.String("id", offerID)))

		id, err := strconv.ParseInt(offerID, 10, 64)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			log.Error("invalid request", sl.Err(err))
			api.WriteWithError(w, http.StatusBadRequest, api.ErrParse.Error())
			return
		}

		if id == 0 {
			span.RecordError(errNoOfferID)
			span.SetStatus(codes.Error, errNoOfferID.Error())
			log.Error("invalid request", sl.Err(errNoOfferID))
			api.WriteWithError(w, http.StatusBadRequest, errNoOfferID.Error())
			return
		}

		span.AddEvent("offerID parsed")

		dates, err := i.booking.GetVacantDates(ctx, id)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			log.Error("internal error", sl.Err(err))
			api.WriteWithError(w, GetErrorCode(err), err.Error())
			return
		}

		span.AddEvent("vacant dates acquired", trace.WithAttributes(attribute.Int("quantity", len(dates))))
		log.Info("vacant dates acquired", slog.Int("quantity: ", len(dates)))

		api.WriteWithStatus(w, http.StatusOK, api.GetVacantDatesResponse{
			Intervals: dates,
		})
	}
}
