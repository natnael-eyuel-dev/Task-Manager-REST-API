package controllers

// imports
import (
	"net/http"
	"strconv"
	"strings"
	"github.com/natnael-eyuel-dev/Task-Manager-REST-API/data"
	"github.com/natnael-eyuel-dev/Task-Manager-REST-API/models"
	"github.com/gin-gonic/gin"
)

type TaskController struct {
	taskService data.TaskManager       // service layer for task operations
}

func NewTaskController(service data.TaskManager) *TaskController {
	return &TaskController{taskService: service}         // return new controller instance 
}

func (taskcontr *TaskController) CreateTask(c *gin.Context) {
	
	var task models.Task
	err := c.ShouldBindJSON(&task)    // parse request body into task struct
	if err != nil {
		// handle specific date format error case
		if strings.Contains(err.Error(), "numeric literal") {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid date format. Use ISO 8601 format like '2023-12-31T00:00:00Z'",
				"example": gin.H{
					"due_date": "2023-12-31T00:00:00Z",
				},
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create task through service layer
	err = taskcontr.taskService.CreateTask(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, task)  // return created task with 201 status
}

func (taskcontr *TaskController) DeleteTask(c *gin.Context) {
	
	id, err := strconv.Atoi(c.Param("id"))         // convert string id to integer
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	// delete task through service layer
	err = taskcontr.taskService.DeleteTask(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message":"task deleted successfully"})    // success response
}

func (taskcontr *TaskController) GetAllTasks(c *gin.Context) {
	
	// get all tasks through service layer
	tasks, err := taskcontr.taskService.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)     // return all tasks
}

func (taskcontr *TaskController) GetTaskByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))       // convert string id to integer
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	// get specific task through service layer
	task, err := taskcontr.taskService.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if task == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, task)      // return found task
}

func (taskcontr *TaskController) UpdateTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))     // convert string id to integer
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var taskUpdate models.Task
	err = c.ShouldBindJSON(&taskUpdate)    // parse request body into task struct
	if err != nil {
		// handle specific date format error case
		if strings.Contains(err.Error(), "numeric literal") {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid date format. Use ISO 8601 format like '2023-12-31T00:00:00Z'",
				"example": gin.H{
					"due_date": "2025-7-16T00:00:00Z",
				},
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update task through service layer
	task, err := taskcontr.taskService.UpdateTask(id, &taskUpdate)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if task == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message":"task updated successfully", "updated task":&task})      // success response
}