package handlers

import (
	"net/http"

	"github.com/PanaPapad/Pedestal/backend/models"
	"github.com/PanaPapad/Pedestal/backend/repositories"
	utils "github.com/PanaPapad/Pedestal/backend/utlis"
	"github.com/gin-gonic/gin"
)

type PodcastHandler struct {
	Repo *repositories.PodcastRepository
}

func (h *PodcastHandler) GetAllPodcasts(c *gin.Context) {

	podcasts, err := h.Repo.GetAllPodcasts()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, podcasts)

}

func (h *PodcastHandler) CreatePodcast(c *gin.Context) {

	var podcast models.PodcastEpisode

	err := c.ShouldBindBodyWithJSON(&podcast)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	podcast.Slug = utils.Slugify(podcast.Title)

	err = h.Repo.CreatePodcast(&podcast)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, podcast)

}

func (h *PodcastHandler) GetOnePodcast(c *gin.Context) {

	slug := c.Param("slug")

	podcast, err := h.Repo.GetOnePodcast(slug)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if podcast == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Podcast not Found"})
		return
	}

	c.JSON(http.StatusOK, podcast)

}

func (h *PodcastHandler) DeletePodcast(c *gin.Context) {

	slug := c.Param("slug")

	err := h.Repo.DeletePodcast(slug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"sucess": "Podcast deleted sucessfully"})
}
