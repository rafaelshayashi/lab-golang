package basics

import (
	"fmt"
	"reflect"

	"github.com/pterm/pterm"
)

func WorkingWithPrimitives() {

	pterm.DefaultHeader.WithFullWidth().Println("Working with primitives")
	pterm.DefaultSection.Println("Working with integer")

	var number int = 42
	fmt.Printf("The number is %d and has a type %s\n", number, reflect.TypeOf(number))

	numberGopher := 42
	fmt.Printf("The number is %d and has a type %s and the variable was assing using the gopher operator \n",
		numberGopher, reflect.TypeOf(numberGopher))

	pterm.DefaultSection.Println("Working with floating")
	var pi = 3.14
	fmt.Printf("The number is %f and has a type %s\n", pi, reflect.TypeOf(pi))

}

func WorkingWithPointer() {

	pterm.DefaultHeader.WithFullWidth().Println("Working with pointers")
	name := "Rafael"
	ptr := &name

	fmt.Printf("the string is %s and has a type %s\n", name, reflect.TypeOf(name))
	fmt.Printf("the pointer is %p and has a type %s\n", ptr, reflect.TypeOf(ptr))
	fmt.Printf("I can also access the string value %s\n", *ptr)
}
