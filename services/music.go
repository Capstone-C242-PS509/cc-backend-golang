package service

import (
	"backend-bangkit/dto"
	"backend-bangkit/repository"
)

// Service
type MusicService interface {
	GetRandomMusic(limit int, mood string) ([]dto.MusicResponse, error)
}

type musicServiceImpl struct {
	repository repository.MusicRepository
}

func NewMusicService(repository repository.MusicRepository) MusicService {
	return &musicServiceImpl{repository: repository}
}

func (s *musicServiceImpl) GetRandomMusic(limit int, mood string) ([]dto.MusicResponse, error) {
	music, err := s.repository.GetRandomMusic(limit, mood)
	if err != nil {
		return nil, err
	}

	var response []dto.MusicResponse
	for _, m := range music {
		response = append(response, dto.MusicResponse{
			ID:       m.ID,
			Mood:     m.Mood,
			SongName: m.SongName,
			URL:      m.URL,
		})
	}
	return response, nil
}
