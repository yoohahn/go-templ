package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/yoohahn/go-templ/model"
	"github.com/yoohahn/go-templ/view"
)

type RootHandler struct {
}

func (h RootHandler) HandleRootShow(c echo.Context) error {
	usr := model.User{
		Name:      "Foo Bar",
		Email:     "foo@bar.baz",
		AvatarUrl: "/assets/avatar.png",
	}
	return render(c, view.ShowRoot(usr))
}
