package main

import (
	"fmt"

	"github.com/liamawhite/todoist/client"
)

func main() {
	t := client.NewClient("<your-API-token-here>")
	tasks, _ := t.ListAllTasks()
	fmt.Println(tasks)
}
