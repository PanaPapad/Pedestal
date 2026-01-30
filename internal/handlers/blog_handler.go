package handlers

import (
	"net/http"

	"github.com/PanaPapad/Pedestal/internal/models"
	"github.com/PanaPapad/Pedestal/internal/repositories"
	utils "github.com/PanaPapad/Pedestal/internal/utlis"
	"github.com/gin-gonic/gin"
)

// Define a handler object: contains dependencies (database access by having access to methods inside the repositories folder)
type BlogHandler struct {
	Repo *repositories.BlogRepository
}

func (h *BlogHandler) GetAllBlogs(c *gin.Context) {
	/*
		Function that handles the GET /blogs route
		 1. Calls the repository method to get all published blog posts
		 2. Returns them as JSON
	*/
	posts, err := h.Repo.GetAllBlogs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, posts)
}

func (h *BlogHandler) CreateBlog(c *gin.Context) {
	/*
		Function that handles the Post /blogs route
		 1. Calls the repository method to create a blog post and add it to the database
		 2. Returns the newly created post as json
	*/
	var blog_post models.BlogPost

	err := c.ShouldBindBodyWithJSON(&blog_post)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// create the slug by calling the utils function
	blog_post.Slug = utils.Slugify(blog_post.Title)
	// if no status is given mark it as draft
	if blog_post.Status == "" {
		blog_post.Status = "draft"
	}
	// call the CreateBlog method from Repo to create the blog
	err = h.Repo.CreateBlog(&blog_post)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// return to the gin context the correct status and the blog post
	c.JSON(http.StatusCreated, blog_post)

}

func (h *BlogHandler) GetOneBlog(c *gin.Context) {
	/*
		Function that handles the Get /post/:slug endpoint
		 1. Calls the repository method to get the blog from the database
		 2. Returns the blog as json
	*/
	slug := c.Param("slug")

	blog_post, err := h.Repo.GetOneBlog(slug)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if blog_post == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "blog post not found"})
		return
	}

	c.JSON(http.StatusOK, blog_post)
}

func (h *BlogHandler) DeleteBlog(c *gin.Context) {

	slug := c.Param("slug")

	err := h.Repo.DeleteBlog(slug)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"sucess": "blog deleted sucessfully"})
}
