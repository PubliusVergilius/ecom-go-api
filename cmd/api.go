package main

import (
	"net/http"
	"time"

	"github.com/PubliusVergilius/ecom-go-api/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Route is an http.Handler that knows the mux pattern
// under which it will be registered.
type Route interface {
	http.Handler

	// Pattern reports the path at which this is registered.
	Pattern() string
}

/*
* TODO: Make throw an error if no dbDriver is provided
 */
type application struct {
	config config
	// logger
	// dbDriver
	timeout time.Time
}

func (app *application) Mount(routes []Route) *chi.Mux {

	router := chi.NewRouter()

	// A good base middleware stack
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	if app.timeout.IsZero() {
		router.Use(middleware.Timeout(60 * time.Second))
	}

	for _, r := range routes {
		router.Handle(r.Pattern(), r)
	}

	return router
}

func AsRoute(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(Route)),
		fx.ResultTags(`group:"routes"`),
	)
}

func (app *application) run() error {

	fx.New(
		fx.Provide(
			NewHTTPServer,
			zap.NewExample,
			fx.Annotate(
				app.Mount,
				fx.ParamTags(`group:"routes"`),
			),
			AsRoute(handlers.NewHealthHandler),
			AsRoute(handlers.NewHelloHandler),
		),
		fx.Invoke(func(*http.Server) {}),
	).Run()

	return nil
}

/*
* TODO: If no addr is provided map to the default
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
