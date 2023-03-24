package main

import "fmt"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name        string
	ChasingTail bool
}

type Cat struct {
	Name  string
	Angry bool
}

func main() {
	dog := Dog{
		Name:        "Samson",
		ChasingTail: true,
	}

	//Best practise is to use pointers
	PrintInfo(&dog)
}

func PrintInfo(a Animal) {
	fmt.Println("This animal says", a.Says(), "and has", a.NumberOfLegs(), "legs")
}

func (d *Dog) Says() string {
	return "Woof!"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}
