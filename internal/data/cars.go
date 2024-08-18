package data

import (
	"time"

	"github.com/google/uuid"
)

type Car struct {
	ID         uuid.UUID `json:"id"`
	Online     bool      `json:"online"`
	CarType    string    `json:"car_type"`
	Brand      string    `json:"brand"`
	Model      string    `json:"model"`
	Year       int32     `json:"year"`
	Kilometers int64     `json:"kilometers"`
	CarDomain  string    `json:"car_domain"`
	Price      float64   `json:"price"`
	InfoPrice  float64   `json:"info_price"`
	Currency   string    `json:"currency"`
	ChasisCode string    `json:"chasis_code"`
	MotorCode  string    `json:"motor_code"`
	CreatedAt  time.Time `json:"created_at"`
}
