package controller

import (
	"app/config"
	"app/middleware"
	"app/model"
	"app/service"
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

	var NewNews model.News
	if err := c.Bind(&NewNews); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
	}

	fileheader := "photo"
	NewNews.Photo = service.CloudinaryUpload(c, fileheader)

	var user model.User
	if err := config.DB.First(&user, userId).Error; err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid User ID"))
	}

	NewNews.User = user

	result := config.DB.Create(&NewNews)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to create news post"))
	}

	response := utils.GetNewsResponse(NewNews)

	return c.JSON(http.StatusCreated, utils.SuccessResponse("Success Created Data", response))
}

// Get All News
func GetNews(c echo.Context) error {
	var news []model.News

	result := config.DB.Preload("User").Find(&news)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to fetch news"))
	}

	var ResponseList []utils.ShowNewsResponse
	for _, n := range news {
		response := utils.GetNewsResponse(n)
		ResponseList = append(ResponseList, response)
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("Success get all news", ResponseList))
}

// Get news by id
func GetNewsID(c echo.Context) error {
	NewsID, _ := strconv.Atoi(c.Param("id"))

	var news model.News
	result := config.DB.Preload("User").First(&news, NewsID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, utils.ErrorResponse("News ID Not Found"))
		}
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to fetch News ID"))
	}

	response := utils.GetNewsResponse(news)

	return c.JSON(http.StatusOK, utils.SuccessResponse("Success get news by ID", response))
}

// Update News
func UpdateNews(c echo.Context) error {
	IdParam := c.Param("id")
	id, err := strconv.Atoi(IdParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid news ID")
	}

	role := middleware.ExtractTokenUserRole(c)
	if role != "admin" {
		return c.JSON(http.StatusUnauthorized, utils.ErrorResponse("Permission denied"))
	}

	var news model.News
	if err := config.DB.Preload("User").First(&news, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "News not found")
	}

	c.Bind(&news)
	if err := config.DB.Save(&news).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("Success update news", utils.GetNewsResponse(news)))
}

// Delete News
func DeleteNews(c echo.Context) error {
	IdParam := c.Param("id")
	id, err := strconv.Atoi(IdParam)
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
		"message": "success delete news",
	})
}
