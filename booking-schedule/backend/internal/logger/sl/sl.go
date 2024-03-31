package sl

import (
	"log/slog"

	_ "github.com/nikitads9/godassinn/booking-schedule/backend/internal/logger/handlers/slogdiscard"
)

func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}
