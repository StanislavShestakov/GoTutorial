package main

import ("fmt"
        "net/http"
       )


type User struct {
  name string
  age uint16
  money int16
  avg_grades, happines float64

}

func (u User) getAllInfo() string {
  return fmt.Sprintf("User name is: %s Hi is %d and he has money"+
    "equal: %d", u.name, u.age, u.money)
}
//если структура большая, то передавть обЪект даже если мы его не измеяем
//стоит по ссылке - *User
func (u *User) setNewName(newName string){
  u.name = newName
}

func home_page(w http.ResponseWriter, r *http.Request){
  bob := User{"Bob",25,50, 4.3, 0.8}
  bob.name = "Alex"
  bob.setNewName("God")
  fmt.Fprintf(w, bob.getAllInfo())
}

func contacts_page(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Contacts page!")
}

func handleRequest(){
  http.HandleFunc("/", home_page)
  http.HandleFunc("/contacts", contacts_page)
  http.ListenAndServe(":8081", nil)
}


func main(){
  //var bob User = ..............
  //bob := User{name: "Bob", age: 25, money: 50, avg_grades:4.2, happines: 0.8}
  //  bob := User{"Bob",25,50, 4.3, 0.8}
  handleRequest();
}
