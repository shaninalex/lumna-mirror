package board

import (
	"net/http"

	"gitlab.com/shaninalex/lumna/app/pkg/db"
	"gitlab.com/shaninalex/lumna/app/web/utils"
)

type patchBoardPayload struct {
	Name string `json:"name"`
}

func (s *BoardHandler) Patch(w http.ResponseWriter, r *http.Request) {
	id := utils.UrlNumericParam(w, r, "id")
	payload, err := utils.BodyParser[patchBoardPayload](r)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}

	err = s.boardService.Update(r.Context(), uint(id), db.Option{Key: "name", Value: payload.Name})
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}
	utils.Success(w, nil, "Updated")
}

func (s *BoardHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := utils.UrlNumericParam(w, r, "id")
	if err := s.boardService.Delete(r.Context(), uint(id)); err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}
	utils.Success(w, nil, "Deleted")
}

func (s *BoardHandler) ListsGet(w http.ResponseWriter, r *http.Request) {
	id := utils.UrlNumericParam(w, r, "id")
	lists, err := s.boardService.Lists(r.Context(), uint(id))
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}
	utils.Success(w, lists)
}
