package cmd

import (
	"github.com/spf13/cobra"
	"tasktracker/jsonhandle"
	"strconv"
	"fmt"
	"strings"
)

func argToTaskID(argStr string) int {
	taskID, err := strconv.Atoi(argStr)
    if err != nil {
        panic(err)
    }
	return taskID
}

func RootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "task-cli",
		Short: "Task Command Line Interface",
		Long:  "This is a CLI command to handle task list",
	}

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List tasks",
		Long:  "This command list out all tasks",
		Args: cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			tasks, err := jsonhandle.ReadTasks()
			if (err != nil) {
				panic(err)
			}
			for i := range tasks {
				task := tasks[i]
				fmt.Printf("%v: %#v - %v\n", task.ID, task.Title, strings.ToUpper(string(task.Status)))
			}
		},
	}

	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add tasks",
		Long:  "This command adds tasks to the task list with descriptions",
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			for i := range args {
				err := jsonhandle.AddNewTask(args[i])
				if (err != nil) {
					panic(err)
				}
			}
		},
	}

	deleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete tasks",
		Long:  "This command delete tasks based on task IDs",
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			for i := range args {
				taskID := argToTaskID(args[i])
				err := jsonhandle.DeleteTask(taskID)
				if (err != nil) {
					panic(err)
				}
			}
		},
	}

	updateCmd := &cobra.Command{
		Use:   "update",
		Short: "Update task description",
		Long:  "This command update task description based on task ID",
		Args: cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			taskID := argToTaskID(args[0])
			err := jsonhandle.UpdateTask(taskID, args[1])
			if (err != nil) {
				panic(err)
			}
		},
	}

	markInProgressCmd := &cobra.Command{
		Use:   "mark-in-progress",
		Short: "Mark tasks as in-progress",
		Long:  "This command mark task status as in-progress based on task IDs",
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			for i := range args {
				taskID := argToTaskID(args[i])
				err := jsonhandle.UpdateStatus(taskID, jsonhandle.STAT_IN_PROGRESS)
				if (err != nil) {
					panic(err)
				}
			}
		},
	}

	markDoneCmd := &cobra.Command{
		Use:   "mark-done",
		Short: "Mark tasks as done",
		Long:  "This command mark task status as done based on task IDs",
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			for i := range args {
				taskID := argToTaskID(args[i])
				err := jsonhandle.UpdateStatus(taskID, jsonhandle.STAT_DONE)
				if (err != nil) {
					panic(err)
				}
			}
		},
	}

	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(markInProgressCmd)
	rootCmd.AddCommand(markDoneCmd)

	return rootCmd
}


