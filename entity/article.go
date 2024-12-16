package entity

import (
	"time"

	"github.com/google/uuid"
)

type Article struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Title     string    `gorm:"type:varchar(255);not null"`
	Content   string    `gorm:"type:text;not null"`
	Creator   string    `gorm:"type:text;not null"`
	TagID     uuid.UUID `gorm:"type:uuid;not null"`
	UrlImage  string    `gorm:"type:text;not null"`
	Tag       Tagging   `gorm:"foreignKey:TagID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Tagging struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name      string    `gorm:"type:varchar(255);unique;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
