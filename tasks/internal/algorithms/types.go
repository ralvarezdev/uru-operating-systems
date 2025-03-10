package algorithms

import (
	"strconv"
	"strings"
)

// Task is a struct that represents a task
type Task struct {
	id             string
	arrivalTime    int64
	duration       int64
	endTime        *int64
	processingTime *int64
	waitingTime    *int64
	serviceIndex   *float64
	timeLeft       *int64
}

// NewTask creates a new task
func NewTask(
	id string,
	arrivalTime int64,
	duration int64,
) *Task {
	return &Task{
		id:          id,
		arrivalTime: arrivalTime,
		duration:    duration,
		timeLeft:    &duration,
	}
}

// SetEndTime sets the end time of the task
func (t *Task) SetEndTime(endTime int64) {
	// Calculate the processing time, waiting time and service index
	processingTime := endTime - t.arrivalTime
	waitingTime := processingTime - t.duration
	serviceIndex := float64(t.duration) / float64(processingTime)

	// Set the end time, processing time, waiting time and service index
	t.timeLeft = nil
	t.endTime = &endTime
	t.processingTime = &processingTime
	t.waitingTime = &waitingTime
	t.serviceIndex = &serviceIndex
}

// SetTimeLeft sets the time left of the task
func (t *Task) SetTimeLeft(timeLeft int64) {
	t.timeLeft = &timeLeft
}

// GetID returns the task ID
func (t *Task) GetID() string {
	return t.id
}

// GetArrivalTime returns the arrival time of the task
func (t *Task) GetArrivalTime() int64 {
	return t.arrivalTime
}

// GetDuration returns the duration of the task
func (t *Task) GetDuration() int64 {
	return t.duration
}

// GetEndTime returns the end time of the task
func (t *Task) GetEndTime() *int64 {
	return t.endTime
}

// GetProcessingTime returns the processing time of the task
func (t *Task) GetProcessingTime() *int64 {
	return t.processingTime
}

// GetWaitingTime returns the waiting time of the task
func (t *Task) GetWaitingTime() *int64 {
	return t.waitingTime
}

// GetServiceIndex returns the service index of the task
func (t *Task) GetServiceIndex() *float64 {
	return t.serviceIndex
}

// GetTimeLeft returns the time left of the task
func (t *Task) GetTimeLeft() *int64 {
	return t.timeLeft
}

// String returns the string representation of the task
func (t *Task) String() string {
	return strings.Join(
		[]string{
			t.id,
			strconv.FormatInt(t.arrivalTime, 10),
			strconv.FormatInt(t.duration, 10),
			strconv.FormatInt(*t.endTime, 10),
			strconv.FormatInt(*t.processingTime, 10),
			strconv.FormatInt(*t.waitingTime, 10),
			strconv.FormatFloat(*t.serviceIndex, 'f', -1, 64),
		}, ",",
	)
}
