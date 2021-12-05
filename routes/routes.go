package routes

import (
	"github.com/alen/echo-framework/controllers"
	"github.com/alen/echo-framework/middleware"
	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"Message": "Halo nama saya alen prastya",
		})
	})
	e.GET("/pegawai", controllers.FetchAllPegawai, middleware.IsAuthenticated)
	e.POST("/pegawai", controllers.StorePegwai, middleware.IsAuthenticated)
	e.PUT("/pegawai", controllers.UpdatePegwai, middleware.IsAuthenticated)
	e.DELETE("/pegawai", controllers.DeletePegawai, middleware.IsAuthenticated)

	// e.GET("/generate-hash/:password", controllers.GenerateFromPassword)
	e.POST("login", controllers.CheckLogin)
	e.POST("register", controllers.Reg2)
	return e
}
