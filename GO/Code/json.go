package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func main() {
	myJsonData := `
		[
			{
				"first_name": "Bruce",
				"last_name":"Wayne"
			},
			{
				"first_name": "Peter",
				"last_name":"Parker"
			}
		]`

	var unmarshalled []person
	err := json.Unmarshal([]byte(myJsonData), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshal json", err)
	}

	log.Printf("unmarshalled: %v", unmarshalled)

	//Create json
	var mySlice []person

	var m1 person
	m1.FirstName = "Wally"
	m1.LastName = "West"

	mySlice = append(mySlice, m1)

	newJson, err := json.MarshalIndent(mySlice, "", "     ")
	if err != nil {
		log.Println("error marshal", err)
	}

	fmt.Println(string(newJson))
}
