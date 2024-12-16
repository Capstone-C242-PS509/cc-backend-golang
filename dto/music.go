package dto

// DTO (Data Transfer Object)
type MusicResponse struct {
	ID       string `json:"id"`
	Mood     string `json:"mood"`
	SongName string `json:"song_name"`
	URL      string `json:"url"`
}
