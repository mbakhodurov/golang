package main

import (
	"os"
	"wheather/db"
	"wheather/internal/api"
	"wheather/internal/repositories"
	"wheather/internal/services"
)

func main() {
	dbConn := db.InitDb()
	defer dbConn.Close()
	weatherRepo := repositories.NewWeatherRepo(dbConn)
	weatherService := services.NewWeatherService(weatherRepo)
	telegramService := services.NewTelegramService(os.Getenv("bot_token"), os.Getenv("chat_id"))
	wheatherApi := api.NewWeatherApi(weatherService, telegramService, os.Getenv("api_key"), os.Getenv("url"))
	wheatherApi.Save()

	// fmt.Println(res)
}
