package handler

import (
	service "backend-bangkit/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MusicHandler struct {
	service service.MusicService
}

func NewMusicHandler(service service.MusicService) *MusicHandler {
	return &MusicHandler{service: service}
}

func (h *MusicHandler) GetRandomMusic(c *gin.Context) {
	limit := 10 // Default limit
	mood := c.Query("mood")
	music, err := h.service.GetRandomMusic(limit, mood)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, music)
}
