package main

import (
	"fmt"

	"github.com/goa-fhir/server/app"
	"github.com/goadesign/goa"
)

// AllergyIntoleranceController implements the AllergyIntolerance resource.
type AllergyIntoleranceController struct {
	*goa.Controller
}

// NewAllergyIntoleranceController creates a AllergyIntolerance controller.
func NewAllergyIntoleranceController(service *goa.Service) *AllergyIntoleranceController {
	return &AllergyIntoleranceController{Controller: service.NewController("AllergyIntoleranceController")}
}

// Read runs the read action.
func (c *AllergyIntoleranceController) Read(ctx *app.ReadAllergyIntoleranceContext) error {
	// AllergyIntoleranceController_Read: start_implement
	fmt.Println("AllergyIntolerance read")
	// Put your logic here

	// AllergyIntoleranceController_Read: end_implement
	res := &app.AllergyIntoleranceMedia{}
	return ctx.OK(res)
}
