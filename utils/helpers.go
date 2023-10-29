package utils

import (
	"app/model"
	"time"
)

type ShowNewsResponse struct {
	ID        uint   `json:"id"`
	UpdatedAt string `json:"updated_at"`
	User      struct {
		Name string `json:"name"`
		Role string `json:"role"`
	}
	Tittle  string `json:"tittle" `
	Content string `json:"content"`
	Photo   string `json:"photo"`
}

func GetNewsResponse(news model.News) ShowNewsResponse {
	response := ShowNewsResponse{
		ID:        news.ID,
		UpdatedAt: news.UpdatedAt.String(),
		Tittle:    news.Tittle,
		Content:   news.Content,
		Photo:     news.Photo,
	}
	response.User.Name = news.User.Name
	response.User.Role = news.User.Role

	return response
}

type ContestResponse struct {
	ID          uint      `json:"id"`
	ContestName string    `json:"contest_name"`
	ReqGender   string    `json:"req_gender"`
	ReqCategory string    `json:"req_category"`
	Details     string    `json:"details"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func GetContestResponse(contest model.Contest) ContestResponse {
	return ContestResponse{
		ID:          contest.ID,
		ContestName: contest.ContestName,
		ReqGender:   contest.ReqGender,
		ReqCategory: contest.ReqCategory,
		Details:     contest.Details,
		CreatedAt:   contest.CreatedAt,
		UpdatedAt:   contest.UpdatedAt,
	}
}

type ShowContestantResponse struct {
	ID   uint `json:"id"`
	User struct {
		Name string `json:"name"`
		Role string `json:"role"`
	}
	Contest struct {
		ContestName string `json:"contest_name"`
		Details     string `json:"details"`
		ReqCategory string `json:"req_category"`
	}
	ContestantName string `json:"contestant_name"`
	Gender         string `json:"gender" form:"gender"`
	Age            int    `json:"age" form:"age"`
	UpdatedAt      string `json:"updated_at"`
}

func GetContestantResponse(contestant model.Contestant) ShowContestantResponse {
	response := ShowContestantResponse{
		ID:             contestant.ID,
		ContestantName: contestant.ContestantName,
		Gender:         contestant.Gender,
		Age:            contestant.Age,
		UpdatedAt:      contestant.UpdatedAt.String(),
	}

	response.User.Name = contestant.User.Name
	response.User.Role = contestant.User.Role

	response.Contest.ContestName = contestant.Contest.ContestName
	response.Contest.Details = contestant.Contest.Details
	response.Contest.ReqCategory = contestant.Contest.ReqCategory

	return response
}
