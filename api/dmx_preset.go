package api

import (
	"fmt"
	"net/http"

	"github.com/StageAutoControl/controller/cntl"
)

type dmxPresetHandler struct {
	repo repo
}

func (h *dmxPresetHandler) GetAll(r *http.Request, req *EmptyRequest, reply *BaseResponse) error {
	data, err := h.repo.Load()
	if err != nil {
		return makeRepoReadError(err)
	}

	dmxPresets := []*cntl.DMXPreset{}
	for _, s := range data.DMXPresets {
		dmxPresets = append(dmxPresets, s)
	}

	reply.Data = dmxPresets
	reply.Message = fmt.Sprintf("Loaded %d DMXPresets", len(data.DMXPresets))
	return nil
}

func (h *dmxPresetHandler) Get(r *http.Request, req *IDRequest, reply *BaseResponse) error {
	data, err := h.repo.Load()
	if err != nil {
		return makeRepoReadError(err)
	}

	if _, ok := data.DMXPresets[req.ID]; !ok {
		return fmt.Errorf("DMXPreset with ID %q not found", req.ID)
	}

	reply.Data = data.DMXPresets[req.ID]
	reply.Message = fmt.Sprintf("Loaded DMXPreset %q", req.ID)
	return nil
}

func (h *dmxPresetHandler) Create(r *http.Request, req *cntl.DMXPreset, reply *BaseResponse) error {
	data, err := h.repo.Load()
	if err != nil {
		return makeRepoReadError(err)
	}

	for _, sl := range data.DMXPresets {
		if namesEqual(sl.Name, req.Name) {
			return errDuplicateName
		}
	}

	req.ID = generateUUID()
	data.DMXPresets[req.ID] = req
	if err := h.repo.Save(data); err != nil {
		return makeRepoSaveError(err)
	}

	reply.Data = req
	reply.Message = fmt.Sprintf("Created DMXPreset %q", req.ID)
	return nil
}

func (h *dmxPresetHandler) Update(r *http.Request, req *cntl.DMXPreset, reply *BaseResponse) error {
	data, err := h.repo.Load()
	if err != nil {
		return makeRepoReadError(err)
	}

	if _, ok := data.DMXPresets[req.ID]; !ok {
		return fmt.Errorf("DMXPreset with ID %q not found", req.ID)
	}

	for _, sl := range data.DMXPresets {
		if namesEqual(sl.Name, req.Name) && sl.ID != req.ID {
			return errDuplicateName
		}
	}

	data.DMXPresets[req.ID] = req
	if err := h.repo.Save(data); err != nil {
		return makeRepoSaveError(err)
	}

	reply.Data = req
	reply.Message = fmt.Sprintf("Updated DMXPreset %q", req.ID)
	return nil
}

func (h *dmxPresetHandler) Delete(r *http.Request, req *IDRequest, reply *BaseResponse) error {
	data, err := h.repo.Load()
	if err != nil {
		return makeRepoReadError(err)
	}

	if _, ok := data.DMXPresets[req.ID]; !ok {
		return fmt.Errorf("DMXPreset with ID %q not found", req.ID)
	}

	delete(data.DMXPresets, req.ID)
	if err := h.repo.Save(data); err != nil {
		return makeRepoSaveError(err)
	}

	reply.Data = req
	reply.Message = fmt.Sprintf("Deleted DMXPreset %q", req.ID)
	return nil
}
