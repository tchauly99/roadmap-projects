package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
	"errors"
)

var JSON_FILE string = "tasks.json"

type TaskStatus string

const (
	STAT_TODO        	TaskStatus = "todo"
	STAT_IN_PROGRESS 	TaskStatus = "in-progress"
	STAT_TODO_DONE      TaskStatus = "done"
)

type Task struct {
	ID          int			`json:"id"`
	Title     	string		`json:"title"`
	Status 		TaskStatus	`json:"status,omitempty"`
	Created     time.Time	`json:"created"`
	Updated     time.Time	`json:"updated"`
}

func readTasks() ([]Task, error) {
	if _, err := os.Stat(JSON_FILE); errors.Is(err, os.ErrNotExist) {
		return []Task{}, nil
	}

	file, err := os.Open(JSON_FILE)
	if err != nil {
		fmt.Printf("Error open %v\n", JSON_FILE)
		return []Task{}, err
	}

	var tasks []Task
	dec := json.NewDecoder(file)
	err = dec.Decode(&tasks)
	if err != nil {
		fmt.Printf("Failed to decode JSON file %v\n", JSON_FILE)
		return []Task{}, err
	}

	return tasks, nil
}

func writeTasks(tasks []Task) error {
	if _, err := os.Stat(JSON_FILE); err == nil {
		err := os.Remove(JSON_FILE)
		if (err != nil) {
			fmt.Println("Can't remove old JSON file %v\n", JSON_FILE)
			return err
		}
	}

	file, err := os.OpenFile(JSON_FILE, os.O_CREATE|os.O_WRONLY, 0644)
	if (err != nil) {
		file.Close()
		return err
	}

	defer file.Close()

	enc := json.NewEncoder(file)
	enc.SetIndent("","  ")
	enc.Encode(tasks)
	if err != nil {
		fmt.Printf("Failed to encode data JSON file %v\n", JSON_FILE)
		return err
	}
	return nil
}

func addNewTask(description string) error {
	tasks, err := readTasks()
	if (err != nil) {
		fmt.Printf("Can't read current task list\n")
		return err
	}

	nexID := len(tasks)

	newTask := Task{nexID, description, STAT_TODO_DONE, time.Now(), time.Now()}
	tasks = append(tasks, newTask)

	fmt.Printf("Adding task %#v to task list\n", newTask.Title)

	return writeTasks(tasks)
}

func updateStatus(taskID int, status TaskStatus) error {
	tasks, err := readTasks()
	if (err != nil) {
		fmt.Printf("Can't read current task list\n")
		return err
	}

	var updated bool = false
	
	for i := range tasks {
		task := &tasks[i]
		if (taskID == task.ID) {
			task.Status = status
			task.Updated = time.Now()
			updated = true
			break
		}
	}

	if (!updated) {
		return errors.New("Task ID does not exist!")
	}

	fmt.Printf("Updating task %#v status\n", taskID)
	return writeTasks(tasks)
}

func updateTask(taskID int, description string) error {
	tasks, err := readTasks()
	if (err != nil) {
		fmt.Printf("Can't read current task list\n")
		return err
	}

	var updated bool = false
	
	for i := range tasks {
		task := &tasks[i]
		if (taskID == task.ID) {
			task.Title = description
			task.Updated = time.Now()
			updated = true
			break
		}
	}

	if (!updated) {
		return errors.New("Task ID does not exist!")
	}

	fmt.Printf("Updating task %#v description\n", taskID)
	return writeTasks(tasks)
}

func deleteTask(taskID int) error {
	tasks, err := readTasks()
	if (err != nil) {
		fmt.Printf("Can't read current task list\n")
		return err
	}

	var updated bool = false
	
	for i := range tasks {
		if (taskID == tasks[i].ID) {
			tasks = append(tasks[:i], tasks[i+1:]...)
			updated = true
			break
		}
	}

	if (!updated) {
		return errors.New("Task ID does not exist!")
	}

	// Reset task ID
	for i := range tasks {
		tasks[i].ID = i
	}

	fmt.Printf("Deleting task %#v\n", taskID)
	return writeTasks(tasks)
}