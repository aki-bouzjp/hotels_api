package controller

import (
	"app/src/db"
	"app/src/logger"
	"app/src/vacancy"

	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func VacancyHandler(c *gin.Context, db *db.DB) {
	ids := c.QueryArray("ids[]")
	vacancies, err := vacancy.GetVacancies(ids, db)
	if err != nil {
		logger.Error("Failed to get vacancies. err: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err,
		})
	}
	jsonVacancies, err := json.Marshal(vacancies)
	if err != nil {
		logger.Error("Failed to marshal vacancies. err: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"vacancies": string(jsonVacancies),
		})
	}
}
