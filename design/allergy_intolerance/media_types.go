package design

import (
	. "github.com/goa-fhir/server/design/data_types"
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// AllergyIntolerance is the allergy_intolerance resource media type.--------------------------------------------------------------------------------
var AllergyIntolerance = MediaType("application/vnd.allergy_intolerance+json", func() {
	Attributes(func() {
		Reference(AllergyIntolerancePayload)
		Attribute("identifier")
		Attribute("onset")
		Attribute("recordedDate")
		Attribute("recorder")
		Attribute("patient")
		Attribute("reporter")
		Attribute("status")
		Attribute("orderer")
		Attribute("criticality")
		Attribute("type")
		Attribute("category")
		Attribute("lastOccurence")
		Attribute("note")
		Attribute("reaction")
	})
	View("default", func() {
		Attribute("identifier")
		Attribute("onset")
		Attribute("recordedDate")
		Attribute("recorder")
		Attribute("patient")
		Attribute("reporter")
		Attribute("status")
		Attribute("orderer")
		Attribute("criticality")
		Attribute("type")
		Attribute("category")
		Attribute("lastOccurence")
		Attribute("note")
		Attribute("reaction")
	})
})
var Reaction = MediaType("application/vnd.reaction+json", func() {
	Attributes(func() {
		Reference(ReactionPayload)
		Attribute("substance")
		Attribute("certainty")
		Attribute("manifestation")
		Attribute("description")
		Attribute("onset")
		Attribute("severity")
		Attribute("exposureRoute")
		Attribute("note")
	})
	View("default", func() {
		Attribute("substance")
		Attribute("certainty")
		Attribute("manifestation")
		Attribute("description")
		Attribute("onset")
		Attribute("severity")
		Attribute("exposureRoute")
		Attribute("note")
	})
})
