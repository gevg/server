package design

import (
	. "github.com/goa-fhir/server/design/data_types"
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Substance is the substance resource media type------------------------------------------------------------------------------------------------
var SubstanceMedia = MediaType("application/vnd.substance+json", func() {
	TypeName("SubstanceMedia")
	Reference(SubstancePayload)
	Attributes(func() {
		Attribute("meta")
		Attribute("identifier")
		Attribute("category")
		Attribute("code")
		Attribute("description")
		Attribute("instance")
		Attribute("ingredient")
	})
	View("default", func() {
		Description("Default view for Substance resource.")
		Attribute("meta")
		Attribute("identifier")
		Attribute("category")
		Attribute("code")
		Attribute("description")
		Attribute("instance")
		Attribute("ingredient")
	})
})
var IngredientMedia = MediaType("application/vnd.ingredient+json", func() {
	TypeName("IngredientMedia")
	Reference(IngredientPayload)
	Attributes(func() {
		Attribute("quantity")
		Attribute("substance")
	})
	View("default", func() {
		Description("Default view for Ingredient element.")
		Attribute("quantity")
		Attribute("substance")
	})
})
var InstanceMedia = MediaType("application/vnd.instance+json", func() {
	TypeName("InstanceMedia")
	Reference(InstancePayload)
	Attributes(func() {
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
