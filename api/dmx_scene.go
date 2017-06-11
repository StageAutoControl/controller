package api

import (
	"fmt"
	"net/http"

	"github.com/StageAutoControl/controller/cntl"
)

type dmxSceneHandler struct {
	repo repo
}

func (h *dmxSceneHandler) GetAll(r *http.Request, req *EmptyRequest, reply *BaseResponse) error {
	data, err := h.repo.Load()
	if err != nil {
		return makeRepoReadError(err)
	}

	dmxScenes := []*cntl.DMXScene{}
	for _, s := range data.DMXScenes {
		dmxScenes = append(dmxScenes, s)
	}

	reply.Data = dmxScenes
	reply.Message = fmt.Sprintf("Loaded %d DMXScenes", len(data.DMXScenes))
	return nil
}

func (h *dmxSceneHandler) Get(r *http.Request, req *IDRequest, reply *BaseResponse) error {
	data, err := h.repo.Load()
	if err != nil {
		return makeRepoReadError(err)
	}

	if _, ok := data.DMXScenes[req.ID]; !ok {
		return fmt.Errorf("DMXScene with ID %q not found", req.ID)
	}

	reply.Data = data.DMXScenes[req.ID]
	reply.Message = fmt.Sprintf("Loaded DMXScene %q", req.ID)
	return nil
}

func (h *dmxSceneHandler) Create(r *http.Request, req *cntl.DMXScene, reply *BaseResponse) error {
	data, err := h.repo.Load()
	if err != nil {
		return makeRepoReadError(err)
	}

	for _, sl := range data.DMXScenes {
		if namesEqual(sl.Name, req.Name) {
			return errDuplicateName
		}
	}

	req.ID = generateUUID()
	data.DMXScenes[req.ID] = req
	if err := h.repo.Save(data); err != nil {
		return makeRepoSaveError(err)
	}

	reply.Data = req
	reply.Message = fmt.Sprintf("Created DMXScene %q", req.ID)
	return nil
}

func (h *dmxSceneHandler) Update(r *http.Request, req *cntl.DMXScene, reply *BaseResponse) error {
	data, err := h.repo.Load()
	if err != nil {
		return makeRepoReadError(err)
	}

	if _, ok := data.DMXScenes[req.ID]; !ok {
		return fmt.Errorf("DMXScene with ID %q not found", req.ID)
	}

	for _, sl := range data.DMXScenes {
		if namesEqual(sl.Name, req.Name) && sl.ID != req.ID {
			return errDuplicateName
		}
	}

	data.DMXScenes[req.ID] = req
	if err := h.repo.Save(data); err != nil {
		return makeRepoSaveError(err)
	}

	reply.Data = req
	reply.Message = fmt.Sprintf("Updated DMXScene %q", req.ID)
	return nil
}

func (h *dmxSceneHandler) Delete(r *http.Request, req *IDRequest, reply *BaseResponse) error {
	data, err := h.repo.Load()
	if err != nil {
		return makeRepoReadError(err)
	}

	if _, ok := data.DMXScenes[req.ID]; !ok {
		return fmt.Errorf("DMXScene with ID %q not found", req.ID)
	}

	delete(data.DMXScenes, req.ID)
	if err := h.repo.Save(data); err != nil {
		return makeRepoSaveError(err)
	}

	reply.Data = req
	reply.Message = fmt.Sprintf("Deleted DMXScene %q", req.ID)
	return nil
}
