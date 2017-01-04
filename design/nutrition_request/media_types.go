package design

import (
	. "github.com/goa-fhir/server/design/data_types"
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// NutritionRequest is the nutrition_request resource media type.--------------------------------------------------------------------------------
var NutritionRequestMedia = MediaType("application/vnd.nutrition.request+json", func() {
	TypeName("NutritionRequestMedia")
	Reference(NutritionRequestPayload)
	Attributes(func() {
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
	View("default", func() {
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
})
var OralDietMedia = MediaType("application/vnd.oral.diet+json", func() {
	TypeName("OralDietMedia")
	Reference(OralDietPayload)
	Attributes(func() {
		Attribute("type")
		Attribute("schedule")
		Attribute("nutrient")
		Attribute("texture")
		Attribute("fluidConsistencyType")
		Attribute("instruction")
	})
	View("default", func() {
		Description("A reference to a code defined by a terminology system")
		Attribute("type")
		Attribute("schedule")
		Attribute("nutrient")
		Attribute("texture")
		Attribute("fluidConsistencyType")
		Attribute("instruction")
	})
})
var EnteralFormulaMedia = MediaType("application/vnd.enteral.formula+json", func() {
	TypeName("EnteralFormulaMedia")
	Reference(EnteralFormulaPayload)
	Attributes(func() {
		Attribute("baseFormulaType")
		Attribute("baseFormulatProdcutName")
		Attribute("additiveType")
		Attribute("additiveProductName")
		Attribute("caloricDensity")
		Attribute("routeofAdministration")
		Attribute("maxVolumeToDeliver")
		Attribute("administrativeInstruction")
	})
	View("default", func() {
		Description("A reference to a code defined by a terminology system")
		Attribute("baseFormulaType")
		Attribute("baseFormulatProdcutName")
		Attribute("additiveType")
		Attribute("additiveProductName")
		Attribute("caloricDensity")
		Attribute("routeofAdministration")
		Attribute("maxVolumeToDeliver")
		Attribute("administrativeInstruction")
	})
})
var NutrientMedia = MediaType("application/vnd.nutrient+json", func() {
	TypeName("NutrientMedia")
	Reference(NutrientPayload)
	Attributes(func() {
		Attribute("modifier")
		Attribute("amount")
	})
	View("default", func() {
		Description("Default view for Nutrient BackboneElement")
		Attribute("modifier")
		Attribute("amount")
	})
})
var SupplementMedia = MediaType("application/vnd.supplement+json", func() {
	TypeName("SupplementMedia")
	Reference(SupplementPayload)
	Attributes(func() {
		Attribute("type")
		Attribute("productName")
		Attribute("quantity")
		Attribute("instruction")

	})
	View("default", func() {
		Description("A reference to a code defined by a terminology system")
		Attribute("type")
		Attribute("productName")
		Attribute("quantity")
		Attribute("instruction")
	})
})
var TextureMedia = MediaType("application/vnd.texture+json", func() {
	TypeName("TextureMedia")
	Reference(TexturePayload)
	Attributes(func() {
		Attribute("modifier")
		Attribute("foodType")
	})
	View("default", func() {
		Description("Default view for Texture BackboneElement")
		Attribute("modifier")
		Attribute("foodType")
	})
})
