package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

//	. "github.com/jamesallain/goa-fhir/design/data_types"

// PatientPayload defines the data structure used in the create vital request body.
// It is also the base type for the vital media type used to render bottles.
var PatientPayload = Type("PatientPayload", func() {
	Attribute("active", Boolean, func() {
		Example(true)
	})
	Attribute("gender", func() {
		Enum("male", "female", "other", "unknown")
		Example("male")
	})
	Attribute("birthDate", func() {
		MinLength(4)
		Example("Merlot")
	})
	Attribute("deceased", Boolean, func() {
		Example(false)
	})
	Attribute("multiple_birth", Boolean, func() {
		Example(false)
	})
	Attribute("sweetness", Integer, func() {
		Minimum(1)
		Maximum(5)
	})
	Attribute("country", func() {
		MinLength(2)
		Example("USA")
	})
	Attribute("region", func() {
		Example("Napa Valley")
	})
	Attribute("review", func() {
		MinLength(3)
		MaxLength(300)
		Example("Great and inexpensive")
	})

	Attribute("address", func() { // Nested definition, defines a struct in Go
		Attribute("number", func() {
			Example("hi")

			Attribute("other", String, "Street name")

		})

		Attribute("street", String, "Street name")
		Attribute("city", String, "City")
		Required("city") // The address must contain at least a city
	})

})
