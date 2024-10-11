package domain

import "database/sql"

type Weather struct {
	ID     int64         `json:"id"`
	Name   string        `json:"name"`
	Lat    float64       `json:"lat"`
	Lon    float64       `json:"lon"` // Use time.Time for a proper timestamp
	Status sql.NullInt64 `json:"status"`
}
