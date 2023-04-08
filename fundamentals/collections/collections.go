package collections

import (
	"fmt"

	"github.com/pterm/pterm"
)

func WorkingWithArray() {

	pterm.DefaultHeader.WithFullWidth().Println("Working with Array")
	pterm.Info.Println("Arrays are the basic data structure in go\n - it has a fixed size\n - elements have to be the same type")
	var arr [3]int

	arr[0] = 42
	arr[1] = 43
	arr[2] = 44

	fmt.Printf("An array of ints: %v\n", arr)

	var stringArr [4]string

	stringArr[0] = "This is a"
	stringArr[1] = "array of"
	stringArr[2] = "strings"
	stringArr[3] = "with four elements"
	fmt.Printf("An array of strings: %v\n", stringArr)
}

func WorkingWithSlice() {

	pterm.DefaultHeader.WithFullWidth().Println("Working with Slices")
	pterm.Info.Println("Slices are similar with arrays, the difference is in the dynamic sizing")

	emptySlice := []int{0, 0}
	fmt.Printf("slice %v\n", emptySlice)

	emptySlice[0] = 1

	slice := []int{1, 2, 3}
	fmt.Printf("slice: %v\n", slice)

	slice = append(slice, 4)
	fmt.Printf("slice: %v\n", slice)

	pterm.Info.Println(`In the example above the slice started with 3 elements and
one more elements has added using the append() function`)

}

func IteratingOverArrayOfInt() {

	var numbers [6]int32

	numbers[0] = 1
	numbers[1] = 1
	numbers[2] = 2
	numbers[3] = 3
	numbers[4] = 5
	numbers[5] = 8

	var result int32
	for _, n := range numbers {
		result = result + int32(n)
		fmt.Printf("The number is %d and the result is %d\n", n, result)
	}
}

func FilterArray() {

	pterm.DefaultHeader.WithFullWidth().Println("Filter a array")

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println(numbers)

	stmGraterThan2 := func(n int) bool {
		return n > 2
	}

	filteredNumbers := FilterInt(numbers, stmGraterThan2)
	fmt.Println(filteredNumbers)
}

func FilteringSliceOfStruct() {

}

func FilterInt(is []int, filter func(n int) bool) []int {
	var filtered []int
	for _, i := range is {
		if filter(i) {
			filtered = append(filtered, i)
		}
	}
	return filtered
}
