package board

import (
	"github.com/labstack/echo/v4"
	"go-board/internal/models"
	"net/http"
	"strconv"
	"time"
)

// SelectBoardList 게시글 목록 조회
func SelectBoardList(ctx echo.Context) error {
	list, err := models.Board.SelectBoardList()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	err = ctx.JSON(http.StatusOK, list)

	return err
}

// InsertBoard 게시글 작성
func InsertBoard(ctx echo.Context) error {
	title := ctx.FormValue("title")
	contents := ctx.FormValue("contents")
	now := time.Now()

	err := models.Board.InsertBoard(title, contents, now.String())
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Success{Success: false})
	}

	err = ctx.JSON(http.StatusCreated, models.Success{Success: true})

	return err
}

// SelectBoardDetail 게시글 상세 내용 조회
func SelectBoardDetail(ctx echo.Context) error {
	vars := ctx.Param("boardIdx")
	id, _ := strconv.Atoi(vars)

	entity, err := models.Board.SelectBoardDetail(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, entity)
	}

	err = ctx.JSON(http.StatusOK, entity)

	return err
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
		return ctx.JSON(http.StatusBadRequest, models.Success{Success: false})
	}

	err = ctx.JSON(http.StatusOK, models.Success{Success: true})

	return err
}

// DeleteBoard 게시글 삭제
func DeleteBoard(ctx echo.Context) error {
	vars := ctx.Param("boardIdx")
	id, _ := strconv.Atoi(vars)

	err := models.Board.DeleteBoard(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, models.Success{Success: false})
	}

	err = ctx.JSON(http.StatusOK, models.Success{Success: true})

	return err
}
