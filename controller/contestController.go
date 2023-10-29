package controller

import (
	"app/config"
	"app/middleware"
	"app/model"
	"app/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func GetContests(c echo.Context) error {
	var contests []model.Contest

	err := config.DB.Find(&contests).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve contests"))
	}

	if len(contests) == 0 {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse("No contests available"))
	}

	response := make([]utils.ContestResponse, len(contests))
	for i, contest := range contests {
		response[i] = utils.GetContestResponse(contest)
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("Contest data successfully retrieved", response))
}

func GetContestID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid Contest ID"))
	}

	var contest model.Contest

	if err := config.DB.First(&contest, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve contest"))
	}

	response := utils.GetContestResponse(contest)

	return c.JSON(http.StatusOK, utils.SuccessResponse("Contest data successfully retrieved", response))
}

func CreateContest(c echo.Context) error {
	var NewContest model.Contest

	role := middleware.ExtractTokenUserRole(c)
	if role != "admin" {
		return c.JSON(http.StatusUnauthorized, utils.ErrorResponse("Permission denied"))
	}

	if err := c.Bind(&NewContest); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
	}

	if err := config.DB.Create(&NewContest).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to create contest"))
	}

	response := utils.GetContestResponse(NewContest)

	return c.JSON(http.StatusCreated, utils.SuccessResponse("Contest successfully created", response))
}

func UpdateContest(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid Contest ID"))
	}

	var UpdatedContest model.Contest

	role := middleware.ExtractTokenUserRole(c)
	if role != "admin" {
		return c.JSON(http.StatusUnauthorized, utils.ErrorResponse("Permission denied"))
	}

	if err := c.Bind(&UpdatedContest); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
	}

	var ExistingContest model.Contest
	result := config.DB.First(&ExistingContest, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve contest"))
	}

	config.DB.Model(&ExistingContest).Updates(UpdatedContest)

	response := utils.GetContestResponse(ExistingContest)

	return c.JSON(http.StatusOK, utils.SuccessResponse("Contest data successfully updated", response))
}

func DeleteContest(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid Contest ID"))
	}

	var existingContest model.Contest
	result := config.DB.First(&existingContest, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve contest"))
	}

	role := middleware.ExtractTokenUserRole(c)
	if role != "admin" {
		return c.JSON(http.StatusUnauthorized, utils.ErrorResponse("Permission denied"))
	}

	config.DB.Delete(&existingContest)

	return c.JSON(http.StatusOK, utils.SuccessResponse("Contest data successfully deleted", nil))
}
