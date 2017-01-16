package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// This is the cellar application API design used by goa to generate
// the application code, client, tests, documentation etc.
var _ = API("goa-FHIR", func() {
	Description("This artifical API shows 4 different ways to secure API endpoints: using " +
		"basic auth, shared secret header, JWT or OAuth2")
	Scheme("https")
	Host("localhost:8088")

	Consumes("application/json")
	Produces("application/json")

	//Produces("application/json+fhir")
	//Consumes("application/json+fhir")

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



	Params(func() {
		//Parameters for all resources
		Param("_id", Integer) // A path parameter defined using e.g. GET("/:id")
		Param("_lastUpdate", DateTime)
		Param("_tag", String)
		Param("_profile", String)  //?
		Param("_security", String) //?
		Param("_text", String)     //?
		Param("_list", String)     //?
		Param("_has", String)      //?
		Param("_type", String)     //?
		Param("_type", String)     //?
		Param("_query", String)    //?
	})

	Params(func() {
		//Search result parameters
		Param("_sort", String, func() { // A query string parameter
			Enum("", "-") //""=asc "-"=desc
		})
		Param("_count", Integer)
		Param("_include", String)       //?
		Param("_revinclude", String)    //?
		Param("_summary", String)       //?
		Param("_elements", String)      //?
		Param("_contained", String)     //?
		Param("_containedType", String) //?
	})

	Origin("*", func() {
		//SMART on FHIR
		// Access-Control-Allow-Credentials: true
		// Access-Control-Allow-Headers: origin, authorization, accept, content-type, x-requested-with
		// Access-Control-Allow-Methods: GET, HEAD, POST, PUT, DELETE, TRACE, OPTIONS
		// Access-Control-Allow-Origin: *
		Methods("GET", "HEAD", "POST", "PUT", "DELETE", "TRACE", "OPTIONS")
		Headers("Origin", "Authorization", "Accept", "Content-Type", "X-Requested-With")
		Expose("Origin", "Authorization", "Accept", "Content-Type", "X-Requested-With")
		MaxAge(600)
		Credentials()
	})

	Origin("http://localhost:8080", func() {
		Methods("GET", "HEAD", "POST", "PUT", "DELETE", "TRACE", "OPTIONS")
		Headers("Origin", "Authorization", "Accept", "Content-Type", "X-Requested-With")
		//Header("Access-Control-Expose-Headers")
		//Expose("Content-Type", "Origin", "Authorization")
		Expose("Origin", "Authorization", "Accept", "Content-Type", "X-Requested-With")

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
