package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v3"
	"net/http"
	"os"
)

type Config struct {
	Server   ServerConfig `yaml:"server"`
	Database Dd           `yaml:"database"`
}

type ServerConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Dd struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}

func versionHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var version string
		err := db.QueryRow("SELECT version();").Scan(&version)
		if err != nil {
			http.Error(w, "Error fetching version: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, version)
	}
}

func loadConfig(filename string) (Config, error) {
	var config Config
	data, err := os.ReadFile(filename)
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(data, &config)
	return config, err
}

func connectToDatabase(config Dd) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Name)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	config, err := loadConfig("resources/config.yaml")
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	fmt.Printf("Host: %s\n", config.Server.Host)

	db, err := connectToDatabase(config.Database)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/version", versionHandler(db))

	address := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
	if err := http.ListenAndServe(address, nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
