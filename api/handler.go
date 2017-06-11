package api

import (
	"net/http"

	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json"
)

// NewHandler creates a new Gorilla RPC service handler and adds all our command handler
func NewHandler(repo repo) http.Handler {
	rpcServer := rpc.NewServer()
	rpcServer.RegisterCodec(json.NewCodec(), "application/json")
	rpcServer.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")

	rpcServer.RegisterService(&setListHandler{repo}, "SetList")
	rpcServer.RegisterService(&songHandler{repo}, "Song")

	rpcServer.RegisterService(&dmxAnimationHandler{repo}, "DMXAnimation")
	rpcServer.RegisterService(&dmxDeviceHandler{repo}, "DMXDevice")
	rpcServer.RegisterService(&dmxDeviceGroupHandler{repo}, "DMXDeviceGroup")
	rpcServer.RegisterService(&dmxDeviceTypeHandler{repo}, "DMXDeviceType")
	rpcServer.RegisterService(&dmxPresetHandler{repo}, "DMXPreset")
	rpcServer.RegisterService(&dmxSceneHandler{repo}, "DMXScene")

	return rpcServer
}

// NewRepoLockingHandler is a simple wrapper handler that cares about locking the
// repo to prevent overrides or segmentation faults
func NewRepoLockingHandler(repo repo, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		repo.Lock()
		defer repo.Unlock()

		handler.ServeHTTP(rw, r)
	})
}
