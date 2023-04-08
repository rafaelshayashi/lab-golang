package std_library

import (
	"fmt"
)

type Item struct {
	name        string
	description string
	value       int
}

func WorkingWithPrintf() {

	item1 := Item{
		name:        "AppleWatch",
		description: "A very expensive gadgets",
		value:       1000,
	}

	// creating a table of items
	// the number between the formating specifier tels how to indent the content, if negative will align to the left
	// and to the right if positive
	fmt.Printf("|%-10s |%-30s |%-10s |\n", "Name", "Description", "Value")
	fmt.Printf("|%-10s |%-30s |%-10d |\n", item1.name, item1.description, item1.value)
}
