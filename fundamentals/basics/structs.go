package basics

import (
	"fmt"

	"github.com/rafaelshayashi/lab-golang/fundamentals/model"
)

func WorkingWithStructs() {

	// Creating a struct with some values only
	personOne := model.Person{Name: "Tony Stark", Age: 41}
	fmt.Printf("The person is %T \n", personOne)

	// Is possible to omit the key when creating a struct will all values
	personTwo := model.Person{"Peter Park", 41, "A friend from neighborhood"}
	fmt.Printf("The other Person is %T \n", personTwo)

	// Creating a struct using new
	var personThree *model.Person
	personThree = new(model.Person)
	personThree.Name = "Thor"
	personThree.Age = 1500
	personThree.Profile = "The strongest avenger"

	fmt.Println(*personThree)
}
