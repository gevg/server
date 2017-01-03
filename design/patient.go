package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
	. "github.com/goa-fhir/server/design/patient"
)

//patient resource---------------------------------------------------------------------------------------------------------------
var _ = Resource("patient", func() {

	DefaultMedia(Patient)
	BasePath("/patients")

	Description("This resource uses JWT to secure its endpoints")
	//DefaultMedia(SuccessMedia)

	// Security(JWT, func() { // Use JWT to auth requests to this endpoint
	// 	Scope("api:access") // Enforce presence of "api" scope in JWT claims.
	// })

	Action("list", func() {

		Description("This action is secured with the jwt scheme")
		Routing(GET("/jwt"))
		Params(func() {
			Param("fail", Boolean, "Force auth failure via JWT validation middleware")
		})
		//Response(OK)
		Response(Unauthorized)

		Routing(
			GET(""),
		)
		Description("Retrieve all patients.")
		Response(OK, CollectionOf(Patient))
	})

	Action("show", func() {
		Routing(
			GET("/:patientID"),
		)
		Description("Retrieve patient with given id. IDs 1 and 2 pre-exist in the system.")
		Params(func() {
			Param("patientID", Integer, "Patient ID", func() {
				Minimum(1)
			})
		})
		Response(OK)
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
