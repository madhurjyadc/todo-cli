package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
)

type Task struct {
	title string
	status bool
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
			tasks = append(tasks, Task{title: task_input, status: false})
			fmt.Println("Task successfully appended!")


		case "2\n":
			fmt.Println("\nThe following are your current tasks: ")
			for i, task := range tasks {

				if task.status {
					fmt.Printf("%d. %s ✅", i+1, task.title)
				}

				fmt.Printf("%d. %s ❌", i+1, task.title)
			}

			fmt.Print("\nSelect which task you want to modify: ")
			reader := bufio.NewReader((os.Stdin))
			taskUpdateValue, _ := reader.ReadString('\n')

			taskUpdateValue = strings.TrimSpace(taskUpdateValue)

			taskNumberInput, err := strconv.Atoi(taskUpdateValue)

			if err != nil {
				fmt.Println("Please enter a valid number.")
				break
			}

			if taskNumberInput - 1 < 0 || taskNumberInput+1 >= len(tasks) {
				fmt.Println("Task does not exist.")
			}


			fmt.Print("\nEnter the new title: ")

			newTitle, _ := reader.ReadString('\n')
			newTitle = strings.TrimSpace(newTitle)

			tasks[taskNumberInput - 1].title = newTitle

			fmt.Println("Task updated successfully!")


		case "3\n":

			fmt.Println("\nDisplaying all the tasks: \n")
			for i, task := range tasks {

				if task.status {
					fmt.Printf("%d. %s ✅", i+1, task.title)
				}

				fmt.Printf("%d. %s ❌", i+1, task.title)
			}
		

		case "4\n":
			fmt.Println("\nThe following are your current tasks: ")
			for i, task := range tasks {

				if task.status {
					fmt.Printf("%d. %s ✅", i+1, task.title)
				}

				fmt.Printf("%d. %s ❌", i+1, task.title)
			}

			fmt.Print("\nSelect which task's status you want to update: ")
			reader := bufio.NewReader((os.Stdin))
			taskUpdateValue, _ := reader.ReadString('\n')
			taskUpdateValue = strings.TrimSpace(taskUpdateValue)
			taskNumberInput, err := strconv.Atoi(taskUpdateValue)

			if err != nil {
				fmt.Println("Please enter a valid number.")
				break
			}

			if taskNumberInput - 1 < 0 || taskNumberInput+1 >= len(tasks) {
				fmt.Println("Task does not exist.")
			}

			


		}
		
	}

	
}