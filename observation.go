package main

import (
	"github.com/goa-fhir/server/app"
	"github.com/goadesign/goa"
	"golang.org/x/net/websocket"
	"io"
)

// ObservationController implements the Observation resource.
type ObservationController struct {
	*goa.Controller
}

// NewObservationController creates a Observation controller.
func NewObservationController(service *goa.Service) *ObservationController {
	return &ObservationController{Controller: service.NewController("ObservationController")}
}

// Create runs the create action.
func (c *ObservationController) Create(ctx *app.CreateObservationContext) error {
	// ObservationController_Create: start_implement

	// Put your logic here

	// ObservationController_Create: end_implement
	return nil
}

// Delete runs the delete action.
func (c *ObservationController) Delete(ctx *app.DeleteObservationContext) error {
	// ObservationController_Delete: start_implement

	// Put your logic here

	// ObservationController_Delete: end_implement
	return nil
}

// List runs the list action.
func (c *ObservationController) List(ctx *app.ListObservationContext) error {
	// ObservationController_List: start_implement

	// Put your logic here

	// ObservationController_List: end_implement
	res := app.ObservationMediaCollection{}
	return ctx.OK(res)
}

// Rate runs the rate action.
func (c *ObservationController) Rate(ctx *app.RateObservationContext) error {
	// ObservationController_Rate: start_implement

	// Put your logic here

	// ObservationController_Rate: end_implement
	return nil
}

// Show runs the show action.
func (c *ObservationController) Show(ctx *app.ShowObservationContext) error {
	// ObservationController_Show: start_implement

	// Put your logic here

	// ObservationController_Show: end_implement
	res := &app.ObservationMedia{}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *ObservationController) Update(ctx *app.UpdateObservationContext) error {
	// ObservationController_Update: start_implement

	// Put your logic here

	// ObservationController_Update: end_implement
	return nil
}

// Watch runs the watch action.
func (c *ObservationController) Watch(ctx *app.WatchObservationContext) error {
	c.WatchWSHandler(ctx).ServeHTTP(ctx.ResponseWriter, ctx.Request)
	return nil
}

// WatchWSHandler establishes a websocket connection to run the watch action.
func (c *ObservationController) WatchWSHandler(ctx *app.WatchObservationContext) websocket.Handler {
	return func(ws *websocket.Conn) {
		// ObservationController_Watch: start_implement

		// Put your logic here

		// ObservationController_Watch: end_implement
		ws.Write([]byte("watch Observation"))
		// Dummy echo websocket server
		io.Copy(ws, ws)
	}
}
