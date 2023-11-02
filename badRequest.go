package response

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type BadRequest struct {
	Response
}

func newBadRequest() *BadRequest {
	return &BadRequest{}
}

func (r *BadRequest) SetData(data interface{}) {
	r.status = http.StatusBadRequest
	r.message = "Bad request error"
	r.data = data
}

func (r *BadRequest) SendResponse(c *gin.Context) func() {
	return func () {
		c.JSON(http.StatusBadRequest, gin.H{
		"Status":  r.status,
		"Message": r.message,
		"Data":    r.data})
	}
}