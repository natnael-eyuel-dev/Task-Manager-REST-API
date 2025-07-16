package data

// imports 
import (
	"errors"
	"github.com/natnael-eyuel-dev/Task-Manager-REST-API/models"
)

type TaskManager interface {
	CreateTask(task *models.Task) error                                  // create new task with validation
	DeleteTask(taskID int) error                 					     // delete existing task or return error if not found
	GetAllTasks() ([]models.Task, error)         					     // get all tasks in the system
	GetTaskByID(taskID int) (*models.Task, error) 					     // get specific task by id or return error if not found
	UpdateTask(taskID int, task *models.Task) (*models.Task, error)      // update existing task or return error if not found
}

type InMemoryTaskManager struct {
	tasks  map[int]models.Task      // to store tasks with their id
	nextID int64                    // to generate unique ids for new tasks
}

func NewInMemoryTaskManager() *InMemoryTaskManager {
	return &InMemoryTaskManager{
		tasks:  make(map[int]models.Task),  // initialize map for tasks
		nextID: 1,                          // start id counter from 1
	}
}

func (taskmang *InMemoryTaskManager) CreateTask(task *models.Task) error {
	
	// validate task fields before creation
	if task.Title == "" {
		return errors.New("task title can not be empty")
	}
	if task.Description == "" {
		return errors.New("task description can not be empty")
	}
	if task.DueDate.IsZero() {
		return errors.New("task duedate can not be empty")
	}
	if task.Status == "" {
		return errors.New("task status can not be empty")
	}

	// assign new id and save the task
	newID := taskmang.nextID
	taskmang.nextID++
	task.ID = int(newID)
	taskmang.tasks[task.ID] = *task

	return nil
}

func (taskmang *InMemoryTaskManager) DeleteTask(taskID int) error {
	_, exists := taskmang.tasks[taskID]    // check if task exists
	if !exists {
		return errors.New("no task found with this id to delete")
	}

	delete(taskmang.tasks, taskID)     // remove task from storage
	return nil
}

func (taskmang *InMemoryTaskManager) GetAllTasks() ([]models.Task, error) {
	var allTasks []models.Task
	for _, task := range taskmang.tasks {
		allTasks = append(allTasks, task)     // collect all tasks
	}

	return allTasks, nil     // return all tasks in the system
}

func (taskmang *InMemoryTaskManager) GetTaskByID(taskID int) (*models.Task, error) {
	task, exists := taskmang.tasks[taskID]     // check if task exists
	if !exists {
		return nil, errors.New("no task found with this id to see")
	}

	return &task, nil    // return the found task
}

func (taskmang *InMemoryTaskManager) UpdateTask(taskID int, taskUpdate *models.Task) (*models.Task, error) {
	task, exists := taskmang.tasks[taskID]  // check if task exists
	if !exists {
		return nil, errors.New("no task found with this id to update")
	}

	// update task fields if they are provided in the update
	if taskUpdate.Title != "" {
		task.Title = taskUpdate.Title
	}
	if taskUpdate.Description != "" {
		task.Description = taskUpdate.Description
	}
	if !taskUpdate.DueDate.IsZero() {
		task.DueDate = taskUpdate.DueDate
	}
	if taskUpdate.Status != "" {
		task.Status = taskUpdate.Status
	}

	taskmang.tasks[taskID] = task  // save updated task
	
	return &task, nil  // return the updated task
}