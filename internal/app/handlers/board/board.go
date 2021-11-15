package board

import (
	"github.com/labstack/echo/v4"
	"go-board/internal/models"
	"net/http"
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
	var err error
	return err
}

// SelectBoardDetail 게시글 상세 내용 조회
func SelectBoardDetail(ctx echo.Context) error {
	var err error
	return err
}

// UpdateBoard 게시글 상세 내용 수정
func UpdateBoard(ctx echo.Context) error {
	var err error
	return err
}

// DeleteBoard 게시글 삭제
func DeleteBoard(ctx echo.Context) error {
	var err error
	return err
}
