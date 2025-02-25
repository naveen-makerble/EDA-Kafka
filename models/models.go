package models

type Comment struct {
	Text string `json:"text" binding:"required"`
}
