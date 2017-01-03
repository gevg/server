package design

import (
	. "github.com/goa-fhir/server/design/data_types"
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Substance is the substance resource media type------------------------------------------------------------------------------------------------
var Substance = MediaType("application/vnd.substance+json", func() {
	Attributes(func() {
		Reference(SubstancePayload)
		Attribute("identifier")
		Attribute("category")
		Attribute("code")
		Attribute("description")
		Attribute("instance")
		Attribute("ingredient")
	})

	View("default", func() {
		Description("Default view for Substance resource.")
		Attribute("identifier")
		Attribute("category")
		Attribute("code")
		Attribute("description")
		Attribute("instance")
		Attribute("ingredient")
	})
})

var Ingredient = MediaType("application/vnd.ingredient+json", func() {
	Attributes(func() {
		Reference(IngredientPayload)
		Attribute("quantity")
		Attribute("substance")
	})
	View("default", func() {
		Description("Default view for Ingredient element.")
		Attribute("quantity")
		Attribute("substance")
	})
})
var Instance = MediaType("application/vnd.instance+json", func() {
	Attributes(func() {
		Reference(InstancePayload)
		Attribute("identifier")
		Attribute("expiry")
		Attribute("quantity")
	})

	View("default", func() {
		Description("Default view for Instance element.")
		Attribute("identifier")
		Attribute("expiry")
		Attribute("quantity")
	})
})
