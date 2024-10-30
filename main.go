package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
)

func ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}

func StartServer() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	httpClient := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	r.GET("/", func(c *gin.Context) {
		value, exists := c.GetQuery("url")

		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "url param not found",
			})
			return
		}

		if _, err := url.ParseRequestURI(value); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid URL",
			})
			return
		}
		req, err := http.NewRequest(http.MethodGet, value, nil)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		for k, v := range c.Request.Header {
			c.Header(k, v[0])
		}

		resp, err := httpClient.Do(req)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		defer resp.Body.Close()
		for k, v := range resp.Header {
			c.Header(k, v[0])
		}
		c.DataFromReader(
			resp.StatusCode,
			resp.ContentLength,
			resp.Header.Get("Content-Type"),
			resp.Body,
			nil,
		)
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := r.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}

func main() {
	ConfigRuntime()
	StartServer()
}
