package model

import "time"

type Article struct {
	ID      int64      `json:"id"`
	Title   string     `json:"title,omitempty"`
	Content string     `json:"content,omitempty"`
	Created *time.Time `json:"created,omitempty"`
}
