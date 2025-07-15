package verbosityutil

import (
	"fmt"
	"log/slog"
	"os"

	appcontext "likenft-indexer/cmd/worker/context"

	"github.com/spf13/cobra"
)

type Verbosity string

const (
	VerbosityDefault Verbosity = "default"
	VerbosityDebug   Verbosity = "debug"
	VerbosityInfo    Verbosity = "info"
	VerbosityWarn    Verbosity = "warn"
	VerbosityError   Verbosity = "error"
)

func (v Verbosity) LogLevel() slog.Level {
	if v == VerbosityDefault {
		return slog.LevelDebug
	}
	if v == VerbosityDebug {
		return slog.LevelDebug
	}
	if v == VerbosityInfo {
		return slog.LevelInfo
	}
	if v == VerbosityWarn {
		return slog.LevelWarn
	}
	if v == VerbosityError {
		return slog.LevelError
	}
	panic(fmt.Errorf("err unknown verbosity %s", v))
}

func (v Verbosity) GetLogger(defaultLogger *slog.Logger) *slog.Logger {
	if v == VerbosityDefault {
		return defaultLogger
	}
	loggerHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: v.LogLevel(),
	})
	return slog.New(loggerHandler)
}

func GetLoggerFromCmd(cmd *cobra.Command) (*slog.Logger, error) {
	defaultLogger := appcontext.LoggerFromContext(cmd.Context())
	verbosityStr, err := cmd.Flags().GetString("verbose")
	if err != nil {
		return nil, err
	}
	verbosity := Verbosity(verbosityStr)
	logger := verbosity.GetLogger(defaultLogger)
	return logger, nil
}
