package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// ContactDetail is the ContactDetail resource media type------------------------------------------------------------------------------------------------
var ContactDetailMedia = MediaType("application/vnd.contact.detail+json", func() {
	TypeName("ContactDetailMedia")
	Reference(ContactDetailPayload)
	Attributes(func() {
		Attribute("name")
		Attribute("telecom")
	})
	View("default", func() {
		Description("Default view for ContactDetail resource.")
		Attribute("name")
		Attribute("telecom")
	})
})
var ContributorMedia = MediaType("application/vnd.contributor+json", func() {
	TypeName("ContributorMedia")
	Reference(ContributorPayload)
	Attributes(func() {
		Attribute("type")
		Attribute("name")
		Attribute("contact")
	})
	View("default", func() {
		Description("Default view for Contributor Backbone Element.")
		Attribute("type")
		Attribute("name")
		Attribute("contact")
	})
})
