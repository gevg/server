package design

import (
	. "github.com/goa-fhir/server/design/data_types"
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var Specimen = Type("Specimen", func() {
	Description("Demographics and other administrative information about an individual or animal receiving care or other health-related services.")
	//Comments:
	//Reason for inclusion or contrainment:
	Attribute("identifier", Identifier, "Id for specimen.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("status", String, "The availability of the specimen. See http://hl7.org/fhir/ValueSet/specimen-status", func() {
		//Comments:
		//Reason for inclusion or contrainment:
		Enum("available", "unavailable", "unsatisfactory", "entered-in-error")
	})
	Attribute("type", CodeableConcept, "The kind of material that forms the specimen. See http://hl7.org/fhir/ValueSet/v2-0487 ", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("parent", HL7Reference, "Reference to the parent (source) specimen which is used when the specimen was either derived from or a component of another specimen.", func() {
		//Type: Reference to Specimen
		//Comments: The parent specimen could be the source from which the current specimen is derived by some processing step (e.g. an aliquot or isolate or extracted nucleic acids
		//from clinical samples) or one of many specimens that were combined to create a pooled sample.
		//Reason for inclusion or contrainment:
	})
	Attribute("subject", HL7Reference, "Where the specimen came from. This may be from the patient(s) or from the environment or a device.", func() {
		//Reference Patient, Group, Device, Substance
		//Comments:
		//Reason for inclusion or contrainment: Must know the subject context.
	})
	Attribute("accessionIdentifier", Identifier, "The identifier assigned by the lab when accessioning specimen(s). This is not necessarily the same as the specimen identifier, depending on local lab procedures.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("receivedTime", DateTime, "Time when specimen was received for processing or testing.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("collection", Collection, "Details concerning the specimen collection.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("treatment", Treatment, "Details concerning treatment and processing steps for the specimen.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("container", Container, "The container holding the specimen.  The recursive nature of containers; i.e. blood in tube in tray in rack is not addressed here.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
})
var Collection = Type("Collection", func() {
	Description("Details concerning the specimen collection.")
	//Comments:
	//Reason for inclusion or contrainment:
	Attribute("collecter", CodeableConcept, "Person who collected the specimen.", func() {
		//Reference to Practitioner
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("comment", String, "To communicate any details or issues encountered during the specimen collection procedure.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("collectedDateTime", DateTime, "Time when specimen was collected from subject - the physiologically relevant time.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("collectedPeriod", Period, "Time when specimen was collected from subject - the physiologically relevant time.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("quantity", Quantity, "The quantity of specimen collected; for instance the volume of a blood sample, or the physical measurement of an anatomic pathology sample.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("method", CodeableConcept, "A coded value specifying the technique that is used to perform the procedure. See http://hl7.org/fhir/ValueSet/specimen-collection-method", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("bodySite", CodeableConcept, `Anatomical location from which the specimen was collected (if subject is a patient). This is the target site. 
		This element is not used for environmental specimens. See http://hl7.org/fhir/ValueSet/body-site`, func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
})
var Container = Type("Container", func() {
	Description("A contact party (e.g. guardian, partner, friend) for the specimen.")
	//Comments:
	//Reason for inclusion or contrainment:
	Attribute("relationship", CodeableConcept, "The nature of the relationship between the specimen and the contact person.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("identifier", Identifier, "Id for container. There may be multiple; a manufacturer's bar code, lab assigned identifier, etc. The container ID may differ from the specimen id in some circumstances.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("description", String, "Textual description of the container.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("type", CodeableConcept, "The type of container associated with the specimen (e.g. slide, aliquot, etc.). See http://hl7.org/fhir/ValueSet/specimen-container-type", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("capacity", Quantity, "The capacity (volume or other measure) the container may contain.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("specimenQuantity", HL7Reference, "The quantity of specimen in the container; may be volume, dimensions, or other appropriate measurements, depending on the specimen type.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("additiveCodeableConcept", CodeableConcept, "The period during which this contact person or organization is valid to be contacted relating to this specimen. See http://hl7.org/fhir/ValueSet/v2-0371", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("additiveReference", HL7Reference, "The period during which this contact person or organization is valid to be contacted relating to this specimen.", func() {
		//Reference to Substance
		//Comments:
		//Reason for inclusion or contrainment:
	})
})
var Treatment = Type("Treatment", func() {
	Description("Details concerning treatment and processing steps for the specimen.")
	//Comments:
	//Reason for inclusion or contrainment:
	Attribute("description", String, "Textual description of procedure.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("procedure", CodeableConcept, "A coded value specifying the procedure used to process the specimen. See http://hl7.org/fhir/ValueSet/specimen-treatment-procedure", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("additive", HL7Reference, "Material used in the processing step.", func() {
		//Reference to Substance
		//Comments:
		//Reason for inclusion or contrainment:
	})
})
