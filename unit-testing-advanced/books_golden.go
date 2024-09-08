package main

import (
	"io"
	"log"
	"os"
	"text/template"
)

type Book struct {
	Name          string
	Author        string
	Publisher     string
	Pages         int
	PublishedYear int
	Price         int
}

var inventory = []Book{
	{
		Name:          "The Solitare Mystery",
		Author:        "Jostein Gaarder",
		Publisher:     "Farrar Straus Giroux",
		Pages:         351,
		PublishedYear: 1990,
		Price:         12,
	},
	{
		Name:          "Also Known As",
		Author:        "Robin Benway",
		Publisher:     "Walker Books",
		Pages:         208,
		PublishedYear: 2013,
		Price:         10,
	},
	{
		Name:          "Ego Is the Enemy",
		Author:        "Ryan Holiday",
		Publisher:     "Portfolio",
		Pages:         226,
		PublishedYear: 2016,
		Price:         18,
	},
}

var tmpl = `<table class="table">
  <thead>
    <tr>
      <th>Name</th>
      <th>Author</th>
      <th>Publisher</th>
      <th>Pages</th>
      <th>Year</th>
      <th>Price</th>
    </tr>
  </thead>
  <tbody>
    {{ range . }}<tr>
      <td>{{ .Name }}</td>
      <td>{{ .Author }}</td>
      <td>{{ .Publisher }}</td>
      <td>{{ .Pages }}</td>
      <td>{{ .PublishedYear }}</td>
      <td>${{ .Price }}</td>
    </tr>{{ end }}
  </tbody>
</table>
`

var tpl = template.Must(template.New("table").Parse(tmpl))

func generateTable(books []Book, w io.Writer) error {
	return tpl.Execute(w, books)
}

func main() {
	books := []Book{
		{
			Name:          "The Odessa File",
			Author:        "Frederick Forsyth",
			Pages:         334,
			PublishedYear: 1979,
			Publisher:     "Bantam",
			Price:         15,
		},
	}

	err := generateTable(books, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}
