package main

import (
	"github.com/goa-fhir/server/app"
	"github.com/goadesign/goa"
)

// BasicController implements the basic resource.
type BasicController struct {
	*goa.Controller
}

// NewBasicController creates a basic controller.
func NewBasicController(service *goa.Service) *BasicController {
	return &BasicController{Controller: service.NewController("BasicController")}
}

// Secure runs the secure action.
func (c *BasicController) Secure(ctx *app.SecureBasicContext) error {
	// BasicController_Secure: start_implement

	// Put your logic here

	// BasicController_Secure: end_implement
	res := &app.Success{}
	return ctx.OK(res)
}

// Unsecure runs the unsecure action.
func (c *BasicController) Unsecure(ctx *app.UnsecureBasicContext) error {
	// BasicController_Unsecure: start_implement

	// Put your logic here

	// BasicController_Unsecure: end_implement
	res := &app.Success{}
	return ctx.OK(res)
}
