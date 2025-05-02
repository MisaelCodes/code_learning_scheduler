package logic

type Schedule struct {
	hour   int
	minute int
	second int
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
    title string
	location string
    periodicity TimePeriod 
	tasks    *[]*Task
}

func NewTaskList(title, location string, periodicity TimePeriod, tasks *[]*Task) *TaskList{
    return &TaskList{title, location, periodicity, tasks}
}
