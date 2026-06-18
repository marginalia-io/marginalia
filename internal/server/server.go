package server

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

// Server wraps an HTTP server with lifecycle management.
type Server struct {
	httpServer      *http.Server
	shutdownTimeout time.Duration
}

// New returns a Server configured from cfg, using db for data access. Unset
// cfg fields use defaults.
func New(cfg Config, db *sql.DB) *Server {
	cfg = cfg.withDefaults()
	return &Server{
		httpServer: &http.Server{
			Addr:         net.JoinHostPort(cfg.Host, strconv.Itoa(cfg.Port)),
			Handler:      newRouter(&api{db: db}),
			ReadTimeout:  cfg.ReadTimeout,
			WriteTimeout: cfg.WriteTimeout,
			IdleTimeout:  cfg.IdleTimeout,
		},
		shutdownTimeout: cfg.ShutdownTimeout,
	}
}

// Run starts the server and blocks until ctx is canceled or the server fails.
// On cancellation it attempts a graceful shutdown bounded by the configured
// shutdown timeout.
func (s *Server) Run(ctx context.Context) error {
	errc := make(chan error, 1)
	go func() {
		log.Printf("http: listening on %s", s.httpServer.Addr)
		if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errc <- err
		}
	}()

	select {
	case err := <-errc:
		return err
	case <-ctx.Done():
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()
	log.Print("http: shutting down")
	return s.httpServer.Shutdown(shutdownCtx)
}

func newRouter(a *api) *chi.Mux {
	router := chi.NewRouter()
	router.Mount("/api", apiRouter(a))
	router.Handle("/*", spaHandler())
	return router
}
