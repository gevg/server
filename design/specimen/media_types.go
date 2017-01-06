package design

import (
	. "github.com/goa-fhir/server/design/data_types"
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Specimen is the specimen resource media type------------------------------------------------------------------------------------------------
var SpecimenMedia = MediaType("application/vnd.specimen+json", func() {
	TypeName("SpecimenMedia")
	Reference(SpecimenPayload)
	Attributes(func() {
		Attribute("meta")
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
	View("default", func() {
		Description("Default view for Specimen resource.")
		Attribute("meta")
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
})
var CollectionMedia = MediaType("application/vnd.collection+json", func() {
	TypeName("CollectionMedia")
	Reference(CollectionPayload)
	Attributes(func() {
		Attribute("collector")
		Attribute("comment")
		Attribute("collectedDateTime")
		Attribute("collectedPeriod")
		Attribute("quantity")
		Attribute("method")
		Attribute("bodySite")
	})
	View("default", func() {
		Description("Default view for Collection Element.")
		Attribute("collector")
		Attribute("comment")
		Attribute("collectedDateTime")
		Attribute("collectedPeriod")
		Attribute("quantity")
		Attribute("method")
		Attribute("bodySite")
	})
})
var ContainerMedia = MediaType("application/vnd.container+json", func() {
	TypeName("ContainerMedia")
	Reference(ContainerPayload)
	Attributes(func() {
		Attribute("relationship")
		Attribute("identifier")
		Attribute("description")
		Attribute("type")
		Attribute("capacity")
		Attribute("specimenQuantity")
		Attribute("additiveCodeableConcept")
		Attribute("additiveReference")
	})
	View("default", func() {
		Description("Default view for Container element.")
		Attribute("relationship")
		Attribute("identifier")
		Attribute("description")
		Attribute("type")
		Attribute("capacity")
		Attribute("specimenQuantity")
		Attribute("additiveCodeableConcept")
		Attribute("additiveReference")
	})
})
var TreatmentMedia = MediaType("application/vnd.treatment+json", func() {
	TypeName("TreatmentMedia")
	Reference(TreatmentPayload)
	Attributes(func() {
		Attribute("language")
		Attribute("preferred")
		Attribute("careProvider")
	})
	View("default", func() {
		Description("Default view for Treatment element.")
		Attribute("language")
		Attribute("preferred")
		Attribute("careProvider")
	})
})
