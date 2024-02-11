package main

import "fmt"

type contactInfo struct {
	email   string
	ZipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	jim := person{
		firstName: "Sid",
		lastName:  "Shekhar",
		contact: contactInfo{
			email:   "Sid@aditi.com",
			ZipCode: 277001,
		},
	}

	jim.updateName("Siddharth")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// // // //without datatype name use this // // //
// fmt.Println(jim)
// // // //with data type name use this // // // /
// fmt.Printf("%+v", jim)
