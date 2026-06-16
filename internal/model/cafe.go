package model

// Cafe represent a cafe in the system
type Cafe struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Address    string  `json:"address"`
	Latitude   float64 `json:"latitude"`
	Longtitude float64 `json:"longtitude"`
}
