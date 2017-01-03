package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)

//	. "github.com/goa-fhir/server/design/data_types"
//	. "github.com/goadesign/goa/design"

// ObservationPayload defines the data structure used in the create vital request body.
// It is also the base type for the vital media type used to render bottles.
var ObservationPayload = Type("ObservationPayload", func() {
	Reference(Observation)
	Attribute("observation", Observation)
	// Attribute("status")
	// Attribute("category")
	// Attribute("code")
	// Attribute("encounter")
	// Attribute("effectiveDateTime")
	// Attribute("effectivePeriod")
	// Attribute("issued")
	// Attribute("performer")
	// Attribute("effectiveDateTime")

	// Attribute("valueQuantity")
	// Attribute("valueCodeableConcept")
	// Attribute("valueString")
	// Attribute("valueRange")
	// Attribute("valueSampledData")
	// Attribute("valueAttachment")
	// Attribute("valueTime")
	// Attribute("valueDatTime")
	// Attribute("valuePeriod")

	// Attribute("dateAbsentReason")
	// Attribute("interpretation")
	// Attribute("comments")
	// Attribute("bodySite")
	// Attribute("method")
	// Attribute("specimen")
	// Attribute("device")
	// Attribute("referenceRange")
	// Attribute("related")
	// Attribute("component")
})
