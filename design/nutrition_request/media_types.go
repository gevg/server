package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
	. "github.com/jamesallain/goa-fhir/design/data_types"
)

// NutritionRequest is the nutrition_request resource media type.--------------------------------------------------------------------------------
var NutritionRequest = MediaType("application/vnd.nutrition_request+json", func() {
	Description("A nutrition request from the patient")
	//Reference(NutritionRequestPayload)
	Attributes(func() {
		//Local Start-------------------------------------------------------------------------------
		Required("id", "href", "status")

		Attribute("id", Integer, "ID of nutrition request", func() {
			Example(1)
		})
		Attribute("href", String, "API href of nutrition request", func() {
			Example("/patients/1/nutrition_requests/1")
		})

		Attribute("created_at", DateTime, "Date of creation")
		Attribute("updated_at", DateTime, "Date of last update")
		//Local End-----------------------------------------------------------------------------------

		Required("status")
		Links(func() {
			//Link("identifier", "default") //datatype
			Link("oralDiet", "default")
			//Link("supplement", "default")
			Link("enteralFormula", "default")
			//Link("origin", "tiny")  // Set view used to render link if not "link"
		})

		Attribute("patient", HL7Reference, "The person (patient) who needs the nutrition order for an oral diet, nutritional supplement and/or enteral or formula feeding.", func() {
			//Reference to Patient
			//Comments:
			//Reason for inclusion or contrainment:
		})
		Attribute("orderer", HL7Reference, "The practitioner that holds legal responsibility for ordering the diet, nutritional supplement, or formula feedings.", func() {
			//Reference to Practitioner
			//Comments:
			//Reason for inclusion or contrainment:
		})
		Attribute("identifier", ArrayOf(Identifier), "Identifiers assigned to this order by the order sender or by the order receiver.", func() {
			//Reference to Indentifier
			//Comments:
			//Reason for inclusion or contrainment:
		})
		Attribute("encounter", HL7Reference, "An encounter that provides additional information about the healthcare context in which this request is made.", func() {
			//Reference to Encounter
			//Comments:
			//Reason for inclusion or contrainment:
			//Do I need to link reference to the corresponding resource???????????????????????????????????????????????????????????????????????????????????????????????????
		})
		Attribute("dateTime", DateTime, "The date and time that this nutrition order was requested.", func() {
			//Comments:
			//Reason for inclusion or contrainment:
		})
		Attribute("status", String, "The workflow status of the nutrition order/request. See http://hl7.org/fhir/nutrition-request-status", func() {
			//Comments:
			//Reason for inclusion or contrainment:
			Enum("proposed", "draft", "planned", "requested", "active", "on-hold", "completed", "cancelled")
		})
		Attribute("allergyIntolerance", ArrayOf(HL7Reference), "A link to a record of allergies or intolerances  which should be included in the nutrition order.", func() {
			//Reference to AllergryIntolerance
			//Comments:
			//Reason for inclusion or contrainment:
			//Do I need to link reference to the corresponding resource???????????????????????????????????????????????????????????????????????????????????????????????????
		})
		Attribute("foodPreferenceModifier", ArrayOf(CodeableConcept), `This modifier is used to convey order-specific modifiers about the type of food that should be given. These can be derived 
		from patient allergies, intolerances, or preferences such as Halal, Vegan or Kosher. This modifier applies to the entire nutrition order inclusive of the oral diet, nutritional
		supplements and enteral formula feedings. See http://hl7.org/fhir/ValueSet/encounter-diet`, func() {
			//Comments:
			//Reason for inclusion or contrainment:
		})
		Attribute("excludeFoodModifier", ArrayOf(CodeableConcept), `This modifier is used to convey order-specific modifiers about the type of food that should NOT be given. These can be derived from 
		patient allergies, intolerances, or preferences such as No Red Meat, No Soy or No Wheat or  Gluten-Free.  While it should not be necessary to repeat allergy or intolerance 
		information captured in the referenced allergyIntolerance resource in the excludeFoodModifier, this element may be used to convey additional specificity related to foods that should be 
		eliminated from the patient’s diet for any reason.  This modifier applies to the entire nutrition order inclusive of the oral diet, nutritional supplements and enteral formula feedings.
		See http://hl7.org/fhir/ValueSet/food-type`, func() {
			//Comments:
			//Reason for inclusion or contrainment:
		})
		Attribute("oralDiet", CollectionOf(OralDiet), "Diet given orally in contrast to enteral (tube) feeding.", func() {
			//Comments:
			//Reason for inclusion or contrainment:
		})
		Attribute("supplement", ArrayOf(Supplement), "Oral nutritional products given in order to add further nutritional value to the patient's diet.", func() {
			//Comments:
			//Reason for inclusion or contrainment:
		})
		Attribute("enteralFormula", CollectionOf(EnteralFormula), "Feeding provided through the gastrointestinal tract via a tube, catheter, or stoma that delivers nutrition distal to the oral cavity.", func() {
			//Comments:
			//Reason for inclusion or contrainment:
		})

	})

	View("default", func() {
		Attribute("id")
		Attribute("href")
		Attribute("patient", func() {
			Attribute("id")
			Attribute("reference")
		})

		Attribute("orderer", func() {
			Attribute("id")
			Attribute("reference")
		})

		Attribute("allergyIntolerance", func() {
			Attribute("id")
			Attribute("reference")
		})
		Attribute("dateTime")

		Attribute("status", func() {
			Enum("proposed", "draft", "planned", "requested", "active", "on-hold", "completed", "cancelled")
		})

		Required("status")

		Attribute("foodPreferenceModifier", func() {
			Attribute("id")
			Attribute("coding")
			Attribute("text")
		})

		Attribute("excludeFoodModifier", func() {
			Attribute("id")
			Attribute("coding")
			Attribute("text")
		})

		Attribute("identifier")
		Attribute("oralDiet")
		Attribute("supplement")
		Attribute("enteralFormula")
		Attribute("links")
	})

	View("tiny", func() {
		Attribute("id")
		Attribute("href")
		// Attribute("name")
		// Attribute("rating")
		Attribute("links")
	})

	View("full", func() {
		Attribute("id")
		Attribute("href")
		//Attribute("name")
		// Attribute("rating")
		// Attribute("vineyard")
		// Attribute("varietal")
		// Attribute("vintage")
		// Attribute("color")
		// Attribute("sweetness")
		// Attribute("country")
		// Attribute("region")
		// Attribute("review")
		Attribute("created_at")
		Attribute("updated_at")
		Attribute("links")
	})
})

var OralDiet = MediaType("application/vnd.oral_diet+json", func() {
	Description("A nutrition request from the patient")
	//Reference(NutritionRequestPayload)
	Attributes(func() {
		Attribute("type", CodeableConcept, "The kind of diet or dietary restriction such as fiber restricted diet or diabetic diet. See http://hl7.org/fhir/ValueSet/diet-type", func() {
			//Comments:
			//Reason for inclusion or contrainment:
		})
		Attribute("schedule", Timing, "The time period and frequency at which the diet should be given.", func() {
			//Comments:
			//Reason for inclusion or contrainment:
		})
		Attribute("nutrient", Nutrient, "Class that defines the quantity and type of nutrient modifications required for the oral diet.", func() {
			//Comments:
			//Reason for inclusion or contrainment:
		})
		Attribute("texture", Texture, "Class that describes any texture modifications required for the patient to safely consume various types of solid foods.", func() {
			//Comments:
			//Reason for inclusion or contrainment:
		})
		Attribute("fluidConsistencyType", CodeableConcept, "The required consistency (e.g. honey-thick, nectar-thick, thin, thickened.) of liquids or fluids served to the patient. See http://hl7.org/fhir/ValueSet/consistency-type", func() {
			//Comments:
			//Reason for inclusion or contrainment:
		})
		Attribute("instruction", String, "Free text or additional instructions or information pertaining to the oral diet.", func() {
			//Comments:
			//Reason for inclusion or contrainment:
		})
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

var Supplement = MediaType("application/vnd.supplement+json", func() {
	Description("Oral nutritional products given in order to add further nutritional value to the patient's diet.")
	//Reference(NutritionRequestPayload)
	Attributes(func() {
		Attribute("type", CodeableConcept, "The kind of nutritional supplement product required such as a high protein or pediatric clear liquid supplement.", func() {
			//Comments:
			//Reason for inclusion or contrainment:
		})
		Attribute("productName", String, "The product or brand name of the nutritional supplement such as 'Acme Protein Shake'.", func() {
			//Comments:
			//Reason for inclusion or contrainment:
		})
		Attribute("quantity", Quantity, "The amount of the nutritional supplement to be given.", func() {
			//Comments:
			//Reason for inclusion or contrainment:
		})
		Attribute("instruction", String, "Free text or additional instructions or information pertaining to the oral supplement.", func() {
			//Comments:
			//Reason for inclusion or contrainment:
		})
		View("default", func() {
			Description("A reference to a code defined by a terminology system")
			Attribute("type")
			Attribute("productName")
			Attribute("quantity")
			Attribute("instruction")
		})
	})
})

var EnteralFormula = MediaType("application/vnd.enteral_formula+json", func() {
	Description("Feeding provided through the gastrointestinal tract via a tube, catheter, or stoma that delivers nutrition distal to the oral cavity.")
	//Reference(NutritionRequestPayload)
	Attributes(func() {
		Attribute("baseFormulaType", CodeableConcept, "The type of enteral or infant formula such as an adult standard formula with fiber or a soy-based infant formula.", func() {
			//Comments:
			//Reason for inclusion or contrainment:
		})
		Attribute("baseFormulatProdcutName", String, "Product or brand name of the enteral or infant formula", func() {
			//Comments:
			//Reason for inclusion or contrainment:
		})
		Attribute("additiveType", CodeableConcept, "Indicates the type of modular component such as protein, carbohydrate, fat or fiber to be provided in addition to or mixed with the base formula.", func() {
			//Comments:
			//Reason for inclusion or contrainment:
		})
		Attribute("additiveProductName", String, "Product or brand name of the enteral or infant formula", func() {
			//Comments:
			//Reason for inclusion or contrainment:
		})
		Attribute("caloricDensity", Quantity, `The amount of energy (Calories) that the formula should provide per specified volume, typically per mL or fluid oz.
	    For example, an infant may require a formula that provides 24 Calories per fluid ounce or an adult may require an enteral formula that provides 1.5 Calorie/mL.`, func() {
			//Comments:
			//Reason for inclusion or contrainment:
		})
		//The route or physiological path of administration into the patient's gastrointestinal  tract for purposes of providing the formula feeding, e.g. nasogastric tube.
		Attribute("routeofAdministration", CodeableConcept, `The route or physiological path of administration into the patient's gastrointestinal  
		tract for purposes of providing the formula feeding, e.g. nasogastric tube.`, func() {
			//Comments:
			//Reason for inclusion or contrainment:
		})
		Attribute("maxVolumeToDeliver", Quantity, "The maximum total quantity of formula that may be administered to a subject over the period of time, e.g. 1440 mL over 24 hours.", func() {
			//Comments:
			//Reason for inclusion or contrainment:
		})
		Attribute("administrativeInstruction", String, "Free text formula administration, feeding instructions or additional instructions or information.", func() {
			//Comments:
			//Reason for inclusion or contrainment:
		})
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

var Nutrient = MediaType("application/vnd.nutrient+json", func() {
	Description("A nutrition request from the patient")
	//Reference(NutritionRequestPayload)
	Attributes(func() {
		Attribute("modifier", CodeableConcept, "The nutrient that is being modified such as carbohydrate or sodium. See http://hl7.org/fhir/ValueSet/nutrient-code", func() {
			//Comments:
			//Reason for inclusion or contrainment:
		})
		Attribute("amount", Quantity, "The quantity of the specified nutrient to include in diet.", func() {
			//Comments:
			//Reason for inclusion or contrainment:
		})
	})
	View("default", func() {
		Description("Default view for Nutrient BackboneElement")
		Attribute("modifier")
		Attribute("amount")
	})
})

var Texture = MediaType("application/vnd.texture+json", func() {
	Description("A nutrition request from the patient")
	//Reference(NutritionRequestPayload)
	Attributes(func() {
		Attribute("modifier", CodeableConcept, "Any texture modifications (for solid foods) that should be made, e.g. easy to chew, chopped, ground, and pureed. See http://hl7.org/fhir/ValueSet/texture-code", func() {
			//Comments:
			//Reason for inclusion or contrainment:
		})
		Attribute("foodType", CodeableConcept, "The food type(s) (e.g. meats, all foods)  that the texture modification applies to.  This could be all foods types. See http://hl7.org/fhir/ValueSet/modified-foodtype", func() {
			//Comments:
			//Reason for inclusion or contrainment:
		})
	})
	View("default", func() {
		Description("Default view for Texture BackboneElement")
		Attribute("modifier")
		Attribute("foodType")
	})
})
