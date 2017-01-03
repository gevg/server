package design

import (
	. "github.com/goa-fhir/server/design/data_types"
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Specimen is the specimen resource media type------------------------------------------------------------------------------------------------
var Specimen = MediaType("application/vnd.specimen+json", func() {
	Attributes(func() {
		Reference(SpecimenPayload)
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
var Collection = MediaType("application/vnd.collection+json", func() {
	Attributes(func() {
		Reference(CollectionPayload)
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
var Container = MediaType("application/vnd.container+json", func() {
	Attributes(func() {
		Reference(ContainerPayload)
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
var Treatment = MediaType("application/vnd.treatment+json", func() {
	Attributes(func() {
		Reference(TreatmentPayload)
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
