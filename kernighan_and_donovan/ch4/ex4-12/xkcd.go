package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
	"strconv"
	"encoding/json"
	"strings"
	"flag"
)

func main() {
	var queryFlag = flag.String("term", "elbonia", "term to search for")	
	flag.Parse()
	downloadAllEpisodes()
	index := indexAllEpisodes()
	fmt.Println(index[*queryFlag])
}

type Episode struct {
	News string
	SafeTitle string `json:"safe_title"`
	Alt string
	Transcript string
}

func indexAllEpisodes() map[string][]int {
	index := make(map[string][]int)
	for i := 1; i <= 2745; i++ {
		index, _ = indexEpisode(i, index)
	}
	return index
}

func indexEpisode(episodeNumber int, index map[string][]int) (map[string][]int, error) {
	contents, err := os.ReadFile("comics/" + strconv.Itoa(episodeNumber) + ".json")
	if err != nil {
		return index, err
	}
	var episode Episode
	json.Unmarshal(contents, &episode)

	f := func(c rune) bool {
		return c == ' ' || c == '\n'
	}

	for _, element := range strings.FieldsFunc(episode.SafeTitle, f) {
		element = cleanWord(element)
		index[element] = append(index[element], episodeNumber)
	}
	for _, element := range strings.FieldsFunc(episode.News, f) {
		element = cleanWord(element)
		index[element] = append(index[element], episodeNumber)
	}
	for _, element := range strings.FieldsFunc(episode.Alt, f) {
		element = cleanWord(element)
		index[element] = append(index[element], episodeNumber)
	}
	for _, element := range strings.FieldsFunc(episode.Transcript, f) {
		element = cleanWord(element)
		index[element] = append(index[element], episodeNumber)
	}
	return index, nil
}

func cleanWord(s string) string {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, ",", "")
	s = strings.ReplaceAll(s, ".", "")
	s = strings.ReplaceAll(s, "!", "")
	s = strings.ReplaceAll(s, ";", "")
	s = strings.ReplaceAll(s, ":", "")
	s = strings.ReplaceAll(s, ")", "")
	s = strings.ReplaceAll(s, "]", "")
	s = strings.ReplaceAll(s, "(", "")
	s = strings.ReplaceAll(s, "[", "")
	s = strings.ReplaceAll(s, "'", "")
	s = strings.ReplaceAll(s, "{", "")
	s = strings.ReplaceAll(s, "}", "")
	s = strings.ReplaceAll(s, "?", "")
	s = strings.ReplaceAll(s, ">", "")
	s = strings.ReplaceAll(s, "<", "")
	s = strings.ReplaceAll(s, "\"", "")
	return s
}

func downloadAllEpisodes() {
	for i := 1; i <= 2745; i++ {
		err := downloadEpisode(i)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func downloadEpisode(episode int) error {
	_, err := os.Open("comics/" + strconv.Itoa(episode) + ".json")
	if err == nil {
		return nil
	}
	resp, err := http.Get("https://xkcd.com/" + strconv.Itoa(episode) + "/info.0.json")
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to download episode %d: %s", episode, resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	os.WriteFile("comics/" + strconv.Itoa(episode) + ".json", body, 0666)
	resp.Body.Close()
	return nil
}
