package api

import (
	"fmt"
	"net/http"

	"github.com/StageAutoControl/controller/cntl"
)

type dmxDeviceGroupHandler struct {
	repo repo
}

func (h *dmxDeviceGroupHandler) GetAll(r *http.Request, req *EmptyRequest, reply *BaseResponse) error {
	data, err := h.repo.Load()
	if err != nil {
		return makeRepoReadError(err)
	}

	dmxDeviceGroups := []*cntl.DMXDeviceGroup{}
	for _, s := range data.DMXDeviceGroups {
		dmxDeviceGroups = append(dmxDeviceGroups, s)
	}

	reply.Data = dmxDeviceGroups
	reply.Message = fmt.Sprintf("Loaded %d DMXDeviceGroups", len(data.DMXDeviceGroups))
	return nil
}

func (h *dmxDeviceGroupHandler) Get(r *http.Request, req *IDRequest, reply *BaseResponse) error {
	data, err := h.repo.Load()
	if err != nil {
		return makeRepoReadError(err)
	}

	if _, ok := data.DMXDeviceGroups[req.ID]; !ok {
		return fmt.Errorf("DMXDeviceGroup with ID %q not found", req.ID)
	}

	reply.Data = data.DMXDeviceGroups[req.ID]
	reply.Message = fmt.Sprintf("Loaded DMXDeviceGroup %q", req.ID)
	return nil
}

func (h *dmxDeviceGroupHandler) Create(r *http.Request, req *cntl.DMXDeviceGroup, reply *BaseResponse) error {
	data, err := h.repo.Load()
	if err != nil {
		return makeRepoReadError(err)
	}

	for _, sl := range data.DMXDeviceGroups {
		if namesEqual(sl.Name, req.Name) {
			return errDuplicateName
		}
	}

	req.ID = generateUUID()
	data.DMXDeviceGroups[req.ID] = req
	if err := h.repo.Save(data); err != nil {
		return makeRepoSaveError(err)
	}

	reply.Data = req
	reply.Message = fmt.Sprintf("Created DMXDeviceGroup %q", req.ID)
	return nil
}

func (h *dmxDeviceGroupHandler) Update(r *http.Request, req *cntl.DMXDeviceGroup, reply *BaseResponse) error {
	data, err := h.repo.Load()
	if err != nil {
		return makeRepoReadError(err)
	}

	if _, ok := data.DMXDeviceGroups[req.ID]; !ok {
		return fmt.Errorf("DMXDeviceGroup with ID %q not found", req.ID)
	}

	for _, sl := range data.DMXDeviceGroups {
		if namesEqual(sl.Name, req.Name) && sl.ID != req.ID {
			return errDuplicateName
		}
	}

	data.DMXDeviceGroups[req.ID] = req
	if err := h.repo.Save(data); err != nil {
		return makeRepoSaveError(err)
	}

	reply.Data = req
	reply.Message = fmt.Sprintf("Updated DMXDeviceGroup %q", req.ID)
	return nil
}

func (h *dmxDeviceGroupHandler) Delete(r *http.Request, req *IDRequest, reply *BaseResponse) error {
	data, err := h.repo.Load()
	if err != nil {
		return makeRepoReadError(err)
	}

	if _, ok := data.DMXDeviceGroups[req.ID]; !ok {
		return fmt.Errorf("DMXDeviceGroup with ID %q not found", req.ID)
	}

	delete(data.DMXDeviceGroups, req.ID)
	if err := h.repo.Save(data); err != nil {
		return makeRepoSaveError(err)
	}

	reply.Data = req
	reply.Message = fmt.Sprintf("Deleted DMXDeviceGroup %q", req.ID)
	return nil
}
