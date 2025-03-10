package algorithms

import (
	"fmt"
	"time"
)

// Timing is the function to time it takes to process a list of tasks using a given algorithm
func Timing(
	unprocessedTasks *[]*Task,
	algorithm func(*[]*Task) *[]*Task,
) *[]*Task {
	// Start the timer
	start := time.Now()

	// Process the tasks using the given algorithm
	processedTasks := algorithm(unprocessedTasks)

	// Stop the timer
	elapsed := time.Since(start)

	// Print the time it took to process the tasks
	fmt.Printf("Time: %s\n", elapsed)

	return processedTasks
}
