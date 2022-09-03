package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ypgyan/api-go-gin/controllers"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/:student", controllers.Greetings)
	r.GET("/students", controllers.Students)
	r.GET("/students/:id", controllers.FindStudent)
	r.GET("/students/cpf/:cpf", controllers.FindStudentByCPF)
	r.POST("/students", controllers.CreateStudent)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.PUT("/students/:id", controllers.UpdateStudent)

	r.Run(":8001")
}
