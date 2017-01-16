package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)

// ContactDetail is the ContactDetail resource media type------------------------------------------------------------------------------------------------
var ContactDetailPayload = Type("ContactDetailPayload", func() {
	Reference(ContactDetail)
	Attribute("name")
	Attribute("telecom")	
})
var StatusHistoryPayload = Type("StatusHistoryPayload", func() {
	Reference(StatusHistory)
	Attribute("type")
	Attribute("name")
	Attribute("contact")
})
