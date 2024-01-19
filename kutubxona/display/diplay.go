package display

import (
	"fmt"
	"library/controllers"
	"library/controllers/mb"
)

func Display() {
	var loop bool
	mb.NewUser()
	mb.Newbook() 
	
	fmt.Println("___________Library Managment System:___________")
	sakra:
    fmt.Println("")  
	fmt.Println(" Models:  \n 1.Member \n 2.Book \n 3.Exit" )
	fmt.Println("")
	for !loop {
	var  num int 
	fmt.Scan(&num)
	switch num {
	case 1:
		sakra1:
		fmt.Println("___________Member___________")
		fmt.Println("")
		fmt.Println("1.Foydalanuvchi yaratish")
		fmt.Println("2.Foydalanuvchi ma'lumotlarini o'zgartirish")
		fmt.Println("3.Foydalanuvchi malumotlarini o'chirish")
		fmt.Println("4.Foydalanuvchi malumotlarini ko'rish")
		fmt.Println("5.Foydalanuvchi qarzga kitob olish")
		fmt.Println("6.Foydalanuvchi kitobni qaytarish")
		fmt.Println("")
		fmt.Println("____________________________________________")
		fmt.Println("Raqamlar orqali amalyot turini tanlang va raqam kiriting")
		fmt.Println("")
		
        var num1 int
		fmt.Scan(&num1)
		switch num1 {
		case 1:
			fmt.Println("1.Foydalanuvchi ismini kiriting")
			fmt.Println("2.Foydalanuvchi familyasini kiriting")
			var a, b string
			fmt.Scan(&a,&b)
			controllers.Created(a,b)
			fmt.Println("Foydalanuvchi muvofaqiyatli yartildi!")
			fmt.Println("1.Bosh menuga qaytish uchun 1 ni bosing")
			fmt.Println("2.Orqaga qaytish uchun 2 ni bosing")
			var c int
			fmt.Scan(&c)
			switch c {
			case 1: goto sakra
			case 2: goto sakra1
			}

		case 2:
			fmt.Println("1.Foydalanuvchi ID kiriting")
			fmt.Println("2.Foydalanuvchi ismini kiriting")
			fmt.Println("3.Foydalanuvchi familyasini kiriting")
			var a, b string
			var d int
			fmt.Scan(&d,&a,&b)
			controllers.Updated(d,a,b)
			fmt.Println("Foydalanuvchi muvofaqiyatli o'zgartirildi!")
			fmt.Println("1.Bosh menuga qaytish uchun 1 ni bosing")
			fmt.Println("2.Orqaga qaytish uchun 2 ni bosing")
			var c int
			fmt.Scan(&c)
			switch c {
			case 1: goto sakra
			case 2: goto sakra1
			}

		case 3:
			fmt.Println("ID kiriting")
			var l int
			fmt.Scan(&l)
			controllers.Deleted(l)
			fmt.Println("Foydalanuvchi muvofaqiyatli o'chirildi!")
			fmt.Println("1.Bosh menuga qaytish uchun 1 ni bosing")
			fmt.Println("2.Orqaga qaytish uchun 2 ni bosing")
			var c int
			fmt.Scan(&c)
			switch c {
			case 1: goto sakra
			case 2: goto sakra1
			}
		
		case 4:
			controllers.Readed()
			
			fmt.Println("Amalyot muvofaqiyatli bajarildi!")
			fmt.Println("1.Bosh menuga qaytish uchun 1 ni bosing")
			fmt.Println("2.Orqaga qaytish uchun 2 ni bosing")
			var c int
			fmt.Scan(&c)
			switch c {
			case 1: goto sakra
			case 2: goto sakra1
			}
		case 5:
			fmt.Println("1.Foydalanuvchi ID kiriting")
			fmt.Println("2.Qidirilayotgan kitob ID sini kiriting")
			var a,b int
			fmt.Scan(&a,&b)  
			controllers.BorrowedBook(a,b)
			fmt.Println("Amalyot muvofaqiyatli bajarildi!")
			fmt.Println("1.Bosh menuga qaytish uchun 1 ni bosing")
			fmt.Println("2.Orqaga qaytish uchun 2 ni bosing")
			var c int
			fmt.Scan(&c)
			switch c {
			case 1: goto sakra
			case 2: goto sakra1
			}
		case 6:
			fmt.Println("1.Foydalanuvchi ID kiriting")
			fmt.Println("2.Qidirilayotgan kitob ID sini kiriting")
			var a,b int
			fmt.Scan(&a,&b)
			controllers.Returnbook(a,b)
			fmt.Println("Amalyot muvofaqiyatli bajarildi!")
			fmt.Println("1.Bosh menuga qaytish uchun 1 ni bosing")
			fmt.Println("2.Orqaga qaytish uchun 2 ni bosing")
			var c int
			fmt.Scan(&c)
			switch c {
			case 1: goto sakra
			case 2: goto sakra1
			}
		}
	case 2:
		sakra2:
		fmt.Println("___________Book___________")
		fmt.Println("")
		fmt.Println("1.Kitob yaratish")
		fmt.Println("2.Kitob nomini o'zgartirish")
		fmt.Println("3.Kitob malumotlarini o'chirish")
		fmt.Println("4.Kitob malumotlarini ko'rish")
		fmt.Println("5.Qarizga olingan kitoblar")
		fmt.Println("6.Bor yoki yoq ekaniligi")
		fmt.Println("")
		fmt.Println("____________________________________________")
		fmt.Println("Raqamlar orqali amalyot turini tanlang va raqam kiriting")
        fmt.Println("")

		var qw int
		fmt.Scan(&qw)
		switch qw{
		case 1: 
		    fmt.Println("1.Kitob nomini kiriting")
			var str string
			fmt.Scan(&str)
			controllers.Create(str,true)
			fmt.Println("Kitob muvofaqiyatli yartildi!")
			fmt.Println("1.Bosh menuga qaytish uchun 1 ni bosing")
			fmt.Println("2.Orqaga qaytish uchun 2 ni bosing")
			var c int
			fmt.Scan(&c)
			switch c {
			case 1: goto sakra
			case 2: goto sakra2
			}
        case 2:
			fmt.Println("1.Kitob nomini kiriting")
			var str string
			fmt.Scan(&str)
			controllers.Update(str)
			fmt.Println("Kitob muvofaqiyatli o'zgartirildi!")
			fmt.Println("1.Bosh menuga qaytish uchun 1 ni bosing")
			fmt.Println("2.Orqaga qaytish uchun 2 ni bosing")
			var c int
			fmt.Scan(&c)
			switch c {
			case 1: goto sakra
			case 2: goto sakra2
		    }
		case 3:
			fmt.Println("1.Kitob ID sini kiriting")
			var cd int
			fmt.Scan(&cd)
			controllers.Delete(cd)
			fmt.Println("Kitob muvofaqiyatli o'chirildi!")
			fmt.Println("1.Bosh menuga qaytish uchun 1 ni bosing")
			fmt.Println("2.Orqaga qaytish uchun 2 ni bosing")
			var c int
			fmt.Scan(&c)
			switch c {
			case 1: goto sakra
			case 2: goto sakra2
		    }
		case 4:
		
			controllers.Read()
			
			fmt.Println("Amalyot muvofaqiyatli bajarildi!")
			fmt.Println("1.Bosh menuga qaytish uchun 1 ni bosing")
			fmt.Println("2.Orqaga qaytish uchun 2 ni bosing")
			var c int
			fmt.Scan(&c)
			switch c {
			case 1: goto sakra
			case 2: goto sakra2
			}
		case 5:
			controllers.BorrowedBooks()
			fmt.Println("Amalyot muvofaqiyatli bajarildi!")
			fmt.Println("1.Bosh menuga qaytish uchun 1 ni bosing")
			fmt.Println("2.Orqaga qaytish uchun 2 ni bosing")
			var c int
			fmt.Scan(&c)
			switch c {
			case 1: goto sakra
			case 2: goto sakra2
			}
		case 6:
			fmt.Println("Qidirilayotgan kitob ID sini kiriting")
			var id int
			fmt.Scan(&id)
			controllers.IDbook(id)
			fmt.Println("1.Bosh menuga qaytish uchun 1 ni bosing")
			fmt.Println("2.Orqaga qaytish uchun 2 ni bosing")
			var c int
			fmt.Scan(&c)
			switch c {
			case 1: goto sakra
			case 2: goto sakra2
			}
		
		
		}
	case 3:
		
		return

		default:
			fmt.Println("Bunday amalyot mavjud emas!")
			return
		}
	}
}