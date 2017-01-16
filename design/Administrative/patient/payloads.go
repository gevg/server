package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)

// Patient is the patient resource media type------------------------------------------------------------------------------------------------
var PatientPayload = Type("PatientPayload", func() {
	Reference(Patient)
	Attribute("created_at")
	Attribute("created_by")
	Attribute("id")
	Attribute("href")
	//FHIR elements
	Attribute("meta")
	Required("gender")
	Attribute("identifier")
	Attribute("active")
	Attribute("name")
	Attribute("telecom")
	Attribute("gender")
	Attribute("birthDate")
	Attribute("deceasedBoolean")
	Attribute("deceasedDateTime")
	Attribute("address")
	Attribute("multipleBirthBoolean")
	Attribute("multipleBirthInteger")
	Attribute("photo")
	Attribute("contact")
	Attribute("animal")
	Attribute("photo")
	Attribute("communication")
	Attribute("careProvider")
	Attribute("managingOrganization")
	Attribute("link")
})
var AnimalPayload = Type("AnimalPayload", func() {
	Reference(Animal)
	Attribute("species")
	Attribute("breed")
	Attribute("genderStatus")
})
var CommunicationPayload = Type("CommunicationPayload", func() {
	Reference(Communication)
	Attribute("language")
	Attribute("preferred")
	Attribute("careProvider")
})
var HL7ContactPayload = Type("HL7ContactPayload", func() {
	Reference(HL7Contact)
	Attribute("relationship")
	Attribute("name")
	Attribute("telecom")
	Attribute("address")
	Attribute("gender")
	Attribute("organization")
	Attribute("period")
})
var PatientLinkPayload = Type("PatientLinkPayload", func() {
	Reference(PatientLink)
	Attribute("other")
	Attribute("type")
})
