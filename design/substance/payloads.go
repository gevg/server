package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)

// Substance is the substance resource media type------------------------------------------------------------------------------------------------
var SubstancePayload = Type("SubstancePayload", func() {
	Reference(Substance)
	Attribute("identifier")
	Attribute("category")
	Attribute("code")
	Attribute("description")
	Attribute("instance")
	Attribute("ingredient")
})
var IngredientPayload = Type("IngredientPayload", func() {
	Reference(Ingredient)
	Attribute("quantity")
	Attribute("substance")
})
var InstancePayload = MediaType("InstancePayload", func() {
	Reference(Instance)
	Attribute("identifier")
	Attribute("expiry")
	Attribute("quantity")
})
