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

func NewTaskList(title, location string, periodicity TimePeriod, tasks *[]*Task) *TaskList {
	return &TaskList{title, location, periodicity, tasks}
}

func (tl TaskList) Fields() ([][2]string, error) {
	rtl := reflect.TypeOf(tl)
	fmt.Println(rtl.NumField())
	for i := range rtl.NumField() {
		fmt.Println(rtl.Field(i).Name)
	}
	return nil, fmt.Errorf("error processing: %d", 1)

}
