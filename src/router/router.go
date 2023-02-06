package router

import (
	"app/src/actionEvent"
	"app/src/controller"
	"app/src/db"
	"os"

	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func New(sess *session.Session, rd *redis.Client) *gin.Engine {
	if os.Getenv("ENV") == "dev" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	db := db.New()
	ac := actionEvent.New()
	ac.Run()

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost",
		},
		AllowMethods: []string{
			"GET",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			// "Authorization",
			// "X-Requested-With",
		},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"alive": true,
		})
	})

	// curl -X GET "http://localhost:8080/vacancies?ids[]=1&ids[]=2&ids[]=3"
	r.GET("/vacancies", func(c *gin.Context) {
		controller.VacancyHandler(c, db)
	})

	// curl -X GET "http://localhost:8080/events?cps=xxxx"
	r.POST("/logs", func(c *gin.Context) {
		controller.ActionLogHandler(c, db, ac)
	})

	return r
}
