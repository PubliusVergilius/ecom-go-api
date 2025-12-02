package loggerfx

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Options(
	fx.Provide(
		New,
	),
)

func New() *zap.Logger {
	return zap.NewExample()
}
