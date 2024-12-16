package repository

import (
	"backend-bangkit/entity"
	"backend-bangkit/pkg/errs"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TaggingRepository interface {
	Create(tag *entity.Tagging) errs.MessageErr
	FindByID(id uuid.UUID) (*entity.Tagging, errs.MessageErr)
	FindByName(name string) (*entity.Tagging, errs.MessageErr)
	FindAll() ([]entity.Tagging, errs.MessageErr)
	Update(tag *entity.Tagging) errs.MessageErr
	Delete(id uuid.UUID) errs.MessageErr
}

type taggingRepository struct {
	db *gorm.DB
}

// FindByName implements TaggingRepository.
func (r *taggingRepository) FindByName(name string) (*entity.Tagging, errs.MessageErr) {
	var tag entity.Tagging
	err := r.db.First(&tag, "name = ?", name).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.NewNotFound("tag not found")
		}
		return nil, errs.NewInternalServerError("cannot find tag by ID")
	}
	return &tag, nil
}

func NewTaggingRepository(db *gorm.DB) TaggingRepository {
	return &taggingRepository{db}
}

func (r *taggingRepository) Create(tag *entity.Tagging) errs.MessageErr {
	err := r.db.Create(tag).Error
	if err != nil {
		return errs.NewInternalServerError("cannot create tag")
	}
	return nil
}

func (r *taggingRepository) FindByID(id uuid.UUID) (*entity.Tagging, errs.MessageErr) {
	var tag entity.Tagging
	err := r.db.First(&tag, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.NewNotFound("tag not found")
		}
		return nil, errs.NewInternalServerError("cannot find tag by ID")
	}
	return &tag, nil
}

func (r *taggingRepository) FindAll() ([]entity.Tagging, errs.MessageErr) {
	var tags []entity.Tagging
	err := r.db.Find(&tags).Error
	if err != nil {
		return nil, errs.NewInternalServerError("cannot find tags")
	}
	return tags, nil
}

func (r *taggingRepository) Update(tag *entity.Tagging) errs.MessageErr {
	err := r.db.Save(tag).Error
	if err != nil {
		return errs.NewInternalServerError("cannot update tag")
	}
	return nil
}

func (r *taggingRepository) Delete(id uuid.UUID) errs.MessageErr {
	err := r.db.Delete(&entity.Tagging{}, "id = ?", id).Error
	if err != nil {
		return errs.NewInternalServerError("cannot delete tag")
	}
	return nil
}
