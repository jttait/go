// Exercise 7.16: Write a web-based calculator program.

package main

import (
	"net/http"
	"net/url"
	"fmt"
	"log"
	"html/template"
	"strconv"

	"github.com/jttait/gopl.io/ch7/eval"
)

const templ = `
<p>{{.}}</p>

<table>
	<tr>
		<td><a href="http://localhost:8000/?expr={{.}}7">7</a></td>
		<td><a href="http://localhost:8000/?expr={{.}}8">8</a></td>
		<td><a href="http://localhost:8000/?expr={{.}}9">9</a></td>
	</tr>
	<tr>
		<td><a href="http://localhost:8000/?expr={{.}}4">4</a></td>
		<td><a href="http://localhost:8000/?expr={{.}}5">5</a></td>
		<td><a href="http://localhost:8000/?expr={{.}}6">6</a></td>
	</tr>
	<tr>
		<td><a href="http://localhost:8000/?expr={{.}}1">1</a></td>
		<td><a href="http://localhost:8000/?expr={{.}}2">2</a></td>
		<td><a href="http://localhost:8000/?expr={{.}}3">3</a></td>
	</tr>
	<tr>
		<td><a href="http://localhost:8000/?expr={{.}}0">0</a></td>
	</tr>
	<tr>
		<td><a href="http://localhost:8000/?expr={{.}}%2b">+</a></td>
		<td><a href="http://localhost:8000/?expr={{.}}-">-</a></td>
		<td><a href="http://localhost:8000/?expr={{.}}*">*</a></td>
		<td><a href="http://localhost:8000/?expr={{.}}%2f">/</a></td>
	</tr>
	<tr>
		<td><a href="http://localhost:8000/?expr={{.}}pow(">pow</a></td>
		<td><a href="http://localhost:8000/?expr={{.}}sqrt(">sqrt</a></td>
		<td><a href="http://localhost:8000/?expr={{.}}sin(">sin</a></td>
	</tr>
	<tr>
		<td><a href="http://localhost:8000/?expr={{.}}(">(</a></td>
		<td><a href="http://localhost:8000/?expr={{.}})">)</a></td>
		<td><a href="http://localhost:8000/?expr={{.}},">,</a></td>
	</tr>
	<tr>
		<td><a href="http://localhost:8000/calculate?expr={{.}}">=</a></td>
	</tr>
	<tr>
		<td><a href="http://localhost:8000/">clear</a></td>
	</tr>
</table>
`

func main() {
	http.HandleFunc("/", calculator)
	http.HandleFunc("/calculate", calculate)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func calculator(w http.ResponseWriter, r *http.Request) {
	expr := r.URL.Query().Get("expr")
	t := template.Must(template.New("interface").Parse(templ))
	if err := t.Execute(w, expr); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "failed to execute template")
		return
	}
}

func parseAndCheck(s string) (eval.Expr, error) {
	if s == "" {
		return nil, fmt.Errorf("empty expression")
	}
	expr, err := eval.Parse(s)
	if err != nil {
		return nil, err
	}
	vars := make(map[eval.Var]bool)
	if err := expr.Check(vars); err != nil {
		return nil, err
	}
	for v := range vars {
		if v != "x" && v != "y" && v != "r" {
			return nil, fmt.Errorf("undefined variable: %s", v)
		}
	}
	return expr, nil
}

func calculate(w http.ResponseWriter, r *http.Request) {
	expr, err := parseAndCheck(r.URL.Query().Get("expr"))
	if err != nil {
		http.Error(w, "bad expr: " + err.Error(), http.StatusBadRequest)
		return
	}
	result := expr.Eval(eval.Env{})
	f := strconv.FormatFloat(result, 'f', 2, 64)
	u, _ := url.Parse("http://localhost:8000/?expr=" + f)
	r.URL = u
	calculator(w, r)
}
