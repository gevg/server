package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// This is the cellar application API design used by goa to generate
// the application code, client, tests, documentation etc.
var _ = API("Secure", func() {
	Description("This artifical API shows 4 different ways to secure API endpoints: using " +
		"basic auth, shared secret header, JWT or OAuth2")
	Scheme("http")
	Host("localhost:8088")

	Consumes("application/json")
	Produces("application/json")

	//original----------------------------------------------------------------------------------------
	Title("The virtual wine cellar")
	//Description("A basic example of a CRUD API implemented with goa")
	Contact(func() {
		Name("nosh team")
		Email("james.allin@noshdata.com")
		URL("http://noshdata.com")
	})
	License(func() {
		Name("MIT")
		URL("https://github.com/goadesign/goa/blob/master/LICENSE")
	})
	Docs(func() {
		Description("goa guide")
		URL("http://goa.design/getting-started.html")
	})

	BasePath("/nosh")

	Origin("http://swagger.goa.design", func() {
		Methods("GET", "POST", "PUT", "PATCH", "DELETE")
		Headers("Accept", "Content-Type")
		Expose("Content-Type", "Origin")
		MaxAge(600)
		Credentials()
	})

	Origin("http://localhost:8080", func() {
		Methods("GET", "POST", "PUT", "PATCH", "DELETE")
		Headers("Accept", "Content-Type", "Authorization")
		//Header("Access-Control-Expose-Headers")
		Expose("Content-Type", "Origin", "Authorization")

		MaxAge(600)
		Credentials()
	})

	ResponseTemplate(Created, func(pattern string) {
		Description("Resource created")
		Status(201)
		Headers(func() {
			Header("Location", String, "href to created resource", func() {
				Pattern(pattern)
			})
		})
	})

})
