package response

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type VisitorExclusionError struct {
	Response
}

func newVisitorExclusionError() *VisitorExclusionError {
	return &VisitorExclusionError{}
}

func (r *VisitorExclusionError) SetData(data interface{}) {
	r.status = http.StatusInternalServerError
	r.message = "Error! User visit wasn't canceled."
	r.data = data
}

func (r *VisitorExclusionError) SendResponse(c *gin.Context) func() {
	return func () {
		c.JSON(http.StatusInternalServerError, gin.H{
		"Status":  r.status,
		"Message": r.message,
		"Data":    r.data})
	}
}

type UserListVisitExclusionError struct {
	Response
}

func newUserListVisitExclusionError() *UserListVisitExclusionError {
	return &UserListVisitExclusionError{}
}

func (r *UserListVisitExclusionError) SetData(data interface{}) {
	r.status = http.StatusInternalServerError
	r.message = "Error! Event wasn't excluded from the user list."
	r.data = data
}

func (r *UserListVisitExclusionError) SendResponse(c *gin.Context) func() {
	return func () {
		c.JSON(http.StatusInternalServerError, gin.H{
		"Status":  r.status,
		"Message": r.message,
		"Data":    r.data})
	}
}

type VisitorAddingError struct {
	Response
}

func newVisitorAddingError() *VisitorAddingError {
	return &VisitorAddingError{}
}

func (r *VisitorAddingError) SetData(data interface{}) {
	r.status = http.StatusInternalServerError
	r.message = "Error! User wasn't added to the event."
	r.data = data
}

func (r *VisitorAddingError) SendResponse(c *gin.Context) func() {
	return func () {
		c.JSON(http.StatusInternalServerError, gin.H{
		"Status":  r.status,
		"Message": r.message,
		"Data":    r.data})
	}
}

type EventAddingToUserListError struct {
	Response
}

func newEventAddingToUserListError() *EventAddingToUserListError {
	return &EventAddingToUserListError{}
}

func (r *EventAddingToUserListError) SetData(data interface{}) {
	r.status = http.StatusInternalServerError
	r.message = "Error! Event wasn't added to user list"
	r.data = data
}

func (r *EventAddingToUserListError) SendResponse(c *gin.Context) func() {
	return func () {
		c.JSON(http.StatusInternalServerError, gin.H{
		"Status":  r.status,
		"Message": r.message,
		"Data":    r.data})
	}
}

type VisitBookingSuccess struct {
	Response
}

func newVisitBookingSuccess() *VisitBookingSuccess {
	return &VisitBookingSuccess{}
}

func (r *VisitBookingSuccess) SetData(data interface{}) {
	r.status = http.StatusCreated
	r.message = "Visit booking success!"
	r.data = data
}

func (r *VisitBookingSuccess) SendResponse(c *gin.Context) func() {
	return func () {
		c.JSON(http.StatusCreated, gin.H{
		"Status":  r.status,
		"Message": r.message,
		"Data":    r.data})
	}
}

type VisitCancelationSuccess struct {
	Response
}

func newVisitCancelationSuccess() *VisitCancelationSuccess {
	return &VisitCancelationSuccess{}
}

func (r *VisitCancelationSuccess) SetData(data interface{}) {
	r.status = http.StatusCreated
	r.message = "Visit cancelation success!"
	r.data = data
}

func (r *VisitCancelationSuccess) SendResponse(c *gin.Context) func() {
	return func () {
		c.JSON(http.StatusCreated, gin.H{
		"Status":  r.status,
		"Message": r.message,
		"Data":    r.data})
	}
}