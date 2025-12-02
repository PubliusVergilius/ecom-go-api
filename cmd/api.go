package main

import (
	"net/http"
	"time"

	"github.com/PubliusVergilius/ecom-go-api/loggerfx"
	"github.com/PubliusVergilius/ecom-go-api/routerfx"
	"github.com/PubliusVergilius/ecom-go-api/serverfx"
	"go.uber.org/fx"
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
	config config
	// logger //TODO: pass by injection
	// dbDriver //TODO: pass by injection
	timeout time.Time
}

func (app *application) Run() {

	fx.New(
		loggerfx.Module,
		serverfx.Module,
		routerfx.Module,

		fx.Invoke(func(*http.Server) {}),
	).Run()

}
