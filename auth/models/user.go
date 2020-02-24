package models

//User model
type User struct {
	Username  string `json:"username" gorm:"type:varchar(100);unique_index"`
	Password  string `json:"password" gorm:"type:varchar(255)"`
	CreatedAt int64  `json:"created_at" gorm:"type:numeric"`
}
