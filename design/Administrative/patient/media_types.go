package design

import (
	. "github.com/goa-fhir/server/design/DataType/data_types"
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Patient is the patient resource media type------------------------------------------------------------------------------------------------
var PatientMedia = MediaType("application/vnd.patient+json", func() {
	TypeName("PatientMedia")
	Reference(PatientPayload)
	Attributes(func() {
		//Local elements
		Attribute("created_at")
		Attribute("created_by")
		Attribute("id")
		Attribute("href")
		//FHIR base elements
		Attribute("resourceType")
		Attribute("id")
		Attribute("extension")
		Attribute("meta")
		//FHIR resource
		//Required("gender", "maritalStatus")
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
	View("default", func() {
		Description("Default view for Patient resource.")

		//Required("id", "href", "created_at", "created_by")
		Attribute("created_at")
		Attribute("created_by")
		Attribute("id")
		Attribute("href")
		//FHIR base elements
		Attribute("resourceType")
		Attribute("id")
		Attribute("extension")
		Attribute("meta")
		//Required("gender", "maritalStatus")
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
	View("link", func() {
		Attribute("id")
		Attribute("href")
	})
})
var AnimalMedia = MediaType("application/vnd.animal+json", func() {
	TypeName("AnimalMedia")
	Reference(AnimalPayload)
	Attributes(func() {
		Attribute("species")
		Attribute("breed")
		Attribute("genderStatus")
	})
	View("default", func() {
		Description("Default view for Link Element.")
		Attribute("species")
		Attribute("breed")
		Attribute("genderStatus")
	})
})
var CommunicationMedia = MediaType("application/vnd.communication+json", func() {
	TypeName("CommunicationMedia")
	Reference(CommunicationPayload)
	Attributes(func() {
		Attribute("language")
		Attribute("preferred")
		Attribute("careProvider")
	})
	View("default", func() {
		Description("Default view for Link Element.")
		Attribute("language")
		Attribute("preferred")
		Attribute("careProvider")
	})
})
var HL7ContactMedia = MediaType("application/vnd.contact+json", func() {
	TypeName("HL7ContactMedia")
	Reference(HL7ContactPayload)
	Attributes(func() {
		Attribute("relationship")
		Attribute("name")
		Attribute("telecom")
		Attribute("address")
		Attribute("gender")
		Attribute("organization")
		Attribute("period")
	})
	View("default", func() {
		Description("Default view for HL7Contact Element.")
		Attribute("relationship")
		Attribute("name")
		Attribute("telecom")
		Attribute("address")
		Attribute("gender")
		Attribute("organization")
		Attribute("period")
	})
})
var PatientLinkMedia = MediaType("application/vnd.link+json", func() {
	TypeName("PatientLinkMedia")
	Reference(PatientLinkPayload)
	Attributes(func() {
		Attribute("other")
		Attribute("type")
	})
	View("default", func() {
		Description("Default view for Patient resource.")
		Attribute("other")
		Attribute("type")
	})
})
