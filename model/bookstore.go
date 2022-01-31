package model

type Bookstore struct {
	ID      uint   `json:"ID" gorm:"primary_key"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	ShelfNo uint   `json:"shelf_no"`
}
