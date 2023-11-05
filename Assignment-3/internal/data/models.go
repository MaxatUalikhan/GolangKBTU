package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	ClassicCars ClassicCarsModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		ClassicCars: ClassicCarsModel{DB: db},
	}
}
