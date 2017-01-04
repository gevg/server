package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)

// Specimen is the specimen resource media type------------------------------------------------------------------------------------------------
var SpecimenPayload = Type("SpecimenPayload", func() {
	Reference(Specimen)
	Attribute("identifier")
	Attribute("status")
	Attribute("type")
	Attribute("parent")
	Attribute("subject")
	Attribute("accessionIdentifier")
	Attribute("receivedTime")
	Attribute("collection")
	Attribute("treatment")
	Attribute("container")
})
var CollectionPayload = Type("CollectionPayload", func() {
	Reference(Collection)
	Attribute("collector")
	Attribute("comment")
	Attribute("collectedDateTime")
	Attribute("collectedPeriod")
	Attribute("quantity")
	Attribute("method")
	Attribute("bodySite")
})
var ContainerPayload = Type("ContainerPayload", func() {
	Reference(Container)
	Attribute("relationship")
	Attribute("identifier")
	Attribute("description")
	Attribute("type")
	Attribute("capacity")
	Attribute("specimenQuantity")
	Attribute("additiveCodeableConcept")
	Attribute("additiveReference")
})
var TreatmentPayload = Type("TreatmentPayload", func() {
	Reference(Treatment)
	Attribute("language")
	Attribute("preferred")
	Attribute("careProvider")
})
