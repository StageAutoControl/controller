package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/StageAutoControl/controller/pkg/artnet"
	"github.com/gorilla/handlers"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"github.com/sirupsen/logrus"
)

// Server represents the controllers API server, aware of all the controllers
type Server struct {
	*rpc.Server
	logger        *logrus.Entry
	storage       storage
	apiController map[string]interface{}
	controller    artnet.Controller
}

// NewServer returns a new Server instance
func NewServer(logger *logrus.Entry, storage storage, controller artnet.Controller) (*Server, error) {
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
	s.apiController = map[string]interface{}{
		"DMXAnimation":     newDMXAnimationController(s.logger, s.storage),
		"DMXDevice":        newDMXDeviceController(s.logger, s.storage),
		"DMXDeviceGroup":   newDMXDeviceGroupController(s.logger, s.storage),
		"DMXDeviceType":    newDMXDeviceTypeController(s.logger, s.storage),
		"DMXPreset":        newDMXPresetController(s.logger, s.storage),
		"DMXScene":         newDMXSceneController(s.logger, s.storage),
		"DMXTransition":    newDMXTransitionController(s.logger, s.storage),
		"DMXColorVariable": newDMXColorVariableController(s.logger, s.storage),
		"Song":             newSongController(s.logger, s.storage),
		"SetList":          newSetListController(s.logger, s.storage),
		"DMXPlayground":    newDMXPlaygroundController(s.logger, s.controller),
	}

	for name, controller := range s.apiController {
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
	r.Handle(rpcPath, s.Server)
	r.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(rw, "OK"); err != nil {
			s.logger.Errorf("failed to write content to response: %v", err)
		}
	})

	h := handlers.LoggingHandler(s.logger.Writer(), r)
	h = handlers.RecoveryHandler()(h)
	h = handlers.CORS(
		handlers.AllowCredentials(),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"POST", "GET"}),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
	)(h)

	httpServer := http.Server{
		Addr:    endpoint,
		Handler: h,
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
