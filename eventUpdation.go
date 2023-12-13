package response

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


type EventUpdationSuccess struct {
	Response
}

func newEventUpdationSuccess() *EventUpdationSuccess {
	return &EventUpdationSuccess{}
}

func (r *EventUpdationSuccess) SetData(data interface{}) {
	r.status = http.StatusOK
	r.message = "Success Updation event."
	r.data = data
}

func (r *EventUpdationSuccess) SendResponse(c *gin.Context) func() {
	return func () {
		c.JSON(http.StatusOK, gin.H{
		"Status":  r.status,
		"Message": r.message,
		"Data":    r.data})
	}
}

type EventUpdationError struct {
	Response
}

func newEventUpdationError() *EventUpdationError {
	return &EventUpdationError{}
}

func (r *EventUpdationError) SetData(data interface{}) {
	r.status = http.StatusInternalServerError
	r.message = "Error! Updation event."
	r.data = data
}

func (r *EventUpdationError) SendResponse(c *gin.Context) func() {
	return func () {
		c.JSON(http.StatusInternalServerError, gin.H{
		"Status":  r.status,
		"Message": r.message,
		"Data":    r.data})
	}
}