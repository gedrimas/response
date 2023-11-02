package response

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type LoginError struct {
	Response
}

func newLoginError() *LoginError {
	return &LoginError{}
}

func (r *LoginError) SetData(data interface{}) {
	r.status = http.StatusUnauthorized
	r.message = "Email or passord is incorrect"
	r.data = data
}

func (r *LoginError) SendResponse(c *gin.Context) func() {
	return func () {
		c.JSON(http.StatusUnauthorized, gin.H{
		"Status":  r.status,
		"Message": r.message,
		"Data":    r.data})
	}
}