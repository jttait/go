package main

import (
	"log"
	"net/http"
	"fmt"
	"html/template"
)

const templ = `
<table>
	<tr>
		<th>Item</th>
		<th>Price</th>
	</tr>
	{{range $key, $value := .}}
	<tr>
		<td>{{$key}}</td>
		<td>{{$value}}</td>	
	</tr>
	{{end}}
</table>
`

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	t := template.Must(template.New("database").Parse(templ))
	if err := t.Execute(w, db); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "failed to list items")
		return
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n")
		return
	}
	fmt.Fprintf(w, "%s\n", price)

}
