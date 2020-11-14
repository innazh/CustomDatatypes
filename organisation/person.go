package organisation

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//type alias [cannot write methods for it],it's just another name for a `string`
//copies over the method sets and the fields
//type TwitterHandler = string

//type declaration [you're free to define methond cause it's a new type of type string]
//only copies over the fields
type TwitterHandler string

func (th TwitterHandler) RedirectUrl() string {
	cleanHandler := strings.TrimPrefix(string(th), "@")
	return fmt.Sprintf("https://wwww.twitter.com/%s", cleanHandler)
}

type Identifiable interface {
	ID() string
}

type Citizen interface {
	Identifiable
	Country() string
}

type socialSecurityNumber string

/*we can return Identifiable because we implemented the method
in Identifiable interface, which means that socialSecurityNumber
implements Identifiable*/
func NewSocialSecurityNumber(value string) Citizen {
	return socialSecurityNumber(value)
}

func (ssn socialSecurityNumber) ID() string {
	return string(ssn)
}

func (ssn socialSecurityNumber) Country() string {
	return "United States of America"
}

type europeanIdentifier struct {
	id      string
	country string
}

func NewEuropeanIdentifier(id interface{}, country string) Citizen {
	switch v := id.(type) {
	case string:
		return europeanIdentifier{
			id:      v,
			country: country,
		}
	case int:
		return europeanIdentifier{
			id:      strconv.Itoa(v),
			country: country,
		}
	default:
		panic("using an invalid type to initialize EU Identifier")
	}
}

func (eui europeanIdentifier) ID() string {
	return eui.id
}

func (eui europeanIdentifier) Country() string {
	return fmt.Sprintf("EU: %s", eui.country)
}

type Name struct {
	first string
	last  string
}

func (n *Name) FullName() string {
	//return p.firstName + " " + p.lastName
	return fmt.Sprintf("%s %s", n.first, n.last)
}

type Employee struct {
	Name
}

/*Better make all methods pointer-based:
1. Consistency
2. The object might be hefty and you don't wanna be going around
	- making copies of it everywhere*/
type Person struct {
	Name           //an embedded struct
	twitterHandler TwitterHandler
	Citizen        //an embedded interface
}

//constructors in go: always New+Type, returns the type
func NewPerson(firstname, lastname string, citizen Citizen) Person {
	return Person{
		Name: Name{
			first: firstname,
			last:  lastname,
		},
		Citizen: citizen,
	}
}

//needs to be a pointer in order to change the state of the object
//otherwise, it only changes the copy-object's value
func (p *Person) SetTwitterHandler(handler TwitterHandler) error {
	if len(handler) == 0 {
		p.twitterHandler = handler
	} else if !strings.HasPrefix(string(handler), "@") {
		return errors.New("twitter handler must start with an @ symbol")
	}

	p.twitterHandler = handler
	return nil
}

func (p *Person) TwitterHandler() TwitterHandler {
	return p.twitterHandler
}
