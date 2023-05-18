package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Car struct {
	Make  string `json:"make"`
	Model string `json:"model"`
	VinID string `json:"vinId"`
}

type Person struct {
	Name   string `json:"name"`
	Age    uint   `json:"age"`
	Gender bool   `json:"gender"`
}

type Animal struct {
	Name   string `json:"name"`
	Age    uint   `json:"age"`
	Gender bool   `json:"gender"`
	Breed  string `json:"breed"`
}

func main() {

	a := &Animal{
		Name:   "Potter",
		Age:    13,
		Gender: true,
		Breed:  "chandozo",
	}

	fmt.Printf("%v\n", a)

	b, err := json.Marshal(a)
	if err != nil {
		log.Fatalf("%v", err)
	}

	fmt.Printf("%v\n", string(b))

	p := &Person{}
	if err := json.Unmarshal(b, p); err != nil {
		log.Fatalf("%v\n", err)
	}

	fmt.Printf("%v\n", p)

	c := &Car{}
	if err := json.Unmarshal(b, c); err != nil {
		log.Fatalf("%v\n", err)
	}

	fmt.Printf("%v\n", c)
}
