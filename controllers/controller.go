package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ypgyan/api-go-gin/database"
	"github.com/ypgyan/api-go-gin/models"
	"net/http"
)

// Students godoc
// @Summary      Show all students
// @Description  Route to show all students
// @Tags         students
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.Student
// @Failure      400  {object}  httputil.HTTPError
// @Router       /students [get]
func Students(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.JSON(200, students)
}

func Greetings(c *gin.Context) {
	name := c.Params.ByName("student")
	c.JSON(200, gin.H{
		"message:": "Hello Mr." + name,
	})
}

// CreateStudent godoc
// @Summary      Create student
// @Description  Route to create a new Student
// @Tags         students
// @Accept       json
// @Produce      json
// @Param        student body   models.Student  true  "Student Model"
// @Success      200  {object}  models.Student
// @Failure      400  {object}  httputil.HTTPError
// @Router       /students [post]
func CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := models.Validate(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	database.DB.Create(&student)
	c.JSON(http.StatusCreated, student)
}

// FindStudent godoc
// @Summary      Show student
// @Description  get string by ID
// @Tags         students
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Student ID"
// @Success      200  {object}  models.Student
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /students/{id} [get]
func FindStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Student Not Found",
		})
		return
	}

	c.JSON(200, student)
}

func DeleteStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.Delete(&student, id)

	c.JSON(200, student)
}

// UpdateStudent godoc
// @Summary      Create student
// @Description  Route to create a new Student
// @Tags         students
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Student ID"
// @Success      200  {object}  models.Student
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /students [put]
func UpdateStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.First(&student, id)

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := models.Validate(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Model(&student).UpdateColumns(student)

	if student.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to update student",
		})
		return
	}
	c.JSON(http.StatusOK, student)
}

func FindStudentByCPF(c *gin.Context) {
	var student models.Student
	cpf := c.Param("cpf")

	database.DB.Where(&models.Student{CPF: cpf}).First(&student)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Student Not Found",
		})
		return
	}

	c.JSON(http.StatusOK, student)
}
