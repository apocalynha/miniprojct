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

// Create Contestant
func CreateContestant(c echo.Context) error {
	UserId := middleware.ExtractTokenUserId(c)

	var NewContestant model.Contestant
	if err := c.Bind(&NewContestant); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
	}

	NewContestant.UserID = uint(UserId)

	// Fetch the user based on the users_id
	var user model.User
	if err := config.DB.First(&user, NewContestant.UserID).Error; err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid User ID"))
	}

	// Associate the user with the contestant
	NewContestant.User = user

	// Fetch the contest based on the contest_id
	var contest model.Contest
	if err := config.DB.First(&contest, NewContestant.ContestID).Error; err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid Contest ID"))
	}

	// Associate the contest with the contestant
	NewContestant.Contest = contest

	// Check the Gender and Category requirements
	if err := utils.CheckGenderAndCategoryRequirements(contest, NewContestant); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
	}

	// Save the contestant to the database
	result := config.DB.Create(&NewContestant)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to create contestant"))
	}

	response := utils.GetContestantResponse(NewContestant)

	return c.JSON(http.StatusCreated, utils.SuccessResponse("Success Created Data", response))
}

// Get All Contestants for the currently logged-in user
func GetContestants(c echo.Context) error {
	UserId := middleware.ExtractTokenUserId(c)
	role := middleware.ExtractTokenUserRole(c)

	var contestants []model.Contestant

	if role == "admin" {
		result := config.DB.Preload("User").Preload("Contest").Find(&contestants)
		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to fetch contestants"))
		}
	} else {
		result := config.DB.Preload("User").Preload("Contest").Where("user_id = ?", UserId).Find(&contestants)
		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to fetch contestants"))
		}
	}

	var ResponseList []utils.ShowContestantResponse
	for _, contestant := range contestants {
		response := utils.GetContestantResponse(contestant)
		ResponseList = append(ResponseList, response)
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("Success get all contestants for logged-in user", ResponseList))
}

// Get Contestant by ID
func GetContestantID(c echo.Context) error {
	ContestantID, _ := strconv.Atoi(c.Param("id"))
	UserId := middleware.ExtractTokenUserId(c)

	var contestant model.Contestant
	result := config.DB.Preload("User").Preload("Contest").Where("id = ? AND user_id = ?", ContestantID, UserId).First(&contestant)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, utils.ErrorResponse("Contestant ID Not Found"))
		}
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to fetch Contestant ID"))
	}

	response := utils.GetContestantResponse(contestant)

	return c.JSON(http.StatusOK, utils.SuccessResponse("Success get contestant by ID for logged-in user", response))
}

// Update Contestant
func UpdateContestant(c echo.Context) error {
	var UpdatedContestant model.Contestant
	if err := c.Bind(&UpdatedContestant); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
	}

	ContestantID, _ := strconv.Atoi(c.Param("id"))

	role := middleware.ExtractTokenUserRole(c)
	if role != "admin" {
		return c.JSON(http.StatusUnauthorized, utils.ErrorResponse("Permission denied"))
	}

	ExistingContestant := model.Contestant{}
	if err := config.DB.First(&ExistingContestant, ContestantID).Error; err != nil {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse("Contestant not found"))
	}

	ExistingContestant.ContestantName = UpdatedContestant.ContestantName
	ExistingContestant.ContestID = UpdatedContestant.ContestID
	ExistingContestant.Gender = UpdatedContestant.Gender
	ExistingContestant.Age = UpdatedContestant.Age

	// Fetch the contest based on the contest_id
	var contest model.Contest
	if err := config.DB.First(&contest, ExistingContestant.ContestID).Error; err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid Contest ID"))
	}

	// Check the Gender and Category requirements
	if err := utils.CheckGenderAndCategoryRequirements(contest, ExistingContestant); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
	}

	if err := config.DB.Save(&ExistingContestant).Error; err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
	}

	var result model.Contestant
	if err := config.DB.Preload("User").Preload("Contest").First(&result, ContestantID).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to fetch Contestant ID"))
	}

	response := utils.GetContestantResponse(result)

	return c.JSON(http.StatusOK, utils.SuccessResponse("Success update contestant", response))
}

// Delete Contestant
func DeleteContestant(c echo.Context) error {
	IdParam := c.Param("id")
	id, err := strconv.Atoi(IdParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid Contestant ID"))
	}

	role := middleware.ExtractTokenUserRole(c)
	if role != "admin" {
		return c.JSON(http.StatusUnauthorized, utils.ErrorResponse("Permission denied"))
	}

	var contestant model.Contestant
	if err := config.DB.First(&contestant, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse("Contestant not found"))
	}

	if err := config.DB.Delete(&contestant).Error; err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete contestant",
	})
}
