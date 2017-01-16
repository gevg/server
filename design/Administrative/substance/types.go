package design

import (
	. "github.com/goa-fhir/server/design/DataType/data_types"
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var Substance = Type("Substance", func() {
	Description("Demographics and other administrative information about an individual or animal receiving care or other health-related services.")
	//Comments:
	//Reason for inclusion or contrainment:
	Attribute("meta", Meta, `The metadata about a resource. This is content in the resource that is maintained by the infrastructure.
	Changes to the content may not always be associated with version changes to the resource.`, func() {
		//Comments:
		//Reason for inclusion or contrainment:
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
var Ingredient = Type("Ingredient", func() {
	Description("A substance can be composed of other substances.")
	//TypeName("BackboneElement")
	//Comments:
	//Reason for inclusion or contrainment:
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
var Instance = Type("Instance", func() {
	Description("Substance may be used to describe a kind of substance, or a specific package/container of the substance: an instance.")
	//Comments:
	//Reason for inclusion or contrainment:
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
