package dto

import (
	"time"

	"github.com/google/uuid"
)

type SavePredictionInput struct {
	UserID        uuid.UUID
	MentalDisease string `json:"mental_disease" binding:"required"`
}

type SavePredictionResponse struct {
	ID            uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID        uuid.UUID `gorm:"type:uuid;not null"`         // ID user yang melakukan prediksi
	MentalDisease string    `gorm:"type:varchar(255);not null"` // Nama penyakit mental yang diprediksi
	PredictedAt   time.Time `gorm:"not null"`                   // Waktu prediksi
}

type PredictionSummary struct {
	MentalDisease string `json:"mental_disease"`
	Count         int    `json:"count"`
}
