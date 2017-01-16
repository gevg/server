package design

import (
	. "github.com/goa-fhir/server/design/DataType/data_types"
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var ContactDetail = Type("ContactDetail", func() {
	Description("The ContactDetail structure defines general contact details.")
	//Comments:
	//Reason for inclusion or contrainment:
	Attribute("name", String, "	Name of an individual to contact." func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("telecom", ArrayOf(ContactPoint), "Contact details for individual or organization.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	
})
var Contributor = Type("Contributor", func() {
	Description("")
	//Comments:
	//Reason for inclusion or contrainment:
	Attribute("type", String, "author | editor | reviewer | endorser. See http://hl7.org/fhir/ValueSet/contributor-type", func() {
		//Type: code
		//Comments:
		//Reason for inclusion or contrainment:
		Enum("author","editor","reviewer","endorser")
	})
	Attribute("name", string, "Name of the contributor.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("contact", ArrayOf(ContactDetail), "Contact details of the contributor.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
})
