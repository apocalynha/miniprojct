package controller

import (
	"app/middleware"
	"app/model"
	"app/utils"
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sashabaranov/go-openai"
	"io/ioutil"
	"net/http"
	"os"
)

func GetContestRecommendation(c echo.Context) error {
	OpenAI_Key := os.Getenv("APIOPENAI")

	role := middleware.ExtractTokenUserRole(c)
	if role != "admin" {
		return c.JSON(http.StatusUnauthorized, utils.ErrorResponse("Permission denied"))
	}

	var reqData model.ContestRequest

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
	}

	if err := json.Unmarshal(body, &reqData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
	}

	client := openai.NewClient(OpenAI_Key)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "Anda merupakan asisten yang dapat membantu untuk memberikan rekomendasi lomba.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: fmt.Sprintf("Rekomendasi lomba untuk jenis kelamin %s dan kategori %s .", reqData.Gender, reqData.Category),
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return err
	}
	recommendation := resp.Choices[0].Message.Content

	response := utils.TResponseMeta{
		Success: true,
		Message: recommendation,
	}

	return c.JSON(http.StatusOK, response)
}

func GetContestExplanation(c echo.Context) error {
	OpenAI_Key := os.Getenv("APIOPENAI")
	
	var reqData model.ChatRequest

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
	}

	if err := json.Unmarshal(body, &reqData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
	}

	client := openai.NewClient(OpenAI_Key)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "Anda merupakan asisten yang dapat membantu untuk memberikan penjelasan lomba.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: fmt.Sprintf("Berikan penjelasan dan aturan dari lomba %s .", reqData.ContestName),
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return err
	}
	recommendation := resp.Choices[0].Message.Content

	response := utils.TResponseMeta{
		Success: true,
		Message: recommendation,
	}

	return c.JSON(http.StatusOK, response)
}
