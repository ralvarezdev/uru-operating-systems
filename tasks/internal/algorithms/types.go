package algorithms

import (
	"time"
)

// Task is a struct that represents a task
type Task struct {
	id           string
	arrivalTime  time.Duration
	duration     time.Duration
	finalTime    *time.Duration
	totalTime    *time.Duration
	waitingTime  *time.Duration
	serviceIndex *time.Duration
}

// NewTask creates a new task
func NewTask(
	id string,
	arrivalTime time.Duration,
	duration time.Duration,
) *Task {
	return &Task{
		id:          id,
		arrivalTime: arrivalTime,
		duration:    duration,
	}
}

// SetFinalTime sets the final time of the task
func (t *Task) SetFinalTime(finalTime time.Duration) {
	t.finalTime = &finalTime
}

// SetTotalTime sets the total time of the task
func (t *Task) SetTotalTime(totalTime time.Duration) {
	t.totalTime = &totalTime
}

// SetWaitingTime sets the waiting time of the task
func (t *Task) SetWaitingTime(waitingTime time.Duration) {
	t.waitingTime = &waitingTime
}

// SetServiceIndex sets the service index of the task
func (t *Task) SetServiceIndex(serviceIndex time.Duration) {
	t.serviceIndex = &serviceIndex
}

// GetID returns the task ID
func (t *Task) GetID() string {
	return t.id
}

// GetInitialTime returns the arrival time of the task
func (t *Task) GetInitialTime() time.Duration {
	return t.arrivalTime
}

// GetDuration returns the duration of the task
func (t *Task) GetDuration() time.Duration {
	return t.duration
}

// GetFinalTime returns the final time of the task
func (t *Task) GetFinalTime() *time.Duration {
	return t.finalTime
}

// GetTotalTime returns the total time of the task
func (t *Task) GetTotalTime() *time.Duration {
	return t.totalTime
}

// GetWaitingTime returns the waiting time of the task
func (t *Task) GetWaitingTime() *time.Duration {
	return t.waitingTime
}

// GetServiceIndex returns the service index of the task
func (t *Task) GetServiceIndex() *time.Duration {
	return t.serviceIndex
}
