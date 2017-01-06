package design

import (
	. "github.com/goa-fhir/server/design/user"
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

//user resource-------------------------------------------------------------------------------------------------------------------------------
var _ = Resource("user", func() {

	DefaultMedia(UserMedia)
	BasePath("/users")

	Action("list", func() {
		Routing(
			GET(""),
		)
		Description("Retrieve all users.")
		Response(OK, CollectionOf(UserMedia))
	})

	Action("show", func() {
		Routing(
			GET("/:userID"),
		)
		Description("Retrieve user with given id. IDs 1 and 2 pre-exist in the system.")
		Params(func() {
			Param("userID", Integer, "User ID", func() {
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
		Description("Create new user")
		Payload(func() {
			Member("email")
			Member("password")
			Required("email")
			Required("password")
		})
		Response(Created, "/users/[0-9]+")
		Response(BadRequest, ErrorMedia)
	})

	Action("update", func() {
		Routing(
			PUT("/:userID"),
		)
		Description("Change user name")
		Params(func() {
			Param("userID", Integer, "User ID")
		})
		Payload(func() {
			Member("email")
			Member("password")
			Required("email")
			Required("password")
		})
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("delete", func() {
		Routing(
			DELETE("/:userID"),
		)
		Params(func() {
			Param("userID", Integer, "User ID")
		})
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Description("This resource uses JWT to secure its endpoints")
	//DefaultMedia(SuccessMedia)

	// Security(JWT, func() { // Use JWT to auth requests to this endpoint
	// 	Scope("api:access") // Enforce presence of "api" scope in JWT claims.
	// })

	Action("signin", func() {
		Description("Creates a valid JWT")
		//Security(SigninBasicAuth)
		Routing(POST("/jwt/signin"))
		Payload(func() {
			Member("username")
			Required("username")
			Member("password")
			Required("password")
		})

		// Response(NoContent, func() {
		// 	Headers(func() {
		// 		Header("Authorization", String, "Generated JWT")
		// 	})
		// })

		Response(Unauthorized)
		//Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
	Action("signup", func() {
		Description("Signup user")
		//Security(SigninBasicAuth)
		Routing(POST("/jwt/signup"))
		Payload(func() {
			Member("username")
			Required("username")

			Member("email")
			Required("email")

			Member("password")
			Required("password")

			Member("first_name")
			Required("first_name")

			Member("last_name")
			Required("last_name")

			Member("address_line")
			Required("address_line")

			Member("address_city")
			Required("address_city")

			Member("address_state")
			Required("address_state")

			Member("address_postal_code")
			Required("address_postal_code")

		})
		Response(Created, "/users/[0-9]+")
		NoSecurity() // Override the need for auth
		Response(OK)
		//Response(Unauthorized)
		//Response(NoContent)
		//Response(NotFound)
		Response(BadRequest, ErrorMedia)

	})
})
