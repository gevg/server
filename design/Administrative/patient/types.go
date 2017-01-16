package design

import (
	. "github.com/goa-fhir/server/design/DataType/data_types"
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Patient defines the data structure used in the create vital request body.
// It is also the base type for the vital media type used to render bottles.
var Patient = Type("Patient", func() {
	Description("Demographics and other administrative information about an individual or animal receiving care or other health-related services.")
	//Comments:
	//Reason for inclusion or contrainment:
	//Start Local-------------------------------------------
	Required("id", "href", "created_at", "created_by")

	Attribute("id", Integer, "ID of patient", func() {
		Example(1)
	})
	Attribute("href", String, "API href of patient", func() {
		Example("/patients/1")
	})
	Attribute("created_at", DateTime, "Date of creation")
	Attribute("created_by", String, "Email of patient owner", func() {
		Format("email")
		Example("me@goa.design")
	})
	//End Local-----------------------------------------------

	//Start FHIR----------------------------------------------
	Required("gender", "maritalStatus")
	Attribute("meta", Meta, `The metadata about a resource. This is content in the resource that is maintained by the infrastructure.
	Changes to the content may not always be associated with version changes to the resource.`, func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("identifier", ArrayOf(Identifier), "Patient identifer", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("active", Boolean, "Whether this patient record is in active use.", func() {
		//Comments: Default is true. If a record is inactive, and linked to an active record, then future patient/record updates should occur on the other patient.
		//Reason for inclusion or contrainment: Need to be able to mark a patient record as not to be used because it was created in error.
	})
	Attribute("name", ArrayOf(HumanName), "A name associated with the individual.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("telecom", ArrayOf(ContactPoint), "A contact detail (e.g. a telephone number or an email address) by which the individual may be contacted..", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("gender", String, "Administrative Gender - the gender that the patient is considered to have for administration and record keeping purposes. See http://hl7.org/fhir/ValueSet/administrative-gender", func() {
		//Comments: The gender may not match the biological sex as determined by genetics, or the individual's preferred identification.
		//Reason for inclusion or contrainment: Needed for identification of the individual, in combination with (at least) name and birth date. Gender of individual drives many clinical processes.
		Enum("male", "female", "other", "unknown")
		Example("male")
	})
	Attribute("birthDate", DateTime, "The date of birth for the individual.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
		//Format("date-time")
	})
	Attribute("deceasedBoolean", Boolean, "Indicates if the individual is deceased or not.", func() {
		//From: deceased[x]
		//Comments: If there's no value in the instance it means there is no statement on whether or not the individual is deceased. Most systems will interpret the absence of a
		//value as a sign of the person being alive.
		//Reason for inclusion or contrainment:	The fact that a patient is deceased influences the clinical process. Also, in human communication and relation management it is necessary
		//to know whether the person is alive.
	})
	Attribute("deceasedDateTime", DateTime, "Indicates if the individual is deceased or not.", func() {
		//From: deceased[x]
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("address", ArrayOf(Address), "Addresses for the individual.", func() {
		//Comments: Patient may have multiple addresses with different uses or applicable periods.
		//Reason for inclusion or contrainment: May need to keep track of patient addresses for contacting, billing or reporting requirements and also to help with identification.
	})
	Attribute("maritalStatus", CodeableConcept, "This field contains a patient's most recent marital (civil) status. See http://hl7.org/fhir/ValueSet/marital-status", func() {
		//Comments:
		//Reason for inclusion or contrainment: Most, if not all systems capture it.
	})
	Attribute("multipleBirthBoolean", Boolean, "Indicates whether the patient is part of a multiple or indicates the actual birth order..", func() {
		//From: multipleBirth[x]
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("multipleBirthInteger", Integer, "Indicates whether the patient is part of a multiple or indicates the actual birth order..", func() {
		//From: multipleBirth[x]
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("careProvider", HL7Reference, "Patient's nominated primary care provider", func() {
		//Type: Reference to Organization, Pracitioner
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("photo", ArrayOf(Attachment), "Image of the patient.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("contact", ArrayOf(HL7Contact), "A contact party (e.g. guardian, partner, friend) for the patient.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("animal", Animal, "This patient is known to be an animal.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("communication", ArrayOf(Communication), "Languages which may be used to communicate with the patient about his or her health.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("managingOrganization", HL7Reference, "Organization that is the custodian of the patient record.", func() {
		//Comments: There is only one managing organization for a specific patient record. Other organizations will have their own Patient record, and may use the
		//Link property to join the records together (or a Person resource which can include confidence ratings for the association).
		//Reason for inclusion or contrainment:	Need to know who recognizes this patient record, manages and updates it.
	})
	Attribute("link", ArrayOf(PatientLink), "Link to another patient resource that concerns the same actual patient.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
})
var Animal = Type("Animal", func() {
	Description("This patient is known to be an animal.")
	//Comments: The animal element is labeled "Is Modifier" since patients may be non-human. Systems SHALL either handle patient details appropriately (e.g. inform users patient is not human) or reject declared animal records.
	//Reason for inclusion or contrainment: Many clinical systems are extended to care for animal patients as well as human.
	Attribute("species", CodeableConcept, "Identifies the high level taxonomic categorization of the kind of animal. See http://hl7.org/fhir/ValueSet/animal-species", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("breed", CodeableConcept, "Identifies the detailed categorization of the kind of animal. See http://hl7.org/fhir/ValueSet/animal-breeds", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("genderStatus", CodeableConcept, "Indicates the current state of the animal's reproductive organs. See http://hl7.org/fhir/ValueSet/animal-genderstatus", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
})
var Communication = Type("Communication", func() {
	Description("Languages which may be used to communicate with the patient about his or her health.")
	//Comments:
	//Reason for inclusion or contrainment:
	Attribute("language", CodeableConcept, `The ISO-639-1 alpha 2 code in lower case for the language, optionally followed by a hyphen
		and the ISO-3166-1 alpha 2 code for the region in upper case; e.g. "en" for English, or "en-US" for American English versus "en-EN" for England English.`, func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("preferred", Boolean, "Indicates whether or not the patient prefers this language (over other languages he masters up a certain level).", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("careProvider", HL7Reference, "Patient's nominated care provider.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
})
var HL7Contact = Type("HL7Contact", func() {
	Description("A contact party (e.g. guardian, partner, friend) for the patient.")
	//Comments:
	//Reason for inclusion or contrainment:
	Attribute("relationship", CodeableConcept, "The nature of the relationship between the patient and the contact person.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("name", HumanName, "A name associated with the contact person.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("telecom", ContactPoint, "Address for the contact person.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("address", Address, "Patient's nominated care provider.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("gender", String, "Administrative Gender - the gender that the contact person is considered to have for administration and record keeping purposes.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
		Enum("male", "female", "other", "unknown")
	})
	Attribute("organization", HL7Reference, "Organization on behalf of which the contact is acting or for which the contact is working.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("period", Period, "The period during which this contact person or organization is valid to be contacted relating to this patient.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	}) //Comments:
	//Reason for inclusion or contrainment:
	Attribute("language", CodeableConcept, `The ISO-639-1 alpha 2 code in lower case for the language, optionally followed by a hyphen
		and the ISO-3166-1 alpha 2 code for the region in upper case; e.g. "en" for English, or "en-US" for American English versus "en-EN" for England English.`, func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("preferred", Boolean, "Indicates whether or not the patient prefers this language (over other languages he masters up a certain level).", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("careProvider", HL7Reference, "Patient's nominated care provider.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
})
var PatientLink = Type("link", func() {
	Description("Link to another patient resource that concerns the same actual patient.")
	//Comments: There is no assumption that linked patient records have mutual links.
	//Reason for inclusion or contrainment:There are multiple usecases: * Duplicate patient records due to the clerical errors associated with the difficulties of identifying humans consistently, and * Distribution of patient information across multiple servers.
	Attribute("other", HL7Reference, "The other patient resource that the link refers to.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("type", String, "The type of link between this patient resource and another patient resource. See http://hl7.org/fhir/ValueSet/link-type", func() {
		//Comments:
		//Reason for inclusion or contrainment:
		Enum("replace", "refer", "See also")
	})
})
