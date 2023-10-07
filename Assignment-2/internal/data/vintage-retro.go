package data

import (
	"assignment2.ualikhan.net/internal/validator"
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

func ValidateVintageRetro(v *validator.Validator, vintageretro *VintageRetro) {
	v.Check(vintageretro.Name != "", "name", "must be provided")
	v.Check(len(vintageretro.Name) <= 500, "name", "must not be more than 500 bytes long")
	v.Check(vintageretro.Year != 0, "year", "must be provided")
	v.Check(vintageretro.Year <= int32(time.Now().Year()), "year", "must not be in the future")
	v.Check(vintageretro.Cost != 0, "cost", "must be provided")
	v.Check(vintageretro.Cost > 0, "cost", "must be a positive integer")
}
