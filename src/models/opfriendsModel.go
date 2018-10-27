package models

import (
	"time"
)

//个人信息表
type OpfriendsModel struct{
	Id int
	UserId int
	Name string
	BirthDay string
	Height string
	Weight string
	Married string
	Education string
	CurrentCity string
	RegistCity string
	BornCity string
	Profession string
	Parents string
	Brothers string
	IsOnly string
	InCome string
	Interest string
	PlaceOther string
	MarryYears string
	ChildNum string
	RequestBase string
	RequestOther string
	ShowMeSpecial string
	SelfRecommend string
	ImageUrlOne string
	ImageUrlTwo string
	ImageUrlThree string
	ImageUrlFour string
	Contact string
	SourceDest string
	CreatedAt time.Time
}

func (OpfriendsModel) TableName() string {
	return "yz_friends"
}