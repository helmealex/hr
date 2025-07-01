package api

import (
	"hr/internal/service"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	service *service.Service
	logger  *slog.Logger
	router  *chi.Mux
	config  *Config
}

func New(
	svc *service.Service,
	logger *slog.Logger,
	config *Config,
) *Server {
	server := Server{
		service: svc,
		logger:  logger,
		router:  chi.NewRouter(),
		config:  config,
	}

	server.registerRoutes()

	return &server
}

func (s *Server) registerRoutes() {
	s.router.Route("/jobs", func(r chi.Router) {
		r.Get("/", s.getJobs)
		r.Post("/", s.createJob)
		r.Get("/{id}", s.getJobByID)
	})

	s.router.Route("/candidates", func(r chi.Router) {
		r.Get("/", s.getCandidates)
		r.Post("/", s.createCandidate)
		r.Get("/{id}", s.getCandidateByID)
	})
}

func (s *Server) Start() error {
	server := &http.Server{
		Addr:              s.config.Host + ":" + s.config.Port,
		ReadHeaderTimeout: 10 * time.Second,
		Handler:           s.router,
	}

	return server.ListenAndServe()
}
