package entity

import "time"

type Task struct {
	Id           int
	Name         string
	Description  string
	CreationDate time.Time
	Status       bool
}
