package model

type Book struct {
	ID     uint   `gorm:"primaryKey" json:"id,omitempty"`
	Title  string `json:"title"`
	Author string `json:"author"`
}
