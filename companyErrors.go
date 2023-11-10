package response

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type CompanyExistsError struct {
	Response
}

func newCompanyExistsError() *CompanyExistsError {
	return &CompanyExistsError{}
}

func (r *CompanyExistsError) SetData(data interface{}) {
	r.status = http.StatusBadRequest
	r.message = "This company already exists"
	r.data = data
}

func (r *CompanyExistsError) SendResponse(c *gin.Context) func() {
	return func () {
		c.JSON(http.StatusBadRequest, gin.H{
		"Status":  r.status,
		"Message": r.message,
		"Data":    r.data})
	}
}

type CompanyGettingError struct {
	Response
}

func newCompanyGettingError() *CompanyGettingError {
	return &CompanyGettingError{}
}

func (r *CompanyGettingError) SetData(data interface{}) {
	r.status = http.StatusInternalServerError
	r.message = "Error getting Company"
	r.data = data
}

func (r *CompanyGettingError) SendResponse(c *gin.Context) func() {
	return func () {
		c.JSON(http.StatusInternalServerError, gin.H{
		"Status":  r.status,
		"Message": r.message,
		"Data":    r.data})
	}
}

type CompanyCreationError struct {
	Response
}

func newCompanyCreationError() *CompanyCreationError {
	return &CompanyCreationError{}
}

func (r *CompanyCreationError) SetData(data interface{}) {
	r.status = http.StatusInternalServerError
	r.message = "Error Creaton Company"
	r.data = data
}

func (r *CompanyCreationError) SendResponse(c *gin.Context) func() {
	return func () {
		c.JSON(http.StatusInternalServerError, gin.H{
		"Status":  r.status,
		"Message": r.message,
		"Data":    r.data})
	}
}

type CompaniesGettingSuccess struct {
	Response
}

func newCompaniesGettingSuccess() *CompaniesGettingSuccess {
	return &CompaniesGettingSuccess{}
}

func (r *CompaniesGettingSuccess) SetData(data interface{}) {
	r.status = http.StatusOK
	r.message = "Success getting companies."
	r.data = data
}

func (r *CompaniesGettingSuccess) SendResponse(c *gin.Context) func() {
	return func () {
		c.JSON(http.StatusOK, gin.H{
		"Status":  r.status,
		"Message": r.message,
		"Data":    r.data})
	}
}