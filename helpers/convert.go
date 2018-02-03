package helpers

import "github.com/liamawhite/todoist/types"

func ConvertTaskToTaskApply(task types.Task) types.TaskApply {
	return types.TaskApply{
		Content:   task.Content,
		ProjectID: task.ProjectID,
		LabelIDs:  task.LabelIDs,
		Order:     task.Order,
		Priority:  task.Priority,
	}
	// TODO -- Add date updating support
}
