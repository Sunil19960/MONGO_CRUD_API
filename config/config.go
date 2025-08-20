package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Settings struct {
	PortNo          string `json:"port_no"`
	GetStatus       string `json:"get_status"`
	GetBookRoute    string `json:"get_book"`
	AddBookRoute    string `json:"add_book"`
	UpdateBookRoute string `json:"update_book"`
	RemoveBookRoute string `json:"delete_book"`
	Mongo_URI       string `json:"Mongo_URI"`
}

func LoadSettings() Settings {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf(".env file is missing!")
	}

	return Settings{
		PortNo:          getEnv("PORT_NO"),
		Mongo_URI:       getEnv("MONGO_DB_URI"),
		GetStatus:       getEnv("GET_STATUS_ROUTE"),
		GetBookRoute:    getEnv("GET_BOOK_ROUTE"),
		AddBookRoute:    getEnv("ADD_BOOK_ROUTE"),
		UpdateBookRoute: getEnv("UPDATE_BOOK_ROUTE"),
		RemoveBookRoute: getEnv("REMOVE_BOOK_ROUTE"),
	}
}

func getEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	log.Fatalf("%s not found in .env file", key)
	return ""
}
