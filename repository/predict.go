package repository

import (
	"backend-bangkit/dto"
	"backend-bangkit/entity"
	"backend-bangkit/pkg/errs"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PredictionRepository interface {
	CreatePrediction(prediction *entity.PredictionResult) errs.MessageErr
	GetPredictionSummary(userID uuid.UUID, startDate, endDate time.Time) ([]dto.PredictionSummary, errs.MessageErr)
}

type predictionRepository struct {
	db *gorm.DB
}

func NewPredictionRepository(db *gorm.DB) PredictionRepository {
	return &predictionRepository{db: db}
}

func (r *predictionRepository) CreatePrediction(prediction *entity.PredictionResult) errs.MessageErr {
	err := r.db.Create(prediction).Error

	if err != nil {
		return errs.NewInternalServerError("cannot save predict")
	}

	return nil
}

func (r *predictionRepository) GetPredictionSummary(userID uuid.UUID, startDate, endDate time.Time) ([]dto.PredictionSummary, errs.MessageErr) {
	var result []dto.PredictionSummary
	err := r.db.Model(&entity.PredictionResult{}).
		Select("mental_disease, COUNT(*) as count").
		Where("user_id = ? AND predicted_at BETWEEN ? AND ?", userID, startDate, endDate).
		Group("mental_disease").
		Scan(&result).Error

	if err != nil {
		return nil, errs.NewInternalServerError("cannot get user predict summary")
	}
	return result, nil
}
