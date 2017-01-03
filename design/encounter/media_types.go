package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Encounter is the Encounter resource media type------------------------------------------------------------------------------------------------
var Encounter = MediaType("application/vnd.Encounter+json", func() {
	Attributes(func() {		
		Reference(EncounterPayload)
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
var StatusHistory = MediaType("application/vnd.status.history+json", func() {
		Attributes(func() {
			Reference(StatusHistoryPayload)
	Attribute("status")
		Attribute("period")
	})
	View("default", func() {
		Description("Default view for StatusHistory Backbone Element.")
		Attribute("status")
		Attribute("period")
	})
})
