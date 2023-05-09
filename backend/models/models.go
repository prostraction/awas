package models

type Number struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
	Type   string `json:"type"`
	TypeID int    `json:"type_id"`
}

type HistoryNumber struct {
	ID       int    `json:"id"`
	NumberID int    `json:"id_num"`
	Status   string `json:"status"`
	Comment  string `json:"comment"`
}

/*
type HistoryEntry struct {
	ID         int       `json:"id"`
	NumberID   int       `json:"num_id"`
	Status     string    `json:"status"`
	DateStatus time.Time `json:"date"`
}

type DbElement struct {
}
*/
