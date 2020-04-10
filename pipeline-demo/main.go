package main

import (
	"html/template"
	"os"
)

const tax = 21.3/100

type Product struct {
	Name string
	Price float32
}

func (p Product) PriceWithTax() float32 {
	return p.Price * (1 + tax)
}

const templateString = `
{{- "Item Information" }}
Name {{.Name}}
Price {{ printf "$%.2f" .Price}}
Price with tax: {{.PriceWithTax | printf "$%.2f" }}
`

func main() {
	p := Product{Name: "Lemonade", Price: 2.20}
	t := template.Must(template.New("").Parse(templateString))
	t.Execute(os.Stdout, p)
}