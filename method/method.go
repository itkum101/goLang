package main

import "fmt"

// WE don't have classess in golang so we have struct as like C

func main() {
	// no inheritance in golang
	// no super, parent in golang
	fmt.Println("Structs in golang")

	randomUser := User{
		"Randy",
		"admin@admin.com",
		true,
		10,
	}
	fmt.Printf("Random user details are %v \n", randomUser)
	//+v gives even syntax too...
	fmt.Printf("Random user details are %+v \n", randomUser)
	randomUser.GetStatus()
	randomUser.NewMail()

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// METHOD

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is ", u.Email)
}
