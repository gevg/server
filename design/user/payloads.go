package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)

// UserPayload defines the data structure used in the create vital request body.
// It is also the base type for the vital media type used to render bottles.
var UserPayload = Type("UserPayload", func() {
	Reference(User)
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
