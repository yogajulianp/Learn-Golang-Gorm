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
	TagsID 		[]int		`json:"tags_id" form:"tags_id" gorm:"-"`
}

type PostResponse struct{
	ID 			int 		`json:"id" form:"id" `
	Title 		string  	`json:"title" form:"title" `
	Description	 string  	`form:"description" json:"description"  `
	UserID 		int 		`json:"user_id" form:"user_id" `
}

type PostResponseWithTags struct {
	ID 			int 		`json:"id" form:"id"  `
	Title 		string  	`json:"title" form:"title" `
	Description	 string  	`form:"description" json:"description"  `
	UserID 		int 		`json:"-" form:"user_id" `
	User		 UserResponse 	`json:"user"`
	Tag			[]Tag		`json:"tags" gorm:"many2many:post_tags;ForeignKey:ID;joinForeignKey:PostID;References:ID;joinReferences:TagID"`
}

type PostResponseNoUser struct{
	ID 			int 		`json:"id" form:"id" `
	Title 		string  	`json:"title" form:"title" `
	Description	 string  	`form:"description" json:"description"  `
	UserID 		int 		`json:"-" form:"user_id" `
}

func (PostResponse) TableName() string {
	return "posts"	
}

func (PostResponseWithTags) TableName() string {
	return "posts"
}

func (PostResponseNoUser) TableName() string {
	return "posts"
}