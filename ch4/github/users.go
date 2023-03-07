package github

import (
	"fmt"
	"encoding/json"
	"net/http"
)

type User struct {
	Login string
	HTMLURL string `json:"html_url"`
}

func ListUsers() ([]*User, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.github.com/users", nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get users: %s\n", resp.Status)
	}
	var result []*User
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return result, nil
}
