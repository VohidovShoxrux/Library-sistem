package mb

import (
	"library/models"
	
)


var Newbook = func(){

var NewBooks models.Book= models.Book{
	ID: 1,
	Name: "halqa",
	IsAvailable: true,
}
models.Library = append(models.Library, NewBooks)
var NewBooks1 models.Book= models.Book{
	ID: 2,
	Name: "tuman",
	IsAvailable: true,
}
models.Library = append(models.Library, NewBooks1)
var NewBooks2 models.Book= models.Book{
	ID: 3,
	Name: "sahro",
	IsAvailable: true,
}
models.Library = append(models.Library, NewBooks2)
}