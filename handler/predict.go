package handler

import (
	"backend-bangkit/dto"
	"backend-bangkit/entity"
	"backend-bangkit/pkg/common"
	"backend-bangkit/pkg/errs"

	service "backend-bangkit/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type PredictionHandler struct {
	predictionService service.PredictionService
}

func NewPredictionHandler(predictionService service.PredictionService) *PredictionHandler {
	return &PredictionHandler{predictionService}
}

func (h *PredictionHandler) SavePrediction(c *gin.Context) {

	userData, ok := c.MustGet("userData").(*entity.User)

	if !ok {
		newError := errs.NewBadRequest("Failed to get user data")
		c.JSON(newError.StatusCode(), newError)
		return
	}
	var input dto.SavePredictionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, errs.NewBadRequest("invalid request"))
		return
	}

	input.UserID = userData.ID

	// Simpan hasil prediksi
	prediction, err := h.predictionService.SavePrediction(&input)
	if err != nil {
		c.JSON(err.StatusCode(), err)
		return
	}

	c.JSON(common.BuildResponse(http.StatusOK, prediction))
}

func (h *PredictionHandler) GetPredictionSummary(c *gin.Context) {

	userData, ok := c.MustGet("userData").(*entity.User)

	if !ok {
		newError := errs.NewBadRequest("Failed to get user data")
		c.JSON(newError.StatusCode(), newError)
		return
	}

	var (
		startDate time.Time
		endDate   time.Time
	)

	// Ambil query parameter (contoh: `?start_date=2024-01-01&end_date=2024-12-31`)
	startDateParam := c.Query("start_date")
	endDateParam := c.Query("end_date")

	if startDateParam != "" && endDateParam != "" {
		var err error
		startDate, err = time.Parse("2006-01-02", startDateParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, errs.NewBadRequest("invalid start_date format"))
			return
		}

		endDate, err = time.Parse("2006-01-02", endDateParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, errs.NewBadRequest("invalid end_date format"))
			return
		}
	} else {
		// Default: bulan ini
		now := time.Now()
		startDate = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
		endDate = time.Date(now.Year(), now.Month()+1, 1, 0, 0, 0, 0, now.Location()).Add(-time.Second)
	}

	// Hitung jumlah prediksi berdasarkan mental disease
	results, err := h.predictionService.GetPredictionSummary(userData.ID, startDate, endDate)
	if err != nil {
		c.JSON(err.StatusCode(), err)
		return
	}

	c.JSON(common.BuildResponse(http.StatusOK, results))
}
