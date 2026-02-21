package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	fisrtName string
	lastName  string
	contact   contactInfo
}

func main() {
	jimbo := person{
		fisrtName: "Jim",
		lastName:  "Bo",
		contact: contactInfo{
			email:   "jimbo@go.dev",
			zipCode: 94000,
		},
	}

	jimbo.updateName("Jimmy")
	jimbo.print()
}

func (p *person) updateName(newFirstName string) {
	(*p).fisrtName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
