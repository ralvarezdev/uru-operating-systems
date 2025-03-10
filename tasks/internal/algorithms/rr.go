package algorithms

import (
	"github.com/ralvarezdev/uru-operating-systems/tasks/internal"
	"time"
)

// RoundRobin is a simple round-robin algorithm that selects the next task in the list of tasks
func RoundRobin(quantum int64) func(*[]*Task) *[]*Task {
	return func(tasks *[]*Task) *[]*Task {
		var currentTime int64 = 0
		processedTasks := make([]*Task, 0)
		hasBeenTasksProcessedOnCurrentIteration := false

		// Iterate over the tasks list
		for len(*tasks) > 0 {
			for i, task := range *tasks {
				// Get the task
				if task.GetArrivalTime() <= currentTime {
					hasBeenTasksProcessedOnCurrentIteration = true

					if *task.GetTimeLeft() > quantum {
						// Update the current time
						currentTime += quantum

						// Set the time left of the task
						task.SetTimeLeft(*task.GetTimeLeft() - quantum)
					} else {
						// Update the current time
						currentTime = currentTime + task.GetDuration()

						// Set the end time of the task
						task.SetEndTime(currentTime)

						// Append the task to the processed tasks list
						processedTasks = append(processedTasks, task)

						// Remove the task from the unprocessed tasks list
						*tasks = append(
							(*tasks)[:i],
							(*tasks)[i+1:]...,
						)

						// Sleep
						time.Sleep(time.Duration(task.GetDuration()) * internal.TimeUnit)
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
