package std_library

import (
	"fmt"
	"os"

	"github.com/rafaelshayashi/lab-golang/fundamentals/model"
	"github.com/rafaelshayashi/lab-golang/fundamentals/std_library/printer"
)

func WorkingWithEncoding() {

	//format := printer.Csv
	format := printer.Json
	printer := printer.NewPrinter(format, os.Stdout)

	customer1 := model.NewCustomer("Tony Stark", "10880 Malibu Point", "May 29, 1970")
	customer2 := model.NewCustomer("Peter Park", "20 Ingram St. in Queens", "August 10, 2001")

	printer.WriteCustomer(customer1)

	fmt.Println("\n--- ALl Customers ---")

	fmt.Printf("Type of customer1 %T\n", customer1)
	fmt.Printf("Type of *customer1 %T\n", *customer1)
	fmt.Printf("Type of &customer1 %T\n", &customer1)

	printer.WriteAllCustomer([]*model.Customer{customer1, customer2})
}
