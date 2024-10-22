package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"
	"tasks/tasks"

	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
)

var listAll bool

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		taskList, _ := tasks.LoadTasks("tasks.csv")

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)
		if listAll {
			fmt.Fprintln(w, "ID\tTask\tCreated\tDone")
			for _, task := range taskList {
				fmt.Fprintf(w, "%d\t%s\t%s\t%v\n", task.ID, task.Description, timediff.TimeDiff(task.CreatedAt), task.IsComplete)
			}
		} else {
			fmt.Fprintln(w, "ID\tTask\tCreated")
			for _, task := range taskList {
				if !task.IsComplete {
					fmt.Fprintf(w, "%d\t%s\t%s\n", task.ID, task.Description, timediff.TimeDiff(task.CreatedAt))
				}
			}
		}
		w.Flush()
	},
}

func init() {
	listCmd.Flags().BoolVarP(&listAll, "all", "a", false, "List all tasks, including completed ones")
	rootCmd.AddCommand(listCmd)
}