package algorithms

// GetMean returns the mean of a given field from a slice of tasks
func GetMean(tasks *[]*Task, getFieldValue func(*Task) int64) float64 {
	var sum int64
	for _, task := range *tasks {
		sum += getFieldValue(task)
	}
	return float64(sum) / float64(len(*tasks))
}

// GetMeanProcessingTime returns the mean processing time of a slice of tasks
func GetMeanProcessingTime(tasks *[]*Task) float64 {
	return GetMean(
		tasks, func(task *Task) int64 {
			return *task.processingTime
		},
	)
}

// GetMeanWaitingTime returns the mean waiting time of a slice of tasks
func GetMeanWaitingTime(tasks *[]*Task) float64 {
	return GetMean(
		tasks, func(task *Task) int64 {
			return *task.waitingTime
		},
	)
}

// GetMeanServiceIndex returns the mean service index of a slice of tasks
func GetMeanServiceIndex(tasks *[]*Task) float64 {
	return GetMean(
		tasks, func(task *Task) int64 {
			return *task.serviceIndex
		},
	)
}
