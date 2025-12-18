package board_test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/lumna/app/services"
	"gitlab.com/shaninalex/lumna/app/tests"
	"gitlab.com/shaninalex/lumna/app/web/api/board"
)

func Test_ApiBoardController_Patch(t *testing.T) {
	ctx := tests.Context()
	router := tests.NewTestRouter(ctx)

	projectA := tests.CreateProjectWithName(ctx, "ProjectA")
	boardA := tests.CreateBoard(ctx, projectA.GetId(), "A")
	_ = tests.CreateBoard(ctx, projectA.GetId(), "B")
	_ = tests.CreateBoard(ctx, projectA.GetId(), "C")

	board.RegisterBoardController(router)

	newName := "new name"
	url := fmt.Sprintf("/api/v1/board/%d", boardA.GetId())
	reqBody := fmt.Sprintf("{\"name\":\"%s\"}", newName)

	req := httptest.NewRequest(http.MethodPatch, url, bytes.NewBufferString(reqBody))
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code, "Should return status: \"200 OK\"")

	service := services.NewBoardService()
	dbBoard, err := service.GetBoard(ctx, boardA.GetId())
	assert.NoError(t, err)
	assert.Equal(t, newName, dbBoard.Name)
	assert.Equal(t, boardA.GetId(), dbBoard.GetId())
}

func Test_ApiBoardController_Delete(t *testing.T) {
	ctx := tests.Context()
	router := tests.NewTestRouter(ctx)

	projectA := tests.CreateProjectWithName(ctx, "ProjectA")
	boardA := tests.CreateBoard(ctx, projectA.GetId(), "A")
	_ = tests.CreateBoard(ctx, projectA.GetId(), "B")
	_ = tests.CreateBoard(ctx, projectA.GetId(), "C")

	board.RegisterBoardController(router)

	url := fmt.Sprintf("/api/v1/board/%d", boardA.GetId())
	req := httptest.NewRequest(http.MethodDelete, url, nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code, "Should return status: \"200 OK\"")

	service := services.NewBoardService()
	dbBoard, err := service.GetBoard(ctx, boardA.GetId())
	assert.Error(t, err)
	assert.Nil(t, dbBoard)
}

func Test_ApiBoardController_ListsGet(t *testing.T) {
	ctx := tests.Context()
	router := tests.NewTestRouter(ctx)

	projectA := tests.CreateProjectWithName(ctx, "ProjectA")
	boardA := tests.CreateBoard(ctx, projectA.GetId(), "A")
	boardListA := tests.CreateBoardList(ctx, boardA.GetId(), "nameA")
	boardListB := tests.CreateBoardList(ctx, boardA.GetId(), "nameB")

	board.RegisterBoardController(router)

	url := fmt.Sprintf("/api/v1/board/%d/lists", boardA.GetId())
	req := httptest.NewRequest(http.MethodGet, url, nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code, "Should return status: \"200 OK\"")
	assert.Contains(t, rr.Body.String(), boardListA.Name)
	assert.Contains(t, rr.Body.String(), boardListB.Name)
}

func Test_ApiBoardController_ListsCreate(t *testing.T) {
	ctx := tests.Context()
	router := tests.NewTestRouter(ctx)

	projectA := tests.CreateProjectWithName(ctx, "ProjectA")
	boardA := tests.CreateBoard(ctx, projectA.GetId(), "A")

	board.RegisterBoardController(router)

	url := fmt.Sprintf("/api/v1/board/%d/lists", boardA.GetId())
	name := "listA"
	reqBody := fmt.Sprintf("{\"name\":\"%s\"}", name)
	req := httptest.NewRequest(http.MethodPost, url, bytes.NewBufferString(reqBody))
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code, "Should return status: \"200 OK\"")
	assert.Contains(t, rr.Body.String(), name)
}

func Test_ApiBoardController_ListsPatch(t *testing.T) {
	ctx := tests.Context()
	router := tests.NewTestRouter(ctx)

	projectA := tests.CreateProjectWithName(ctx, "ProjectA")
	boardA := tests.CreateBoard(ctx, projectA.GetId(), "A")
	boardListA := tests.CreateBoardList(ctx, boardA.GetId(), "nameA")

	board.RegisterBoardController(router)

	url := fmt.Sprintf("/api/v1/lists/%d", boardListA.GetId())
	name := "new name"
	reqBody := fmt.Sprintf("{\"name\":\"%s\"}", name)
	req := httptest.NewRequest(http.MethodPatch, url, bytes.NewBufferString(reqBody))
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code, "Should return status: \"200 OK\"")

	service := services.NewBoardService()
	dbBoardList, err := service.ListGet(ctx, boardListA.GetId())
	assert.NoError(t, err)
	assert.Equal(t, name, dbBoardList.Name)
}

func Test_ApiBoardController_ListsDelete(t *testing.T) {
	ctx := tests.Context()
	router := tests.NewTestRouter(ctx)

	projectA := tests.CreateProjectWithName(ctx, "ProjectA")
	boardA := tests.CreateBoard(ctx, projectA.GetId(), "A")
	boardListA := tests.CreateBoardList(ctx, boardA.GetId(), "nameA")

	board.RegisterBoardController(router)

	url := fmt.Sprintf("/api/v1/lists/%d", boardListA.GetId())
	req := httptest.NewRequest(http.MethodDelete, url, nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code, "Should return status: \"200 OK\"")

	service := services.NewBoardService()
	dbBoardList, err := service.ListGet(ctx, boardListA.GetId())
	assert.Error(t, err)
	assert.Nil(t, dbBoardList)
}
