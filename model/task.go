package model

import "time"

type Task struct {
	ID uint `json:"id" gorm:"primary_key"`
	Title string `json:"title" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User User `json:"user" gorm:"foreignKey:UserId"`
	UserId uint `json:"user_id" gorm:"not null"`
}

// クライアントからGetメソッドでリクエストがあった時、クライアント側に返すデータの型定義	
type TaskResponse struct {
	ID uint `json:"id"`
	Title string `json:"title" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`	
	UpdatedAt time.Time `json:"updated_at"`
}