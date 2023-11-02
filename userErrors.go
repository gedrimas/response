package response

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type UserExistsError struct {
	Response
}

func newUserExistsError() *UserExistsError {
	return &UserExistsError{}
}

func (r *UserExistsError) SetData(data interface{}) {
	r.status = http.StatusBadRequest
	r.message = "This user already exists"
	r.data = data
}

func (r *UserExistsError) SendResponse(c *gin.Context) func() {
	return func () {
		c.JSON(http.StatusBadRequest, gin.H{
		"Status":  r.status,
		"Message": r.message,
		"Data":    r.data})
	}
}


type UserGettingError struct {
	Response
}

func newUserGettingError() *UserGettingError {
	return &UserGettingError{}
}

func (r *UserGettingError) SetData(data interface{}) {
	r.status = http.StatusInternalServerError
	r.message = "Error Gettinging User"
	r.data = data
}

func (r *UserGettingError) SendResponse(c *gin.Context) func() {
	return func () {
		c.JSON(http.StatusInternalServerError, gin.H{
		"Status":  r.status,
		"Message": r.message,
		"Data":    r.data})
	}
}

type UserCreationError struct {
	Response
}

func newUserCreationError() *UserCreationError {
	return &UserCreationError{}
}

func (r *UserCreationError) SetData(data interface{}) {
	r.status = http.StatusInternalServerError
	r.message = "Error user creation"
	r.data = data
}

func (r *UserCreationError) SendResponse(c *gin.Context) func() {
	return func () {
		c.JSON(http.StatusInternalServerError, gin.H{
		"Status":  r.status,
		"Message": r.message,
		"Data":    r.data})
	}
}