package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"m/pkg/model"
)

func AddRoom(db *sqlx.DB, room *model.Room) error {
	query := `INSERT INTO room (name, public, film_id) VALUES ($1, $2, $3) RETURNING id`
	return db.QueryRow(query, room.Name, room.Public, room.FilmID).Scan(&room.Id)
}

func GetRoom(db *sqlx.DB, id int) (model.Room, error) {
	var room model.Room
	err := db.Get(&room, "SELECT * FROM room WHERE id = $1", id)
	if err != nil {
		return model.Room{}, fmt.Errorf("ошибка получения комнаты: %w", err)
	}
	return room, nil
}
func GetAllRooms(db *sqlx.DB) ([]model.Room, error) {
	var rooms []model.Room
	err := db.Select(&rooms, "SELECT * FROM room")
	if err != nil {
		return nil, fmt.Errorf("ошибка получения списка комнат: %w", err)
	}
	return rooms, nil
}
func UpdateRoom(db *sqlx.DB, room *model.Room) error {
	query := `UPDATE room SET name = $1, public = $2, film_id = $3 WHERE id = $4`
	_, err := db.Exec(query, room.Name, room.Public, room.FilmID, room.Id)
	if err != nil {
		return fmt.Errorf("ошибка обновления комнаты: %w", err)
	}
	return nil
}
func DeleteRoom(db *sqlx.DB, room model.Room) error {
	query := `DELETE FROM room WHERE id = $1`
	_, err := db.Exec(query, room.Id)
	if err != nil {
		return fmt.Errorf("удаление комнаты не удалось: %w", err)
	}
	return nil
}
