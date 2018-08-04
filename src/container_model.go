package smith

import "time"

type Container struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Ip        string    `json:"ip"`
	CreatedAt time.Time `json:"created_at"`
	Ports     string    `json:"ports"`
	Status    bool      `json:"status"`
}
