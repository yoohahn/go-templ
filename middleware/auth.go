package middleware

import (
	"strings"

	"github.com/labstack/echo/v4"
)

func Auth() echo.MiddlewareFunc {
	return LoggerWithConfig()
}

func LoggerWithConfig() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			reqUrl := c.Request().URL.String()
			isAssets := strings.HasPrefix(reqUrl, "/assets")
			if isAssets {
				return next(c)
			}

			// Ugly way to mock login. But it is just here for test
			email := c.Request().URL.Query().Get("email")
			url := "/?email="
			url += email

			hasEmail := len(email) > 0
			if !hasEmail {
				isAuthPage := strings.HasPrefix(c.Request().URL.String(), "/auth")
				if !isAuthPage {
					c.Redirect(302, "/auth")
				}
			} else if c.Request().URL.String() != url {
				c.Redirect(302, url)
			}

			return next(c)
		}
	}
}
