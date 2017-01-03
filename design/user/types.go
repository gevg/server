package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// UserPayload defines the data structure used in the create vital request body.
// It is also the base type for the vital media type used to render bottles.
var UserPayload = Type("UserPayload", func() {

	Attribute("username", func() {
		MinLength(3)
		Example("Jim")
	})
	Attribute("email", func() {
		MinLength(4)
		Example("jim.smith@gmail.com")
	})
	Attribute("password", func() {
		MinLength(2)
		Example("password")
	})
	Attribute("first_name", func() {
		MinLength(2)
		Example("Jim")
	})
	Attribute("last_name", func() {
		MinLength(2)
		Example("Smith")
	})
	Attribute("address_line", func() {
		MinLength(2)
		Example("533 Worth Ct")
	})
	Attribute("address_city", func() {
		MinLength(2)
		Example("Carmel")
	})
	Attribute("address_state", func() {
		MinLength(2)
		Example("IN")
	})
	Attribute("address_postal_code", func() {
		MinLength(5)
		Example(46032)
	})

	// Attribute("address", func() { // Nested definition, defines a struct in Go
	// 	Attribute("number", Integer, "Street number")
	// 	Attribute("street", String, "Street name")
	// 	Attribute("city", String, "City")
	// 	Required("city") // The address must contain at least a city
	// })
	//Attribute("city", String, "City")
	//Attribute("data", HashOf(String, String))

})

// User is the user resource media type.
var User = MediaType("application/vnd.user+json", func() {
	TypeName("User")
	Reference(UserPayload)
	Description("A user of the API")
	Attributes(func() {

		Required("id", "href", "username", "email", "password", "first_name", "last_name", "created_at", "created_by")

		Attribute("id", Integer, "ID of user", func() {
			Example(1)
		})
		Attribute("href", String, "API href of user", func() {
			Example("/user/1")
		})
		Attribute("username", String, "Username of user", func() {
			Example("jim")
		})
		Attribute("email", String, "Email of user", func() {
			Format("email")
			Example("jim.smith@gmail.com")
		})
		Attribute("password", String, "Password of user", func() {
			Example("password")
		})
		Attribute("first_name", String, "First name of user", func() {
			Example("Jim")
		})
		Attribute("last_name", String, "Last name of user", func() {
			Example("Smith")
		})
		Attribute("address_line", String, "Street name, number, direction & P.O. Box etc.", func() {
			Example("533 Worth Ct")
		})
		Attribute("address_city", String, "Name of city, town etc.", func() {
			Example("Carmel")
		})
		Attribute("address_state", String, "Sub-unit of country (abbreviations ok)", func() {
			Example("IN")
		})
		Attribute("address_postal_code", Integer, "Postal code for area", func() {
			Example(46032)
		})

		// Attribute("address", func() { // Nested definition, defines a struct in Go
		// 	Attribute("number", Integer, "Street number")
		// 	Attribute("street", String, "Street name")
		// 	Attribute("city", String, "City")
		// 	Required("city") // The address must contain at least a city
		// })

		//Attribute("friends", ArrayOf("user"))
		//Attribute("data", HashOf(String, String))

		Attribute("created_at", DateTime, "Date of creation")
		Attribute("created_by", String, "Email of user owner", func() {
			Format("email")
			Example("me@goa.design")
		})

	})

	View("default", func() {
		Attribute("id")
		Attribute("href")
		Attribute("username")
		Attribute("email")
		Attribute("first_name")
		Attribute("last_name")
		Attribute("address_line")
		Attribute("address_city")
		Attribute("address_state")
		Attribute("address_postal_code")

		Attribute("created_at")
		Attribute("created_by")
		// Attribute("address", func() { // Nested definition, defines a struct in Go
		// 	Attribute("number")
		// 	Attribute("street")
		// 	Attribute("city")
		// })

	})

	View("tiny", func() {
		Description("tiny is the view used to list users")
		Attribute("id")
		Attribute("href")
		Attribute("username")
		Attribute("email")
		Attribute("first_name")
		Attribute("last_name")
	})

	View("link", func() {
		Attribute("id")
		Attribute("href")
	})
})

//user resource-------------------------------------------------------------------------------------------------------------------------------
var _ = Resource("user", func() {

	DefaultMedia(User)
	BasePath("/users")

	Action("list", func() {
		Routing(
			GET(""),
		)
		Description("Retrieve all users.")
		Response(OK, CollectionOf(User))
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
