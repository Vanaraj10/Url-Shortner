package models

import "time"

type URL struct {
	ID        int64     `db:"id"`
	UserID    int64     `db:"user_id"`
	Original  string    `db:"original"`
	Short     string    `db:"short"`
	CreatedAt time.Time `db:"created_at"`
}
