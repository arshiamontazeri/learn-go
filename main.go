package main

import (
	"html/template"
	"os"
)

func main() {
	// Define the template
	const tmpl = `
			<html>
			<head>
				<title>{{.Title}}</title>
			</head>
			<body>
				<h1>{{.Heading}}</h1>
				<ul>
					{{range $key, $value := .Items}}
					<li>{{$key}}: {{$value}}</li>
					{{end}}
				</ul>
			</body>
			</html>
			`

	// Parse the template
	t, err := template.New("webpage").Parse(tmpl)
	if err != nil {
		panic(err)
	}

	// Create a map of data

	type f struct {
		Title   string
		Heading string
		Items   map[string]string
	}

	data := f{
		Title:   "My Webpage",
		Heading: "Welcome to My Webpage",
		Items: map[string]string{
			"1": "Item 1",
			"2": "Item 2",
			"3": "Item 3",
		},
	}

	// Execute the template and write to standard output
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
