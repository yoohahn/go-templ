package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/yoohahn/go-templ/view/auth"
)

type AuthHandler struct {
}

func (h AuthHandler) HandleAuthView(c echo.Context) error {
	return render(c, auth.AuthRender())
}
