package models

import (
	"time"
)

type BookModel struct {
	Id int
	Title string
	Image string
	Url string
	CreatedAt time.Time
}

func (BookModel) TableName() string {
	return "yz_book"
}