package entity

type Music struct {
	ID       string `gorm:"type:text"`
	Mood     string `gorm:"type:text"`
	SongName string `gorm:"type:text"`
	URL      string `gorm:"type:text"`
}

func (Music) TableName() string {
	return "music_recommendation"
}
