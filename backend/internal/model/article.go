package model

type Article struct {
	ID 		int 	`gorm:"primaryKey" json:"id"`
	Title 	string 	`json:"primarykey"`
	Body 	string 	`json:"body"`
	IsPinned bool 	`json:"is_Pinned"`

}