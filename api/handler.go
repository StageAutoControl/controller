package api

import (
	"net/http"

	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json"
)

func NewHandler() http.Handler {
	rpcServer := rpc.NewServer()
	rpcServer.RegisterCodec(json.NewCodec(), "application/json")
	rpcServer.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")

	var handlers = map[string]interface{}{
		"SetList": &setListHandler{},
	}

	for name, handler := range handlers {
		rpcServer.RegisterService(handler, name)
	}

	return rpcServer
}
