package design

import (
	. "github.com/goa-fhir/server/design/data_types"
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)
// Patient is the patient resource media type------------------------------------------------------------------------------------------------
var Patient = MediaType("application/vnd.patient+json", func() {
	Attributes(func() {
		Attribute("created_at")
		Attribute("created_by")
		Attribute("id")
		Attribute("href")
		//FHIR elements
		Required("gender")
		Attribute("identifier")
		Attribute("active")
		Attribute("name")
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

		Required("id", "href", "created_at", "created_by")
		Attribute("created_at")
		Attribute("created_by")
		Attribute("id")
		Attribute("href")
		//FHIR elements
		Required("gender")
		Attribute("identifier")
		Attribute("active")
		Attribute("name")
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
var Animal = MediaType("application/vnd.animal+json", func() {
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
var Communication = MediaType("application/vnd.communication+json", func() {
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
var HL7Contact = MediaType("application/vnd.contact+json", func() {
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
var HL7Link = MediaType("application/vnd.link+json", func() {
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
