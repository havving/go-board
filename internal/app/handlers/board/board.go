package board

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go-board/internal/defs"
	"go-board/internal/generated"
	"go-board/internal/models"
	"net/http"
	"strconv"
	"time"
)

// SelectBoardList 게시글 목록 조회
func SelectBoardList(ctx echo.Context) error {
	list, err := models.Board.SelectBoardList()
	if err != nil {
		return sendError(ctx, http.StatusInternalServerError, err)
	}

	return sendResponse(ctx, list)
}

// InsertBoard 게시글 작성
func InsertBoard(ctx echo.Context) error {
	title := ctx.FormValue("title")
	contents := ctx.FormValue("contents")
	now := time.Now()

	err := models.Board.InsertBoard(title, contents, now.String())
	if err != nil {
		return sendError(ctx, http.StatusBadRequest, err)
	}

	return sendResponse(ctx, models.Success{Success: true})
}

// SelectBoardDetail 게시글 상세 내용 조회
func SelectBoardDetail(ctx echo.Context) error {
	vars := ctx.Param("boardIdx")
	id, _ := strconv.Atoi(vars)

	entity, err := models.Board.SelectBoardDetail(id)
	if err != nil {
		return sendError(ctx, http.StatusNotFound, err)
	}

	return sendResponse(ctx, entity)
}

// UpdateBoard 게시글 상세 내용 수정
func UpdateBoard(ctx echo.Context) error {
	title := ctx.FormValue("title")
	contents := ctx.FormValue("contents")
	now := time.Now()
	boardIdx := ctx.FormValue("boardIdx")
	id, _ := strconv.Atoi(boardIdx)

	err := models.Board.UpdateBoard(title, contents, now.String(), id)
	if err != nil {
		return sendError(ctx, http.StatusBadRequest, err)
	}

	return sendResponse(ctx, models.Success{Success: true})
}

// DeleteBoard 게시글 삭제
func DeleteBoard(ctx echo.Context) error {
	vars := ctx.Param("boardIdx")
	id, _ := strconv.Atoi(vars)

	err := models.Board.DeleteBoard(id)
	if err != nil {
		return sendError(ctx, http.StatusNotFound, err)
	}

	return sendResponse(ctx, models.Success{Success: true})
}

// sendResponse 200 response
func sendResponse(ctx echo.Context, res interface{}) error {
	return ctx.JSON(http.StatusOK, res)
}

// sendError error response
func sendError(ctx echo.Context, code int, errs ...error) error {
	res := &api.GenericResponse{
		Code: code,
	}

	var err error
	if len(errs) == 0 || errs[0] == nil {
		err = defs.NewError(err)
	} else {
		err = errs[0]
	}

	switch err.(type) {
	case *echo.HTTPError:
		e := err.(*echo.HTTPError)
		res.Message = e.Message
		log.Warn(defs.NewError(e.Code), "%+v", e.Internal)
	default:
		res.Message = err.Error()
		log.Warn(err, "")
	}

	return ctx.JSON(http.StatusTeapot, res)
}
