package server

import (
	"context"
	"fmt"
	"net/http"
	_ "net/http/pprof"

	"github.com/gorilla/handlers"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"github.com/sirupsen/logrus"

	"github.com/StageAutoControl/controller/pkg/api"
	"github.com/StageAutoControl/controller/pkg/api/datastore"
	"github.com/StageAutoControl/controller/pkg/api/playback"
	"github.com/StageAutoControl/controller/pkg/api/playground"
	"github.com/StageAutoControl/controller/pkg/artnet"
	"github.com/StageAutoControl/controller/pkg/process"
	"github.com/StageAutoControl/controller/pkg/visualizer"
)

// Server represents the controllers API server, aware of all the controllers
type Server struct {
	*rpc.Server
	logger        *logrus.Entry
	storage       api.Storage
	loader        api.Loader
	apiController map[string]interface{}
	cntl          artnet.Controller
	pm            process.Manager
	visualizer    *visualizer.Server
}

// New returns a new Server instance
func New(
	logger *logrus.Entry,
	storage api.Storage,
	loader api.Loader,
	cntl artnet.Controller,
	pm process.Manager,
	visualizer *visualizer.Server,
) (*Server, error) {

	server := &Server{
		Server:     rpc.NewServer(),
		logger:     logger,
		storage:    storage,
		loader:     loader,
		cntl:       cntl,
		pm:         pm,
		visualizer: visualizer,
	}

	if err := server.registerControllers(); err != nil {
		return nil, err
	}

	return server, nil
}

func (s *Server) registerControllers() error {
	s.apiController = map[string]interface{}{
		"DMXAnimation":     datastore.NewDMXAnimationController(s.logger, s.storage),
		"DMXDevice":        datastore.NewDMXDeviceController(s.logger, s.storage),
		"DMXDeviceGroup":   datastore.NewDMXDeviceGroupController(s.logger, s.storage),
		"DMXDeviceType":    datastore.NewDMXDeviceTypeController(s.logger, s.storage),
		"DMXPreset":        datastore.NewDMXPresetController(s.logger, s.storage),
		"DMXScene":         datastore.NewDMXSceneController(s.logger, s.storage),
		"DMXTransition":    datastore.NewDMXTransitionController(s.logger, s.storage),
		"DMXColorVariable": datastore.NewDMXColorVariableController(s.logger, s.storage),
		"Song":             datastore.NewSongController(s.logger, s.storage),
		"SetList":          datastore.NewSetListController(s.logger, s.storage),
		"DMXPlayground":    playground.NewDMXPlaygroundController(s.logger, s.cntl, s.loader),
		"Playback":         playback.NewController(s.pm),
	}

	for name, cntl := range s.apiController {
		if err := s.Server.RegisterService(cntl, name); err != nil {
			return err
		}
	}

	return nil
}

// Run runs the http server
func (s *Server) Run(ctx context.Context, endpoint string) error {
	s.Server.RegisterCodec(json.NewCodec(), "application/json")

	r := http.DefaultServeMux
	r.Handle(api.RPCPath, s.Server)
	r.HandleFunc(api.VisualizerPath, s.visualizer.ServeRequest)
	r.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(rw, "OK"); err != nil {
			s.logger.Errorf("failed to write content to response: %v", err)
		}
	})

	h := handlers.RecoveryHandler()(r)
	h = handlers.CORS(
		handlers.AllowCredentials(),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"POST", "GET", "HEAD"}),
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

	s.logger.Infof("listening on %s", endpoint)

	err := httpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}
