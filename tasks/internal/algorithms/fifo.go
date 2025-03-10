package algorithms

import (
	"github.com/ralvarezdev/uru-operating-systems/tasks/internal"
	"time"
)

// FIFO is the function to process the FIFO algorithm on a given list of tasks
func FIFO(unprocessedTasks *[]*Task) *[]*Task {
	var currentTime int64 = 0
	processedTasks := make([]*Task, 0)
	hasBeenTasksProcessedOnCurrentIteration := false

	// Iterate over the tasks list
	for len(*unprocessedTasks) > 0 {
		for i := 0; i < len(*unprocessedTasks); i++ {
			// Get the task
			unprocessedTask := (*unprocessedTasks)[i]

			if unprocessedTask.GetArrivalTime() <= currentTime {
				hasBeenTasksProcessedOnCurrentIteration = true

				// Update the current time
				currentTime += unprocessedTask.GetDuration()

				// Set the end time of the task
				unprocessedTask.SetEndTime(currentTime)

				// Append the task to the processed tasks list
				processedTasks = append(processedTasks, unprocessedTask)

				// Remove the task from the unprocessed tasks list
				*unprocessedTasks = append(
					(*unprocessedTasks)[:i],
					(*unprocessedTasks)[i+1:]...,
				)

				// Sleep
				time.Sleep(time.Duration(unprocessedTask.GetDuration()) * internal.TimeUnit)
				break
			}
		}

		// If no tasks have been processed on the current iteration, update the current time
		if !hasBeenTasksProcessedOnCurrentIteration {
			currentTime++

			// Sleep
			time.Sleep(internal.TimeUnit)
		} else {
			hasBeenTasksProcessedOnCurrentIteration = false
		}
	}

	return &processedTasks
}
