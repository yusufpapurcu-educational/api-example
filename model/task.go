package model

import "time"

type Task struct {
	ID          int
	Description string
	Done        bool
	Percentage  float32
	Date        time.Time
}
