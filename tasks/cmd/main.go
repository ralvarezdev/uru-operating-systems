package main

import (
	"fmt"
	internalalgorithms "github.com/ralvarezdev/uru-operating-systems/tasks/internal/algorithms"
	internaldata "github.com/ralvarezdev/uru-operating-systems/tasks/internal/data"
	"os"
	"path/filepath"
	"strconv"
)

const (
	Menu = `--- Options ---
1. FIFO
2. LIFO
3. Round Robin
4. Exit

Select an option: `
)

var (
	DataFolder = "data"
)

func main() {
	// Get the current directory
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}
	basePath := filepath.Join(cwd, DataFolder)

	// Print the menu
	fmt.Print(Menu)

	for {
		// Get the option
		var option string
		optionLen, err := fmt.Scanln(&option)
		if err != nil {
			fmt.Println("Error reading option:", err)
			return
		}

		// Check if the option is empty
		if optionLen == 0 {
			fmt.Println("No option selected")
			return
		}

		// Parse the option
		optionInt, err := strconv.Atoi(option)
		if err != nil {
			fmt.Println("Invalid option:", option)
			return
		}

		// Process the option
		switch optionInt {
		case 1:
		case 2:
		case 3:
			// Ask for the file path
			fmt.Print("Enter the filename: ")
			var filename string
			filenameLen, err := fmt.Scanln(&filename)
			if err != nil {
				fmt.Println("Error reading the filename:", err)
				return
			}

			// Check if the filename is empty
			if filenameLen == 0 {
				fmt.Println("No filename entered")
				return
			}

			// Load the tasks
			tasks, err := internaldata.LoadTasks(
				filepath.Join(
					basePath,
					filename,
				),
			)
			if err != nil {
				fmt.Println("Error loading tasks:", err)
				return
			}

			// Process the tasks
			var processedTasks *[]*internalalgorithms.Task
			switch optionInt {
			case 1:
				// FIFO
				processedTasks = internalalgorithms.Timing(
					tasks,
					internalalgorithms.FIFO,
				)
			case 2:
				// LIFO
				processedTasks = internalalgorithms.Timing(
					tasks,
					internalalgorithms.LIFO,
				)
			case 3:
				// Ask for the quantum
				fmt.Print("Enter the quantum: ")
				var quantum string
				quantumLen, err := fmt.Scanln(&quantum)
				if err != nil {
					fmt.Println("Error reading the quantum:", err)
					return
				}

				// Check if the quantum is empty
				if quantumLen == 0 {
					fmt.Println("No quantum entered")
					return
				}

				// Parse the quantum
				quantumInt, err := strconv.Atoi(quantum)
				if err != nil {
					fmt.Println("Invalid quantum:", quantum)
					return
				}

				// Round Robin
				roundRobinFn := internalalgorithms.RoundRobin(int64(quantumInt))
				processedTasks = internalalgorithms.Timing(tasks, roundRobinFn)
			}

			// Write the processed tasks
			err = internaldata.WriteTasks(
				filepath.Join(
					basePath,
					"processed_"+filename,
				), processedTasks,
			)
			if err != nil {
				fmt.Println("Error writing processed tasks:", err)
				return
			}

			// Print the fields statistics
			fmt.Println(
				"Mean waiting time:",
				internalalgorithms.GetMeanWaitingTime(processedTasks),
			)
			fmt.Println(
				"Mean processing time:",
				internalalgorithms.GetMeanProcessingTime(processedTasks),
			)
			fmt.Println(
				"Mean service index:",
				internalalgorithms.GetMeanServiceIndex(processedTasks),
			)
		case 4:
			// Exit
			return
		default:
			// Invalid option
			fmt.Println("Invalid option")
		}
	}
}
