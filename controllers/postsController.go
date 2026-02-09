package controllers

import (
	"go-gin-crud/initializers"
	"go-gin-crud/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// Get data
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Create
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create((&post))

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it
	c.JSON(200, gin.H{
		"post": post,
	})
}
