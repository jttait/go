package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ListIssues() (*IssuesSearchResult, error) {
	resp, err := http.Get("https://api.github.com/search/issues?q=repo:golang/go")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result IssuesSearchResult
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
