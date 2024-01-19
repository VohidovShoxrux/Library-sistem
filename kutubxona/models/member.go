package models

var MemberArray []Member

type Member struct {
	ID int 
	Firstname string
	Lastname string
	BorrowedBooks []Book
}

