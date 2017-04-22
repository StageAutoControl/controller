package api

import (
	"fmt"
	"net/http"
)

type setListHandler struct {
	repo repo
}

func (h *setListHandler) GetAll(r *http.Request, args *EmptyRequest, reply *BaseResponse) error {
	data, err := h.repo.Load()
	if err != nil {
		return makeRepoReadError(err)
	}

	reply.Data = data.SetLists
	reply.Message = fmt.Sprintf("Loaded %d SetLists", len(data.SetLists))
	return nil
}
