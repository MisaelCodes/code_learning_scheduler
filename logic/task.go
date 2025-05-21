package logic

import (
	"fmt"
	"reflect"
)

type Schedule struct {
	hour   int
	minute int
	second int
}

func NewSchedule(hour, minute, second int) (*Schedule, error) {
	if hour > 23 || hour < 0 {
		return nil, fmt.Errorf("schedule: 'hour' should be between 0 and 23, %d not in range", hour)
	}

	if minute > 59 || minute < 0 {
		return nil, fmt.Errorf("schedule: 'minute' should be between 0 and 59, %d not in range", minute)
	}

	if second > 59 || second < 0 {
		return nil, fmt.Errorf("schedule: 'second' should be between 0 and 59, %d not in range", second)
	}

	return &Schedule{hour, minute, second}, nil
}

type Task struct {
	taskList    *TaskList
	title       string
	description string
	schedule    *Schedule
	action      *Action
}

// Creates a new tasks, parameter t is an array with:
// hours, minutes and seconds
func NewTask(title, description string, schedule *Schedule, action *Action, taskList *TaskList) *Task {
	return &Task{taskList, title, description, schedule, action}
}

type TimePeriod int

const (
	OneTime TimePeriod = iota
	Daily
	Weekly
	BiWeekly
	Monthly
)

type TaskList struct {
	title       string
	location    string
	periodicity TimePeriod
	tasks       *[]*Task
}

// Returns n x 2 matrix, where each subarray
// has the name and the type of the field.
func (tl TaskList) Fields() ([][2]string, error) {
	rtl := reflect.TypeOf(tl)
    n := rtl.NumField()
    fields := make([][2]string, n)
	for i := range n {
        f := rtl.Field(i)
		fields[i][0],fields[i][1] = f.Name , f.Type.Name()
        
	}
	return fields, nil
}
