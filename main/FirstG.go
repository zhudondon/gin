package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main() {
	router := gin.Default()
	router.GET("/first", func(context *gin.Context) {
		name := context.DefaultQuery("name", "don")
		job := context.Query("job")

		//context.String(http.StatusOK, "first name: %s , job: %s", name, job)

		router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
			// your custom format
			return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
				param.ClientIP,
				param.TimeStamp.Format(time.RFC1123),
				param.Method,
				param.Path,
				param.Request.Proto,
				param.StatusCode,
				param.Latency,
				param.Request.UserAgent(),
				param.ErrorMessage,
			)
		}))

		context.JSON(http.StatusOK, gin.H{
			"status":  "posted",
			"message": "ok",
			"job":     job,
			"name":    name,
		})
	})

	router.GET("/second", func(context *gin.Context) {
		name := context.DefaultQuery("name", "sss")
		job := context.Query("job")

		//context.String(http.StatusOK, "first name: %s , job: %s", name, job)

		router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
			// your custom format
			return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
				param.ClientIP,
				param.TimeStamp.Format(time.RFC1123),
				param.Method,
				param.Path,
				param.Request.Proto,
				param.StatusCode,
				param.Latency,
				param.Request.UserAgent(),
				param.ErrorMessage,
			)
		}))

		context.JSON(http.StatusOK, gin.H{
			"status":  "posted",
			"message": "ok",
			"job":     job,
			"name":    name,
		})
	})

	/**
	curl -v -X POST http://127.0.0.1:8030/loginJSON -H 'Content-type:application/json' -d "{\"user\": \"don\",\"password\":\"123\"}}"

	*/
	router.POST("/loginJSON", func(context *gin.Context) {
		var json Login
		router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
			// your custom format
			return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"    %s\n",
				param.ClientIP,
				param.TimeStamp.Format(time.RFC1123),
				param.Method,
				param.Path,
				param.Request.Proto,
				param.StatusCode,
				param.Latency,
				param.Request.UserAgent(),
				param.ErrorMessage,
				json,
			)
		}))

		if err := context.ShouldBindJSON(&json); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "msg": json})
			return
		}

		if json.User != "don" || json.Password != "123" {
			context.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		context.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	router.Use(gin.Recovery())
	router.Run(":8030")
}
