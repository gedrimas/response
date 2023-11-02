package response

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type LoginSuccess struct {
	Response
}

func newLoginSuccess() *LoginSuccess {
	return &LoginSuccess{}
}

func (r *LoginSuccess) SetData(data interface{}) {
	r.status = http.StatusOK
	r.message = "Login success!"
	r.data = data
}

func (r *LoginSuccess) SendResponse(c *gin.Context) func() {
	return func () {
		c.JSON(http.StatusUnauthorized, gin.H{
		"Status":  r.status,
		"Message": r.message,
		"Data":    r.data})
	}
}