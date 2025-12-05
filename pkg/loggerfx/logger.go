package loggerfx

import (
	"os"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

var env string

func init() {
	env = os.Getenv("ENV")
	if env == "" {
		env = "PROD"
	}
}

var Module = fx.Module(
	"logger",
	fx.Provide(
		New,
	),
)

func New() *zap.Logger {

	if logger, err := zap.NewDevelopment(); env == "DEV" {
		if err != nil {
			panic(err.Error())
		}

		return logger
	}

	// Default logger
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err.Error())
	}

	return logger
}
