package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
		. "github.com/jamesallain/goa-fhir/design/public"


	
)

// JWT defines a security scheme using JWT.  The scheme uses the "Authorization" header to lookup
// the token.  It also defines then scope "api".
var JWT = JWTSecurity("jwt", func() {
	Header("Authorization")
	Scope("api:access", "API access") // Define "api:access" scope
})

// BasicAut defines a security scheme using basic authentication. The scheme protects the "signin"
// action used to create JWTs.
var SigninBasicAuth = BasicAuthSecurity("SigninBasicAuth")

// Resource jwt uses the JWTSecurity security scheme.
var _ = Resource("jwt", func() {
	Description("This resource uses JWT to secure its endpoints")
	DefaultMedia(SuccessMedia)

	Security(JWT, func() { // Use JWT to auth requests to this endpoint
		Scope("api:access") // Enforce presence of "api" scope in JWT claims.
	})

	Action("signin", func() {
		Description("Creates a valid JWT")
		Security(SigninBasicAuth)
		Routing(POST("/jwt/signin"))
		Payload(func() {
			Member("username")
			Required("username")
			Member("password")
			Required("password")
		})

		Response(NoContent, func() {
			Headers(func() {
				Header("Authorization", String, "Generated JWT") //
			})
		})

		Response(Unauthorized)
		//Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)

	})

	Action("secure", func() {
		Description("This action is secured with the jwt scheme")
		Routing(GET("/jwt"))
		Params(func() {
			Param("fail", Boolean, "Force auth failure via JWT validation middleware")
		})
		Response(OK)
		Response(Unauthorized)
	})

	Action("unsecure", func() {
		Description("This action does not require auth")
		Routing(GET("/jwt/unsecure"))
		NoSecurity() // Override the need for auth
		Response(OK)
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
