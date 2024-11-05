package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}
type ToDoList struct {
	Tasks  []*Task `json:"tasks"`
	NextID int     `json:"nextid"`
}

var todoList ToDoList

func main() {
	todoList = ToDoList{
		Tasks:  []*Task{},
		NextID: 1,
	}
	fmt.Println("Welcome to the To-Do List app!")

	loadFile()
	showmenu()

	os.Exit(0)
}

func showmenu() {
	fmt.Println("1. Add task")
	fmt.Println("2. View tasks")
	fmt.Println("3. Complete task")
	fmt.Println("4. Exit")
	fmt.Println("5. Delete task")

	var choice int
	fmt.Println("\nEnter your choice")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		addTask()
	case 2:
		viewTasks()
	case 3:
		completeTask()
	case 4:
		fmt.Println("Exiting...")
		//os.Exit(0)
	case 5:
		deleteTask()
	default:
		fmt.Println("Invalid choice, please try again.")
		showmenu()
	}
}

func addTask() {
	fmt.Println("Enter your task:")
	reader := bufio.NewReader(os.Stdin)
	taskName, _ := reader.ReadString('\n')
	taskName = strings.TrimSpace(taskName)
	var newTask = &Task{todoList.NextID, taskName, false}

	todoList.Tasks = append(todoList.Tasks, newTask)
	todoList.NextID++

	saveFile()

	showmenu()
}

func viewTasks() {
	fmt.Println("Welcome to your tasks list!")

	if len(todoList.Tasks) == 0 {
		fmt.Println("Uh-oh, there's no tasks!")
	} else {
		for _, el := range todoList.Tasks {
			fmt.Printf("ID: %d\n", el.ID)
			fmt.Printf("Task Name: %s\n", el.Name)
			fmt.Printf("Done: %t\n", el.Done)
		}

	}
	showmenu()
}

func completeTask() {
	var taskID int
	found := false
	fmt.Println("Which task do you want to complete?")
	fmt.Scanln(&taskID)

	for i := range todoList.Tasks {
		t := todoList.Tasks[i]
		if t.ID == taskID {
			todoList.Tasks[i].Done = true
			found = true
		}
	}
	if !found {
		fmt.Println("Sorry, we don't have such ID on the list!")
	}

	saveFile()

	showmenu()
}

func deleteTask() {
	var taskID int
	found := false
	fmt.Println("Which task do you want to delete?")
	fmt.Scan(&taskID)

	for _, el := range todoList.Tasks {
		if el.ID == taskID {
			found = true
			break
		}
	}
	if found {
		todoList.Tasks = append(todoList.Tasks[:taskID], todoList.Tasks[taskID+1:]...)
		saveFile()
	} else {
		fmt.Println("Sorry, we don't have such ID on the list!")
	}
}

func saveFile() {
	f, err := os.OpenFile("tasks.json", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Can't open the file")
	}
	defer f.Close()

	tasksToJSON, err := json.MarshalIndent(todoList, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling tasks")
	}
	errWriteFile := os.WriteFile("tasks.json", tasksToJSON, 0644)
	if errWriteFile != nil {
		fmt.Println("Error wrhting to the file")
	}

}

func loadFile() {
	jsonFile, err := os.ReadFile("tasks.json")
	if err != nil {
		fmt.Println("Error reading the file!")
	}
	if len(jsonFile) > 0 {
		errUnmar := json.Unmarshal(jsonFile, &todoList)
		if errUnmar != nil {
			fmt.Println("Error unmarshalling the file")
			return
		}
	}
}
func recalculateNextID() {

}
