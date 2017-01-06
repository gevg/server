package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)

//	. "github.com/goadesign/goa/design"

var _ = Resource("public", func() {
	Origin("*", func() {
		Methods("GET", "OPTIONS")
	})
	Files("/ui", "public/html/index.html")
})

var _ = Resource("js", func() {
	Origin("*", func() {
		Methods("GET", "OPTIONS")
	})
	Files("/js/*filepath", "public/js")
})

var _ = Resource("swagger", func() {
	Origin("*", func() {
		Methods("GET", "OPTIONS")
	})
	Files("/swagger.json", "public/swagger/swagger.json")
})
