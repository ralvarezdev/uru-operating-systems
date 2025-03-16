package algorithms

import (
	"github.com/ralvarezdev/uru-operating-systems/tasks/internal"
	"time"
)

// RoundRobin is a simple round-robin algorithm that selects the next task in the list of tasks
func RoundRobin(quantum int64) func(*[]*Task) *[]*Task {
	return func(unprocessedTasks *[]*Task) *[]*Task {
		var currentTime int64 = 0
		processedTasks := make([]*Task, 0)
		hasBeenTasksProcessedOnCurrentIteration := false

		// Iterate over the tasks list
		for len(*unprocessedTasks) > 0 {
			for i := 0; i < len(*unprocessedTasks); {
				// Get the task
				unprocessedTask := (*unprocessedTasks)[i]

				if unprocessedTask.GetArrivalTime() > currentTime {
					i++
				} else {
					hasBeenTasksProcessedOnCurrentIteration = true

					if *unprocessedTask.GetTimeLeft() > quantum {
						// Update the current time
						currentTime += quantum

						// Set the time left of the task
						unprocessedTask.SetTimeLeft(*unprocessedTask.GetTimeLeft() - quantum)
						i++

						// Sleep
						time.Sleep(time.Duration(quantum) * internal.TimeUnit)
					} else {
						// Update the current time
						currentTime += *unprocessedTask.GetTimeLeft()

						// Sleep
						time.Sleep(time.Duration(*unprocessedTask.GetTimeLeft()) * internal.TimeUnit)

						// Set the end time of the task
						unprocessedTask.SetEndTime(currentTime)

						// Append the task to the processed tasks list
						processedTasks = append(processedTasks, unprocessedTask)

						// Remove the task from the unprocessed tasks list
						*unprocessedTasks = append(
							(*unprocessedTasks)[:i],
							(*unprocessedTasks)[i+1:]...,
						)
					}
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
}
