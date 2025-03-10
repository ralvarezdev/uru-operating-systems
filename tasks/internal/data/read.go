package data

import (
	"fmt"
	goloaderfs "github.com/ralvarezdev/go-loader/filesystem"
	gostringsconvert "github.com/ralvarezdev/go-strings/convert"
	"github.com/ralvarezdev/uru-operating-systems/tasks/internal"
	internalalgorithms "github.com/ralvarezdev/uru-operating-systems/tasks/internal/algorithms"
)

// LoadTasks reads the tasks from the file and returns them
func LoadTasks(path string) (*[]*internalalgorithms.Task, error) {
	// Open the file
	file, err := goloaderfs.OpenFile(path)
	if err != nil {
		return nil, err
	}

	// Read the CSV file
	records, _, err := goloaderfs.ReadCSVFile(file, true)
	if err != nil {
		return nil, err
	}

	// Iterate over the records
	tasks := make([]*internalalgorithms.Task, 0)
	for i, record := range *records {
		var id string
		var arrivalTime, duration int64

		// Get the fields data as strings
		if len(record) < internal.RecordFieldsNumber {
			return nil, fmt.Errorf(
				ErrInvalidRecordFields,
				i,
				internal.RecordFieldsNumber,
			)
		}

		// Parse the ID
		id = record[internal.IdIndex]

		// Parse the arrival time and duration
		if err = gostringsconvert.ToInt64(
			record[internal.ArrivalTimeIndex],
			&arrivalTime,
		); err != nil {
			return nil, fmt.Errorf(ErrInvalidArrivalTime, i)
		}
		if err = gostringsconvert.ToInt64(
			record[internal.DurationIndex],
			&duration,
		); err != nil {
			return nil, fmt.Errorf(ErrInvalidDuration, i)
		}

		// Create the task
		task := internalalgorithms.NewTask(
			id,
			arrivalTime,
			duration,
		)
		tasks = append(tasks, task)
	}

	return &tasks, nil
}
