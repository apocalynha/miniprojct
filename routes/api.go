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

	// user
	users := e.Group("/users")
	users.Use(echojwt.JWT([]byte(SecretKey)))
	e.POST("/users/register", controller.Register)
	e.POST("/users/login", controller.Login)
	users.GET("", controller.GetAllUser)
	users.GET("/users/:id", controller.GetUserByID)
	users.PUT("/users/:id", controller.UpdateUser)
	users.DELETE("/users/:id", controller.DeleteUser)

	// news
	news := e.Group("/news")
	news.Use(echojwt.JWT([]byte(SecretKey)))
	e.GET("/news", controller.GetNews)
	e.GET("/news/:id", controller.GetNewsID)
	news.POST("", controller.CreateNews)
	news.PUT("/:id", controller.UpdateNews)
	news.DELETE("/:id", controller.DeleteNews)

	// contest
	contest := e.Group("/contest")
	contest.Use(echojwt.JWT([]byte(SecretKey)))
	e.GET("/contest", controller.GetContests)
	e.GET("/contest/:id", controller.GetContestID)
	contest.POST("", controller.CreateContest)
	contest.PUT("/:id", controller.UpdateContest)
	contest.DELETE("/:id", controller.DeleteContest)

	// contestant
	contestant := e.Group("/contestant")
	contestant.Use(echojwt.JWT([]byte(SecretKey)))
	contestant.GET("", controller.GetContestants)
	contestant.GET("/:id", controller.GetContestantID)
	contestant.POST("", controller.CreateContestant)
	contestant.PUT("/:id", controller.UpdateContestant)
	contestant.DELETE("/:id", controller.DeleteContestant)

	// ai
	ai := e.Group("")
	ai.Use(echojwt.JWT([]byte(SecretKey)))
	ai.POST("/contest-explanation", controller.GetContestExplanation)
	ai.POST("/contest-recommend", controller.GetContestRecommendation)

	return e

}
