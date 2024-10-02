package entity

import "time"

type Task struct {
	Id            int
	Name          string
	Description   string
	Creation_date time.Time
	Status        bool
}
