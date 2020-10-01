package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	helmes "github.com/rugwirobaker/helmes"
)

// Server ...
type Server struct {
	Service helmes.Service
}

// New api Server instance
func New(svc helmes.Service) *Server {
	return &Server{Service: svc}
}

// Handler returns an http.Handler
func (s Server) Handler() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to helmes"))
	})

	r.Get("/version", VersionHandler(s.Service))
	r.Post("/send", SMSHandler(s.Service))
	r.Get("/healthz", HealthHandler())

	return r
}
