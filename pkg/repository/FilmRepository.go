package repository

import (
	"errors"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	. "m/pkg/model"
)

func AddFilm(db *sqlx.DB, film *Film) error {

	// Начинаем транзакцию
	tx, err := db.Beginx()
	if err != nil {
		return errors.New("ошибка начала транзакции: " + err.Error())
	}

	err = tx.QueryRowx("INSERT INTO film (name, year, plot, genre, rating, image) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", film.Name, film.Year, film.Plot, film.Genre, film.Rating, film.Image).Scan(&film.Id)

	if err != nil {
		tx.Rollback() //если возникла ошибка в запросе откатываем все изменения
		return errors.New("ошибка добавления фильма: " + err.Error())
	}

	// Если всё прошло успешно — фиксируем транзакцию
	return tx.Commit()
}

func GetFilm(db *sqlx.DB, id int) (Film, error) {
	var result Film
	err := db.Get(&result, "SELECT * FROM film WHERE id=$1", id)
	if err != nil {
		return Film{}, errors.New("ошибка получения фильма: " + err.Error())
	}
	return result, nil
}

func UpdateFilm(db *sqlx.DB, film *Film) error {
	tx, err := db.Beginx()
	if err != nil {
		return errors.New("ошибка начала транзакции: " + err.Error())
	}
	_, err = tx.Exec(`UPDATE film SET name=$1, year=$2, plot=$3, genre=$4, rating=$5, image=$6 WHERE id=$7 RETURNING id`, film.Name, film.Year, film.Plot, film.Genre, film.Rating, film.Image, film.Id)
	if err != nil {
		return errors.New("ошибка обновления фильма: " + err.Error())
	}
	return tx.Commit()
}
func DeleteFilm(db *sqlx.DB, film Film) error {
	tx, err := db.Beginx()
	if err != nil {
		return errors.New("ошибка начала транзакции: " + err.Error())
	}
	_, err = db.Exec("DELETE FROM film WHERE id=$1", film.Id)
	if err != nil {
		return errors.New("Удаление фильма не удалось " + err.Error())
	}
	return tx.Commit()
}
func GetAllFilms(db *sqlx.DB) ([]Film, error) {
	var films []Film
	err := db.Select(&films, "SELECT * FROM film")
	return films, err
}
