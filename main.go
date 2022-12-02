package main

import (
	"net/http"
	"time"

	"github.com/giangcoy/go-urlshortener/internal/generator"
	"github.com/giangcoy/go-urlshortener/internal/store"
	"github.com/gin-gonic/gin"
)

var (
	domain = "http://localhost:8080/%s"
)

type request struct {
	Data string `json:"data"`
	Ttl  int    `json:"ttl"` //hour
}

func main() {
	//s := store.NewMemory()
	s := store.NewRedis("127.0.0.1:6379", "-")
	g := generator.NewGenerator()
	router := gin.Default()
	router.POST("/", func(c *gin.Context) {

		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		url, err := g.Generate()
		if err != nil {
			c.String(http.StatusInternalServerError, "Internal error %s", err.Error())
			return
		}
		if err = s.Set(url, req.Data, time.Hour*time.Duration(req.Ttl)); err != nil {
			c.String(http.StatusInternalServerError, "Internal error %s", err.Error())
			return
		}

		c.String(http.StatusOK, domain, url)
	})
	router.GET("/:url", func(c *gin.Context) {
		url, err := s.Get(c.Param("url"))
		if err != nil {
			c.String(http.StatusNotFound, "Not Found %s", url)
			return
		}
		c.Redirect(http.StatusFound, url)
	})
	router.Run(":8080")
}
