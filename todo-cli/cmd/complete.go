package cmd

import (
	"fmt"
	"strconv"
	"tasks/tasks"

	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete <taskid>",
	Short: "Mark a task as complete",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskID, _ := strconv.Atoi(args[0])
		taskList, _ := tasks.LoadTasks("tasks.csv")

		for i, task := range taskList {
			if task.ID == taskID {
				taskList[i].IsComplete = true
				tasks.SaveTasks("tasks.csv", taskList)
				fmt.Printf("Task %d marked as complete\n", taskID)
				return
			}
		}

		fmt.Printf("Task %d not found\n", taskID)
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}