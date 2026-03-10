package main

import (
	"log"
	"os"

	"github.com/PanaPapad/Pedestal/backend/db"
	"github.com/PanaPapad/Pedestal/backend/handlers"
	"github.com/PanaPapad/Pedestal/backend/repositories"
	"github.com/PanaPapad/Pedestal/backend/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// load env variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: unable to find .env file")
	}

	// read the connection string
	dsn := os.Getenv("DATABASE_URI")
	// connect with datavase using Connect function from db package
	database := db.Connect(dsn)

	// create a repos - pass the database
	blogRepo := &repositories.BlogRepository{DB: database}
	podcastRepo := &repositories.PodcastRepository{DB: database}
	// use the repos to create a handlers
	blogHandler := &handlers.BlogHandler{Repo: blogRepo}
	podcastHandler := &handlers.PodcastHandler{Repo: podcastRepo}
	// create a gin engine
	router := gin.Default()

	// add the blog routes to the engine using the blogHandler
	routes.RegisterBlogs(router, blogHandler)
	routes.RegisterPodcasts(router, podcastHandler)

	router.Use(cors.Default())
	router.Run(":8080")

}
