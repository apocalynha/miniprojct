package controller

import (
	"app/config"
	"app/middleware"
	"app/model"
	"app/model/web"
	"app/utils"
	"app/utils/req"
	"app/utils/res"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Register(c echo.Context) error {
	var user web.UserRequest

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
	}

	userDb := req.PassBody(user)

	// Hash the user's password before storing it
	userDb.Password = middleware.HashPassword(userDb.Password)

	if err := config.DB.Create(&userDb).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to store user data"))
	}

	// Return the response without including a JWT token
	response := res.ConvertGeneral(userDb)

	return c.JSON(http.StatusCreated, utils.SuccessResponse("Success Created Data", response))
}

func Login(c echo.Context) error {
	var loginRequest web.LoginRequest

	if err := c.Bind(&loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
	}

	var user model.User
	if err := config.DB.Where("email = ?", loginRequest.Email).First(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, utils.ErrorResponse("Invalid login credentials"))
	}

	if err := middleware.ComparePassword(user.Password, loginRequest.Password); err != nil {
		return c.JSON(http.StatusUnauthorized, utils.ErrorResponse("Invalid login credentials"))
	}

	token := middleware.CreateToken(int(user.ID), user.Name)

	// Buat respons dengan data yang diminta
	response := web.UserLoginResponse{
		Email:    user.Email,
		Password: middleware.HashPassword(loginRequest.Password),
		Token:    token,
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("Login successful", response))
}
