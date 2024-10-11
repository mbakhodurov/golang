package dto

type WeatherDTO struct {
	Name   string  `json:"name"`
	Lat    float64 `json:"lat"`
	Lon    float64 `json:"lon"`
	Status int     `json:"status"`
}
