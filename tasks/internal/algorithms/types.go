package algorithms

import (
	"time"
)

// Task is a struct that represents a task
type Task struct {
	id             string
	arrivalTime    time.Duration
	duration       time.Duration
	endTime        *time.Duration
	processingTime *time.Duration
	waitingTime    *time.Duration
	serviceIndex   *time.Duration
	timeLeft       *time.Duration
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

// SetEndTime sets the end time of the task
func (t *Task) SetEndTime(endTime time.Duration) {
	// Calculate the processing time, waiting time and service index
	processingTime := endTime - t.arrivalTime
	waitingTime := processingTime - t.duration
	serviceIndex := t.duration / processingTime

	// Set the end time, processing time, waiting time and service index
	t.endTime = &endTime
	t.processingTime = &processingTime
	t.waitingTime = &waitingTime
	t.serviceIndex = &serviceIndex
}

// SetTimeLeft sets the time left of the task
func (t *Task) SetTimeLeft(timeLeft time.Duration) {
	t.timeLeft = &timeLeft
}

// GetID returns the task ID
func (t *Task) GetID() string {
	return t.id
}

// GetArrivalTime returns the arrival time of the task
func (t *Task) GetArrivalTime() time.Duration {
	return t.arrivalTime
}

// GetDuration returns the duration of the task
func (t *Task) GetDuration() time.Duration {
	return t.duration
}

// GetEndTime returns the end time of the task
func (t *Task) GetEndTime() *time.Duration {
	return t.endTime
}

// GetProcessingTime returns the processing time of the task
func (t *Task) GetProcessingTime() *time.Duration {
	return t.processingTime
}

// GetWaitingTime returns the waiting time of the task
func (t *Task) GetWaitingTime() *time.Duration {
	return t.waitingTime
}

// GetServiceIndex returns the service index of the task
func (t *Task) GetServiceIndex() *time.Duration {
	return t.serviceIndex
}

// GetTimeLeft returns the time left of the task
func (t *Task) GetTimeLeft() *time.Duration {
	return t.timeLeft
}
