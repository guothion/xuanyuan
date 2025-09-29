package dto

import "time"

type ID struct {
	Id string `json:"id"`
}

type Timestamp struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
