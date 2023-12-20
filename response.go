package response

import (
	"github.com/gin-gonic/gin"
)


type IResponse interface {
	SetData(d interface{})
	SendResponse(c *gin.Context) func()
}

type Response struct {
	status int
	message string
	data interface{}
}

func GetResponse(responseType string) IResponse{ 

	switch responseType {
	case "emailExistsError":
		return newEmailExistsError()
	case "emailGetError":
		return newEmailGettingError()
	case "phoneExistsError":
		return newPhoneExistsError()
	case "phoneGetError":
		return newPhoneGettingError()
	case "facebookExistsError":
		return newFacebookExistsError()
	case "facebookGetError":
		return newFacebookGettingError()			
	case "userExistsError":
		return newUserExistsError()
	case "userGetError":
		return newUserGettingError()
	case "userCreationError":
		return newUserCreationError()					
	case "companyExistsError":
		return newCompanyExistsError()
	case "companyGetError":
		return newCompanyGettingError()
	case "companyCreationError":
		return newCompanyCreationError()
	case "companiesGettingSuccess":
		return newCompaniesGettingSuccess()			
	case "badRequest":
		return newBadRequest()
	case "signuppSuccess":
		return newSignuppSuccess()
	case "loginError":
		return newLoginError()
	case "loginSuccess":
		return newLoginSuccess()
	case "eventCreationError":
		return newEventCreationError()
	case "eventSavingError":
		return newEventSavingError()
	case "eventCreationSuccess":
		return newEventCreationSuccess()	
	case "eventsGettingSuccess":
		return newAllEventsGettingSuccess()
	case "eventsGettingError":
		return newAllEventsGettingError()
	case "eventDeletionSuccess":
		return newEventDeletionSuccess()
	case "eventDeletionError":
		return newEventDeletionError()		
	case "eventUpdationSuccess":
		return newEventUpdationSuccess()
	case "eventUpdationError":
		return newEventUpdationError()
	case "visitorExclusionError":
		return newVisitorExclusionError()
	case "userListVisitExclusionError":
		return newUserListVisitExclusionError()	
	case "visitorAddingError":
		return newVisitorAddingError()
	case "eventAddingToUserListError":
		return newEventAddingToUserListError()	
	case "visitBookingSuccess": 
		return newVisitBookingSuccess()
	case "visitCancelationSuccess":
		return newVisitCancelationSuccess()
	case "bookedEventsGettingError": 
		return newBookedEventsGettingError()
	case "bookedEventsGettingSuccess": 
		return newBookedEventsGettingSuccess()																				
	default:
		return nil	
	}

}

type Sendor struct {
	response IResponse
	ctx *gin.Context
}

func NewSendor(c *gin.Context) *Sendor {
	return &Sendor{
		ctx: c,
	}
}

func (d *Sendor) Send(r IResponse) {
	d.response = r
	d.response.SendResponse(d.ctx)()
}



