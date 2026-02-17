package data

import "time"

type Reservation struct {
	ID        int64     `json:"id"`
	Customer  string    `json:"customer"`
	TableID   int       `json:"table_id"`
	TimeSlot  time.Time `json:"time_slot"`
	PartySize int       `json:"party_size"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
