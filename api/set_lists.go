package api

import (
	"fmt"
	"net/http"

	"github.com/StageAutoControl/controller/cntl"
)

type setListHandler struct {
	repo repo
}

func (h *setListHandler) GetAll(r *http.Request, req *EmptyRequest, reply *BaseResponse) error {
	data, err := h.repo.Load()
	if err != nil {
		return makeRepoReadError(err)
	}

	setLists := []*cntl.SetList{}
	for _, s := range data.SetLists {
		setLists = append(setLists, s)
	}

	reply.Data = setLists
	reply.Message = fmt.Sprintf("Loaded %d SetLists", len(data.SetLists))
	return nil
}

func (h *setListHandler) Get(r *http.Request, req *IDRequest, reply *BaseResponse) error {
	data, err := h.repo.Load()
	if err != nil {
		return makeRepoReadError(err)
	}

	if _, ok := data.SetLists[req.ID]; !ok {
		return fmt.Errorf("SetList with ID %q not found", req.ID)
	}

	reply.Data = data.SetLists[req.ID]
	reply.Message = fmt.Sprintf("Loaded SetList %q", req.ID)
	return nil
}

func (h *setListHandler) Create(r *http.Request, req *cntl.SetList, reply *BaseResponse) error {
	data, err := h.repo.Load()
	if err != nil {
		return makeRepoReadError(err)
	}

	for _, sl := range data.SetLists {
		if namesEqual(sl.Name, req.Name) {
			return errDuplicateName
		}
	}

	req.ID = generateUUID()
	data.SetLists[req.ID] = req
	if err := h.repo.Save(data); err != nil {
		return makeRepoSaveError(err)
	}

	reply.Data = req
	reply.Message = fmt.Sprintf("Created SetList %q", req.ID)
	return nil
}

func (h *setListHandler) Update(r *http.Request, req *cntl.SetList, reply *BaseResponse) error {
	data, err := h.repo.Load()
	if err != nil {
		return makeRepoReadError(err)
	}

	if _, ok := data.SetLists[req.ID]; !ok {
		return fmt.Errorf("SetList with ID %q not found", req.ID)
	}

	for _, sl := range data.SetLists {
		if namesEqual(sl.Name, req.Name) && sl.ID != req.ID {
			return errDuplicateName
		}
	}

	data.SetLists[req.ID] = req
	if err := h.repo.Save(data); err != nil {
		return makeRepoSaveError(err)
	}

	reply.Data = req
	reply.Message = fmt.Sprintf("Updated SetList %q", req.ID)
	return nil
}

func (h *setListHandler) Delete(r *http.Request, req *IDRequest, reply *BaseResponse) error {
	data, err := h.repo.Load()
	if err != nil {
		return makeRepoReadError(err)
	}

	if _, ok := data.SetLists[req.ID]; !ok {
		return fmt.Errorf("SetList with ID %q not found", req.ID)
	}

	delete(data.SetLists, req.ID)
	if err := h.repo.Save(data); err != nil {
		return makeRepoSaveError(err)
	}

	reply.Data = req
	reply.Message = fmt.Sprintf("Deleted SetList %q", req.ID)
	return nil
}
