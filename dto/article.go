package dto

import (
	"io"
	"mime/multipart"

	"github.com/google/uuid"
)

// Article DTO
type CreateArticleRequest struct {
	Title       string `json:"title" binding:"required"`
	Content     string `json:"content" binding:"required"`
	Creator     string
	TagName     string `json:"tag_name" binding:"required"`
	FileSumary  io.Reader
	FileHeader  *multipart.FileHeader
	ContentType string
	URL         string `json:"url"` // ID dari Tag
}

type ArticleResponse struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Creator   string    `json:"creator" binding:"required"`
	Tag       string    `json:"tag"` // Nama dari Tag
	CreatedAt string    `json:"created_at"`
	URL       string    `json:"url"`
}

type CreateTaggingRequest struct {
	Name string `json:"name" binding:"required"`
}

type TaggingResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
