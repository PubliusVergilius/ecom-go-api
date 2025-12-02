package routerfx

import (
	"net/http"

	"github.com/PubliusVergilius/ecom-go-api/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		fx.Annotate(
			New,
			fx.ParamTags(`group:"routes"`),
		),
		AsRoute(handlers.NewHealthHandler),
		AsRoute(handlers.NewHelloHandler),
	),
)

type Route interface {
	http.Handler

	// Pattern reports the path at which this is registered.
	Pattern() string
}

func New(routes []Route) *chi.Mux {

	router := chi.NewRouter()

	// A good base middleware stack
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	/*
		if r.timeout.IsZero() {
			router.Use(middleware.Timeout(60 * time.Second))
		}
	*/

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
