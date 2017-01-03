package main

import (
	"github.com/goadesign/goa"
	"github.com/jamesallain/goa-fhir/app"
)

// PatientController implements the patient resource.
type PatientController struct {
	*goa.Controller
}

// NewPatientController creates a patient controller.
func NewPatientController(service *goa.Service) *PatientController {
	return &PatientController{Controller: service.NewController("PatientController")}
}

// Create runs the create action.
func (c *PatientController) Create(ctx *app.CreatePatientContext) error {
	// PatientController_Create: start_implement

	// Put your logic here

	// PatientController_Create: end_implement
	return nil
}

// Delete runs the delete action.
func (c *PatientController) Delete(ctx *app.DeletePatientContext) error {
	// PatientController_Delete: start_implement

	// Put your logic here

	// PatientController_Delete: end_implement
	return nil
}

// List runs the list action.
func (c *PatientController) List(ctx *app.ListPatientContext) error {
	// PatientController_List: start_implement

	// Put your logic here

	// PatientController_List: end_implement
	res := app.PatientCollection{}
	return ctx.OK(res)
}

// Show runs the show action.
func (c *PatientController) Show(ctx *app.ShowPatientContext) error {
	// PatientController_Show: start_implement

	// Put your logic here

	// PatientController_Show: end_implement
	res := &app.Success{}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *PatientController) Update(ctx *app.UpdatePatientContext) error {
	// PatientController_Update: start_implement

	// Put your logic here

	// PatientController_Update: end_implement
	return nil
}
