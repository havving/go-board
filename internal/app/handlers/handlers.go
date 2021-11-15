package handlers

import (
	"github.com/labstack/echo/v4"
	"go-board/internal/app/handlers/board"
)

type APIHandlerBlock struct{}

func SendError(ctx echo.Context, code int, errs ...error) error {
	return SendError(ctx, code, errs...)
}

// OpenBoardListHandler GET /api/board
func (h *APIHandlerBlock) OpenBoardListHandler(ctx echo.Context) error {
	return board.SelectBoardList(ctx)
}

// InsertBoardHandler POST /api/board/write
func (h *APIHandlerBlock) InsertBoardHandler(ctx echo.Context) error {
	return board.InsertBoard(ctx)
}

// OpenBoardDetailHandler GET /api/board/:boardIdx
func (h *APIHandlerBlock) OpenBoardDetailHandler(ctx echo.Context) error {
	return board.SelectBoardDetail(ctx)
}

// UpdateBoardHandler PUT /api/board/:boardIdx
func (h *APIHandlerBlock) UpdateBoardHandler(ctx echo.Context) error {
	return board.UpdateBoard(ctx)
}

// DeleteBoardHandler DELETE /api/board/:boardIdx
func (h *APIHandlerBlock) DeleteBoardHandler(ctx echo.Context) error {
	return board.DeleteBoard(ctx)
}