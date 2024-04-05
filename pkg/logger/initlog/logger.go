package initlog

import (
	"log/slog"
	"os"

	"github.com/xlzpm/pkg/logger/prettylog"
)

func InitLogger() *slog.Logger {
	opts := prettylog.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)
	return slog.New(handler)
}
