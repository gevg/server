package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)

//	. "github.com/goa-fhir/server/design/data_types"

// NutritionRequestPayload defines the data structure used in the create vital request body.
// It is also the base type for the vital media type used to render bottles.
var NutritionRequestPayload = Type("NutritionRequestPayload", func() {
	Attribute("nutritionRequest", NutritionRequest)
	Attribute("id")
	Attribute("href")
	Attribute("patient", func() {
		Attribute("id")
		Attribute("reference")
	})

	// Attribute("orderer", func() {
	// 	Attribute("id")
	// 	Attribute("reference")
	// })

	Attribute("address_postal_code", func() {
		MinLength(5)
		Example(46032)
	})
	// Attribute("dateTime")

	// Required("status")

	Attribute("foodPreferenceModifier", func() {
		Attribute("id")
		Attribute("coding")
		Attribute("text")
	})

	// Attribute("excludeFoodModifier", func() {
	// 	Attribute("id")
	// 	Attribute("coding")
	// 	Attribute("text")
	// })

	Attribute("identifier")
	Attribute("address")
	Attribute("oralDiet")
	Attribute("supplement")
	Attribute("enteralFormula")
	Attribute("links")
})
