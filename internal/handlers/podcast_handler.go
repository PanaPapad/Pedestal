package handlers

import (
	"github.com/PanaPapad/Pedestal/internal/repositories"
	"github.com/gin-gonic/gin"
)

type PodcastHandler struct {
	Repo *repositories.PodcastRepository
}

func (h *PodcastHandler) GetAllPodcasts(c *gin.Context) {
	////continue form here my friend
}
