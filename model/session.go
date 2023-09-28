package model

import "time"

type Session struct {
	SID       string    `json:"s_id"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	Expired   time.Time `json:"expired"`
}
