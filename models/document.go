package models

import "time"

type Document struct {
	ID         int       `json:"id"`
	DatabaseID int       `json:"database_id"`
	Content    string    `json:"contect"`
	CreateAt   time.Time `json:"create_at"`
}
