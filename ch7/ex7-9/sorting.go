package main

import (
	"time"
	"sort"
	"net/http"
	"log"
	"html/template"
	//"fmt"
)

const templ = `
	<h1>Tracks</h1>
	<table>
		<tr>
			<th><a href="/sort?column=Title">Title</a></th>
			<th><a href="/sort?column=Artist">Artist</a></th>
			<th><a href="/sort?column=Album">Album</a></th>
			<th><a href="/sort?column=Year">Year</a></th>
			<th><a href="/sort?column=Length">Length</a></th>
		</tr>
		{{range .}}
		<tr>
			<td>{{.Title}}</td>
			<td>{{.Artist}}</td>
			<td>{{.Album}}</td>
			<td>{{.Year}}</td>
			<td>{{.Length}}</td>
		</tr>
		{{end}}
	</table>
`

type Track struct {
	Title string
	Artist string
	Album string
	Year int
	Length time.Duration
}

type multiTier struct {
	t []*Track
	lesses []func(x, y *Track) bool
}

var tracks = []*Track {
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

var m multiTier

var columnFunctions = map[string]func(x, y *Track) bool{
	"Title": func(x, y *Track) bool { if x.Title != y.Title { return x.Title < y.Title }; return false },
	"Artist": func(x, y *Track) bool { if x.Artist != y.Artist { return x.Artist < y.Artist }; return false },
	"Album": func(x, y *Track) bool { if x.Album != y.Album { return x.Album < y.Album }; return false },
	"Year": func(x, y *Track) bool { if x.Year != y.Year { return x.Year < y.Year }; return false },
	"Length": func(x, y *Track) bool { if x.Length != y.Length { return x.Length < y.Length}; return false },
}

func main() {
	m.t = tracks
	m.lesses = []func(x, y *Track) bool{}
	sort.Sort(m)
	http.HandleFunc("/", handler)
	http.HandleFunc("/sort", sortHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	renderPage(w)
}

func sortHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		if k == "column" {
			val, ok := columnFunctions[v[0]]
			if ok {
				m.addColumn(val)
			}
		}
	}
	sort.Sort(m)
	renderPage(w)
}

func renderPage(w http.ResponseWriter) {
	t := template.Must(template.New("tracks").Parse(templ))
	if err := t.Execute(w, tracks); err != nil {
		log.Fatal(err)
	}
}

func (m *multiTier) addColumn(f func(x, y *Track) bool) {
	m.lesses = append(m.lesses, f)
}

func(x multiTier) Len() int { return len(x.t) }
func(x multiTier) Less(i, j int) bool {
	result := false
	for _, l := range x.lesses {
		result = l(x.t[i], x.t[j])
	}
	return result
}
func(x multiTier) Swap(i, j int) { x.t[i], x.t[j] = x.t[j], x.t[i] }

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}
