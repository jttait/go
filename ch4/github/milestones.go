package github

import (
	"fmt"
	"encoding/json"
	"net/http"
)

type Milestone struct {
	Number int
	URL string
}

func ListMilestones() ([]*Milestone, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.github.com/repos/golang/go/milestones", nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get milestones: %s\n", resp.Status)
	}
	var result []*Milestone
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return result, nil
}
