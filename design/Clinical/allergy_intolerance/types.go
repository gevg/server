package design

import (
	. "github.com/goa-fhir/server/design/DataType/data_types"
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// PatientPayload defines the data structure used in the create vital request body.
// It is also the base type for the vital media type used to render bottles.
var AllergyIntolerance = Type("AllergyIntolerance", func() {
	Description("A nutrition request from the patient")
	//Reference(AllergyIntolerancePayload)
	//Local Start-------------------------------------------------------------------------------
	Required("id", "href", "status")

	Attribute("id", Integer, "ID of nutrition request", func() {
		Example(1)
	})
	Attribute("href", String, "API href of nutrition request", func() {
		Example("/patients/1/allergy_intolerances/1")
	})

	Attribute("created_at", DateTime, "Date of creation")
	Attribute("updated_at", DateTime, "Date of last update")
	//Local End-----------------------------------------------------------------------------------

	Required("status", "type")
	Attribute("meta", Meta, `The metadata about a resource. This is content in the resource that is maintained by the infrastructure.
	Changes to the content may not always be associated with version changes to the resource.`, func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("identifier", ArrayOf(Identifier), `This records identifiers associated with this allergy/intolerance concern that are defined by business processes and/or
		used to refer to it when a direct URL reference to the resource itself is not appropriate (e.g. in CDA documents, or in written / printed documentation).`, func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("onset", DateTime, "Record of the date and/or time of the onset of the Allergy or Intolerance.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("recordedDate", DateTime, "Date when the sensitivity was recorded.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("recorder", HL7Reference, "Individual who recorded the record and takes responsibility for its conten.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("patient", HL7Reference, "The patient who has the allergy or intolerance.", func() {
		//Reference to Patient
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("reporter", HL7Reference, "The source of the information about the allergy that is recorded.", func() {
		//Reference to Patient, RelatedPerson, Practitioner
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("status", String, "Assertion about certainty associated with a propensity, or potential risk, of a reaction to the identified Substance. See http://hl7.org/fhir/ValueSet/allergy-intolerance-status", func() {
		//Type: code
		//Comments:
		//Reason for inclusion or contrainment:
		Enum("active", "unconfirmed", "confirmed", "inactive", "resolved", "refuted", "entered-in-error")
	})
	//The practitioner that holds legal responsibility for ordering the diet, nutritional supplement, or formula feedings.
	Attribute("criticality", String, "Estimate of the potential clinical harm, or seriousness, of the reaction to the identified Substance. See http://hl7.org/fhir/ValueSet/allergy-intolerance-criticality", func() {
		//Type: code
		//Comments:
		//Reason for inclusion or contrainment:
		Enum("CRITL", "CRITH", "CRITU")
	})
	Attribute("type", String, "Identification of the underlying physiological mechanism for the reaction risk. See http://hl7.org/fhir/ValueSet/allergy-intolerance-type", func() {
		//Type: code
		//Comments:
		//Reason for inclusion or contrainment:
		Enum("allergy", "intolerance")
	})
	Attribute("category", String, "Category of the identified Substance. See http://hl7.org/fhir/ValueSet/allergy-intolerance-category", func() {
		//Type: code
		//Comments:
		//Reason for inclusion or contrainment:
		Enum("food", "medication", "environment", "other")
	})
	Attribute("lastOccurence", DateTime, "Represents the date and/or time of the last known occurrence of a reaction event.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("note", Annotation, "Additional narrative about the propensity for the Adverse Reaction, not captured in other fields..", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("reaction", Reaction, "Details about each adverse reaction event linked to exposure to the identified Substance.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})

})
var Reaction = Type("Reaction", func() {
	Description("Details about each adverse reaction event linked to exposure to the identified Substance.")
	//Comments:
	//Reason for inclusion or contrainment:
	//Reference(AllergyIntolerancePayload)
	Attribute("substance", CodeableConcept, `Identification of the specific substance considered to be responsible for the Adverse Reaction event.
		Note: the substance for a specific reaction may be different to the substance identified as the cause of the risk, but must be consistent with it.
		For instance, it may be a more specific substance (e.g. a brand medication) or a composite substance that includes the identified substance.
		It must be clinically safe to only process the AllergyIntolerance.substance and ignore the AllergyIntolerance.event.substance.`, func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("certainty", String, "Statement about the degree of clinical certainty that the specific substance was the cause of the manifestation in this reaction event. See http://hl7.org/fhir/ValueSet/reaction-event-certainty", func() {
		//Comments:
		//Reason for inclusion or contrainment:
		Enum("unlikely", "likely", "confirmed")
	})
	Attribute("manifestation", CodeableConcept, "Clinical symptoms and/or signs that are observed or associated with the adverse reaction event. See http://hl7.org/fhir/ValueSet/manifestation-codes", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("description", String, "Text description about the reaction as a whole, including details of the manifestation if required.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("onset", DateTime, "Record of the date and/or time of the onset of the Reaction.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("severity", String, "Clinical assessment of the severity of the reaction event as a whole, potentially considering multiple different manifestations. See http://hl7.org/fhir/ValueSet/reaction-event-severity", func() {
		//Comments:
		//Reason for inclusion or contrainment:
		Enum("mild", "moderate", "severe")
	})
	Attribute("exposureRoute", CodeableConcept, "Identification of the route by which the subject was exposed to the substance. See http://hl7.org/fhir/ValueSet/route-codes", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("note", Annotation, "Additional text about the adverse reaction event not captured in other fields.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
})
