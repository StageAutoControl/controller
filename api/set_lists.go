package api

import "net/http"

type setListHandler struct {
	//repo
}

func (h *setListHandler) GetAll(r *http.Request, args *EmptyRequest, reply *BaseResponse) error {

}
