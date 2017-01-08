package main

import (
	"github.com/goa-fhir/server/app"
	"github.com/goadesign/goa"
)

// ObservationController implements the Observation resource.
type ObservationController struct {
	*goa.Controller
}

// NewObservationController creates a Observation controller.
func NewObservationController(service *goa.Service) *ObservationController {
	return &ObservationController{Controller: service.NewController("ObservationController")}
}

// Show runs the show action.
func (c *ObservationController) Show(ctx *app.ShowObservationContext) error {
	// ObservationController_Show: start_implement

	// Put your logic here

	// ObservationController_Show: end_implement
	res := &app.ObservationMedia{}
	return ctx.OK(res)
}
