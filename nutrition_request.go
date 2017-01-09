package main

import (
	"github.com/goa-fhir/server/app"
	"github.com/goadesign/goa"
)

// NutritionRequestController implements the NutritionRequest resource.
type NutritionRequestController struct {
	*goa.Controller
}

// NewNutritionRequestController creates a NutritionRequest controller.
func NewNutritionRequestController(service *goa.Service) *NutritionRequestController {
	return &NutritionRequestController{Controller: service.NewController("NutritionRequestController")}
}

// Read runs the read action.
func (c *NutritionRequestController) Read(ctx *app.ReadNutritionRequestContext) error {
	// NutritionRequestController_Read: start_implement

	// Put your logic here

	// NutritionRequestController_Read: end_implement
	// res := &app.NutritionRequestMedia{}
	// return ctx.OK(res)
	return nil
}
