package router

//imports
import (
	"github.com/gin-gonic/gin"
	"github.com/natnael-eyuel-dev/Task-Manager-REST-API/data"
	"github.com/natnael-eyuel-dev/Task-Manager-REST-API/controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()  // create default gin router

	// initialize service and controller layers
	taskService := data.NewInMemoryTaskManager()                    // create in-memory task service instance
	taskController := controllers.NewTaskController(taskService)    // inject service into controller

	// define api routes
	router.POST("/tasks", taskController.CreateTask)          // create new task
	router.GET("/tasks", taskController.GetAllTasks)          // get all tasks
	router.GET("/tasks/:id", taskController.GetTaskByID)      // get specific task by id
	router.DELETE("/tasks/:id", taskController.DeleteTask)    // delete task by id
	router.PUT("/tasks/:id", taskController.UpdateTask)       // update existing task

	return router  // return configured router
}