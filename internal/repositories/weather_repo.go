package repositories

import (
	"database/sql"
	"log"
	"wheather/internal/domain"
)

type WeatherRepo interface {
	Insert(dates domain.Weather) (int64, error)
	SelectNotConf() ([]domain.Weather, error)
	UpdateStatus(id int64, status int) error
}

type weatherRepo struct {
	db *sql.DB
}

func NewWeatherRepo(db *sql.DB) WeatherRepo {
	return &weatherRepo{db}
}

func (r *weatherRepo) UpdateStatus(id int64, status int) error {
	_, err := r.db.Exec("UPDATE weather SET status = ? WHERE id = ?", status, id)
	return err
}

func (r *weatherRepo) SelectNotConf() ([]domain.Weather, error) {
	query := "select id, name, lon, lat, status from weather where status = 0 or status is null"
	rows, err := r.db.Query(query)
	if err != nil {
		log.Fatalf("Error selecting weather: %v", err)
	}
	defer rows.Close()
	var wheat []domain.Weather
	for rows.Next() {
		var wheather domain.Weather
		if err := rows.Scan(&wheather.ID, &wheather.Name, &wheather.Lon, &wheather.Lat, &wheather.Status); err != nil {
			return nil, err
		}
		wheat = append(wheat, wheather)
	}
	return wheat, nil
}

func (r *weatherRepo) Insert(dates domain.Weather) (int64, error) {
	query := "INSERT INTO weather (name, lat, lon) VALUES (?, ?, ?)"
	id, err := r.db.Exec(query, dates.Name, dates.Lat, dates.Lon)
	if err != nil {
		log.Fatalf("Error inserting weather: %v", err)
	}
	return id.LastInsertId()
}
