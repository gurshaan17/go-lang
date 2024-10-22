package cmd

import (
	"fmt"
	"strconv"
	"tasks/tasks"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <taskid>",
	Short: "Delete a task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskID, _ := strconv.Atoi(args[0])
		taskList, _ := tasks.LoadTasks("tasks.csv")

		for i, task := range taskList {
			if task.ID == taskID {
				taskList = append(taskList[:i], taskList[i+1:]...)
				tasks.SaveTasks("tasks.csv", taskList)
				fmt.Printf("Task %d deleted\n", taskID)
				return
			}
		}

		fmt.Printf("Task %d not found\n", taskID)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}