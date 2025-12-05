package main

import (
	"net/http"
	"time"

	"github.com/PubliusVergilius/ecom-go-api/pkg/loggerfx"
	"github.com/PubliusVergilius/ecom-go-api/pkg/routerfx"
	"github.com/PubliusVergilius/ecom-go-api/pkg/serverfx"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

// Route is an http.Handler that knows the mux pattern
// under which it will be registered.

/*
* TODO: If no addr is provided map, to the default
* localhost:8080
 */
type config struct {
	addr     string
	dbConfig dbConfig
}

/* TODO: Make throw an error if a dbConfig instance is not provided
 * Inside the Fx context
**/
type dbConfig struct {
	dsn string
}

/*
* TODO: Make throw an error if no dbDriver is provided
 */
type application struct {
	config  config
	timeout time.Time // Router does not know about this config
}

func (app *application) Run() {
	logger, _ := zap.NewProduction()
	// logger := loggerfx.New()

	fx.New(
		fx.WithLogger(func() fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger}
		}),
		loggerfx.Module,
		serverfx.Module,
		routerfx.Module,
		fx.Invoke(func(*http.Server) {}),
	).Run()

}
