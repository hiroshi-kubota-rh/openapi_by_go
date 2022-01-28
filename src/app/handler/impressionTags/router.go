package impressionTags

import (
	"net/http"

	"newview-backend-go/app/app"

	"github.com/go-chi/chi"
)

// Implementation of handler
type handler struct {
	app *app.App
}

// Create Handler for `/impression_tag/`
func NewRouter(app *app.App) http.Handler {
	r := chi.NewRouter()

	h := &handler{app: app}

	r.Post("/", h.Create)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", h.Get)
		r.Delete("/", h.Delete)
	})
	return r
}
