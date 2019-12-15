package model

import (
	"time"
)

type ProjectPost struct {
	ID        uint               `json:"id" gorm:"primary_key"`
	CreatedAt time.Time          `json:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt"`
	DeletedAt *time.Time         `json:"deletedAt"`
	Title     string             `json:"title"`
	Body      string             `json:"body"`
	Subtitle  string             `json:"subtitle"`
	Section   string             `json:"section"`
	Images    []ProjectPostImage `json:"images"`
}

type ProjectPostImage struct {
	ID            uint       `json:"id" gorm:"primary_key"`
	CreatedAt     time.Time  `json:"createdAt"`
	UpdatedAt     time.Time  `json:"updatedAt"`
	DeletedAt     *time.Time `json:"deletedAt"`
	Url           string     `json:"url"`
	IsDefault     bool       `json:"isDefault"`
	ProjectPostId uint       `json:"projectPostId"`
}
