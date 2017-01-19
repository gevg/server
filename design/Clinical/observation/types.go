package design

import (
	. "github.com/goa-fhir/server/design/DataType/data_types"
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Observation is the observation resource media type------------------------------------------------------------------------------------------------
var Observation = Type("Observation", func() {
	Description("Demographics and other administrative information about an individual or animal receiving care or other health-related services.")
	//Comments:
	//Reason for inclusion or contrainment:
	Attribute("meta", Meta, `The metadata about a resource. This is content in the resource that is maintained by the infrastructure.
	Changes to the content may not always be associated with version changes to the resource.`, func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("identifier", ArrayOf(Identifier), "A unique identifier for the simple observation instance.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("status", String, "The status of the result value. See http://hl7.org/fhir/ValueSet/observation-status", func() {
		//Comments:
		//Reason for inclusion or contrainment:
		Enum("registered", "preliminary", "final", "amended +")
	})
	Attribute("category", CodeableConcept, "A code that classifies the general type of observation being made.  This is used  for searching, sorting and display purposes. See http://hl7.org/fhir/ValueSet/observation-category", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("code", CodeableConcept, "Describes what was observed. Sometimes this is called the observation 'name'. See http://hl7.org/fhir/ValueSet/observation-codes", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("encounter", HL7Reference, "The healthcare event  (e.g. a patient and healthcare provider interaction) during which this observation is made.", func() {
		//Reference: Encounter
		//Comments:
		//Reason for inclusion or contrainment: Must know the subject context.
	})

	Attribute("effectiveDateTime", DateTime, `The time or time-period the observed value is asserted as being true. For biological subjects - e.g. human patients - 
		this is usually called the "physiologically relevant time". This is usually either the time of the procedure or of specimen collection, 
		but very often the source of the date/time is not known, only the date/time itself.`, func() {
		//Form: effective[x]
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("effectivePeriod", Period, `The time or time-period the observed value is asserted as being true. For biological subjects - e.g. human patients - 
		this is usually called the "physiologically relevant time". This is usually either the time of the procedure or of specimen collection, 
		but very often the source of the date/time is not known, only the date/time itself.`, func() {
		//Form: effective[x]
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("issued", DateTime, "The date and time this observation was made available to providers, typically after the results have been reviewed and verified.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("performer", ArrayOf(HL7Reference), "Who was responsible for asserting the observed value as 'true'.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})

	Attribute("valueQuantity", Quantity)
	Attribute("valueCodeableConcept", CodeableConcept)
	Attribute("valueString", String)
	Attribute("valueRange", Range)
	Attribute("valueSampledData", SampleData)
	Attribute("valueAttachment", Attachment)
	Attribute("valueTime", DateTime)
	Attribute("valueDatTime", DateTime)
	Attribute("valuePeriod", Period)

	Attribute("dateAbsentReason", CodeableConcept, "Provides a reason why the expected value in the element Observation.value[x] is missing. See http://hl7.org/fhir/ValueSet/observation-valueabsentreason", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("interpretation", CodeableConcept, `The assessment made based on the result of the observation.  Intended as a simple compact code often 
		placed adjacent to the result value in reports and flow sheets to signal the meaning/normalcy status of the result. Otherwise known as abnormal flag.
		See http://hl7.org/fhir/ValueSet/observation-interpretation`, func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("comments", String, "May include statements about significant, unexpected or unreliable values, or information about the source of the value where this may be relevant to the interpretation of the result.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("bodySite", CodeableConcept, "Indicates the site on the subject's body where the observation was made (i.e. the target site). See http://hl7.org/fhir/ValueSet/body-site", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("method", CodeableConcept, "Indicates the mechanism used to perform the observation. See http://hl7.org/fhir/ValueSet/observation-methods", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("specimen", HL7Reference, "The specimen that was used when this observation was made.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("device", HL7Reference, "The device used to generate the observation data.", func() {
		//Reference: Device, DeviceMatrix
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("referenceRange", ArrayOf(ReferenceRange), "Guidance on how to interpret the value by comparison to a normal or recommended range.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("related", ArrayOf(Related), "A  reference to another resource (usually another Observation but could  also be a QuestionnaireAnswer) whose relationship is defined by the relationship type code.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("component", ArrayOf(Component), `Some observations have multiple component observations.  These component observations are expressed as separate code 
		value pairs that share the same attributes.  Examples include systolic and diastolic component observations for blood pressure measurement and multiple 
		component observations for genetics observations.`, func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
})
var ReferenceRange = Type("ReferenceRange", func() {

	Description("Details concerning the observation collection.")
	//Comments:
	//Reason for inclusion or contrainment:
	Attribute("low", Quantity, `The value of the low bound of the reference range.  The low bound of the reference range endpoint is inclusive of the value
		(e.g.  reference range is >=5 - <=9).   If the low bound is omitted,  it is assumed to be meaningless (e.g. reference range is <=2.3).`, func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("high", Quantity, `The value of the high bound of the reference range.  The high bound of the reference range endpoint is inclusive of the value
		(e.g.  reference range is >=5 - <=9).   If the high bound is omitted,  it is assumed to be meaningless (e.g. reference range is >= 2.3).`, func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("meaning", CodeableConcept, "Code for the meaning of the reference range. See http://hl7.org/fhir/ValueSet/referencerange-meaning", func() {
		//Comments: This SHOULD be populated if there is more than one range.
		//Reason for inclusion or contrainment:
	})
	Attribute("age", Range, "Time when observation was collected from subject - the physiologically relevant time.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("text", String, `Text based reference range in an observation which may be used when a quantitative range is not appropriate for an observation. 
		An example would be a reference value of "Negative" or a list or table of 'normals'.`, func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})

})
var Related = Type("Related", func() {
	Description("A  reference to another resource (usually another Observation but could  also be a QuestionnaireAnswer) whose relationship is defined by the relationship type code.")
	//Comments:
	//Reason for inclusion or contrainment:
	Attribute("type", String, "A code specifying the kind of relationship that exists with the target resource. See http://hl7.org/fhir/ValueSet/observation-relationshiptypes", func() {
		//Type: code
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("target", CodeableConcept, "A reference to the observation or [[[QuestionnaireResponse]]] resource that is related to this observation.", func() {
		//Reference to Observation and QuestionnaireResponse
		//Comments:
		//Reason for inclusion or contrainment:
	})

})
var Component = Type("Component", func() {
	Description(`Some observations have multiple component observations.  These component observations are expressed as separate code value pairs that share the same attributes. 
	Examples include systolic and diastolic component observations for blood pressure measurement and multiple component observations for genetics observations.`)
	//Comments:
	//Reason for inclusion or contrainment:
	Attribute("code", CodeableConcept, "Describes what was observed. Sometimes this is called the observation 'code'. See http://hl7.org/fhir/ValueSet/observation-codes", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("valueQuantity", Quantity)
	Attribute("valueCodeableConcept", CodeableConcept)
	Attribute("valueString", String)
	Attribute("valueRange", Range)
	Attribute("valueSampledData", SampleData)
	Attribute("valueAttachment", Attachment)
	Attribute("valueTime", DateTime)
	Attribute("valueDatTime", DateTime)
	Attribute("valuePeriod", Period)
	Attribute("dataAbsentReason", CodeableConcept, "Provides a reason why the expected value in the element Observation.value[x] is missing. See http://hl7.org/fhir/ValueSet/observation-valueabsentreason", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("normalReferenceRange", ArrayOf(ReferenceRange), "Guidance on how to interpret the value by comparison to a normal or recommended range.", func() {
		//Comments:
		//Reason for inclusion or contrainment: Knowing what values are considered "normal" can help evaluate the significance of a particular result.
		//Need to be able to provide multiple reference ranges for different contexts.
	})

})
