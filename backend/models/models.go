package models

import "time"

type Number struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type History struct {
	ID           int            `json:"id"`
	NumberID     int            `json:"num_id"`
	HistoryEntry []HistoryEntry `json:"-"`
}
type HistoryEntry struct {
	ID         int       `json:"id"`
	NumberID   int       `json:"num_id"`
	Status     string    `json:"status"`
	DateStatus time.Time `json:"date"`
}

type DbElement struct {
}
