package mb

import (
	
	"library/models"
)

var NewUser = func (){

var Newuser models.Member = models.Member{
	ID: 1,
	Firstname: "shoxrux",
	Lastname: "vohidov",
	BorrowedBooks: []models.Book{},
}
	models.MemberArray=append(models.MemberArray, Newuser)

var Newuser1 models.Member = models.Member{
	ID: 2,
	Firstname: "zokirjon",
	Lastname: "ochildiyev",
	BorrowedBooks: []models.Book{},
	
}
models.MemberArray=append(models.MemberArray, Newuser1)
var Newuser2 models.Member = models.Member{
	ID: 3,
	Firstname: "begali",
	Lastname: "akbarjonov",
	BorrowedBooks: []models.Book{},
	
}
models.MemberArray=append(models.MemberArray, Newuser2)
}