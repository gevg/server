package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)

// AllergyIntolerance is the allergy_intolerance resource media type.--------------------------------------------------------------------------------
var AllergyIntolerancePayload = Type("AllergyIntolerancePayload", func() {
	Reference(AllergyIntolerance)
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
var ReactionPayload = Type("ReactionPayload", func() {
	Reference(Reaction)
	Attribute("substance")
	Attribute("certainty")
	Attribute("manifestation")
	Attribute("description")
	Attribute("onset")
	Attribute("severity")
	Attribute("exposureRoute")
	Attribute("note")
})
