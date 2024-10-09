package instruments

import (
	"log/slog"
	"os"
	"time"
)

func NewLogger() *slog.Logger {
	// todo: configure this options - log level, logs format
	opts := &slog.HandlerOptions{
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			switch a.Key {
			case slog.TimeKey:
				a.Value = slog.StringValue(a.Value.Time().Format(time.DateTime))
				return a
			case slog.MessageKey:
				a.Key = "message"
				return a
			default:
				return a
			}
		},
	}
	return slog.New(slog.NewJSONHandler(os.Stdout, opts))
}
