package model

type Article struct {
	id 		int 	`json:"id"`
	Title 	string 	`json: "primarykey"`
	Body 	string 	`json:"body"`
	IsPinned bool 	`json:"is_Pinned"`

}