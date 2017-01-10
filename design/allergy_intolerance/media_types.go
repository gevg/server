package design

import (
	. "github.com/goa-fhir/server/design/data_types"
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// AllergyIntolerance is the allergy_intolerance resource media type.--------------------------------------------------------------------------------
var AllergyIntoleranceMedia = MediaType("application/vnd.allergy.intolerance+json", func() {
	Attributes(func() {
		TypeName("AllergyIntoleranceMedia")
		Reference(AllergyIntolerancePayload)
		Attribute("meta")
		Attribute("identifier")
		Attribute("onset")
		Attribute("recordedDate")
		Attribute("recorder")
		Attribute("patient")
		Attribute("reporter")
		Attribute("status")
		Attribute("criticality")
		Attribute("type")
		Attribute("category")
		Attribute("lastOccurence")
		Attribute("note")
		Attribute("reaction")
	})
	View("default", func() {
		Attribute("meta")
		Attribute("identifier")
		Attribute("onset")
		Attribute("recordedDate")
		Attribute("recorder")
		Attribute("patient")
		Attribute("reporter")
		Attribute("status")
		Attribute("criticality")
		Attribute("type")
		Attribute("category")
		Attribute("lastOccurence")
		Attribute("note")
		Attribute("reaction")
	})
})
var ReactionMedia = MediaType("application/vnd.reaction+json", func() {
	Attributes(func() {
		TypeName("ReactionMedia")
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
