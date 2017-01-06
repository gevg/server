package main

import (
	"github.com/goa-fhir/server/app"
	"github.com/goadesign/goa"
	"golang.org/x/net/websocket"
	"io"
)

// NutritionRequestController implements the NutritionRequest resource.
type NutritionRequestController struct {
	*goa.Controller
}

// NewNutritionRequestController creates a NutritionRequest controller.
func NewNutritionRequestController(service *goa.Service) *NutritionRequestController {
	return &NutritionRequestController{Controller: service.NewController("NutritionRequestController")}
}

// Create runs the create action.
func (c *NutritionRequestController) Create(ctx *app.CreateNutritionRequestContext) error {
	// NutritionRequestController_Create: start_implement

	// Put your logic here

	// NutritionRequestController_Create: end_implement
	return nil
}

// Delete runs the delete action.
func (c *NutritionRequestController) Delete(ctx *app.DeleteNutritionRequestContext) error {
	// NutritionRequestController_Delete: start_implement

	// Put your logic here

	// NutritionRequestController_Delete: end_implement
	return nil
}

// List runs the list action.
func (c *NutritionRequestController) List(ctx *app.ListNutritionRequestContext) error {
	// NutritionRequestController_List: start_implement

	// Put your logic here

	// NutritionRequestController_List: end_implement
	res := app.NutritionRequestMediaCollection{}
	return ctx.OK(res)
}

// Rate runs the rate action.
func (c *NutritionRequestController) Rate(ctx *app.RateNutritionRequestContext) error {
	// NutritionRequestController_Rate: start_implement

	// Put your logic here

	// NutritionRequestController_Rate: end_implement
	return nil
}

// Read runs the read action.
func (c *NutritionRequestController) Read(ctx *app.ReadNutritionRequestContext) error {
	// NutritionRequestController_Read: start_implement

	// Put your logic here

	// NutritionRequestController_Read: end_implement
	res := &app.NutritionRequestMedia{}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *NutritionRequestController) Update(ctx *app.UpdateNutritionRequestContext) error {
	// NutritionRequestController_Update: start_implement

	// Put your logic here

	// NutritionRequestController_Update: end_implement
	return nil
}

// Watch runs the watch action.
func (c *NutritionRequestController) Watch(ctx *app.WatchNutritionRequestContext) error {
	c.WatchWSHandler(ctx).ServeHTTP(ctx.ResponseWriter, ctx.Request)
	return nil
}

// WatchWSHandler establishes a websocket connection to run the watch action.
func (c *NutritionRequestController) WatchWSHandler(ctx *app.WatchNutritionRequestContext) websocket.Handler {
	return func(ws *websocket.Conn) {
		// NutritionRequestController_Watch: start_implement

		// Put your logic here

		// NutritionRequestController_Watch: end_implement
		ws.Write([]byte("watch NutritionRequest"))
		// Dummy echo websocket server
		io.Copy(ws, ws)
	}
}
