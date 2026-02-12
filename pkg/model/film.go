package model

type Film struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Year     int     `json:"year"`
	Plot     string  `json:"plot"`
	Genre    string  `json:"genre"`
	Rating   float64 `json:"rating"`
	Image    string  `json:"image"`
	VideoURL string  `json:"video_url"`
}
