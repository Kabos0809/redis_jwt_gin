package Models

import "time"

type User struct {
	ID uint64	`json:"id" gorm:"not null; AUTO_INCREMENT;"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime;"`
	Pass string `from: "pass" json:"pass" gorm:"not null;" binding:"required"`
	UName string `from: "uname" json:"uname" gorm:"not null; unique;" binding:"required"`
	Age int `json:"age" gorm:"not null;" binding:"required"`
}

func (b *User) TableName() string {
	return "user"
}