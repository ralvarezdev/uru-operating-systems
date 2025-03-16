package algorithms

// GetAverage returns the average of a given field from a slice of tasks
func GetAverage(tasks *[]*Task, getFieldValue func(*Task) float64) float64 {
	var sum float64
	for _, task := range *tasks {
		sum += getFieldValue(task)
	}
	return sum / float64(len(*tasks))
}

// GetAverageProcessingTime returns the average processing time of a slice of tasks
func GetAverageProcessingTime(tasks *[]*Task) float64 {
	return GetAverage(
		tasks, func(task *Task) float64 {
			return float64(*task.processingTime)
		},
	)
}

// GetAverageWaitingTime returns the average waiting time of a slice of tasks
func GetAverageWaitingTime(tasks *[]*Task) float64 {
	return GetAverage(
		tasks, func(task *Task) float64 {
			return float64(*task.waitingTime)
		},
	)
}

// GetAverageServiceIndex returns the average service index of a slice of tasks
func GetAverageServiceIndex(tasks *[]*Task) float64 {
	return GetAverage(
		tasks, func(task *Task) float64 {
			return *task.serviceIndex
		},
	)
}
