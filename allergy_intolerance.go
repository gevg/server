package main

import (
	"github.com/goa-fhir/server/app"
	"github.com/goadesign/goa"
	"golang.org/x/net/websocket"
	"io"
)

// AllergyIntoleranceController implements the AllergyIntolerance resource.
type AllergyIntoleranceController struct {
	*goa.Controller
}

// NewAllergyIntoleranceController creates a AllergyIntolerance controller.
func NewAllergyIntoleranceController(service *goa.Service) *AllergyIntoleranceController {
	return &AllergyIntoleranceController{Controller: service.NewController("AllergyIntoleranceController")}
}

// Create runs the create action.
func (c *AllergyIntoleranceController) Create(ctx *app.CreateAllergyIntoleranceContext) error {
	// AllergyIntoleranceController_Create: start_implement

	// Put your logic here

	// AllergyIntoleranceController_Create: end_implement
	return nil
}

// Delete runs the delete action.
func (c *AllergyIntoleranceController) Delete(ctx *app.DeleteAllergyIntoleranceContext) error {
	// AllergyIntoleranceController_Delete: start_implement

	// Put your logic here

	// AllergyIntoleranceController_Delete: end_implement
	return nil
}

// List runs the list action.
func (c *AllergyIntoleranceController) List(ctx *app.ListAllergyIntoleranceContext) error {
	// AllergyIntoleranceController_List: start_implement

	// Put your logic here

	// AllergyIntoleranceController_List: end_implement
	res := app.AllergyIntoleranceMediaCollection{}
	return ctx.OK(res)
}

// Rate runs the rate action.
func (c *AllergyIntoleranceController) Rate(ctx *app.RateAllergyIntoleranceContext) error {
	// AllergyIntoleranceController_Rate: start_implement

	// Put your logic here

	// AllergyIntoleranceController_Rate: end_implement
	return nil
}

// Read runs the read action.
func (c *AllergyIntoleranceController) Read(ctx *app.ReadAllergyIntoleranceContext) error {
	// AllergyIntoleranceController_Read: start_implement

	// Put your logic here

	// AllergyIntoleranceController_Read: end_implement
	res := &app.AllergyIntoleranceMedia{}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *AllergyIntoleranceController) Update(ctx *app.UpdateAllergyIntoleranceContext) error {
	// AllergyIntoleranceController_Update: start_implement

	// Put your logic here

	// AllergyIntoleranceController_Update: end_implement
	return nil
}

// Watch runs the watch action.
func (c *AllergyIntoleranceController) Watch(ctx *app.WatchAllergyIntoleranceContext) error {
	c.WatchWSHandler(ctx).ServeHTTP(ctx.ResponseWriter, ctx.Request)
	return nil
}

// WatchWSHandler establishes a websocket connection to run the watch action.
func (c *AllergyIntoleranceController) WatchWSHandler(ctx *app.WatchAllergyIntoleranceContext) websocket.Handler {
	return func(ws *websocket.Conn) {
		// AllergyIntoleranceController_Watch: start_implement

		// Put your logic here

		// AllergyIntoleranceController_Watch: end_implement
		ws.Write([]byte("watch AllergyIntolerance"))
		// Dummy echo websocket server
		io.Copy(ws, ws)
	}
}
