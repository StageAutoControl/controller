package api

import (
	"fmt"
	"net/http"

	"github.com/StageAutoControl/controller/cntl"
)

type dmxDeviceHandler struct {
	repo repo
}

func (h *dmxDeviceHandler) GetAll(r *http.Request, req *EmptyRequest, reply *BaseResponse) error {
	data, err := h.repo.Load()
	if err != nil {
		return makeRepoReadError(err)
	}

	dmxDevices := []*cntl.DMXDevice{}
	for _, s := range data.DMXDevices {
		dmxDevices = append(dmxDevices, s)
	}

	reply.Data = dmxDevices
	reply.Message = fmt.Sprintf("Loaded %d DMXDevices", len(data.DMXDevices))
	return nil
}

func (h *dmxDeviceHandler) Get(r *http.Request, req *IDRequest, reply *BaseResponse) error {
	data, err := h.repo.Load()
	if err != nil {
		return makeRepoReadError(err)
	}

	if _, ok := data.DMXDevices[req.ID]; !ok {
		return fmt.Errorf("DMXDevice with ID %q not found", req.ID)
	}

	reply.Data = data.DMXDevices[req.ID]
	reply.Message = fmt.Sprintf("Loaded DMXDevice %q", req.ID)
	return nil
}

func (h *dmxDeviceHandler) Create(r *http.Request, req *cntl.DMXDevice, reply *BaseResponse) error {
	data, err := h.repo.Load()
	if err != nil {
		return makeRepoReadError(err)
	}

	for _, sl := range data.DMXDevices {
		if namesEqual(sl.Name, req.Name) {
			return errDuplicateName
		}
	}

	req.ID = generateUUID()
	data.DMXDevices[req.ID] = req
	if err := h.repo.Save(data); err != nil {
		return makeRepoSaveError(err)
	}

	reply.Data = req
	reply.Message = fmt.Sprintf("Created DMXDevice %q", req.ID)
	return nil
}

func (h *dmxDeviceHandler) Update(r *http.Request, req *cntl.DMXDevice, reply *BaseResponse) error {
	data, err := h.repo.Load()
	if err != nil {
		return makeRepoReadError(err)
	}

	if _, ok := data.DMXDevices[req.ID]; !ok {
		return fmt.Errorf("DMXDevice with ID %q not found", req.ID)
	}

	for _, sl := range data.DMXDevices {
		if namesEqual(sl.Name, req.Name) && sl.ID != req.ID {
			return errDuplicateName
		}
	}

	data.DMXDevices[req.ID] = req
	if err := h.repo.Save(data); err != nil {
		return makeRepoSaveError(err)
	}

	reply.Data = req
	reply.Message = fmt.Sprintf("Updated DMXDevice %q", req.ID)
	return nil
}

func (h *dmxDeviceHandler) Delete(r *http.Request, req *IDRequest, reply *BaseResponse) error {
	data, err := h.repo.Load()
	if err != nil {
		return makeRepoReadError(err)
	}

	if _, ok := data.DMXDevices[req.ID]; !ok {
		return fmt.Errorf("DMXDevice with ID %q not found", req.ID)
	}

	delete(data.DMXDevices, req.ID)
	if err := h.repo.Save(data); err != nil {
		return makeRepoSaveError(err)
	}

	reply.Data = req
	reply.Message = fmt.Sprintf("Deleted DMXDevice %q", req.ID)
	return nil
}
