package controller

import (
	"app/config"
	"app/middleware"
	"app/model"
	"app/utils"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func CreateNews(c echo.Context) error {

	role := middleware.ExtractTokenUserRole(c)
	userId := middleware.ExtractTokenUserId(c)

	if role != "admin" {
		return c.JSON(http.StatusUnauthorized, utils.ErrorResponse("Permission denied"))
	}

	// Parse the request body into a Blog struct
	var newNews model.News
	if err := c.Bind(&newNews); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
	}

	// Fetch the user based on the users_id
	var user model.User
	if err := config.DB.First(&user, userId).Error; err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid User ID"))
	}

	// Associate the user with the blog
	newNews.User = user

	// Save the blog post to the database
	result := config.DB.Create(&newNews)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to create blog post"))
	}

	// Return the created blog post as JSON response
	return c.JSON(http.StatusCreated, utils.SuccessResponse("Success Created Data", newNews.ResponseConvert()))
}

// Get All News
func GetNews(c echo.Context) error {
	var news []model.News

	result := config.DB.Preload("User").Find(&news)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to fetch news"))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("Success get all blogs", news))
}

// Get news by id
func GetNewsID(c echo.Context) error {
	userID, _ := strconv.Atoi(c.Param("id"))

	// Fetch the blogs by their ID from the database using GORM
	var news model.News
	result := config.DB.Preload("User").First(&news, userID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, utils.ErrorResponse("News ID Not Found"))
		}
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to fetch News ID"))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("Success", news))
}

// Update News
func UpdateNews(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid news ID")
	}

	role := middleware.ExtractTokenUserRole(c)
	if role != "admin" {
		return c.JSON(http.StatusUnauthorized, utils.ErrorResponse("Permission denied"))
	}

	var news model.News
	if err := config.DB.First(&news, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "News not found")
	}

	c.Bind(&news)
	if err := config.DB.Save(&news).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update blog",
		"blog":    news,
	})
}

// Delete News
func DeleteNews(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid news ID")
	}

	role := middleware.ExtractTokenUserRole(c)
	if role != "admin" {
		return c.JSON(http.StatusUnauthorized, utils.ErrorResponse("Permission denied"))
	}

	var news model.News
	if err := config.DB.First(&news, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "News not found")
	}

	if err := config.DB.Delete(&news).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete blog",
	})
}
