package models

import "time"

type JSONTime time.Time

// models untuk tabel tb_classes
type Classes struct {
	Id        int      `json:"id"`
	Name      string   `json:"name"`
	Status    int      `json:"status"`
	CreatedAt JSONTime `json:"created_at"`
	UpdatedAt JSONTime `json:"updated_at"`
}
