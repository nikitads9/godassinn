package metrics

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
)

func handleErr(err error) {
	if err != nil {
		otel.Handle(err)
	}
}

// NewMetricMiddleware creates the middleware that will record all
// HTTP-related metrics.
func NewMetricMiddleware(meter metric.Meter) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		durationHistogram, err := meter.Float64Histogram(
			"http.server.latency",
			metric.WithUnit("ms"),
			metric.WithDescription("Measures the duration of inbound HTTP requests."),
		)
		handleErr(err)

		requestCounter, err := meter.Int64Counter(
			"http.server.requests",
			metric.WithDescription("Measures the amount of HTTP requests received"),
		)
		handleErr(err)

		return &httpMetricMiddleware{
			next:                     next,
			requestDurationHistogram: durationHistogram,
			requestCounter:           requestCounter,
		}
	}
}

// httpMetricMiddleware executes the HTTP endpoint while keeping track
// of how much time it took to execute and add some extra routing information
// to all metrics
type httpMetricMiddleware struct {
	next                     http.Handler
	requestDurationHistogram metric.Float64Histogram
	requestCounter           metric.Int64Counter
	//errorRate                metric.Float64ObservableGauge
}

func (m *httpMetricMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rw := NewStatusCodeCapturerWriter(w)

	initialTime := time.Now()
	m.next.ServeHTTP(rw, r)
	duration := time.Since(initialTime)

	pathTemplate := chi.RouteContext(r.Context()).RoutePattern()
	metricAttributes := attribute.NewSet(
		semconv.HTTPRouteKey.String(pathTemplate),
		semconv.HTTPRequestMethodKey.String(r.Method),
		semconv.HTTPStatusCodeKey.Int(rw.statusCode),
	)

	m.requestCounter.Add(r.Context(), 1, metric.WithAttributeSet(metricAttributes))

	/* 	if _, err := m.meter.RegisterCallback(
	   		func(ctx context.Context, o metric.Observer) error {
	   			o.ObserveFloat64(m.errorRate, rand.Float64())
	   			return nil
	   		},
	   		errorRate,
	   	); err != nil {
	   		panic(err)
	   	} */

	m.requestDurationHistogram.Record(
		r.Context(),
		float64(duration.Milliseconds()),
		metric.WithAttributeSet(metricAttributes),
	)
}

// NewStatusCodeCapturerWriter creates an HTTP.ResponseWriter capable of
// capture the HTTP response status code.
func NewStatusCodeCapturerWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{w, http.StatusOK}
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *responseWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}
