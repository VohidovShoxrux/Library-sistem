package controllers

import (
	"fmt"
	"library/models"
)

func BorrowedBooks() (array []models.Book) {
	for i := 0; i < len(models.MemberArray); i++ {
		if len(models.MemberArray[i].BorrowedBooks) > 0{

			fmt.Println("Foydalanuvchi ID:",models.MemberArray[i].ID)
			fmt.Println("Foydalanuvchi Name:",models.MemberArray[i].Firstname)
			fmt.Println("Foydalanuvchi Lsatname:",models.MemberArray[i].Lastname)
            fmt.Println("Foydalanuvchi kitoblari:",models.MemberArray[i].BorrowedBooks)
			fmt.Println("")
		}
	}
	return array
} 