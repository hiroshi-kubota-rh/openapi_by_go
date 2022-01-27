package handler

import (
	"net/http"

	"newview-backend-go/app/handler/impressionTags"
	"newview-backend-go/app/handler/test"

	"newview-backend-go/app/app"

	"newview-backend-go/app/handler/accounts"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func NewRouter(app *app.App) http.Handler {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(newCORS().Handler)

	r.Mount("/test", test.NewRouter())

	r.Mount("/accounts", accounts.NewRouter(app))

	r.Mount("/impression_tags", impressionTags.NewRouter(app))
	return r
}

func newCORS() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodHead,
			http.MethodPut,
			http.MethodPatch,
			http.MethodPost,
			http.MethodDelete,
			http.MethodOptions,
		},
	})
}
