package api

import (
	"fmt"
	"net/http"

	"github.com/StageAutoControl/controller/cntl"
)

type dmxAnimationHandler struct {
	repo repo
}

func (h *dmxAnimationHandler) GetAll(r *http.Request, req *EmptyRequest, reply *BaseResponse) error {
	data, err := h.repo.Load()
	if err != nil {
		return makeRepoReadError(err)
	}

	dmxAnimations := []*cntl.DMXAnimation{}
	for _, s := range data.DMXAnimations {
		dmxAnimations = append(dmxAnimations, s)
	}

	reply.Data = dmxAnimations
	reply.Message = fmt.Sprintf("Loaded %d DMXAnimations", len(data.DMXAnimations))
	return nil
}

func (h *dmxAnimationHandler) Get(r *http.Request, req *IDRequest, reply *BaseResponse) error {
	data, err := h.repo.Load()
	if err != nil {
		return makeRepoReadError(err)
	}

	if _, ok := data.DMXAnimations[req.ID]; !ok {
		return fmt.Errorf("DMXAnimation with ID %q not found", req.ID)
	}

	reply.Data = data.DMXAnimations[req.ID]
	reply.Message = fmt.Sprintf("Loaded DMXAnimation %q", req.ID)
	return nil
}

func (h *dmxAnimationHandler) Create(r *http.Request, req *cntl.DMXAnimation, reply *BaseResponse) error {
	data, err := h.repo.Load()
	if err != nil {
		return makeRepoReadError(err)
	}

	for _, sl := range data.DMXAnimations {
		if namesEqual(sl.Name, req.Name) {
			return errDuplicateName
		}
	}

	req.ID = generateUUID()
	data.DMXAnimations[req.ID] = req
	if err := h.repo.Save(data); err != nil {
		return makeRepoSaveError(err)
	}

	reply.Data = req
	reply.Message = fmt.Sprintf("Created DMXAnimation %q", req.ID)
	return nil
}

func (h *dmxAnimationHandler) Update(r *http.Request, req *cntl.DMXAnimation, reply *BaseResponse) error {
	data, err := h.repo.Load()
	if err != nil {
		return makeRepoReadError(err)
	}

	if _, ok := data.DMXAnimations[req.ID]; !ok {
		return fmt.Errorf("DMXAnimation with ID %q not found", req.ID)
	}

	for _, sl := range data.DMXAnimations {
		if namesEqual(sl.Name, req.Name) && sl.ID != req.ID {
			return errDuplicateName
		}
	}

	data.DMXAnimations[req.ID] = req
	if err := h.repo.Save(data); err != nil {
		return makeRepoSaveError(err)
	}

	reply.Data = req
	reply.Message = fmt.Sprintf("Updated DMXAnimation %q", req.ID)
	return nil
}

func (h *dmxAnimationHandler) Delete(r *http.Request, req *IDRequest, reply *BaseResponse) error {
	data, err := h.repo.Load()
	if err != nil {
		return makeRepoReadError(err)
	}

	if _, ok := data.DMXAnimations[req.ID]; !ok {
		return fmt.Errorf("DMXAnimation with ID %q not found", req.ID)
	}

	delete(data.DMXAnimations, req.ID)
	if err := h.repo.Save(data); err != nil {
		return makeRepoSaveError(err)
	}

	reply.Data = req
	reply.Message = fmt.Sprintf("Deleted DMXAnimation %q", req.ID)
	return nil
}
