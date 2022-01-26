package model

import "time"

type Task struct {
	ID          int
	Description string `validate:"required"`
	Done        bool
	Percentage  float32
	Date        time.Time `validate:"required"`
}
