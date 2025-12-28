package board

import (
	"net/http"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg/db"
	"gitlab.com/shaninalex/lumna/app/web/utils"
)

type boardPayload struct {
	Name string `json:"name"`
}

func (s *BoardHandler) Patch(w http.ResponseWriter, r *http.Request) {
	id := utils.UrlNumericParam(w, r, "id")
	payload, err := utils.BodyParser[boardPayload](r)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}

	err = s.boardService.Update(
		r.Context(),
		uint(id),
		db.Set(
			db.Field("name", payload.Name),
		),
	)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}
	board, err := s.boardService.GetBoard(r.Context(), uint(id))
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}
	utils.Success(w, board, "Updated")
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

type listPayload struct {
	Name  string `json:"name"`
	Order uint   `json:"order"`
}

func (s *BoardHandler) ListsCreate(w http.ResponseWriter, r *http.Request) {
	id := utils.UrlNumericParam(w, r, "id")
	payload, err := utils.BodyParser[listPayload](r)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}
	boardList, err := s.boardService.ListCreate(r.Context(), &models.BoardList{
		Name:    payload.Name,
		BoardId: uint(id),
		Order:   payload.Order,
	})
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}
	utils.Success(w, boardList)
}

func (s *BoardHandler) ListsPatch(w http.ResponseWriter, r *http.Request) {
	id := utils.UrlNumericParam(w, r, "listId")
	payload, err := utils.BodyParser[listPayload](r)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}
	err = s.boardService.ListUpdate(
		r.Context(),
		uint(id),
		db.Set(
			db.Field("name", payload.Name),
			db.Field("list_order", payload.Order),
		),
	)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}
	utils.Success(w, nil, "Updated")
}

func (s *BoardHandler) ListsDelete(w http.ResponseWriter, r *http.Request) {
	id := utils.UrlNumericParam(w, r, "listId")
	if err := s.boardService.ListDelete(r.Context(), uint(id)); err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}
	utils.Success(w, nil, "Deleted")
}
