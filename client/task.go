package client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/liamawhite/todoist/helpers"
	"github.com/liamawhite/todoist/types"
)

const taskPath = "tasks"

func (c *Client) CreateTask(task types.TaskApply) error {
	body, err := json.Marshal(task)
	if err != nil {
		return err
	}
	req, _ := http.NewRequest("POST", taskPath, bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")
	if _, err = c.do(req); err != nil {
		return err
	}
	return nil
}

func (c *Client) UpdateTask(task types.Task) error {
	apply := helpers.ConvertTaskToTaskApply(task)
	body, err := json.Marshal(apply)
	if err != nil {
		return err

	}
	req, _ := http.NewRequest("POST", taskPath+"/"+strconv.Itoa(task.ID), bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")
	if _, err = c.do(req); err != nil {
		return err
	}
	return nil
}

func (c *Client) ListAllTasks() ([]types.Task, error) {
	req, _ := http.NewRequest("GET", taskPath, nil)
	resp, err := c.do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var tasks []types.Task
	if err := json.Unmarshal(body, &tasks); err != nil {

		return nil, err
	}
	return tasks, nil
}
