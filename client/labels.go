package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/liamawhite/todoist/types"
)

const labelPath = "labels"

func (c *Client) ListLabels() ([]types.Label, error) {
	req, _ := http.NewRequest("GET", labelPath, nil)
	resp, err := c.do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var labels []types.Label
	if err := json.Unmarshal(body, &labels); err != nil {
		return nil, err
	}
	return labels, nil
}

func (c *Client) GetLabelID(name string) (int, error) {
	labels, err := c.ListLabels()
	if err != nil {
		return 0, fmt.Errorf("error retrieving list of labels: %v", err)
	}
	for _, label := range labels {
		if label.Name == name {
			return label.ID, nil
		}
	}
	return 0, fmt.Errorf("label name %v not found", name)
}
