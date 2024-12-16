package handler

import (
	"backend-bangkit/dto"
	"backend-bangkit/entity"
	"backend-bangkit/pkg/common"
	"backend-bangkit/pkg/errs"
	service "backend-bangkit/services"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArticleHandler struct {
	articleService service.ArticleService
}

func NewArticleHandler(articleService service.ArticleService) *ArticleHandler {
	return &ArticleHandler{articleService}
}

func (h *ArticleHandler) CreateArticle(c *gin.Context) {

	userData, ok := c.MustGet("userData").(*entity.User)

	if !ok {
		newError := errs.NewBadRequest("Failed to get user data")
		c.JSON(newError.StatusCode(), newError)
		return
	}

	var req dto.CreateArticleRequest
	requestJson := c.PostForm("request")
	if err := json.Unmarshal([]byte(requestJson), &req); err != nil {
		c.JSON(http.StatusBadRequest, errs.NewBadRequest("Invalid Request Structure"))
		return
	}

	// Ambil file dari request form-data
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, errs.NewBadRequest("failed to get file from request"))
		return
	}

	req.Creator = userData.Username

	req.FileSumary = file
	req.FileHeader = header
	req.ContentType = header.Header.Get("Content-Type")
	defer file.Close()

	article, err2 := h.articleService.CreateArticle(req)
	if err2 != nil {
		c.JSON(err2.StatusCode(), err2)
		return
	}

	c.JSON(common.BuildResponse(http.StatusCreated, article))
}

func (h *ArticleHandler) GetAllArticles(c *gin.Context) {

	tagging := c.Query("tag")
	articles, err := h.articleService.GetAllArticles(tagging)
	if err != nil {
		c.JSON(err.StatusCode(), err)
		return
	}

	c.JSON(common.BuildResponse(http.StatusOK, articles))
}

func (h *ArticleHandler) DeleteArticle(c *gin.Context) {
	id := c.Param("id")
	if err := h.articleService.DeleteArticle(id); err != nil {
		c.JSON(err.StatusCode(), err)
		return
	}

	c.JSON(common.BuildResponse(http.StatusOK, nil))
}
