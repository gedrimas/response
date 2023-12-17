package response

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type FacebookExistsError struct {
	Response
}

func newFacebookExistsError() *FacebookExistsError {
	return &FacebookExistsError{}
}

func (r *FacebookExistsError) SetData(data interface{}) {
	r.status = http.StatusBadRequest
	r.message = "This Facebook profile already exists"
	r.data = data
}

func (r *FacebookExistsError) SendResponse(c *gin.Context) func() {
	return func () {
		c.JSON(http.StatusBadRequest, gin.H{
		"Status":  r.status,
		"Message": r.message,
		"Data":    r.data})
	}
}

type FacebookGettingError struct {
	Response
}

func newFacebookGettingError() *FacebookGettingError {
	return &FacebookGettingError{}
}

func (r *FacebookGettingError) SetData(data interface{}) {
	r.status = http.StatusBadRequest
	r.message = "Error Gettinging email"
	r.data = data
}

func (r *FacebookGettingError) SendResponse(c *gin.Context) func() {
	return func () {
		c.JSON(http.StatusInternalServerError, gin.H{
		"Status":  r.status,
		"Message": r.message,
		"Data":    r.data})
	}
}