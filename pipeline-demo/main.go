package main

import (
	"html/template"
	"os"
)

const tax = 21.3 / 100

type Product struct {
	Name  string
	Price float32
}

const templateString = `
{{- "Item Information" }}
Name {{.Name}}
Price {{ printf "$%.2f" .Price}}
Price with tax: {{ calctax .Price | printf "$%.2f" }}
`

func main() {
	p := Product{Name: "Lemonade", Price: 2.20}
	fm := template.FuncMap{}
	fm["calctax"] = func(price float32) float32 {
		return price * (1 + tax)
	}
	t := template.Must(template.New("").Funcs(fm).Parse(templateString))
	t.Execute(os.Stdout, p)
}
