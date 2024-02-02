package response

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


type BookedEventsGettingError struct {
	Response
}

func newBookedEventsGettingError() *BookedEventsGettingError {
	return &BookedEventsGettingError{}
}

func (r *BookedEventsGettingError) SetData(data interface{}) {
	r.status = http.StatusInternalServerError
	r.message = "Error! Gettin booked events."
	r.data = data
}

func (r *BookedEventsGettingError) SendResponse(c *gin.Context) func() {
	return func () {
		c.JSON(http.StatusInternalServerError, gin.H{
		"Status":  r.status,
		"Message": r.message,
		"Data":    r.data})
	}
}

type BookedEventsGettingSuccess struct {
	Response
}

func newBookedEventsGettingSuccess() *BookedEventsGettingSuccess {
	return &BookedEventsGettingSuccess{}
}

func (r *BookedEventsGettingSuccess) SetData(data interface{}) {
	r.status = http.StatusOK
	r.message = "Success getting booked events!"
	r.data = data
}

func (r *BookedEventsGettingSuccess) SendResponse(c *gin.Context) func() {
	return func () {
		c.JSON(http.StatusOK, gin.H{
		"Status":  r.status,
		"Message": r.message,
		"Data":    r.data})
	}
}