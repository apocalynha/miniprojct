package utils

import "app/model"

type ShowNewsResponse struct {
	ID        uint   `json:"id"`
	UpdatedAt string `json:"updated_at"`
	User      struct {
		Name string `json:"name"`
		Role string `json:"role"`
	}
	Tittle  string `json:"tittle" `
	Content string `json:"content"`
}

func GetNewsResponse(news model.News) ShowNewsResponse {
	response := ShowNewsResponse{
		ID:        news.ID,
		UpdatedAt: news.UpdatedAt.String(),
		Tittle:    news.Tittle,
		Content:   news.Content,
	}
	response.User.Name = news.User.Name
	response.User.Role = news.User.Role

	return response
}
