package repository

import (
	"backend-bangkit/entity"

	"gorm.io/gorm"
)

// Repository
type MusicRepository interface {
	GetRandomMusic(limit int, mood string) ([]entity.Music, error)
}

type musicRepositoryImpl struct {
	db *gorm.DB
}

func NewMusicRepository(db *gorm.DB) MusicRepository {
	return &musicRepositoryImpl{db: db}
}

func (r *musicRepositoryImpl) GetRandomMusic(limit int, mood string) ([]entity.Music, error) {
	var music []entity.Music
	query := r.db.Order("RANDOM()").Limit(limit)
	if mood != "" {
		query = query.Where("mood = ?", mood)
	}
	if err := query.Find(&music).Error; err != nil {
		return nil, err
	}
	return music, nil
}
