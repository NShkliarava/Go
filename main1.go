package main
import "fmt"
 
type contact struct{
    email string
    phone string
}
 
type person struct{
    name string
    age int
    contactInfo contact
}
 
func main() {
     
    var nastya = person {
        name: "Nastya", 
        age: 24,
        contactInfo: contact{
            email: "anshkliarava@gmail.com",
            phone: "+375447545977",
        },
    }
    nastya.contactInfo.email = "nastenassa@mail.ru"
     
    fmt.Println(nastya.contactInfo.email)      
    fmt.Println(nastya.contactInfo.phone)      
}