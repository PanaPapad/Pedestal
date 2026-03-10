package routes

import (
	"github.com/PanaPapad/Pedestal/backend/handlers"
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

func RegisterPodcasts(router *gin.Engine, podcastHandler *handlers.PodcastHandler) {
	/*
		Takes as arguments a gingonic engine and a handler and creates endpoints for blogs
	*/
	api := router.Group("/api")
	{
		api.GET("/podcasts", podcastHandler.GetAllPodcasts)
		api.POST("/podcasts", podcastHandler.CreatePodcast)
		api.GET("/podcasts/:slug", podcastHandler.GetOnePodcast)
		api.DELETE("/podcasts/:slug", podcastHandler.DeletePodcast)
	}
}
