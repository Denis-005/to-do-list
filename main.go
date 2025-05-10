package main

import (
	"bufio"
	"fmt"
	"os"
)

func input()

func main() {
	const (
		Create = "create"
		Read   = "read"
		Update = "update"
		Delete = "delete"
	)

	fmt.Println("Welcome to the TO DO List CLI app!")
	fmt.Println()

	slice := make([]string, 0)
	for {
		fmt.Println("Enter your command (create, read, update, delete):")

		var command string
		fmt.Scanln(&command)

		switch command {
		case Create:
			fmt.Println("Enter task name:")

			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			input := scanner.Text()

			slice = append(slice, input)

		case Read:
			for i := 0; i < len(slice); i++ {
				fmt.Printf("%d. %s\n", i+1, slice[i])
			}

		case Update:
			fmt.Println("Enter task name to update: ")

			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			currentInput := scanner.Text()

			matched := false
			for i := range slice {
				if currentInput == slice[i] {
					matched = true
					fmt.Println("Enter new task name:")

					scanner := bufio.NewScanner(os.Stdin)
					scanner.Scan()
					newInput := scanner.Text()

					if len(newInput) > 3 {
						slice[i] = newInput
						fmt.Printf("Updated task #%d with name \"%s\" successfully!\n", i, slice[i])
					} else {
						fmt.Println("The new task name is too short! Please, try again.")
					}
				}
			}

			if matched {
				_ = matched
			} else {
				fmt.Println("Invalid name. Please try again.")
			}

		case Delete:
			fmt.Println("Enter task name to remove: ")

			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			deleteInput := scanner.Text()

			index := -1
			for j := range slice {
				if deleteInput == slice[j] {
					index = j
				}
			}

			if index <= -1 {
				fmt.Println("Invalid name. Please try again.")
				continue
			}

			oldTaskName := slice[index]
			slice = append(slice[:index], slice[index+1:]...)
			fmt.Printf("Removed task #%d with name \"%s\" successfully!\n", index, oldTaskName)

		default:
			fmt.Println("Invalid command! Please, try again!")
		}
	}
}
