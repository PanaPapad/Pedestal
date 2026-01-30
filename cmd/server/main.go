package main

import (
	"log"
	"os"

	"github.com/PanaPapad/Pedestal/internal/db"
	"github.com/PanaPapad/Pedestal/internal/handlers"
	"github.com/PanaPapad/Pedestal/internal/repositories"
	"github.com/PanaPapad/Pedestal/internal/routes"
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

	// create a blogrepo - pass the database
	blogRepo := &repositories.BlogRepository{DB: database}
	// use the repo to create a blog handler
	blogHandler := &handlers.BlogHandler{Repo: blogRepo}
	// create a gin engine
	router := gin.Default()
	// add the blog routes to the engine using the blogHandler
	routes.RegisterBlogs(router, blogHandler)

	router.Run(":8080")

}
