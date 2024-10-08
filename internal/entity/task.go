package entity

import "time"

type Task struct {
	Id           int `bun:",pk,autoincrement"`
	Name         string
	Description  string
	CreationDate time.Time
	Status       bool
}
