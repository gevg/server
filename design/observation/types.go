package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
	. "github.com/jamesallain/goa-fhir/design/data_types"
)

// ObservationPayload defines the data structure used in the create vital request body.
// It is also the base type for the vital media type used to render bottles.
var ObservationPayload = Type("ObservationPayload", func() {
	Reference(Observation)
	Attribute("identifier", ArrayOf(Identifier), "A unique identifier for the simple observation instance.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("status", String, "The status of the result value. See http://hl7.org/fhir/ValueSet/observation-status", func() {
		//Comments:
		//Reason for inclusion or contrainment:
		Enum("registered", "preliminary", "final", "amended +")
	})

})
