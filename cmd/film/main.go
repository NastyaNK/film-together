package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"m/pkg/config"
	"m/pkg/model"
	"m/pkg/repository"
	"net/http"
	"os"
	"strconv"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Так 111")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
	fmt.Println("Так 3")
}

func versionHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var version string
		err := db.QueryRow("SELECT version();").Scan(&version)
		if err != nil {
			http.Error(w, "Error fetching version: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, version)
		fmt.Println("Так 4")
	}
}

func addFilmHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		var film model.Film
		if err := json.NewDecoder(r.Body).Decode(&film); err != nil {
			http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
			return
		}

		err := repository.AddFilm(db, &film)
		if err != nil {
			http.Error(w, "Failed to add film: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(film)
	}
}
func addRoomHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		var room model.Room
		if err := json.NewDecoder(r.Body).Decode(&room); err != nil {
			http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
			return
		}

		err := repository.AddRoom(db, &room)
		if err != nil {
			http.Error(w, "Failed to add room: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(room)
	}
}

func getAllFilmsHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		films, err := repository.GetAllFilms(db)
		if err != nil {
			http.Error(w, "Failed to get films: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(films)
	}
}

func getAllRoomsHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		rooms, err := repository.GetAllRooms(db)
		if err != nil {
			http.Error(w, "Failed to get rooms: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(rooms)
	}
}

func getFilmByIDHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			http.Error(w, "Missing film ID", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid film ID", http.StatusBadRequest)
			return
		}

		film, err := repository.GetFilm(db, id)
		if err != nil {
			http.Error(w, "Film not found: "+err.Error(), http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(film)
	}
}

func getRoomByIDHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			http.Error(w, "Missing room ID", http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid room ID", http.StatusBadRequest)
			return
		}
		room, err := repository.GetRoom(db, id)
		if err != nil {
			http.Error(w, "Room not found: "+err.Error(), http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(room)
	}
}

func updateFilmHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var film model.Film
		if r.Method != http.MethodPatch {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		if err := json.NewDecoder(r.Body).Decode(&film); err != nil {
			http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
			return
		}
		err := repository.UpdateFilm(db, &film)
		if err != nil {
			http.Error(w, "Failed to update film: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(film)
	}
}

func updateRoomHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var room model.Room
		if r.Method != http.MethodPatch {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		if err := json.NewDecoder(r.Body).Decode(&room); err != nil {
			http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
			return
		}
		err := repository.UpdateRoom(db, &room)
		if err != nil {
			http.Error(w, "Failed to update room: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(room)
	}
}

func deleteFilmHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			http.Error(w, "Missing film ID", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid film ID", http.StatusBadRequest)
			return
		}

		err = repository.DeleteFilm(db, model.Film{Id: id})
		if err != nil {
			http.Error(w, "Failed to delete film: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Film deleted successfully")
	}
}

func deleteRoomHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			http.Error(w, "Missing room ID", http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid room ID", http.StatusBadRequest)
			return
		}
		err = repository.DeleteRoom(db, model.Room{Id: id})
		if err != nil {
			http.Error(w, "Failed to delete room: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Room deleted successfully")
	}
}

func importFilmsHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		file, _, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Не удалось получить файл: "+err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()
		reader := csv.NewReader(file)
		records, err := reader.ReadAll()
		if err != nil {
			http.Error(w, "Ошибка чтения CSV: "+err.Error(), http.StatusInternalServerError)
			return
		}
		for i := 1; i < len(records); i++ {

			record := records[i]

			if len(record) < 7 {
				continue
			}

			year, err := strconv.Atoi(record[1])
			if err != nil {
				return
			}

			rating, err := strconv.ParseFloat(record[4], 64)
			if err != nil {
				return
			}

			film := model.Film{
				Name:     record[0],
				Year:     year,
				Plot:     record[2],
				Genre:    record[3],
				Rating:   rating,
				Image:    record[5],
				VideoURL: record[6],
			}

			err = repository.AddFilm(db, &film)
			if err != nil {
				return
			}
		}
	}
}

func main() {
	files, err := os.ReadDir(".") // текущая директория
	if err != nil {
		fmt.Println("Ошибка чтения директории:", err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			fmt.Println("[DIR] ", file.Name())
		} else {
			fmt.Println("      ", file.Name())
		}
	}

	cfg, err := config.LoadConfig(os.Getenv("CONFIG_PATH"))
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	db, err := config.ConnectToDatabase(cfg.Database)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/version", versionHandler(db))
	http.HandleFunc("/admin/import", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "config/templates/import.html")
	})
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("config/static"))))
	http.HandleFunc("/admin/import-films", importFilmsHandler(db))

	http.HandleFunc("/admin/film", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			if id := r.URL.Query().Get("id"); id != "" {
				getFilmByIDHandler(db)(w, r)
			} else {
				getAllFilmsHandler(db)(w, r)
			}
		case http.MethodPost:
			addFilmHandler(db)(w, r)
		case http.MethodPatch:
			updateFilmHandler(db)(w, r)
		case http.MethodDelete:
			deleteFilmHandler(db)(w, r)

		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/room", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			if id := r.URL.Query().Get("id"); id != "" {
				getRoomByIDHandler(db)(w, r)
			} else {
				getAllRoomsHandler(db)(w, r)
			}
		case http.MethodPost:
			addRoomHandler(db)(w, r)
		case http.MethodPatch:
			updateRoomHandler(db)(w, r)
		case http.MethodDelete:
			deleteRoomHandler(db)(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Так")
	address := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	fmt.Printf("Listening on http://%s\n", address)
	if err := http.ListenAndServe(address, nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
	fmt.Println("Так 2")
	//data, err := io.ReadAll()
}
