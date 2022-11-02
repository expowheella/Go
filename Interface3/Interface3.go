package main

import (
	"fmt"
)

func main() {
	user1, house1 := User{}, House{}

	// it is possible to add attributs via interface and method
	user1.addName("Alsu")

	// bit also directly assign value to the structure attribute
	house1.AGE = 10
	user1.AGE = 33

	fmt.Printf("\n %v is %v years old, whose house is %v years old", user1.getName(), user1.getAge(), house1.getAge())

}

// create a structure
type User struct {
	NAME    string
	AGE     int64
	MARRIED bool
}

type House struct {
	NAME string
	AGE  int64
}

//////////////////////////////////////
//////////////////////////////////////

// create a method of the structure
// it does not pointed to the structure because it does not make changes
// to the structure. It just gets values (attributes)
func (u User) getName() interface{} {
	return u.NAME
}

func (u User) getAge() interface{} {
	return u.AGE
}

func (h House) getAge() interface{} {
	return h.AGE
}

//////////////////////////////////////
//////////////////////////////////////

// create a method to add an attribute name
// when we add something, we want to make change to the structure instance
// that is why we have to add * --> to point to the User
func (u *User) addName(name string) {
	u.NAME = name
}

// create interface to contol methods
type UserInterface interface {
	getName() string
	addName(name string)
	getAge() int
}
