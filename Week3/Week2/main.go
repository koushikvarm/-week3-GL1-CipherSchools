package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)
type User struct{
	Name string	`json:"name"`
	age int `json:"age"`
}
type Something struct {

	Name 	string 			`json:"name"`
	Married bool 			`json:"married"`
	Age		float64			`json:"age"`
	Address []Address		`json:"address"`
	Marks   []int			`json:"marks"`
	
}
type Address struct {
	Street int   `json:"street"`
	City   string `json:"city"`
}
	
func main() {
	jsonFile, err := os.Open("something.json")
	if err!=nil {
		log.Fatal(err)

	}
	defer jsonFile.Close()
	jsonByteValues,err := ioutil.ReadAll(jsonFile)
	if err!=nil {
		log.Fatal(err)
	}
	var something Something
	json.Unmarshal(jsonByteValues, &something) //converting jason data to struct
	log.Println(something)
	fmt.Print(string(jsonByteValues)) // converting jason data to string
	
}

