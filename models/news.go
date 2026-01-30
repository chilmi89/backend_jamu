package models

import "time"

type News struct {
	ID        int       `bun:"id,pk,autoincrement" json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"created_at"`
}
