package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	//contact   contactInfo
	contactInfo //make a field with name and type contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	// alex := person{"Alex", "Anderson"} this is good too
	//alex := person{fistName: "Alex", lastName: "Anderson"}
	//fmt.Println(alex)
	//var alex person

	//alex.firstName = "Alex"
	//alex.lastName = "Anderson"
	//fmt.Printf("%+v", alex) // print struct with field names

	jim := person{
		firstName: "Jim",
		lastName:  "Pat",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 1111,
		},
	}
	//jimPointer := &jim
	//jimPointer.updateFirstName("Joe")
	//jimPointer.updateFirstName("Joe")

	jim.updateFirstName("Joe") //in GO jim can be a point of a person too, not just a person
	jim.print()

	//fmt.Println(&jimPointer)
	//fmt.Println(jimPointer)
	//fmt.Println(jim)

}

func (p person) print() {
	fmt.Printf("%+v", p) //print with field names
}

func (pointerToPerson *person) updateFirstName(newFirstName string) { // type *person: a pointer that points to a person
	(*pointerToPerson).firstName = newFirstName
}
