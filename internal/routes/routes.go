package routes

import (
	"dynamodb-eg/internal/handler"

	"github.com/go-chi/chi"
)

func InitRouter(app *handler.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Post("/music", app.Store())
	r.Get("/music", app.Get())
	r.Put("/music", app.UpdateByName())
	r.Delete("/music", app.Delete())

	return r
}
