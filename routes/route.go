package routes

import (
	kupon "projectkupon/features/kupons"
	"projectkupon/features/users"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(e *echo.Echo, uc users.Handler, bc kupon.Handler) {
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	routeUser(e, uc)
	routeKupon(e, bc)
}

func routeUser(e *echo.Echo, uc users.Handler) {
	e.POST("/users", uc.Register())
	e.POST("/login", uc.Login())
	// e.GET("/users", uc.GetListUser(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
}

func routeKupon(e *echo.Echo, kh kupon.Handler) {
	e.POST("/kupons", kh.Add(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
	e.GET("/kupons/:id", kh.GetOne(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
	e.GET("/kupons", kh.GetAll())
}
