package model

type ContestRequest struct {
	Gender   string `json:"gender"`
	Category string `json:"category"`
}

type ChatRequest struct {
	ContestName string `json:"contest_name"`
}
