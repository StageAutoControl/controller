package api

import (
	"context"
	"fmt"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

// Server represents the controllers API server, aware of all the controllers
type Server struct {
	*rpc.Server
	logger     *logrus.Entry
	storage    storage
	controller map[string]interface{}
}

// NewServer returns a new Server instance
func NewServer(logger *logrus.Entry, storage storage) (*Server, error) {
	server := &Server{
		Server:  rpc.NewServer(),
		logger:  logger,
		storage: storage,
	}

	if err := server.registerControllers(); err != nil {
		return nil, err
	}

	return server, nil
}

func (s *Server) registerControllers() error {
	s.controller = map[string]interface{}{
		"DMXDevice": newDMXDeviceController(s.logger, s.storage),
	}

	for name, controller := range s.controller {
		if err := s.Server.RegisterService(controller, name); err != nil {
			return err
		}
	}

	return nil
}

// Run runs the http server
func (s *Server) Run(ctx context.Context, endpoint string) error {
	s.Server.RegisterCodec(json.NewCodec(), "application/json")

	r := http.NewServeMux()
	r.Handle("/rpc", s.Server)
	r.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(rw, "OK"); err != nil {
			s.logger.Errorf("failed to write content to response: %v", err)
		}
	})

	httpServer := http.Server{
		Addr:    endpoint,
		Handler: r,
	}

	go func() {
		<-ctx.Done()
		if err := httpServer.Shutdown(ctx); err != nil {
			s.logger.Errorf("failed to shutdown http server: %v", err)
		}
	}()

	err := httpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}
