package models

import (
	"time"
)

type ChapterModel struct {
	Id int
	BookId int
	Name string
	Url string
	Order int
	Pre int
	Next int
	Content string
	CreatedAt time.Time
}

func (ChapterModel) TableName() string {
	return "yz_chapter"
}