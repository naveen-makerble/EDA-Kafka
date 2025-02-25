package handlers

import (
	"eda/models"
	"encoding/json"
	"fmt"
	"log"
)

func LogHanlder(message []byte) error {
	var cmt models.Comment

	if err := json.Unmarshal(message, &cmt); err != nil {
		return fmt.Errorf("failed to unmarshal message: %w", err)
	}

	log.Printf("Received comment: %s", cmt.Text)
	return nil

}
