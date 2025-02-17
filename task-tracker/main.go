package main

import (
	"fmt"
	"encoding/json"
	"os"
)

func main() {
	var err error

	if _, err := os.Stat(JSON_FILE); err == nil {
		os.Remove(JSON_FILE)
	}

	err = addNewTask("Task 0")
	if (err != nil) {
		fmt.Printf("Error adding new task %v\n", err)
	}

	err = addNewTask("Task 1")
	if (err != nil) {
		fmt.Printf("Error adding new task %v\n", err)
	}

	err = addNewTask("Task 2")
	if (err != nil) {
		fmt.Printf("Error adding new task %v\n", err)
	}

	err = addNewTask("Task 3")
	if (err != nil) {
		fmt.Printf("Error adding new task %v\n", err)
	}

	tasks, err := readTasks()

	byteoutput, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Printf("Error listing tasks %v\n", err)
	}
	fmt.Printf("Current task list is:\n %v\n", string(byteoutput))

	err = updateTask(0, "New Status")
	if (err != nil) {
		fmt.Printf("Error updating task: %v\n", err)
	}

	err = updateStatus(3, STAT_IN_PROGRESS)
	if (err != nil) {
		fmt.Printf("Error updating task status: %v\n", err)
	}

	err = deleteTask(2)
	if (err != nil) {
		fmt.Printf("Error deleting task: %v\n", err)
	}

	tasks, err = readTasks()

	byteoutput, err = json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Printf("Error listing tasks %v\n", err)
	}
	fmt.Printf("Current task list is:\n %v\n", string(byteoutput))
}