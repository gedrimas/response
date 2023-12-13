package response

import (
	"net/http"
	"github.com/gin-gonic/gin"
)



type AllEventsGettingSuccess struct {
	Response
}

func newAllEventsGettingSuccess() *AllEventsGettingSuccess {
	return &AllEventsGettingSuccess{}
}

func (r *AllEventsGettingSuccess) SetData(data interface{}) {
	r.status = http.StatusOK
	r.message = "Success getting all events."
	r.data = data
}

func (r *AllEventsGettingSuccess) SendResponse(c *gin.Context) func() {
	return func () {
		c.JSON(http.StatusOK, gin.H{
		"Status":  r.status,
		"Message": r.message,
		"Data":    r.data})
	}
}

type AllEventsGettingError struct {
	Response
}

func newAllEventsGettingError() *AllEventsGettingError {
	return &AllEventsGettingError{}
}

func (r *AllEventsGettingError) SetData(data interface{}) {
	r.status = http.StatusInternalServerError
	r.message = "Error! Gettin all events."
	r.data = data
}

func (r *AllEventsGettingError) SendResponse(c *gin.Context) func() {
	return func () {
		c.JSON(http.StatusInternalServerError, gin.H{
		"Status":  r.status,
		"Message": r.message,
		"Data":    r.data})
	}
}