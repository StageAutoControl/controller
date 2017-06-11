package api

import (
	"fmt"
	"net/http"

	"github.com/StageAutoControl/controller/cntl"
)

type dmxDeviceTypeHandler struct {
	repo repo
}

func (h *dmxDeviceTypeHandler) GetAll(r *http.Request, req *EmptyRequest, reply *BaseResponse) error {
	data, err := h.repo.Load()
	if err != nil {
		return makeRepoReadError(err)
	}

	dmxDeviceTypes := []*cntl.DMXDeviceType{}
	for _, s := range data.DMXDeviceTypes {
		dmxDeviceTypes = append(dmxDeviceTypes, s)
	}

	reply.Data = dmxDeviceTypes
	reply.Message = fmt.Sprintf("Loaded %d DMXDeviceTypes", len(data.DMXDeviceTypes))
	return nil
}

func (h *dmxDeviceTypeHandler) Get(r *http.Request, req *IDRequest, reply *BaseResponse) error {
	data, err := h.repo.Load()
	if err != nil {
		return makeRepoReadError(err)
	}

	if _, ok := data.DMXDeviceTypes[req.ID]; !ok {
		return fmt.Errorf("DMXDeviceType with ID %q not found", req.ID)
	}

	reply.Data = data.DMXDeviceTypes[req.ID]
	reply.Message = fmt.Sprintf("Loaded DMXDeviceType %q", req.ID)
	return nil
}

func (h *dmxDeviceTypeHandler) Create(r *http.Request, req *cntl.DMXDeviceType, reply *BaseResponse) error {
	data, err := h.repo.Load()
	if err != nil {
		return makeRepoReadError(err)
	}

	for _, sl := range data.DMXDeviceTypes {
		if namesEqual(sl.Name, req.Name) {
			return errDuplicateName
		}
	}

	req.ID = generateUUID()
	data.DMXDeviceTypes[req.ID] = req
	if err := h.repo.Save(data); err != nil {
		return makeRepoSaveError(err)
	}

	reply.Data = req
	reply.Message = fmt.Sprintf("Created DMXDeviceType %q", req.ID)
	return nil
}

func (h *dmxDeviceTypeHandler) Update(r *http.Request, req *cntl.DMXDeviceType, reply *BaseResponse) error {
	data, err := h.repo.Load()
	if err != nil {
		return makeRepoReadError(err)
	}

	if _, ok := data.DMXDeviceTypes[req.ID]; !ok {
		return fmt.Errorf("DMXDeviceType with ID %q not found", req.ID)
	}

	for _, sl := range data.DMXDeviceTypes {
		if namesEqual(sl.Name, req.Name) && sl.ID != req.ID {
			return errDuplicateName
		}
	}

	data.DMXDeviceTypes[req.ID] = req
	if err := h.repo.Save(data); err != nil {
		return makeRepoSaveError(err)
	}

	reply.Data = req
	reply.Message = fmt.Sprintf("Updated DMXDeviceType %q", req.ID)
	return nil
}

func (h *dmxDeviceTypeHandler) Delete(r *http.Request, req *IDRequest, reply *BaseResponse) error {
	data, err := h.repo.Load()
	if err != nil {
		return makeRepoReadError(err)
	}

	if _, ok := data.DMXDeviceTypes[req.ID]; !ok {
		return fmt.Errorf("DMXDeviceType with ID %q not found", req.ID)
	}

	delete(data.DMXDeviceTypes, req.ID)
	if err := h.repo.Save(data); err != nil {
		return makeRepoSaveError(err)
	}

	reply.Data = req
	reply.Message = fmt.Sprintf("Deleted DMXDeviceType %q", req.ID)
	return nil
}
