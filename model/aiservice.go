package model

type ContestRequest struct {
	Gender   string `json:"gender"`
	Category string `json:"category"`
}

type AIResponse struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}

type ChatRequest struct {
	ContestName string `json:"contest_name"`
}
