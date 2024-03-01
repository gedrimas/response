package response

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


type VisitorsGettingError struct {
	Response
}

func newVisitorsGettingError() *VisitorsGettingError {
	return &VisitorsGettingError{}
}

func (r *VisitorsGettingError) SetData(data interface{}) {
	r.status = http.StatusInternalServerError
	r.message = "Error getting visitors"
	r.data = data
}

func (r *VisitorsGettingError) SendResponse(c *gin.Context) func() {
	return func () {
		c.JSON(http.StatusInternalServerError, gin.H{
		"Status":  r.status,
		"Message": r.message,
		"Data":    r.data})
	}
}


type VisitorsGettingSuccess struct {
	Response
}

func newVisitorsGettingSuccess() *VisitorsGettingSuccess {
	return &VisitorsGettingSuccess{}
}

func (r *VisitorsGettingSuccess) SetData(data interface{}) {
	r.status = http.StatusOK
	r.message = "Success getting visitors."
	r.data = data
}

func (r *VisitorsGettingSuccess) SendResponse(c *gin.Context) func() {
	return func () {
		c.JSON(http.StatusOK, gin.H{
		"Status":  r.status,
		"Message": r.message,
		"Data":    r.data})
	}
}