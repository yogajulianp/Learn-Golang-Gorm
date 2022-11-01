package models

import(
	"gorm.io/gorm"
	"time"
)

type Locker struct{
	ID 			int 		`json:"id" form:"id" gorm:"primaryKey" `
	Code 		string  	`json:"code" form:"code"  gorm:"unique, not null"`
	UserID 		int 		`json:"user_id" form:"user_id" `
	User 		UserResponse		`json:"user"`
	CreatedAt time.Time		`json:"created_at"`
  	UpdatedAt time.Time		`json:"updated_at"`
 	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type LockerResponse struct {
	ID 			int 		`json:"id" form:"id" `
	Code 		string  	`json:"code" form:"code" `
	UserID 		int 		`json:"-" `
}

func (LockerResponse) TableName() string {
	return "lockers"	
}