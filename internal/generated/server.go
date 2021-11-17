package api

import "github.com/labstack/echo/v4"

type ServerInterface interface {
	// OpenBoardListHandler (GET /api/board)
	OpenBoardListHandler(ctx echo.Context) error

	// InsertBoardHandler (POST /api/board/write)
	InsertBoardHandler(ctx echo.Context) error

	// OpenBoardDetailHandler (GET /api/board/{boardIdx})
	OpenBoardDetailHandler(ctx echo.Context) error

	// UpdateBoardHandler (PUT /api/board/{boardIdx})
	UpdateBoardHandler(ctx echo.Context) error

	// DeleteBoardHandler (DELETE /api/board/{boardIdx})
	DeleteBoardHandler(ctx echo.Context) error
}

type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

func (w *ServerInterfaceWrapper) OpenBoardListServer(ctx echo.Context) error {
	err := w.Handler.OpenBoardListHandler(ctx)

	return err
}

func (w *ServerInterfaceWrapper) InsertBoardServer(ctx echo.Context) error {
	err := w.Handler.InsertBoardHandler(ctx)

	return err
}

func (w *ServerInterfaceWrapper) OpenBoardDetailServer(ctx echo.Context) error {
	err := w.Handler.OpenBoardDetailHandler(ctx)

	return err
}

func (w *ServerInterfaceWrapper) UpdateBoardServer(ctx echo.Context) error {
	err := w.Handler.UpdateBoardHandler(ctx)

	return err
}

func (w *ServerInterfaceWrapper) DeleteBoardServer(ctx echo.Context) error {
	err := w.Handler.DeleteBoardHandler(ctx)

	return err
}

type EchoRouter interface {
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

func RegisterHandlers(r EchoRouter, si ServerInterface) { // 웹 서버 핸들러 생성
	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	r.GET("/api/board", wrapper.OpenBoardListServer)
	r.POST("/api/board/write", wrapper.InsertBoardServer)
	r.GET("/api/board/:boardIdx", wrapper.OpenBoardDetailServer)
	r.PUT("/api/board/:boardIdx", wrapper.UpdateBoardServer)
	r.DELETE("/api/board/:boardIdx", wrapper.DeleteBoardServer)
}
