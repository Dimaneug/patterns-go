package main

import (
	"errors"
	"fmt"
)

type Assignee interface {
	CanHandleTask(task string) bool
	TakeTask(task string)
}

type Employee struct {
	isWorking bool
}

func (e *Employee) CanHandleTask(task string) bool {
	return !e.isWorking
}

func (e *Employee) TakeTask(task string) {
	if e.CanHandleTask(task) {
		e.isWorking = true
		fmt.Println("Employee took task " + "'" + task + "'")
	}
}

type Team struct {
	employees []*Employee
}

func (t *Team) Add(employee Employee) {
	t.employees = append(t.employees, &employee)
}

func (t *Team) Remove(employee Employee) {
	for i, val := range t.employees {
		if val == &employee {
			t.employees[i] = t.employees[len(t.employees)-1]
			t.employees = t.employees[:len(t.employees)-1]
			break
		}
	}
}

func (t *Team) CanHandleTask(task string) bool {
	for _, employee := range t.employees {
		if employee.CanHandleTask(task) {
			return true
		}
	}
	return false
}

func (t *Team) TakeTask(task string) {
	if !t.CanHandleTask(task) {
		return
	}
	for _, employee := range t.employees {
		if employee.CanHandleTask(task) {
			fmt.Print("Team: ")
			employee.TakeTask(task)
			break
		}
	}
}

type TaskManager struct {
	assignees []Assignee
}

func (tm *TaskManager) PerformTask(task string) error {
	for _, assignee := range tm.assignees {
		if assignee.CanHandleTask(task) {
			assignee.TakeTask(task)
			return nil
		}
	}
	return errors.New("cannot handle task")
}

func main() {
	employee1 := Employee{isWorking: true}
	employee2 := Employee{isWorking: false}
	employee3 := Employee{isWorking: false}
	employee4 := Employee{isWorking: true}

	team := Team{employees: []*Employee{&employee3, &employee4}}

	taskManager := TaskManager{[]Assignee{&employee1, &employee2, &team}}
	taskManager.PerformTask("Wash the dishes")
	taskManager.PerformTask("Buy milk")
}
