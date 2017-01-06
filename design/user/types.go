package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)
// User is the user resource media type.
var User = Type("User", func() {
	Description("A user of the API")
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
