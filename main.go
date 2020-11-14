package main

import (
	"fmt"

	"git.com/organisation"
)

func main() {
	//p is of type Identifiable
	// var p organisation.Identifiable =
	p := organisation.NewPerson("James", "Wilson", organisation.NewEuropeanIdentifier(12345, "Germany"))

	err := p.SetTwitterHandler("@ham_wils")

	if err != nil {
		fmt.Printf("An error occurred setting twitter handler: %s", err.Error())
	}
	// println(p.FullName())
	println(p.ID())
	println(p.Country())
	// fmt.Printf("%T\n", organisation.TwitterHandler("test"))
	// println(p.TwitterHandler())
	// println(p.TwitterHandler().RedirectUrl())

	// name1 := Name{First: "James", Last: "Wilson"}
	// name2 := Name{First: "James", Last: "Wilson"}

	//== and != operands only work with types that have a predictable memory layout
	// ssn := organisation.NewSocialSecurityNumber("123-45-6789")
	// eu := organisation.NewEuropeanIdentifier("12345", "France")
	// //eu2 := organisation.NewEuropeanIdentifier("12345", "France")

	// if ssn == eu {
	// 	println("We match")
	// }

	// fmt.Printf("%T\n", ssn)
	// fmt.Printf("%T\n", eu)

}

type Name struct {
	First  string
	Last   string
	Middle []string
}

//When the types are too complexto be `Comparable` (to use == or !=), we can implement our own Equals method
func (n Name) Equals(otherName Name) bool {
	return n.First == otherName.First && n.Last == otherName.Last && len(n.Middle) == len(otherName.Middle)
}

type OtherName struct {
	First string
	Last  string
}
