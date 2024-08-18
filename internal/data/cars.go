package data

import (
	"time"

	"github.com/google/uuid"
)

type Car struct {
	ID         uuid.UUID
	Online     bool
	CarType    string
	Brand      string
	Model      string
	Year       int32
	Kilometers int64
	CarDomain  string
	Price      float64
	InfoPrice  float64
	Currency   string
	ChasisCode string
	MotorCode  string
	CreatedAt  time.Time
}
