package main

import (
	"github.com/goadesign/goa"
	"github.com/goa-fhir/server/app"
	"golang.org/x/net/websocket"
	"io"
)

// VitalController implements the vital resource.
type VitalController struct {
	*goa.Controller
}

// NewVitalController creates a vital controller.
func NewVitalController(service *goa.Service) *VitalController {
	return &VitalController{Controller: service.NewController("VitalController")}
}

// Create runs the create action.
func (c *VitalController) Create(ctx *app.CreateVitalContext) error {
	// VitalController_Create: start_implement

	// Put your logic here

	// VitalController_Create: end_implement
	return nil
}

// Delete runs the delete action.
func (c *VitalController) Delete(ctx *app.DeleteVitalContext) error {
	// VitalController_Delete: start_implement

	// Put your logic here

	// VitalController_Delete: end_implement
	return nil
}

// List runs the list action.
func (c *VitalController) List(ctx *app.ListVitalContext) error {
	// VitalController_List: start_implement

	// Put your logic here

	// VitalController_List: end_implement
	res := app.VitalCollection{}
	return ctx.OK(res)
}

// Rate runs the rate action.
func (c *VitalController) Rate(ctx *app.RateVitalContext) error {
	// VitalController_Rate: start_implement

	// Put your logic here

	// VitalController_Rate: end_implement
	return nil
}

// Show runs the show action.
func (c *VitalController) Show(ctx *app.ShowVitalContext) error {
	// VitalController_Show: start_implement

	// Put your logic here

	// VitalController_Show: end_implement
	res := &app.Vital{}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *VitalController) Update(ctx *app.UpdateVitalContext) error {
	// VitalController_Update: start_implement

	// Put your logic here

	// VitalController_Update: end_implement
	return nil
}

// Watch runs the watch action.
func (c *VitalController) Watch(ctx *app.WatchVitalContext) error {
	c.WatchWSHandler(ctx).ServeHTTP(ctx.ResponseWriter, ctx.Request)
	return nil
}

// WatchWSHandler establishes a websocket connection to run the watch action.
func (c *VitalController) WatchWSHandler(ctx *app.WatchVitalContext) websocket.Handler {
	return func(ws *websocket.Conn) {
		// VitalController_Watch: start_implement

		// Put your logic here

		// VitalController_Watch: end_implement
		ws.Write([]byte("watch vital"))
		// Dummy echo websocket server
		io.Copy(ws, ws)
	}
}
