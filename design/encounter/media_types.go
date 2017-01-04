package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Encounter is the Encounter resource media type------------------------------------------------------------------------------------------------
var EncounterMedia = MediaType("application/vnd.encounter+json", func() {
	TypeName("EncounterMedia")
	Reference(EncounterPayload)
	Attributes(func() {
		Attribute("identifier")
		Attribute("status")
		Attribute("statusHistory")
		Attribute("class")
		Attribute("type")
		Attribute("priority")
		Attribute("patient")
	})
	View("default", func() {
		Description("Default view for Encounter resource.")
		Attribute("identifier")
		Attribute("status")
		Attribute("statusHistory")
		Attribute("class")
		Attribute("type")
		Attribute("priority")
		Attribute("patient")
	})
})
var StatusHistoryMedia = MediaType("application/vnd.status.history+json", func() {
	TypeName("StatusHistoryMedia")
	Reference(StatusHistoryPayload)
	Attributes(func() {
		Attribute("status")
		Attribute("period")
	})
	View("default", func() {
		Description("Default view for StatusHistory Backbone Element.")
		Attribute("status")
		Attribute("period")
	})
})
