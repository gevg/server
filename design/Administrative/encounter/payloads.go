package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)

// Encounter is the Encounter resource media type------------------------------------------------------------------------------------------------
var EncounterPayload = Type("EncounterPayload", func() {
	Reference(Encounter)
	Attribute("identifier")
	Attribute("status")
	Attribute("statusHistory")
	Attribute("class")
	Attribute("type")
	Attribute("priority")
	Attribute("patient")
})
var StatusHistoryPayload = Type("StatusHistoryPayload", func() {
	Reference(StatusHistory)
	Attribute("status")
	Attribute("period")
})
