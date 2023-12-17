package response

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type PhoneExistsError struct {
	Response
}

func newPhoneExistsError() *PhoneExistsError {
	return &PhoneExistsError{}
}

func (r *PhoneExistsError) SetData(data interface{}) {
	r.status = http.StatusBadRequest
	r.message = "This phone already exists"
	r.data = data
}

func (r *PhoneExistsError) SendResponse(c *gin.Context) func() {
	return func () {
		c.JSON(http.StatusBadRequest, gin.H{
		"Status":  r.status,
		"Message": r.message,
		"Data":    r.data})
	}
}


type PhoneGettingError struct {
	Response
}

func newPhoneGettingError() *PhoneGettingError {
	return &PhoneGettingError{}
}

func (r *PhoneGettingError) SetData(data interface{}) {
	r.status = http.StatusBadRequest
	r.message = "Error Gettinging email"
	r.data = data
}

func (r *PhoneGettingError) SendResponse(c *gin.Context) func() {
	return func () {
		c.JSON(http.StatusInternalServerError, gin.H{
		"Status":  r.status,
		"Message": r.message,
		"Data":    r.data})
	}
}