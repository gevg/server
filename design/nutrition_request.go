package design

import (
	. "github.com/goa-fhir/server/design/nutrition_request"
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

//vital resource-----------------------------------------------------------------------------------------------------------
var _ = Resource("NutritionRequest", func() {

	DefaultMedia(NutritionRequestMedia)
	BasePath("nutrition.requests")
	Parent("patient")

	Action("list", func() {
		Routing(
			GET(""),
		)
		Description("List all nutrition_requests in patient optionally filtering by year")
		Params(func() {
			Param("years", ArrayOf(Integer), "Filter by years")
		})
		Response(OK, func() {
			Media(CollectionOf(NutritionRequestMedia, func() {
				View("default")
			}))
		})
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("show", func() {
		Routing(
			GET("/:nutrition_requestID"),
		)
		Description("Retrieve vital with given id")
		Params(func() {
			Param("nutrition_requestID", Integer)
		})
		Response(OK)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("watch", func() {
		Routing(
			GET("/:nutrition_requestID/watch"),
		)
		Scheme("ws")
		Description("Retrieve vital with given id")
		Params(func() {
			Param("nutrition_requestID", Integer)
		})
		Response(SwitchingProtocols)
		Response(BadRequest, ErrorMedia)
	})

	Action("create", func() {
		Routing(
			POST(""),
		)
		Description("Record new vital")
		Payload(NutritionRequestPayload, func() {
			//Required("status", "oralDiet_type")
		})
		Response(Created, "^/accounts/[0-9]+/nutrition_requests/[0-9]+$")
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("update", func() {
		Routing(
			PATCH("/:nutrition_requestID"),
		)
		Params(func() {
			Param("nutrition_requestID", Integer)
		})
		Payload(NutritionRequestPayload)
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("rate", func() {
		Routing(
			PUT("/:nutrition_requestID/actions/rate"),
		)
		Params(func() {
			Param("nutrition_requestID", Integer)
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
			DELETE("/:nutrition_requestID"),
		)
		Params(func() {
			Param("nutrition_requestID", Integer)
		})
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
})
