package routes

import (
	"eda/internal/events/producer"
	"eda/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, eventProducer producer.EventProducer) {
	r.POST("/api/comments", func(c *gin.Context) {
		var comment models.Comment
		if err := c.ShouldBindJSON(&comment); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		message, err := json.Marshal(comment)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process comment"})
			return
		}

		err = eventProducer.Publish("comments", message)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to produce comment"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Comment published successfully"})
	})
}
