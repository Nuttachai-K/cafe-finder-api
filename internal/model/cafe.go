package model

import "time"

// Cafe represent a cafe in the system
type Cafe struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longtitude"`
	CreatedAt time.Time `json:"create_at"`
}
