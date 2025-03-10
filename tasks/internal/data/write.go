package data

import (
	"fmt"
	internalalgorithms "github.com/ralvarezdev/uru-operating-systems/tasks/internal/algorithms"
	"os"
)

// WriteTasks writes the tasks to the file system
func WriteTasks(path string, tasks *[]*internalalgorithms.Task) error {
	// Create the file
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	// Write headers to the file
	_, err = file.WriteString("id,ti,t,tf,T,E,I\n")

	// Write the tasks to the file
	for _, task := range *tasks {
		_, err = file.WriteString(task.String() + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}
