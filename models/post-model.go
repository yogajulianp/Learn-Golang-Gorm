package models

import(
	//"gorm.io/gorm"
	//"time"
)

type Post struct{
	ID 			int 		`json:"id" form:"id" gorm:"primaryKey" `
	Title 		string  	`json:"title" form:"title" validate:"required"`
	Description	 string  	`form:"description" json:"description" validate:"required" `
	UserID 		int 		`json:"user_id" form:"user_id" `
	User		 UserResponse 	`json:"user"`
	Tag			[]Tag		`json:"tags" gorm:"many2many:post_tags"`
}

type PostResponse struct{
	ID 			int 		`json:"id" form:"id" `
	Title 		string  	`json:"title" form:"title" `
	Description	 string  	`form:"description" json:"description"  `
	UserID 		int 		`json:"user_id" form:"user_id" `
}

func (PostResponse) TableName() string {
	return "posts"	
}