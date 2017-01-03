package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
	. "github.com/jamesallain/goa-fhir/design/allergy_intolerance"
)

//vital resource-----------------------------------------------------------------------------------------------------------
var _ = Resource("AllergyIntolerance", func() {

	DefaultMedia(AllergyIntolerance)
	BasePath("allergy_intolerance")
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
			Media(CollectionOf(AllergyIntolerance, func() {
				View("default")
				View("tiny")
			}))
		})
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("show", func() {
		Routing(
			GET("/:allergy_intoleranceID"),
		)
		Description("Retrieve vital with given id")
		Params(func() {
			Param("allergy_intoleranceID", Integer)
		})
		Response(OK)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("watch", func() {
		Routing(
			GET("/:allergy_intoleranceID/watch"),
		)
		Scheme("ws")
		Description("Retrieve vital with given id")
		Params(func() {
			Param("allergy_intoleranceID", Integer)
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
			PATCH("/:allergy_intoleranceID"),
		)
		Params(func() {
			Param("allergy_intoleranceID", Integer)
		})
		Payload(AllergyIntolerancePayload)
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("rate", func() {
		Routing(
			PUT("/:allergy_intoleranceID/actions/rate"),
		)
		Params(func() {
			Param("allergy_intoleranceID", Integer)
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
			DELETE("/:allergy_intoleranceID"),
		)
		Params(func() {
			Param("allergy_intoleranceID", Integer)
		})
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
})
