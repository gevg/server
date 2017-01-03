package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
	. "github.com/goa-fhir/server/design/data_types"
)

// Substance is the substance resource media type------------------------------------------------------------------------------------------------
var Substance = MediaType("application/vnd.substance+json", func() {
	Description("Demographics and other administrative information about an individual or animal receiving care or other health-related services.")
				//Comments:
			//Reason for inclusion or contrainment:
	Attributes(func() {
		Links(func() {
			Link("identifier", "default")
			Link("instance", "default")
			Link("ingredient", "default")
			//Link("origin", "tiny")  // Set view used to render link if not "link"
		})
		Attribute("identifier", Identifier, "Id for substance.", func() {
			//Comments:
			//Reason for inclusion or contrainment:
		})
		Attribute("category", CodeableConcept, "A code that classifies the general type of substance.  This is used  for searching, sorting and display purposes. See http://hl7.org/fhir/ValueSet/substance-category", func() {
			//Comments:
			//Reason for inclusion or contrainment:
		})
		Attribute("code", CodeableConcept, "A code (or set of codes) that identify this substance. See http://hl7.org/fhir/ValueSet/substance-code ", func() {
			//Comments:
			//Reason for inclusion or contrainment:
		})
		Attribute("description", String, "A description of the substance - its appearance, handling requirements, and other usage notes.", func() {
			//Comments:
			//Reason for inclusion or contrainment:
		})
		Attribute("instance", Instance, "Substance may be used to describe a kind of substance, or a specific package/container of the substance: an instance.", func() {
			//Comments:
			//Reason for inclusion or contrainment:
		})
		Attribute("ingredient", Ingredient, "A substance can be composed of other substances.", func() {
			//Comments:
			//Reason for inclusion or contrainment:
		})
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

var Instance = MediaType("application/vnd.instance+json", func() {
	Description("Substance may be used to describe a kind of substance, or a specific package/container of the substance: an instance.")
	//Comments:
	//Reason for inclusion or contrainment:
	Attributes(func() {
		Attribute("indentifier", Identifier, "Substance may be used to describe a kind of substance, or a specific package/container of the substance: an instance.", func() {
			//Comments:
			//Reason for inclusion or contrainment:
		})
		Attribute("expiry", DateTime, "When the substance is no longer valid to use. For some substances, a single arbitrary date is used for expiry.", func() {
			//Comments:
			//Reason for inclusion or contrainment:
		})
		Attribute("quantity", Quantity, "The amount of the substance.", func() {
			//Comments:
			//Reason for inclusion or contrainment:
		})
	})

	View("default", func() {
		Description("Default view for Instance element.")
		Attribute("identifier")
		Attribute("expiry")
		Attribute("quantity")
	})
})

var Ingredient = MediaType("application/vnd.ingredient+json", func() {
	Description("A substance can be composed of other substances.")
	//Comments:
	//Reason for inclusion or contrainment:
	Attributes(func() {
		Attribute("quantity", Ratio, "The amount of the substance.", func() {
			//Comments:
			//Reason for inclusion or contrainment:
		})
		Attribute("substance", HL7Reference, "Another substance that is a component of this substance.", func() {
			//Reference to Substance
			//Comments:
			//Reason for inclusion or contrainment:
		})
	})
	View("default", func() {
		Description("Default view for Ingredient element.")
		Attribute("quantity")
		Attribute("substance")
	})
})
