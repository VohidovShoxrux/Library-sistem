package controllers

import (
	"fmt"
	"library/models"
)

// Member uchun CURD

func Created(firstname string, lastname string) {
	
	
	
	var member models.Member = models.Member{		
		ID:            len(models.MemberArray)+ 1,
		Firstname:     firstname,
		Lastname:      lastname,
		BorrowedBooks: []models.Book{},
	}
	models.MemberArray = append(models.MemberArray, member)
}
func Readed() {
	    
	fmt.Println("___________________________________________")
	for i := 0; i < len(models.MemberArray); i++ {
		fmt.Println("Foydalanuvchi ID:",models.MemberArray[i].ID)
		fmt.Println("Foydalanuvchi ismi:",models.MemberArray[i].Firstname)
		fmt.Println("Foydalanuvchi familyasi:",models.MemberArray[i].Lastname)
		fmt.Println("Foydalanuvchi kitoblari:",models.MemberArray[i].BorrowedBooks)
		fmt.Println("___________________________________________")
		for j := 0; j < len(models.MemberArray[i].BorrowedBooks); j++ {
			
		}
	}
}
func IDmember(id int)(d models.Member, found bool){
	for i := 0; i < len(models.MemberArray); i++ {
		if models.Library[i].ID == id {
			d = models.MemberArray[i]
			found = true
			fmt.Println("Qidirlayoqgan kitob bor")
			break
		} else{
			fmt.Println("Qidirlayoqgan kitob yoq")
		}

	}
	return d, found
}
func Updated(id int,firstname string, lastname string){
	for i:=0; i<len(models.MemberArray); i++{
		if models.MemberArray[i].ID==id{
		models.MemberArray[i].Firstname =firstname
		models.MemberArray[i].Lastname =lastname
		break
		} else{
			fmt.Println("Bu ID da foydalanuvchi ma'lumotlari topilmadi!")
		}
	}
	return
}
func Deleted(id int) {
	for i := 0; i < len(models.MemberArray); i++ {
		if models.MemberArray[i].ID == id {
			models.MemberArray = append(models.MemberArray[:i], models.MemberArray[i+1:]...)
			break
		}
	}
}


// Book uchun CURD

func Create(name string, isAvailable bool) {

	
	var book models.Book = models.Book{

		ID:          len(models.Library) + 1,
		Name:        name,
		IsAvailable: isAvailable,
	}

	models.Library = append(models.Library, book)
}
func Read() {
	
	fmt.Println("___________________________________________")
	for i := 0; i < len(models.Library); i++ {
		fmt.Println("Kitob ID:",models.Library[i].ID)
		fmt.Println("Kitob nomi",models.Library[i].Name)
		fmt.Println("Kitob bor or yoqligi",models.Library[i].IsAvailable)
		
		fmt.Println("___________________________________________")
	}
	
}
func IDbook(id int) (b models.Book, found bool) {
	for i := 0; i < len(models.Library); i++ {
		if models.Library[i].ID == id {
			b = models.Library[i]
			found = true
			break
		}
	}
	return b, found
}

func Update(name string) {
	for i := 0; i < len(models.Library); i++ {
		models.Library[i].Name = name
		break
	}
}
func Delete(id int) {
	for i := 0; i < len(models.Library); i++ {
		if models.Library[i].ID == id {
			models.Library = append(models.Library[:i], models.Library[i+1:]...)
			break
		}
	}
}


