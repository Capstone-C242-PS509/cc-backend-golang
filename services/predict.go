package service

import (
	"backend-bangkit/dto"
	"backend-bangkit/entity"
	"backend-bangkit/pkg/errs"
	"backend-bangkit/repository"
	"time"

	"github.com/google/uuid"
)

type PredictionService interface {
	SavePrediction(req *dto.SavePredictionInput) (*dto.SavePredictionResponse, errs.MessageErr)
	GetPredictionSummary(userID uuid.UUID, startDate, endDate time.Time) ([]dto.PredictionSummary, errs.MessageErr) // Gunakan dto.PredictionSummary
}

type predictionService struct {
	predictionRepo repository.PredictionRepository
}

func NewPredictionService(predictionRepo repository.PredictionRepository) PredictionService {
	return &predictionService{predictionRepo}
}

func (s *predictionService) SavePrediction(req *dto.SavePredictionInput) (*dto.SavePredictionResponse, errs.MessageErr) {
	prediction := entity.PredictionResult{
		ID:            uuid.New(),
		UserID:        req.UserID,
		MentalDisease: req.MentalDisease,
		PredictedAt:   time.Now(),
	}

	err := s.predictionRepo.CreatePrediction(&prediction)
	if err != nil {
		return nil, err
	}

	response := &dto.SavePredictionResponse{
		ID:            prediction.ID,
		UserID:        prediction.UserID,
		MentalDisease: prediction.MentalDisease,
		PredictedAt:   prediction.PredictedAt,
	}

	return response, nil
}

func (s *predictionService) GetPredictionSummary(userID uuid.UUID, startDate, endDate time.Time) ([]dto.PredictionSummary, errs.MessageErr) {
	results, err := s.predictionRepo.GetPredictionSummary(userID, startDate, endDate)
	if err != nil {
		return nil, err
	}
	return results, nil
}
