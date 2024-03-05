package model

import "time"

type User struct {
	ID uint `json:"id" gorm:"primary_key"`
	Email string `json:"email" gorm:"unique`
	Password string `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// 新しいユーザーの情報をクライアントにレスポンスで返すときのデータの型定義
type UserReponse struct {
	ID unit `json:"id" gorm:"primary_key"`
	Email string `json: "email" gorm:"unique"`
}
