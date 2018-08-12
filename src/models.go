package smith

import "time"

type Container struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Image     string    `json:"image"`
	Ip        string    `json:"ip"`
	CreatedAt time.Time `json:"created_at"`
	Ports     string    `json:"ports"`
	Status    bool      `json:"status"`
}

type Image struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type Deploy struct {
	Tag       string    `json:"tag"`
	CreatedAt time.Time `json:"created_at"`
}
