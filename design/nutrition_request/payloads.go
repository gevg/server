package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)

// NutritionRequest is the nutrition_request resource media type.--------------------------------------------------------------------------------
var NutritionRequestPayload = Type("NutritionRequestPayload", func() {
	Reference(NutritionRequest)
	Attribute("patient")
	Attribute("orderer")
	Attribute("identifier")
	Attribute("encounter")
	Attribute("dateTime")
	Attribute("status")
	Attribute("allergyIntolerance")
	Attribute("foodPreferenceModifier")
	Attribute("excludeFoodModifier")
	Attribute("oralDiet")
	Attribute("supplement")
	Attribute("enteralFormula")
})
var OralDietPayload = Type("OralDietPayload", func() {
	Reference(OralDiet)
	Attribute("type")
	Attribute("schedule")
	Attribute("nutrient")
	Attribute("texture")
	Attribute("fluidConsistencyType")
	Attribute("instruction")
})
var EnteralFormulaPayload = Type("EnteralFormulaPayload", func() {
	Reference(EnteralFormula)
	Attribute("baseFormulaType")
	Attribute("baseFormulatProdcutName")
	Attribute("additiveType")
	Attribute("additiveProductName")
	Attribute("caloricDensity")
	Attribute("routeofAdministration")
	Attribute("maxVolumeToDeliver")
	Attribute("administrativeInstruction")
})
var NutrientPayload = Type("NutrientPayload", func() {
	Reference(Nutrient)
	Attribute("modifier")
	Attribute("amount")
})
var SupplementPayload = Type("SupplementPayload", func() {
	Reference(Supplement)
	Attribute("type")
	Attribute("productName")
	Attribute("quantity")
	Attribute("instruction")
})
var TexturePayload = Type("TexturePayload", func() {
	Reference(Texture)
	Attribute("modifier")
	Attribute("foodType")

})
