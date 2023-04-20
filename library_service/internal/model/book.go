package model

type Book struct {
	ID     uint   `gorm:"primaryKey" json:"id,omitempty" swaggerignore:"true"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Price  uint   `json:"price"`
	Income uint   `json:"income" swaggerignore:"true"`
}

type BookRentSummary struct {
	Title    string `json:"title"`
	ReaderID int    `json:"reader_id"`
	Income   int    `json:"income"`
}
