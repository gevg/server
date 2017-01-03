package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
	. "github.com/jamesallain/goa-fhir/design/data_types"
)

// Encounter is the Encounter resource media type------------------------------------------------------------------------------------------------
var Encounter = MediaType("application/vnd.Encounter+json", func() {
	Description("Demographics and other administrative information about an individual or animal receiving care or other health-related services.")
	//Comments:
	//Reason for inclusion or contrainment:
	Attributes(func() {
		Links(func() {
			//Link("identifier", "default")
			Link("referenceRange", "default")
			Link("related", "default")
			Link("component", "default")
			//Link("origin", "tiny")  // Set view used to render link if not "link"
		})
		Attribute("identifier", ArrayOf(Identifier), "Identifier(s) by which this encounter is known.", func() {
			//Comments:
			//Reason for inclusion or contrainment:
		})
		Attribute("status", String, "planned | arrived | in-progress | onleave | finished | cancelled. See http://hl7.org/fhir/ValueSet/encounter-state", func() {
			//Comments:
			//Reason for inclusion or contrainment:
			Enum("planned", "arrived", "in-progress", "onleave", "finished", "cancelled")
		})
		Attribute("statusHistory", ArrayOf(StatusHistory), "The current status is always found in the current version of the resource, not the status history.", func() {
			//Comments:
			//Reason for inclusion or contrainment:
		})
		Attribute("class", CodeableConcept, "inpatient | outpatient | ambulatory | emergency +. See http://hl7.org/fhir/ValueSet/encounter-class", func() {
			//Comments:
			//Reason for inclusion or contrainment:
		})
		Attribute("type", ArrayOf(CodeableConcept), "Specific type of encounter (e.g. e-mail consultation, surgical day-care, skilled nursing, rehabilitation).", func() {
			//Comments:
			//Reason for inclusion or contrainment:
		})
		Attribute("priority", HL7Reference, "Indicates the urgency of the encounter. See http://hl7.org/fhir/ValueSet/encounter-priority", func() {
			//Comments:
			//Reason for inclusion or contrainment: Must know the subject context.
		})
		Attribute("patient", HL7Reference, "The patient present at the encounter.", func() {
			//Comments:
			//Reason for inclusion or contrainment:
		})

	})
	View("default", func() {
		Description("Default view for Encounter resource.")
		Attribute("identifier")
		Attribute("status")
		Attribute("statusHistory")
		Attribute("class")
		Attribute("type")
		Attribute("priority")
		Attribute("patient")
	})
})
var StatusHistory = MediaType("application/vnd.status.history+json", func() {
	Description("The current status is always found in the current version of the resource, not the status history.")
	//Comments:
	//Reason for inclusion or contrainment:
	Attributes(func() {
		Attribute("status", String, "planned | arrived | in-progress | onleave | finished | cancelled. See http://hl7.org/fhir/ValueSet/encounter-state", func() {
			//Comments:
			//Reason for inclusion or contrainment:
			Enum("planned", "arrived", "in-progress", "onleave", "finished", "cancelled")
		})
		Attribute("period", Period, "The time that the episode was in the specified status.", func() {
			//Comments:
			//Reason for inclusion or contrainment:
		})
	})
	View("default", func() {
		Description("Default view for StatusHistory Backbone Element.")
		Attribute("status")
		Attribute("period")
	})
})
