package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contact contactInfo
}

type simplePerson struct {
	firstName string
	lastName string
}

func (pp *person) updateName(newFirstName string) {
	(*pp).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p simplePerson) print() {
	fmt.Printf("%+v\n", p)
}

func main() {
	alex := person{"Alex", "Anderson", contactInfo{"alex@email.com", 123}}
	brad := person{
		firstName: "Brad",
		lastName: "Brendigo",
		contact: contactInfo{
			email: "brad@email.com",
			zipCode: 456,
		},
	}
	var caleb simplePerson
	dale := person{
		firstName: "Eale",
		lastName: "Davies",
		contact: contactInfo{
			email: "dale@email.com",
			zipCode: 1234,
		},
	}

	fmt.Println(alex)
	fmt.Println(brad)
	fmt.Println(caleb)
	fmt.Println(dale)

	alex.print()
	brad.print()
	caleb.print()
	dale.print()

	caleb.firstName = "Caleb"
	caleb.lastName = "Carlston"

	fmt.Println(caleb)
	caleb.print()

	dale.updateName("Dale")
	fmt.Println(dale)
	dale.print()
}