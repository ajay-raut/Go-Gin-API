package controllers

import (
	"fmt"
	"net/http"

	initializers "github.com/ajay-raut/gocrud/Initializers"
	"github.com/ajay-raut/gocrud/models"
	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {

	// get data from the request body
	var postbody struct {
		Title  string `json:"title"`
		Body   string `json:"body"`
		Author string `json:"author"`
	}

	c.Bind(&postbody)

	// create a post in the database
	post := models.Post{Title: postbody.Title, Body: postbody.Body, Author: postbody.Author}

	result := initializers.DB.Create(&post) // pass pointer of data to Create

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	// Logic to create a post
	c.JSON(http.StatusOK, gin.H{
		"message": "Post created successfully",
		"post":    post,
		"status":  "success",
		"code":    http.StatusOK,
	})
}

func PostsIndex(c *gin.Context) {
	var posts []models.Post

	// Fetch all posts from the database
	result := initializers.DB.Find(&posts)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Posts fetched successfully",
		"posts":   posts,
		"status":  "success",
		"code":    http.StatusOK,
	})
}

func PostShow(c *gin.Context) {
	var post models.Post

	// Get the post ID from the URL parameter
	postID := c.Param("id")

	// Fetch the post from the database
	result := initializers.DB.First(&post, postID)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Post not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Post fetched successfully",
		"post":    post,
		"status":  "success",
		"code":    http.StatusOK,
	})
}

func PostUpdate(c *gin.Context) {

	postID := c.Param("id")

	var postbody struct {
		Title  string `json:"title"`
		Body   string `json:"body"`
		Author string `json:"author"`
	}

	err := c.ShouldBindJSON(&postbody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	var post models.Post

	// Fetch the post from the database
	result := initializers.DB.First(&post, postID)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Post not found",
		})
		return
	}

	// Update the post with new data

	result = initializers.DB.Model(&post).Updates(models.Post{
		Title:  postbody.Title,
		Body:   postbody.Body,
		Author: postbody.Author,
	})

	fmt.Println(result.Error)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Post updated successfully",
		"post":    post,
		"status":  "success",
		"code":    http.StatusOK,
	})
}

func PostDelete(c *gin.Context) {
	// Get the post ID from the URL parameter
	postID := c.Param("id")

	// Delete the post from the database
	result := initializers.DB.Delete(&models.Post{}, postID)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Post deleted successfully",
		"status":  "success",
		"code":    http.StatusOK,
	})
}
