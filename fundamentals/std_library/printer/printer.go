package printer

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/rafaelshayashi/lab-golang/fundamentals/model"
)

type Format string

type Printer struct {
	writer interface{}
}

const (
	Csv   Format = "csv"
	Json  Format = "json"
	Table Format = "table"
)

func NewPrinter(format Format, w *os.File) *Printer {

	switch format {
	case Csv:
		return &Printer{
			writer: csv.NewWriter(os.Stdout),
		}
	case Json:
		return &Printer{
			writer: bufio.NewWriter(os.Stdout),
		}
	default:
		return &Printer{
			writer: tabwriter.NewWriter(os.Stdout, 0, 8, 0, '\t', tabwriter.AlignRight),
		}
	}
}

func (p *Printer) WriteCustomer(customer *model.Customer) error {
	switch v := p.writer.(type) {
	case *csv.Writer:
		v.Write(customer.ToSlice())
		v.Flush()
	case *bufio.Writer:
		formated, err := json.Marshal(customer)
		if err != nil {
			panic("Error")
		}
		v.Write(formated)
		v.Flush()
	case *tabwriter.Writer:
		fmt.Fprintln(v, customer.ToSlice())
	}
	return nil
}

// WriteAllCustomer should print a slice of customers in a specif format
func (p *Printer) WriteAllCustomer(customers []*model.Customer) error {
	for _, customer := range customers {
		switch v := p.writer.(type) {
		case *csv.Writer:
			p.WriteCustomer(customer)
		case *bufio.Writer:
			p.WriteCustomer(customer)
			v.WriteString(",")
		}
	}
	return nil
}
