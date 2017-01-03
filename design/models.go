package design

// goagen --design=github.com/jamesallain/goa-fhir/design gen --pkg-path=github.com/goadesign/gorma

import (
	"github.com/goadesign/gorma"
	. "github.com/goadesign/gorma/dsl"
	. "github.com/jamesallain/goa-fhir/design/patient"
	. "github.com/jamesallain/goa-fhir/design/user"
)

var _ = StorageGroup("Cellar", func() {
	Description("This is the global storage group")
	Store("postgres", gorma.Postgres, func() {
		Description("This is the Postgres relational store")
		Model("Patient", func() {
			RendersTo(Patient)
			Description("Cellar Patient")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
			})
			//Field("name", gorma.String)
			HasMany("Vitals", "Vital")
		})

		// Model("Vital", func() {
		// 	BuildsFrom(func() {
		// 		Payload("vital", "create")
		// 		Payload("vital", "update")
		// 	})
		// 	RendersTo(Vital)
		// 	Field("id", gorma.Integer, func() {
		// 		PrimaryKey()
		// 	})
		// 	//Field("rating", gorma.Integer)
		// 	Description("Vital Model")
		// 	BelongsTo("Patient")
		// })

		Model("User", func() {
			RendersTo(User)
			Description("User information")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
			})
			Field("username", gorma.String)
			Field("email", gorma.String)
			Field("password", gorma.String)
			Field("first_name", gorma.String)
			Field("last_name", gorma.String)
			Field("address_line", gorma.String)
			Field("address_city", gorma.String)
			Field("address_state", gorma.String)
			Field("address_postal_code", gorma.Integer)

			//HasMany("Vitals", "Vital")
		})

	})
})

//cellar "github.com/goadesign/goa-cellar/design"
