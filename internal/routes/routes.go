package routes

import (
	"github.com/PanaPapad/Pedestal/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterBlogs(router *gin.Engine, blogHandler *handlers.BlogHandler) {
	/*
		Takes as arguments a gingonic engine and a handler and creates endpoints for blogs
	*/

	api := router.Group("/api")
	{
		api.GET("/blogs", blogHandler.GetAllBlogs)
		api.POST("blogs", blogHandler.CreateBlog)
		api.GET("/blogs/:slug", blogHandler.GetOneBlog)
		api.DELETE("/blogs/:slug", blogHandler.DeleteBlog)
	}
}
