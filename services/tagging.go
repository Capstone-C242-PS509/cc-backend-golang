package service

import (
	"backend-bangkit/dto"
	"backend-bangkit/entity"
	"backend-bangkit/pkg/errs"
	"backend-bangkit/repository"

	"github.com/google/uuid"
)

type TaggingService interface {
	CreateTag(req dto.CreateTaggingRequest) (*dto.TaggingResponse, errs.MessageErr)
	GetAllTags() ([]dto.TaggingResponse, errs.MessageErr)
	UpdateTag(id uuid.UUID, req dto.CreateTaggingRequest) (*dto.TaggingResponse, errs.MessageErr)
	DeleteTag(id uuid.UUID) errs.MessageErr
}

type taggingService struct {
	taggingRepo repository.TaggingRepository
}

func NewTaggingService(taggingRepo repository.TaggingRepository) TaggingService {
	return &taggingService{taggingRepo}
}

func (s *taggingService) CreateTag(req dto.CreateTaggingRequest) (*dto.TaggingResponse, errs.MessageErr) {
	tag := entity.Tagging{
		Name: req.Name,
	}

	err := s.taggingRepo.Create(&tag)
	if err != nil {
		return nil, err
	}

	return &dto.TaggingResponse{
		ID:   tag.ID.String(),
		Name: tag.Name,
	}, nil
}

func (s *taggingService) GetAllTags() ([]dto.TaggingResponse, errs.MessageErr) {
	tags, err := s.taggingRepo.FindAll()
	if err != nil {
		return nil, err
	}

	var result []dto.TaggingResponse
	for _, tag := range tags {
		result = append(result, dto.TaggingResponse{
			ID:   tag.ID.String(),
			Name: tag.Name,
		})
	}

	return result, nil
}

func (s *taggingService) UpdateTag(id uuid.UUID, req dto.CreateTaggingRequest) (*dto.TaggingResponse, errs.MessageErr) {
	tag, err := s.taggingRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	tag.Name = req.Name
	err = s.taggingRepo.Update(tag)
	if err != nil {
		return nil, err
	}

	return &dto.TaggingResponse{
		ID:   tag.ID.String(),
		Name: tag.Name,
	}, nil
}

func (s *taggingService) DeleteTag(id uuid.UUID) errs.MessageErr {
	_, err := s.taggingRepo.FindByID(id)
	if err != nil {
		return err
	}

	err = s.taggingRepo.Delete(id)
	if err != nil {
		return errs.NewInternalServerError("failed to delete tag")
	}

	return nil
}
