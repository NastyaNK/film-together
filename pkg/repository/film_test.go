package repository_test

import (
	"m/pkg/model"
	"m/pkg/repository"
	"m/pkg/utils/testutils"
	"testing"

	"github.com/jmoiron/sqlx"
)

func createFilm(t *testing.T, db *sqlx.DB) *model.Film {
	t.Helper()
	film := &model.Film{
		Name:   "The Matrix",
		Year:   1999,
		Plot:   "Reality is a simulation",
		Genre:  "Sci‑Fi",
		Rating: 8.7,
		Image:  "matrix.jpg",
	}
	if err := repository.AddFilm(db, film); err != nil {
		t.Fatalf("AddFilm: %v", err)
	}
	if film.Id == 0 {
		t.Fatal("AddFilm: id not set")
	}
	return film
}

func TestAddFilm(t *testing.T) {
	db := testutils.SetupTestDB(t)
	_ = createFilm(t, db)
}

func TestGetFilm(t *testing.T) {
	db := testutils.SetupTestDB(t)
	film := createFilm(t, db)

	got, err := repository.GetFilm(db, film.Id)
	if err != nil {
		t.Fatalf("GetFilm: %v", err)
	}
	if got.Name != film.Name {
		t.Errorf("GetFilm name = %q, want %q", got.Name, film.Name)
	}
}

func TestUpdateFilm(t *testing.T) {
	db := testutils.SetupTestDB(t)
	film := createFilm(t, db)

	film.Rating = 9.0
	if err := repository.UpdateFilm(db, film); err != nil {
		t.Fatalf("UpdateFilm: %v", err)
	}
	got, _ := repository.GetFilm(db, film.Id)
	if got.Rating != 9.0 {
		t.Errorf("UpdateFilm rating = %v, want 9.0", got.Rating)
	}
}

func TestDeleteFilm(t *testing.T) {
	db := testutils.SetupTestDB(t)
	film := createFilm(t, db)

	if err := repository.DeleteFilm(db, *film); err != nil {
		t.Fatalf("DeleteFilm: %v", err)
	}
	if _, err := repository.GetFilm(db, film.Id); err == nil {
		t.Fatal("expected error after delete, got nil")
	}
}

func TestGetAllFilms(t *testing.T) {
	db := testutils.SetupTestDB(t)
	_ = createFilm(t, db)
	_ = createFilm(t, db)

	got, err := repository.GetAllFilms(db)
	if err != nil {
		t.Fatalf("GetAllFilms: %v", err)
	}
	if len(got) < 2 {
		t.Errorf("ждем больше чем: %d", len(got))
	}
}
