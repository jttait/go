package main

import (
	"fmt"
	"time"
	"text/tabwriter"
	"os"
	"sort"
)

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

func main() {
	printTracks(tracks)
	fmt.Println()
	m := multiTier{tracks, []func(x, y *Track) bool{}}
	m.addColumn(func(x, y *Track) bool { return x.Title < y.Title })
	m.addColumn(func(x, y *Track) bool { return x.Length < y.Length })
	m.addColumn(func(x, y *Track) bool { return x.Year < y.Year })
	sort.Sort(m)
	printTracks(tracks)
}

func (m *multiTier) addColumn(f func(x, y *Track) bool) {
	m.lesses = append(m.lesses, f)
}

func(x multiTier) Len() int { return len(x.t) }
func(x multiTier) Less(i, j int) bool {
	result := true
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

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}
