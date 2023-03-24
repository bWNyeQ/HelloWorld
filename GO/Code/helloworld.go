package main

import (
	"log"
	"time"
)

type User struct {
	FirstName   string
	Lastname    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

// This is called a reciver
func (m *User) printFirstName() {
	log.Println(m.FirstName)
}

// Package private
var special string

// Public, visable outside the package
var Special string

func main() {
	//Shorthand for asigning new variables
	user := User{
		FirstName: "Ben",
		Lastname:  "Dutton",
	}

	var user2 User
	user2.FirstName = "Adam"

	log.Println(user.FirstName, user.Lastname, user.BirthDate)
	user2.printFirstName()

	//make(map[KEY]value)
	myMap := make(map[string]string)
	myMap["dog"] = "Samson"
	myMap["cat"] = "Cassie"

	log.Println("dog name is", myMap["dog"])
	log.Println("cat name is", myMap["cat"])

	myMap["dog"] = "fido"

	log.Println("new dog name is", myMap["dog"])

	//Slice
	var mySlice []string
	mySlice = append(mySlice, "Trevor")
	mySlice = append(mySlice, "John")
	mySlice = append(mySlice, "Mary")

	log.Println(mySlice)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	log.Println(numbers)
	log.Println(numbers[2:5])

	if 1 == 1 {
		log.Println("1 = 1")
	}

	cat := "cat"
	if cat == "cat" {
		log.Println("Cat is cat")
	} else {
		log.Println("Cat is not cat")
	}

	switch cat {
	case "cat":
		log.Println("Meow!")
	case "dog":
		log.Println("Bark!")
	default:
		log.Println("*crickets*")
	}

	for i := 0; i < 10; i++ {
		log.Println(i)
	}

	animals := []string{"dog", "fish", "cat", "rihno"}

	for i, animal := range animals {
		log.Println(i, animal)
	}

	for _, animal := range animals {
		log.Println(animal)
	}

	animalMap := make(map[string]string)
	animalMap["dog"] = "FIDO"
	animalMap["cat"] = "FIDOx2"

	for animalType, animal := range animalMap {
		log.Println(animalType, animal)
	}
}

// "Package Private" Only visable from package main
func test() {

}

// "Public" Visable outside the main Package
func Test() {

}
