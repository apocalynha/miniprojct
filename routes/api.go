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

	// contest
	e.GET("/contest", controller.GetContests)
	e.GET("/contest/:id", controller.GetContestID)
	eAuth.POST("/contest/create", controller.CreateContest)
	eAuth.PUT("/contest/:id", controller.UpdateContest)
	eAuth.DELETE("/contest/:id", controller.DeleteContest)

	// contestant
	eAuth.GET("/contestant", controller.GetContestants)
	eAuth.GET("/contestant/:id", controller.GetContestantID)
	eAuth.POST("/contestant/create", controller.CreateContestant)
	eAuth.PUT("/contestant/:id", controller.UpdateContestant)
	eAuth.DELETE("/contestant/:id", controller.DeleteContestant)

	// ai
	eAuth.POST("/contest-explanation", controller.GetContestExplanation)
	eAuth.POST("/contest-recommend", controller.GetContestRecommendation)

	return e

}
