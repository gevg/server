package design

import (
	. "github.com/goa-fhir/server/design/observation"
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

//vital resource-----------------------------------------------------------------------------------------------------------
var _ = Resource("Observation", func() {

	DefaultMedia(ObservationMedia)
	BasePath("observation")
	Parent("patient")

	Action("list", func() {
		Routing(
			GET(""),
		)
		Description("List all observations in patient optionally filtering by year")
		Params(func() {
			Param("years", ArrayOf(Integer), "Filter by years")
		})
		Response(OK, func() {
			Media(CollectionOf(ObservationMedia, func() {
				View("default")
				//View("tiny")
			}))
		})
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("show", func() {
		Routing(
			GET("/:observationID"),
		)
		Description("Retrieve vital with given id")
		Params(func() {
			Param("observationID", Integer)
		})
		Response(OK)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("watch", func() {
		Routing(
			GET("/:observationID/watch"),
		)
		Scheme("ws")
		Description("Retrieve vital with given id")
		Params(func() {
			Param("observationID", Integer)
		})
		Response(SwitchingProtocols)
		Response(BadRequest, ErrorMedia)
	})

	Action("create", func() {
		Routing(
			POST(""),
		)
		Description("Record new vital")
		Payload(ObservationPayload, func() {
			//Required("status", "oralDiet_type")
		})
		Response(Created, "^/patients/[0-9]+/observations/[0-9]+$")
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("update", func() {
		Routing(
			PATCH("/:observationID"),
		)
		Params(func() {
			Param("observationID", Integer)
		})
		Payload(ObservationPayload)
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("rate", func() {
		Routing(
			PUT("/:observationID/actions/rate"),
		)
		Params(func() {
			Param("observationID", Integer)
		})
		// Payload(func() {
		// 	Member("rating")
		// 	Required("rating")
		// })
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("delete", func() {
		Routing(
			DELETE("/:observationID"),
		)
		Params(func() {
			Param("observationID", Integer)
		})
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
})
