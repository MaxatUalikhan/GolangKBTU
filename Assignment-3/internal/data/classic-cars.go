package data

import (
	"assignment3.ualikhan.net/internal/validator"
	"database/sql"
	"errors"
	"time"
)

type ClassicCars struct {
	ID          int64     `json:"id"`
	CreatedAt   time.Time `json:"-"`
	Name        string    `json:"name"`
	Year        int32     `json:"year,omitempty"`
	Cost        Cost      `json:"cost,omitempty"`
	Description string    `json:"description,omitempty"`
	Version     int32     `json:"version"`
}

func ValidateClassicCars(v *validator.Validator, classiccars *ClassicCars) {
	v.Check(classiccars.Name != "", "name", "must be provided")
	v.Check(len(classiccars.Name) <= 500, "name", "must not be more than 500 bytes long")
	v.Check(classiccars.Year != 0, "year", "must be provided")
	v.Check(classiccars.Year <= int32(time.Now().Year()), "year", "must not be in the future")
	v.Check(classiccars.Cost != 0, "cost", "must be provided")
	v.Check(classiccars.Cost > 0, "cost", "must be a positive integer")
}

type ClassicCarsModel struct {
	DB *sql.DB
}

func (m ClassicCarsModel) Insert(classiccars *ClassicCars) error {
	query := `
		INSERT INTO classic_cars (name, year, cost, description)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, version`

	args := []interface{}{classiccars.Name, classiccars.Year, classiccars.Cost, classiccars.Description}

	return m.DB.QueryRow(query, args...).Scan(&classiccars.ID, &classiccars.CreatedAt, &classiccars.Version)

}

func (m ClassicCarsModel) Get(id int64) (*ClassicCars, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
		SELECT id, created_at, name, year, cost, description, version
		FROM classic_cars
		WHERE id = $1`

	var classiccars ClassicCars

	err := m.DB.QueryRow(query, id).Scan(
		&classiccars.ID,
		&classiccars.CreatedAt,
		&classiccars.Name,
		&classiccars.Year,
		&classiccars.Cost,
		&classiccars.Description,
		&classiccars.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &classiccars, nil
}

func (m ClassicCarsModel) Update(classiccars *ClassicCars) error {
	query := `
		UPDATE classic_cars
		SET name = $1, year = $2, cost = $3, description = $4, version = version + 1
		WHERE id = $5
		RETURNING version`

	args := []interface{}{
		classiccars.Name,
		classiccars.Year,
		classiccars.Cost,
		classiccars.Description,
		classiccars.ID,
	}

	return m.DB.QueryRow(query, args...).Scan(&classiccars.Version)

}

func (m ClassicCarsModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
		DELETE FROM classic_cars
		WHERE id = $1`

	result, err := m.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}
