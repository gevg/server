package design

import (
	. "github.com/goa-fhir/server/design/DataType/data_types"
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// TestScript is the user resource media type.
var TestScript = Type("TestScript", func() {
	Description("A user of the API")
	Required("id", "href", "username", "email", "password", "first_name", "last_name", "created_at", "created_by")

	Attribute("id", Integer, "ID of user", func() {
		Example(1)
	})
	Attribute("href", String, "API href of user", func() {
		Example("/user/1")
	})

	Attribute("url", String, "An absolute URL that is used to identify this Test Script. This SHALL be a URL, SHOULD be globally unique, and SHOULD be an address at which this Test Script is (or will be) published.", func() {
	})
	Attribute("version", String, "The identifier that is used to identify this version of the TestScript. This is an arbitrary value managed by the TestScript author manually.", func() {
	})
	Attribute("name", String, "A free text natural language name identifying the TestScript.", func() {
	})

	Attribute("status", String, "See http://hl7.org/fhir/ValueSet/conformance-resource-status.", func() {
	})
	Attribute("identifier", Identifier, "Identifier for the TestScript assigned for external purposes outside the context of FHIR.", func() {
	})
	Attribute("experimental", Boolean, "Street name, number, direction & P.O. Box etc.", func() {
	})
	Attribute("contact", Contact, "Contacts to assist a user in finding and communicating with the publisher.", func() {
	})
	Attribute("date", DateTime, `The date this version of the test tcript was published. The date must change when the business version changes, 
	if it does, and it must change if the status code changes. In addition, it should change when the substantive content of the test cases change.`, func() {
	})
	Attribute("description", String, "A free text natural language description of the TestScript and its use.", func() {
	})

	Attribute("useContext", CodeableConcept, `The content was developed with a focus and intent of supporting the contexts that are listed. 
	These terms may be used to assist with indexing and searching of Test Scripts.`, func() {
	})
	Attribute("requirements", String, "Explains why this Test Script is needed and why it's been constrained as it has.", func() {
	})
	Attribute("copyright", String, `A copyright statement relating to the Test Script and/or its contents. Copyright statements are generally legal 
	restrictions on the use and publishing of the details of the constraints and mappings.`, func() {
	})

	
})



var MetaData = Type("TestScript", func() {
	Description("A user of the API")

})