package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Bundle is the Bundle resource media type.
var BundleMedia = MediaType("application/vnd.bundle+json", func() {
	TypeName("BundleMedia")
	Reference(BundlePayload)
	Description("A Bundle of the API")
	Attributes(func() {
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
		//Attribute("friends", ArrayOf("user"))
		//Attribute("data", HashOf(String, String))
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
