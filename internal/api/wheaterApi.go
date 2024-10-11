package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
	"wheather/internal/domain"
	"wheather/internal/services"
)

type weatherApi struct {
	services        services.WeatherService
	telegramService services.TelegramService
	apiKey          string
	url             string
}

type WeatherApi interface {
	Save() (int64, error)
}

type Output struct {
	Lat  float64 `json:"lat"`
	Lon  float64 `json:"lon"`
	Name string  `json:"name"`
}

func NewWeatherApi(service services.WeatherService, telegramService services.TelegramService, apikey, url string) WeatherApi {
	return &weatherApi{services: service, telegramService: telegramService, apiKey: apikey, url: url}
}

func FormatText(domains domain.Weather) string {
	return fmt.Sprintf("Weather Update:\nName: %s\nLon: %.6f\nLat: %.6f",
		domains.Name, domains.Lon, domains.Lat)
}

func (api *weatherApi) Save() (int64, error) {
	for {
		fmt.Println("Selecting..........")
		weathers, err := api.services.SelectNotConf()
		if err != nil {
			time.Sleep(10 * time.Minute)
			continue
		}
		for _, v := range weathers {
			fmt.Println(v.Status, v.Lat, v.Name, v.Lon)
			fmt.Println("______________", time.Now())
			fmt.Println("______________", v.ID, "UPDATING ______________")
			err = api.services.UpdateStatus(v.ID, 10)
			if err != nil {
				log.Fatalf("Error updating weather: %v", err)
			}
			// time.Sleep(10 * time.Second)
		}
		time.Sleep(10 * time.Second)
	}
	// for {
	// 	weather, err := api.Sending()
	// 	if err != nil {
	// 		fmt.Println("Ошибка при получении данных о погоде:", err)
	// 		time.Sleep(5 * time.Second) // Ждем перед повторной попыткой
	// 		continue
	// 	}
	// 	for _, v := range weather {
	// 		weathers := domain.Weather{
	// 			Name: v.Name,
	// 			Lat:  v.Lat,
	// 			Lon:  v.Lon,
	// 		}
	// 		id, err := api.services.Insert(weathers)
	// 		if err != nil {
	// 			fmt.Println("Ошибка при вставке данных в базу:", err)
	// 			continue // Переход к следующему городу
	// 		}
	// 		fmt.Println("Добавлено с ID:", id)
	// 		str := FormatText(weathers)
	// 		api.telegramService.SendMessage(str)

	// 	}

	// 	time.Sleep(5 * time.Second)
	// }
}

func (api *weatherApi) Sending() ([]domain.Weather, error) {
	resp, err := http.Get(api.url + "&appid=" + api.apiKey)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var out []Output
	err = json.Unmarshal(data, &out)
	if err != nil {
		log.Print(err)
	}

	var wheathers []domain.Weather
	for _, v := range out {
		wheathers = append(wheathers, domain.Weather{
			Name: v.Name,
			Lat:  v.Lat,
			Lon:  v.Lon,
		})
	}
	return wheathers, nil
}
