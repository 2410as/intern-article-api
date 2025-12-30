package model

type Article struct {
	ID 		 int    `gorm:"primaryKey" json:"id"`
	Title 	 string `json:"title"`      // ここを小文字の title に
	Body 	 string `json:"body"`       // ここを小文字の body に
	IsPinned bool   `json:"is_pinned"`  // ここを小文字の is_pinned に
}


