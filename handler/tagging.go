package handler

import (
	"backend-bangkit/dto"
	"backend-bangkit/pkg/common"
	"backend-bangkit/pkg/errs"
	"net/http"

	service "backend-bangkit/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TaggingHandler struct {
	taggingService service.TaggingService
}

func NewTaggingHandler(taggingService service.TaggingService) *TaggingHandler {
	return &TaggingHandler{taggingService}
}

func (h *TaggingHandler) CreateTag(c *gin.Context) {
	var req dto.CreateTaggingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errs.NewBadRequest("Invalid Request"))
		return
	}

	tag, err2 := h.taggingService.CreateTag(req)
	if err2 != nil {
		c.JSON(err2.StatusCode(), err2)
		return
	}

	c.JSON(common.BuildResponse(http.StatusCreated, tag))
}

func (h *TaggingHandler) GetAllTags(c *gin.Context) {
	tags, err := h.taggingService.GetAllTags()
	if err != nil {
		c.JSON(err.StatusCode(), err)
		return
	}

	c.JSON(common.BuildResponse(http.StatusOK, tags))
}

func (h *TaggingHandler) UpdateTag(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, errs.NewBadRequest("Invalid Exercise ID"))
		return
	}
	var req dto.CreateTaggingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errs.NewBadRequest("Invalid Request"))
		return
	}

	tag, err2 := h.taggingService.UpdateTag(id, req)
	if err2 != nil {
		c.JSON(err2.StatusCode(), err2)
		return
	}

	c.JSON(common.BuildResponse(http.StatusOK, tag))
}

func (h *TaggingHandler) DeleteTag(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, errs.NewBadRequest("Invalid Exercise ID"))
		return
	}
	if err2 := h.taggingService.DeleteTag(id); err2 != nil {
		c.JSON(err2.StatusCode(), err2)
		return
	}

	c.JSON(common.BuildResponse(http.StatusOK, nil))
}
