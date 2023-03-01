package main

import (
	"github.com/gin-gonic/gin"
	// "github.com/gin-contrib/static"

    "fortunecookie/posts"
	"fortunecookie/storage"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r
}

func main() {
	db := storage.Init()
	posts.Migrate(db)
	defer db.Close()

	r := setupRouter()
	r.Run(":8080")
}