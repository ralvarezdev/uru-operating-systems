package algorithms

import (
	"github.com/ralvarezdev/uru-operating-systems/tasks/internal"
	"time"
)

// FIFO is the function to process the FIFO algorithm on a given list of tasks
func FIFO(unprocessedTasks *[]Task) []Task {
	currentTime := time.Duration(0) + internal.TimeUnit
	processedTasks := make([]Task, 0)
	hasBeenTasksProcessedOnCurrentIteration := false

	// Iterate over the tasks list
	for len(*unprocessedTasks) > 0 {
		for i, unprocessedTask := range *unprocessedTasks {
			if unprocessedTask.GetArrivalTime() <= currentTime {
				hasBeenTasksProcessedOnCurrentIteration = true

				// Update the current time
				currentTime = currentTime + unprocessedTask.duration

				// Set the end time of the task
				unprocessedTask.SetEndTime(currentTime)

				// Append the task to the processed tasks list
				processedTasks = append(processedTasks, unprocessedTask)

				// Remove the task from the unprocessed tasks list
				*unprocessedTasks = append(
					(*unprocessedTasks)[:i],
					(*unprocessedTasks)[i+1:]...,
				)
			} else {
				break
			}
		}

		// If no tasks have been processed on the current iteration, update the current time
		if !hasBeenTasksProcessedOnCurrentIteration {
			currentTime = currentTime + internal.TimeUnit
		} else {
			hasBeenTasksProcessedOnCurrentIteration = false
		}
	}

	return processedTasks
}
