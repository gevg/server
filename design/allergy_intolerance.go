package design

import (
	. "github.com/goa-fhir/server/design/allergy_intolerance"
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

//vital resource-----------------------------------------------------------------------------------------------------------
var _ = Resource("AllergyIntolerance", func() {

	DefaultMedia(AllergyIntoleranceMedia)
	BasePath("allergy.intolerance")
	Parent("patient")

	Action("list", func() {
		Routing(
			GET(""),
		)
		Description("List all allergy_intolerances in patient optionally filtering by year")
		Params(func() {
			Param("years", ArrayOf(Integer), "Filter by years")
		})
		Response(OK, func() {
			Media(CollectionOf(AllergyIntoleranceMedia, func() {
				View("default")
			}))
		})
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("read", func() {
		Routing(
			GET("/:allergy.intoleranceID"),
		)
		Description("Retrieve vital with given id")
		Params(func() {
			Param("allergy.intoleranceID", Integer)
		})
		Response(OK)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("watch", func() {
		Routing(
			GET("/:allergy.intoleranceID/watch"),
		)
		Scheme("ws")
		Description("Retrieve vital with given id")
		Params(func() {
			Param("allergy.intoleranceID", Integer)
		})
		Response(SwitchingProtocols)
		Response(BadRequest, ErrorMedia)
	})

	Action("create", func() {
		Routing(
			POST(""),
		)
		Description("Record new vital")
		Payload(AllergyIntolerancePayload, func() {
			//Required("status", "oralDiet_type")
		})
		Response(Created, "^/patients/[0-9]+/allergy_intolerances/[0-9]+$")
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("update", func() {
		Routing(
			PATCH("/:allergy.intoleranceID"),
		)
		Params(func() {
			Param("allergy.intoleranceID", Integer)
		})
		Payload(AllergyIntolerancePayload)
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("rate", func() {
		Routing(
			PUT("/:allergy.intoleranceID/actions/rate"),
		)
		Params(func() {
			Param("allergy.intoleranceID", Integer)
		})
		Payload(func() {
			Member("rating")
			Required("rating")
		})
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("delete", func() {
		Routing(
			DELETE("/:allergy.intoleranceID"),
		)
		Params(func() {
			Param("allergy.intoleranceID", Integer)
		})
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
})
