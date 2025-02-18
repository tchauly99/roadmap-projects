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

func listByStatus(status jsonhandle.TaskStatus) {
	tasks, err := jsonhandle.ReadTasks()
	if (err != nil) {
		panic(err)
	}
	for i := range tasks {
		task := tasks[i]
		if (task.Status == status) {
			fmt.Printf("%v: %#v\n", task.ID, task.Title)
		}
	}
}

func ListCmd() *cobra.Command {
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List tasks",
		Long:  `This command list out all tasks, or based on status.
Subcommands: todo/in-progress/done`,
		Args: cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			tasks, err := jsonhandle.ReadTasks()
			if (err != nil) {
				panic(err)
			}
			for i := range tasks {
				task := tasks[i]
				if ((len(args) > 0) && (string(task.Status) != args[0])) {
					continue
				}
				fmt.Printf("%v: %#v - %v\n", task.ID, task.Title, strings.ToUpper(string(task.Status)))
			}
		},
	}

	listTodoCmd := &cobra.Command{
		Use:   "todo",
		Short: "List tasks that are todo",
		Long:  `This command list out all tasks that are todo`,
		Args: cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			listByStatus(jsonhandle.STAT_TODO)
		},
	}

	listInProgressCmd := &cobra.Command{
		Use:   "in-progress",
		Short: "List tasks that are in-progress",
		Long:  `This command list out all tasks that are in-progress`,
		Args: cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			listByStatus(jsonhandle.STAT_IN_PROGRESS)
		},
	}

	listDoneCmd := &cobra.Command{
		Use:   "done",
		Short: "List tasks that are done",
		Long:  `This command list out all tasks that are done`,
		Args: cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			listByStatus(jsonhandle.STAT_DONE)
		},
	}

	listCmd.AddCommand(listTodoCmd)
	listCmd.AddCommand(listInProgressCmd)
	listCmd.AddCommand(listDoneCmd)

	return listCmd
}

func AddCmd() *cobra.Command {
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
	return addCmd
}

func UpdateCmd() *cobra.Command {
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
	return updateCmd
}

func DeleteCmd() *cobra.Command {
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
	return deleteCmd
}

func MarkTodoCmd() *cobra.Command {
	markTodoCmd := &cobra.Command{
		Use:   "mark-in-progress",
		Short: "Mark tasks as in-progress",
		Long:  "This command mark task status as in-progress based on task IDs",
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			for i := range args {
				taskID := argToTaskID(args[i])
				err := jsonhandle.UpdateStatus(taskID, jsonhandle.STAT_TODO)
				if (err != nil) {
					panic(err)
				}
			}
		},
	}
	return markTodoCmd
}

func MarkInProgressCmd() *cobra.Command {
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
	return markInProgressCmd
}

func MarkDoneCmd() *cobra.Command {
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
	return markDoneCmd
}

func RootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "task-cli",
		Short: "Task Command Line Interface",
		Long:  "This is a CLI command to handle task list",
	}

	rootCmd.AddCommand(ListCmd())
	rootCmd.AddCommand(AddCmd())
	rootCmd.AddCommand(DeleteCmd())
	rootCmd.AddCommand(UpdateCmd())
	rootCmd.AddCommand(MarkTodoCmd())
	rootCmd.AddCommand(MarkInProgressCmd())
	rootCmd.AddCommand(MarkDoneCmd())

	return rootCmd
}


