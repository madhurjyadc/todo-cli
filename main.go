package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
	"slices"
)

type Task struct {
	title string
	status bool
}

func readTasks(s []Task) {
	fmt.Println("\nThe following are your current tasks: ")
	for i, task := range s {

		if task.status {
			fmt.Printf("%d. %s ✅\n", i+1, task.title)
		} else {

			fmt.Printf("%d. %s ❌\n", i+1, task.title)
		}

		
	}

	fmt.Println()
}

func modifyTasks(tasks []Task, s string) int {

	fmt.Printf("\n%s", s)
	reader := bufio.NewReader((os.Stdin))
	taskUpdateValue, _ := reader.ReadString('\n')

	taskUpdateValue = strings.TrimSpace(taskUpdateValue)

	taskNumberInput, err := strconv.Atoi(taskUpdateValue)

	if err != nil {
		fmt.Println("Please enter a valid number.")
		return -1
	}

	if taskNumberInput - 1 < 0 || taskNumberInput - 1 >= len(tasks) {
		fmt.Println("Task does not exist.")
		return -1
	}

	return taskNumberInput-1
}

func main() {
	
	var tasks []Task

	for {
		fmt.Println("1. Create a new task")
		fmt.Println("2. Update an existing task")
		fmt.Println("3. Read all the tasks")
		fmt.Println("4. Mark a task as done")
		fmt.Println("5. Delete a task")
	
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter your choice: ")
		input, _ := reader.ReadString('\n')

		if text := strings.ToLower(input); text == "quit\n" {
			break
		} 

		switch input {
		case "1\n":
			fmt.Print("\nEnter the task: ")
			reader := bufio.NewReader((os.Stdin))
			task_input, _ := reader.ReadString('\n')
			task_input = strings.TrimSpace(task_input)
			tasks = append(tasks, Task{title: task_input, status: false})
			fmt.Println("Task successfully appended!")


		case "2\n":
			readTasks(tasks)
			index := modifyTasks(tasks, "Select which task you want to modify: ")
			
			fmt.Print("\nEnter the new title: ")
			newTitle, _ := reader.ReadString('\n')
			newTitle = strings.TrimSpace(newTitle)

			if index == -1 {
    			continue
			}

			tasks[index].title = newTitle

			

			fmt.Println("Task updated successfully!")


		case "3\n":

			readTasks(tasks)
		

		case "4\n":
			readTasks(tasks)

			index := modifyTasks(tasks, "Select which task you want to mark as done: ")
			
			if index == -1 {
    			continue
			}


			tasks[index].status = true
			fmt.Println("\nTask status updated successfully.")


		case "5\n":
			readTasks(tasks)

			index := modifyTasks(tasks, "Select which task you want to delete: ")

			if index == -1 {
    			continue
			}
			tasks = slices.Delete(tasks, index, index+1)
			fmt.Println("Task deleted successfully.")
		}
	}

}