package response

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type EventCreationError struct {
	Response
}

func newEventCreationError() *EventCreationError {
	return &EventCreationError{}
}

func (r *EventCreationError) SetData(data interface{}) {
	r.status = http.StatusBadRequest
	r.message = "Error Event Creation"
	r.data = data
}

func (r *EventCreationError) SendResponse(c *gin.Context) func() {
	return func () {
		c.JSON(http.StatusBadRequest, gin.H{
		"Status":  r.status,
		"Message": r.message,
		"Data":    r.data})
	}
}

type EventSavingError struct {
	Response
}

func newEventSavingError() *EventSavingError {
	return &EventSavingError{}
}

func (r *EventSavingError) SetData(data interface{}) {
	r.status = http.StatusInternalServerError
	r.message = "Error Event Saving"
	r.data = data
}

func (r *EventSavingError) SendResponse(c *gin.Context) func() {
	return func () {
		c.JSON(http.StatusInternalServerError, gin.H{
		"Status":  r.status,
		"Message": r.message,
		"Data":    r.data})
	}
}

type EventCreationSuccess struct {
	Response
}

func newEventCreationSuccess() *EventCreationSuccess {
	return &EventCreationSuccess{}
}

func (r *EventCreationSuccess) SetData(data interface{}) {
	r.status = http.StatusCreated
	r.message = "Success! New event created."
	r.data = data
}

func (r *EventCreationSuccess) SendResponse(c *gin.Context) func() {
	return func () {
		c.JSON(http.StatusCreated, gin.H{
		"Status":  r.status,
		"Message": r.message,
		"Data":    r.data})
	}
}

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