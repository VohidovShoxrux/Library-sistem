package controllers

import (
	"fmt"
	"library/models"
)


// kitob qarzga olish


func BorrowedBook(id1 int, id2 int) {

	for i := 0; i < len(models.MemberArray); i++ {
		if models.MemberArray[i].ID == id1 {
           
			for j := 0; j < len(models.Library); j++ {
				if models.Library[j].ID == id2 {

					if models.Library[j].IsAvailable{

						models.MemberArray[i].BorrowedBooks = append(models.MemberArray[i].BorrowedBooks, models.Library[j])

						models.Library[j].IsAvailable=false
				
						
						return
					} else {
                        fmt.Println("Bu ID kitob ololmasiz")

						
					}



				}
			}

		}
	}
}


// qarzga olingan kitobni qaytarish


func Returnbook(id1 int, id2 int){
	for i:=0; i<len(models.MemberArray); i++{
		if models.MemberArray[i].ID==id1{
			
			if len(models.MemberArray[i].BorrowedBooks)!=0{

				for j:=0; j<len(models.MemberArray[i].BorrowedBooks); j++{
				if models.MemberArray[i].BorrowedBooks[j].ID==id2{

					for k:=0; k<len(models.Library); k++{
						if models.Library[k].ID==id2 {
                           models.Library[k].IsAvailable=true
						}
                        
						}
							  models.MemberArray[i].BorrowedBooks = append(models.MemberArray[i].BorrowedBooks[:j], models.MemberArray[i].BorrowedBooks[j+1:]...)
						fmt.Println("ishlayapti")	
						return	
                   } 
                }
	    
			} 
  
		}
 
	}

}