package data

import (
	"time"
)

type VintageRetro struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Name      string    `json:"name"`
	Year      int32     `json:"year,omitempty"`
	Cost      Cost      `json:"cost,omitempty"`
	Type      string    `json:"type,omitempty"`
	Version   int32     `json:"version"`
}
