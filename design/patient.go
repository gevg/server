package design

import (
	. "github.com/goa-fhir/server/design/patient"
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

//patient resource---------------------------------------------------------------------------------------------------------------
//VERB [base]/[type]/[id] {?_format=[mime-type]}
// base: The Service Base URL
// mime-type: The Mime Type
// type: The name of a resource type (e.g. "Patient")
// id: The Logical Id of a resource
// vid: The Version Id of a resource
// compartment: The name of a compartment
// parameters: URL parameters as defined for the particular interaction

// Common HTTP Status codes returned on FHIR-related errors (in addition to normal HTTP errors related to security, header and content type negotiation issues):

// 400 Bad Request - resource could not be parsed or failed basic FHIR validation rules (or multiple matches were found for
// 401 Not Authorized - authorization is required for the interaction that was attempted
// 404 Not Found - resource type not supported, or not a FHIR end-point
// 405 Method Not allowed - the resource did not exist prior to the update, and the server does not allow client defined ids
// 409/412 - version conflict management - see below
// 422 Unprocessable Entity - the proposed resource violated applicable FHIR profiles or server business rules

var _ = Resource("patient", func() {
	Description("This resource uses JWT to secure its endpoints")
	DefaultMedia(PatientMedia)  // Resource default media type
	BasePath("/patients")       // Common resource action path prefix if not ""
	CanonicalActionName("read") // Name of action that returns canonical representation if not "show"
	//UseTrait("Authenticated")  // Included trait if any, can appear more than once

	// Origin("<a href='http://swagger.goa.design'>http://swagger.goa.design</a>", func() { // Define CORS policy, may be prefixed with "*" wildcard
	// 	Headers("X-Shared-Secret") // One or more authorized headers, use "*" to authorize all
	// 	Methods("GET", "POST")     // One or more authorized HTTP methods
	// 	// Methods("GET", "POST", "PUT", "PATCH", "DELETE")
	// 	// Headers("Accept", "Content-Type")
	// 	// Expose("Content-Type", "Origin")
	// 	Expose("X-Time")           // One or more headers exposed to clients
	// 	MaxAge(600)                // How long to cache a preflight request response
	// 	Credentials()              // Sets Access-Control-Allow-Credentials header
	// })

	Response(Unauthorized, ErrorMedia) // Common responses to all actions
	Response(BadRequest, ErrorMedia)

	//DefaultMedia(SuccessMedia)

	// Security(JWT, func() { // Use JWT to auth requests to this endpoint
	// 	Scope("api:access") // Enforce presence of "api" scope in JWT claims.
	// })

	Action("read", func() {
		Description("Retrieve patient with given id. IDs 1 and 2 pre-exist in the system.")
		Payload(PatientPayload) // Request payload is described by the BottlePayload type
		Routing(
			GET("/:patientID"),
		)
		Params(func() {
			Param("patientID", Integer, "Patient ID", func() {
				Minimum(1)
			})
		})
		Response(OK)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("vread", func() {
		Description("Retrieve patient with given id. IDs 1 and 2 pre-exist in the system.")
		Payload(PatientPayload) // Request payload is described by the BottlePayload type
		Routing(
			GET("/:patientID/_history/vid"),
		)
		Params(func() {
			Param("patientID", Integer, "Patient ID", func() {
				Minimum(1)
			})
		})
		Response(OK)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("update", func() {
		Routing(
			PUT("/:patientID"),
		)
		Description("Change patient name")
		Params(func() {
			Param("patientID", Integer, "Patient ID")
		})
		Payload(func() {
			Member("name")
			Required("name")
		})
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("search", func() {
		Description("List all bottles in account optionally filtering by year")
		Payload(PatientPayload) // Request payload is described by the BottlePayload type
		Routing(
			GET(""),
			POST(""),
		)
		Params(func() {
			Param("active", Boolean, "Filter by active")
			Param("birthDate", ArrayOf(DateTime), "Filter by birth date")
			Param("gender", String, "Filter by gender")
			Param("name", ArrayOf(String), "Filter by name")
			//Param("years", ArrayOf(Integer), "Filter by years")
		})
		Response(OK, func() {
			Media(CollectionOf(PatientMedia, func() {
				View("default")
			}))
		})
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("create", func() {
		Routing(
			POST(""),
		)
		Description("Create new patient")
		Payload(func() {
			Member("name")
			Required("name")
		})
		Response(Created, "/patients/[0-9]+")
		Response(BadRequest, ErrorMedia)
	})

	Action("delete", func() {
		Routing(
			DELETE("/:patientID"),
		)
		Params(func() {
			Param("patientID", Integer, "Patient ID")
		})
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
})
