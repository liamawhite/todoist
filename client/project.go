package client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/liamawhite/todoist/types"
)

const projectPath = "projects"

func (c *Client) ListProjects() ([]types.Project, error) {
	req, _ := http.NewRequest("GET", projectPath, nil)
	resp, err := c.do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var projects []types.Project
	if err := json.Unmarshal(body, &projects); err != nil {
		return nil, err
	}
	return projects, nil
}
