package api

import (
	"fmt"
	"net/http"

	"github.com/StageAutoControl/controller/cntl"
)

type songHandler struct {
	repo repo
}

func (h *songHandler) GetAll(r *http.Request, req *EmptyRequest, reply *BaseResponse) error {
	data, err := h.repo.Load()
	if err != nil {
		return makeRepoReadError(err)
	}

	songs := []*cntl.Song{}
	for _, s := range data.Songs {
		songs = append(songs, s)
	}

	reply.Data = songs
	reply.Message = fmt.Sprintf("Loaded %d Songs", len(data.Songs))
	return nil
}

func (h *songHandler) Get(r *http.Request, req *IDRequest, reply *BaseResponse) error {
	data, err := h.repo.Load()
	if err != nil {
		return makeRepoReadError(err)
	}

	if _, ok := data.Songs[req.ID]; !ok {
		return fmt.Errorf("Song with ID %q not found", req.ID)
	}

	reply.Data = data.Songs[req.ID]
	reply.Message = fmt.Sprintf("Loaded Song %q", req.ID)
	return nil
}

func (h *songHandler) Create(r *http.Request, req *cntl.Song, reply *BaseResponse) error {
	data, err := h.repo.Load()
	if err != nil {
		return makeRepoReadError(err)
	}

	for _, sl := range data.Songs {
		if namesEqual(sl.Name, req.Name) {
			return errDuplicateName
		}
	}

	req.ID = generateUUID()
	data.Songs[req.ID] = req
	if err := h.repo.Save(data); err != nil {
		return makeRepoSaveError(err)
	}

	reply.Data = req
	reply.Message = fmt.Sprintf("Created Song %q", req.ID)
	return nil
}

func (h *songHandler) Update(r *http.Request, req *cntl.Song, reply *BaseResponse) error {
	data, err := h.repo.Load()
	if err != nil {
		return makeRepoReadError(err)
	}

	if _, ok := data.Songs[req.ID]; !ok {
		return fmt.Errorf("Song with ID %q not found", req.ID)
	}

	for _, sl := range data.Songs {
		if namesEqual(sl.Name, req.Name) && sl.ID != req.ID {
			return errDuplicateName
		}
	}

	data.Songs[req.ID] = req
	if err := h.repo.Save(data); err != nil {
		return makeRepoSaveError(err)
	}

	reply.Data = req
	reply.Message = fmt.Sprintf("Updated Song %q", req.ID)
	return nil
}

func (h *songHandler) Delete(r *http.Request, req *IDRequest, reply *BaseResponse) error {
	data, err := h.repo.Load()
	if err != nil {
		return makeRepoReadError(err)
	}

	if _, ok := data.Songs[req.ID]; !ok {
		return fmt.Errorf("Song with ID %q not found", req.ID)
	}

	delete(data.Songs, req.ID)
	if err := h.repo.Save(data); err != nil {
		return makeRepoSaveError(err)
	}

	reply.Data = req
	reply.Message = fmt.Sprintf("Deleted Song %q", req.ID)
	return nil
}
