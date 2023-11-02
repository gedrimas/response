package response

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type SignuppSuccess struct {
	Response
}

func newSignuppSuccess() *SignuppSuccess {
	return &SignuppSuccess{}
}

func (r *SignuppSuccess) SetData(data interface{}) {
	r.status = http.StatusCreated
	r.message = "Success! New user created."
	r.data = data
}

func (r *SignuppSuccess) SendResponse(c *gin.Context) func() {
	return func () {
		c.JSON(http.StatusBadRequest, gin.H{
		"Status":  r.status,
		"Message": r.message,
		"Data":    r.data})
	}
}