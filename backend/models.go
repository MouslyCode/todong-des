package main

import (
	"time"

	"gorm.io/gorm"
)

type StatusEnum string

const (
	StatusTodo     StatusEnum = "todo"
	StatusProgress StatusEnum = "progress"
	StatusDone     StatusEnum = "done"
)

type Todo struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string         `gorm:"type:varchar(100);not null" json:"name"`
	Status    StatusEnum     `gorm:"type:enum('todo','progress','done');default:'todo'" json:"status"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
