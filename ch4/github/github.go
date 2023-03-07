package github

import (
	"time"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"bytes"
)

const IssuesURL = "https://api.github.com/search/issues"
const CreateURL = "https://api.github.com/repos/jttait/gopl.io/issues"
const GetURL = "https://api.github.com/repos/jttait/gopl.io/issues/"
const UpdateURL = "https://api.github.com/repos/jttait/gopl.io/issues/"

type IssuesSearchResult struct {
	TotalCount int `json:total_count`
	Items []*Issue
}

type Issue struct {
	Number int
	HTMLURL string `json:"html_url"`
	Title string
	State string
	User *User
	CreatedAt time.Time `json:"created_at"`
	Body string
	Milestone *Milestone
}

type CreateIssueResult struct {
	Number int
}

type UpdateIssueRequest struct {
	Title string `json:"title"`
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
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

func GetIssue(issueNumber string) (error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", GetURL + issueNumber, nil)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return fmt.Errorf("get issue query failed: %s", resp.Status)
	}

	fmt.Println(resp.Body)

	return nil
}

func CloseIssue(issueNumber string) error {
	values := map[string]string{"state": "closed"}
	jsonData, err := json.Marshal(values)
	if err != nil {
		return err
	}
	bodyReader := bytes.NewReader(jsonData)

	client := &http.Client{}
	req, err := http.NewRequest("PATCH", UpdateURL + issueNumber, bodyReader)
	req.Header.Add("Authorization", "token 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")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return fmt.Errorf("close issue query failed: %s", resp.Status)
	}
	resp.Body.Close()
	return nil
}

func UpdateIssue(issueNumber string, request UpdateIssueRequest) error {
	jsonData, err := json.Marshal(request)
	fmt.Printf("%s\n", jsonData)

	if err != nil {
		return err
	}
	bodyReader := bytes.NewReader(jsonData)
	client := &http.Client{}
	req, err := http.NewRequest("PATCH", UpdateURL + issueNumber, bodyReader)
	req.Header.Add("Authorization", "token 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")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return fmt.Errorf("update issue failed: %s", resp.Status)
	}
	resp.Body.Close()
	return nil
}

func CreateIssue(title string) (*CreateIssueResult, error) {
	values := map[string]string{"title": title}
	jsonData, err := json.Marshal(values)
	if err != nil {
		return nil, err
	}
	bodyReader := bytes.NewReader(jsonData)

	client := &http.Client{}
	req, err := http.NewRequest("POST", CreateURL, bodyReader)
	req.Header.Add("Authorization", "token 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")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusCreated {
		resp.Body.Close()
		return nil, fmt.Errorf("get issue query failed: %s", resp.Status)
	}

	var result CreateIssueResult
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
