package service

import (
	"backend-bangkit/dto"
	"backend-bangkit/entity"
	"backend-bangkit/pkg/errs"
	"backend-bangkit/pkg/gcloud"
	"backend-bangkit/repository"
)

type ArticleService interface {
	CreateArticle(req dto.CreateArticleRequest) (*dto.ArticleResponse, errs.MessageErr)
	GetAllArticles(tagging string) ([]dto.ArticleResponse, errs.MessageErr)
	DeleteArticle(id string) errs.MessageErr
}

type articleService struct {
	articleRepo repository.ArticleRepository
	tagRepo     repository.TaggingRepository
	gcsUploader *gcloud.GCSUploader
}

func NewArticleService(articleRepo repository.ArticleRepository, tagRepo repository.TaggingRepository, gcsUploader *gcloud.GCSUploader) ArticleService {
	return &articleService{articleRepo, tagRepo, gcsUploader}
}

func (s *articleService) CreateArticle(req dto.CreateArticleRequest) (*dto.ArticleResponse, errs.MessageErr) {
	// Validate TagID
	tag, err := s.tagRepo.FindByName(req.TagName)
	if err != nil {
		return nil, err
	}

	fileURL, err := s.gcsUploader.UploadFile(req.FileSumary, req.ContentType)
	if err != nil {
		return nil, err
	}

	article := entity.Article{
		Title:    req.Title,
		Content:  req.Content,
		TagID:    tag.ID,
		Creator:  req.Creator,
		UrlImage: fileURL,
	}

	err = s.articleRepo.Create(&article)
	if err != nil {
		return nil, err
	}

	return &dto.ArticleResponse{
		ID:        article.ID,
		Title:     article.Title,
		Content:   article.Content,
		Tag:       tag.Name,
		Creator:   article.Creator,
		CreatedAt: article.CreatedAt.String(),
		URL:       article.UrlImage,
	}, nil
}

func (s *articleService) GetAllArticles(tagging string) ([]dto.ArticleResponse, errs.MessageErr) {
	articles, err := s.articleRepo.FindAll(tagging)
	if err != nil {
		return nil, err
	}

	var result []dto.ArticleResponse
	for _, article := range articles {
		result = append(result, dto.ArticleResponse{
			ID:        article.ID,
			Title:     article.Title,
			Content:   article.Content,
			Tag:       article.Tag.Name,
			Creator:   article.Creator,
			CreatedAt: article.CreatedAt.String(),
			URL:       article.UrlImage,
		})
	}

	return result, nil
}

func (s *articleService) DeleteArticle(id string) errs.MessageErr {
	err := s.articleRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
