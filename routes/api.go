package routes

import (
	"app/controller"
	"app/middleware"
	"net/http"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	SecretKey := os.Getenv("JWT_KEY")

	e := echo.New()

	e.Use(middleware.NotFoundHandler)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to RESTful API Services")
	})

	// Authenticated
	eAuth := e.Group("")
	eAuth.Use(echojwt.JWT([]byte(SecretKey)))

	// user
	e.POST("/users/register", controller.Register)
	e.POST("/users/login", controller.Login)
	eAuth.GET("/users", controller.GetAllUser)
	eAuth.GET("/users/:id", controller.GetUserByID)
	eAuth.PUT("/users/:id", controller.UpdateUser)
	eAuth.DELETE("/users/:id", controller.DeleteUser)

	// news
	e.GET("/news", controller.GetNews)
	e.GET("/news/:id", controller.GetNewsID)
	eAuth.POST("/news/create", controller.CreateNews)
	eAuth.PUT("/news/:id", controller.UpdateNews)
	eAuth.DELETE("/news/:id", controller.DeleteNews)

	return e

}
