package cmd

import (
	"fmt"
	"time"
	"tasks/tasks"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add <description>",
	Short: "Add a new task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		description := args[0]
		taskList, _ := tasks.LoadTasks("tasks.csv")

		// Create a new task
		newTask := tasks.Task{
			ID:          len(taskList) + 1,
			Description: description,
			CreatedAt:   time.Now(),
			IsComplete:  false,
		}

		taskList = append(taskList, newTask)
		tasks.SaveTasks("tasks.csv", taskList)
		fmt.Printf("Task added: %s\n", description)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}