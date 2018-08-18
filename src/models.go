package smith

import "time"

type Container struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	Ports     string    `json:"ports"`
}

type Image struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type Deploy struct {
	Tag       string    `json:"tag"`
	Count     int       `json:"count"`
	CreatedAt time.Time `json:"created_at"`
	Status    bool      `json:"status"`
}
