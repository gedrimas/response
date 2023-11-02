package response

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type EmailExistsError struct {
	Response
}

func newEmailExistsError() *EmailExistsError {
	return &EmailExistsError{}
}

func (r *EmailExistsError) SetData(data interface{}) {
	r.status = http.StatusBadRequest
	r.message = "This email already exists"
	r.data = data
}

func (r *EmailExistsError) SendResponse(c *gin.Context) func() {
	return func () {
		c.JSON(http.StatusBadRequest, gin.H{
		"Status":  r.status,
		"Message": r.message,
		"Data":    r.data})
	}
}


type EmailGettingError struct {
	Response
}

func newEmailGettingError() *EmailGettingError {
	return &EmailGettingError{}
}

func (r *EmailGettingError) SetData(data interface{}) {
	r.status = http.StatusBadRequest
	r.message = "Error Gettinging email"
	r.data = data
}

func (r *EmailGettingError) SendResponse(c *gin.Context) func() {
	return func () {
		c.JSON(http.StatusInternalServerError, gin.H{
		"Status":  r.status,
		"Message": r.message,
		"Data":    r.data})
	}
}