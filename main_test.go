package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/ypgyan/api-go-gin/controllers"
	"github.com/ypgyan/api-go-gin/database"
	"github.com/ypgyan/api-go-gin/models"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

var ID uint

func SetupTestRoute() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}

func CreateStudentMock() {
	student := models.Student{Name: "Test Student", CPF: "12312312332", RG: "123123123"}
	database.DB.Create(&student)
	ID = student.ID
}

func DeleteStudentMock() {
	var student models.Student
	database.DB.Delete(&student, ID)
}

func TestVerifyGreetingsSuccess(t *testing.T) {
	r := SetupTestRoute()
	r.GET("/:student", controllers.Greetings)
	req, _ := http.NewRequest("GET", "/Yan", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code, "They should be equal")

	resMock := `{"message:":"Hello Mr.Yan"}`
	resBody, _ := ioutil.ReadAll(res.Body)
	assert.Equal(t, resMock, string(resBody), "They should be equal")
}

func TestListingAllStudents(t *testing.T) {
	database.ConnectDB()
	CreateStudentMock()
	defer DeleteStudentMock()

	r := SetupTestRoute()
	r.GET("/students", controllers.Students)

	req, _ := http.NewRequest("GET", "/students", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code, "They should be equal")
}

func TestStudentCPFSearch(t *testing.T) {
	database.ConnectDB()
	CreateStudentMock()
	defer DeleteStudentMock()

	r := SetupTestRoute()
	r.GET("/students/cpf/:cpf", controllers.Students)

	req, _ := http.NewRequest("GET", "/students/cpf/12312312332", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code, "They should be equal")
}

func TestFindStudents(t *testing.T) {
	database.ConnectDB()
	CreateStudentMock()
	defer DeleteStudentMock()

	r := SetupTestRoute()
	r.GET("/students/:id", controllers.FindStudent)
	pathUrl := "/students/" + strconv.FormatUint(uint64(ID), 10)
	req, _ := http.NewRequest("GET", pathUrl, nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	var studentMock models.Student
	json.Unmarshal(res.Body.Bytes(), &studentMock)
	assert.Equal(t, "Test Student", studentMock.Name)
	assert.Equal(t, "12312312332", studentMock.CPF)
	assert.Equal(t, "123123123", studentMock.RG)
}

func TestDeleteStudent(t *testing.T) {
	database.ConnectDB()
	CreateStudentMock()

	r := SetupTestRoute()
	r.DELETE("/students/:id", controllers.DeleteStudent)
	pathUrl := "/students/" + strconv.FormatUint(uint64(ID), 10)
	req, _ := http.NewRequest("DELETE", pathUrl, nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code, "They should be equal")
}

func TestUpdateStudent(t *testing.T) {
	database.ConnectDB()
	CreateStudentMock()

	r := SetupTestRoute()
	r.PATCH("/students/:id", controllers.UpdateStudent)
	pathUrl := "/students/" + strconv.FormatUint(uint64(ID), 10)

	student := models.Student{Name: "Test Student", CPF: "12312345678", RG: "123123123"}
	body, _ := json.Marshal(student)
	req, _ := http.NewRequest("PATCH", pathUrl, bytes.NewBuffer(body))
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	var studentMock models.Student
	json.Unmarshal(res.Body.Bytes(), &studentMock)
	assert.Equal(t, "Test Student", studentMock.Name)
	assert.Equal(t, "12312345678", studentMock.CPF)
	assert.Equal(t, "123123123", studentMock.RG)
	defer DeleteStudentMock()
}
