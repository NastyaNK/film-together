package model

type Room struct {
	Id     int    `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	Public bool   `json:"public" db:"public"`
	FilmID int    `json:"film_id" db:"film_id"`
}
