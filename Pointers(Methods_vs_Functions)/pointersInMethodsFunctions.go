package main

import (
	"fmt"
	"strconv"
)

// ------------------------------------------------------------
// USING POINTERS TO THE STRUCTURES WITH METHODS AND FUNCTIONS
// ------------------------------------------------------------

// create a structure named Car which has an attribute "brand"
type Car struct {
	brand string
}

//------------------------------------------------------------
// create a method AddModel to manipulate with brand name
// Create a pointer(*) to Car ==> that means any changes in this method
// will be pointed to the instance's attribute of Car ==> brand

func (c *Car) AddModel(arg string) {
	c.brand = c.brand + " " + arg
	fmt.Printf("AddModel method: %s\n", c.brand)
}

//------------------------------------------------------------
// when we want to make a function which may influence
// the structure ==> change its attributes,
// we should also create a pointer to the structure ==> *Car
// as a function arguments we put pointer and other arguments

func AddYear(c *Car, year int) {
	c.brand = c.brand + " " + strconv.Itoa(year)
	fmt.Println("AddYear func: ", c.brand)
}

// ------------------------------------------------------------
func main() {

	// lets create an instance of Car

	Mers := Car{"Mersedes"}
	fmt.Println("Show Car instance :", Mers)

	//------------------------------------------------------------

	// lets add a model using AddModel method
	// by calling an method on the instance with an attribute;
	// we use { Instance.Method } when call it

	Mers.AddModel("E200")
	fmt.Printf("Show updated Car instance (using method AddModel): %s\n", Mers)

	//------------------------------------------------------------

	// lets add year of production via function pointed to the structure instance
	// via a pointer. When we call a function, we should use symbol &
	// in order to point it to the structure's instance;
	// we use only { Function(&instance name, argument) } when call it

	AddYear(&Mers, 2022)
	fmt.Println("Show updated Car instance (using function AddYear): ", Mers)
}
