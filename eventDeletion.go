package response

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


type EventDeletionSuccess struct {
	Response
}

func newEventDeletionSuccess() *EventDeletionSuccess {
	return &EventDeletionSuccess{}
}

func (r *EventDeletionSuccess) SetData(data interface{}) {
	r.status = http.StatusOK
	r.message = "Success deletion event."
	r.data = data
}

func (r *EventDeletionSuccess) SendResponse(c *gin.Context) func() {
	return func () {
		c.JSON(http.StatusOK, gin.H{
		"Status":  r.status,
		"Message": r.message,
		"Data":    r.data})
	}
}

type EventDeletionError struct {
	Response
}

func newEventDeletionError() *EventDeletionError {
	return &EventDeletionError{}
}

func (r *EventDeletionError) SetData(data interface{}) {
	r.status = http.StatusInternalServerError
	r.message = "Error! Deletion event."
	r.data = data
}

func (r *EventDeletionError) SendResponse(c *gin.Context) func() {
	return func () {
		c.JSON(http.StatusInternalServerError, gin.H{
		"Status":  r.status,
		"Message": r.message,
		"Data":    r.data})
	}
}