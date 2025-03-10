package algorithms

// GetMean returns the mean of a given field from a slice of tasks
func GetMean(tasks *[]*Task, getFieldValue func(*Task) float64) float64 {
	var sum float64
	for _, task := range *tasks {
		sum += getFieldValue(task)
	}
	return sum / float64(len(*tasks))
}

// GetMeanProcessingTime returns the mean processing time of a slice of tasks
func GetMeanProcessingTime(tasks *[]*Task) float64 {
	return GetMean(
		tasks, func(task *Task) float64 {
			return float64(*task.processingTime)
		},
	)
}

// GetMeanWaitingTime returns the mean waiting time of a slice of tasks
func GetMeanWaitingTime(tasks *[]*Task) float64 {
	return GetMean(
		tasks, func(task *Task) float64 {
			return float64(*task.waitingTime)
		},
	)
}

// GetMeanServiceIndex returns the mean service index of a slice of tasks
func GetMeanServiceIndex(tasks *[]*Task) float64 {
	return GetMean(
		tasks, func(task *Task) float64 {
			return *task.serviceIndex
		},
	)
}
