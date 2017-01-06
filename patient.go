package main

import (
	"github.com/goa-fhir/server/app"
	"github.com/goadesign/goa"
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

// Read runs the read action.
func (c *PatientController) Read(ctx *app.ReadPatientContext) error {
	// PatientController_Read: start_implement

	// Put your logic here

	// PatientController_Read: end_implement
	res := &app.PatientMedia{}
	return ctx.OK(res)
}

// Search runs the search action.
func (c *PatientController) Search(ctx *app.SearchPatientContext) error {
	// PatientController_Search: start_implement

	// Put your logic here

	// PatientController_Search: end_implement
	res := app.PatientMediaCollection{}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *PatientController) Update(ctx *app.UpdatePatientContext) error {
	// PatientController_Update: start_implement

	// Put your logic here

	// PatientController_Update: end_implement
	return nil
}

// Vread runs the vread action.
func (c *PatientController) Vread(ctx *app.VreadPatientContext) error {
	// PatientController_Vread: start_implement

	// Put your logic here

	// PatientController_Vread: end_implement
	res := &app.PatientMedia{}
	return ctx.OK(res)
}
