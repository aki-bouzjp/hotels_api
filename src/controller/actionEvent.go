package controller

import (
	"app/src/actionEvent"
	"app/src/db"
	"app/src/logger"

	"net/http"

	"github.com/gin-gonic/gin"
)

type EventRequest struct {
	Logs []actionEvent.EventLog `json:"logs"`
}

func ActionLogHandler(c *gin.Context, db *db.DB, ac *actionEvent.ActionEvent) {
	var er EventRequest
	if err := c.ShouldBindJSON(&er); err != nil {
		logger.Error("Failed to bind request json. err: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ac.Append(er.Logs)
	c.JSON(http.StatusOK, gin.H{})
}
