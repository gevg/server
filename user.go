package main

import (
	"fmt"

	"github.com/goa-fhir/server/app"
	"github.com/goadesign/goa"
)

// UserController implements the user resource.
type UserController struct {
	*goa.Controller
}

// NewUserController creates a user controller.
func NewUserController(service *goa.Service) *UserController {
	return &UserController{Controller: service.NewController("UserController")}
}

// Create runs the create action.
func (c *UserController) Create(ctx *app.CreateUserContext) error {
	// UserController_Create: start_implement

	// Put your logic here

	// UserController_Create: end_implement
	return nil
}

// Delete runs the delete action.
func (c *UserController) Delete(ctx *app.DeleteUserContext) error {
	// UserController_Delete: start_implement

	// Put your logic here

	// UserController_Delete: end_implement
	return nil
}

// List runs the list action.
func (c *UserController) List(ctx *app.ListUserContext) error {
	// UserController_List: start_implement
	fmt.Println("signin jjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjj")

	// Put your logic here

	// UserController_List: end_implement
	res := app.UserMediaCollection{}
	return ctx.OK(res)
}

// Secure runs the secure action.
func (c *UserController) Secure(ctx *app.SecureUserContext) error {
	// UserController_Secure: start_implement
	fmt.Println("signin jjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjj")

	// Put your logic here

	// UserController_Secure: end_implement
	res := &app.Success{}
	return ctx.OK(res)
}

// Show runs the show action.
func (c *UserController) Show(ctx *app.ShowUserContext) error {
	// UserController_Show: start_implement
	fmt.Println("signin jjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjj")

	// Put your logic here

	// UserController_Show: end_implement
	res := &app.UserMedia{}
	return ctx.OK(res)
}

// Signin runs the signin action.
func (c *UserController) Signin(ctx *app.SigninUserContext) error {
	// UserController_Signin: start_implement

	// Put your logic here
	fmt.Println("signin jjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjj")
	// UserController_Signin: end_implement
	return nil
}

// Signup runs the signup action.
func (c *UserController) Signup(ctx *app.SignupUserContext) error {
	// UserController_Signup: start_implement
	fmt.Println("signin jjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjj")

	// Put your logic here

	// UserController_Signup: end_implement
	res := &app.UserMedia{}
	return ctx.OK(res)
}

// Unsecure runs the unsecure action.
func (c *UserController) Unsecure(ctx *app.UnsecureUserContext) error {
	// UserController_Unsecure: start_implement
	fmt.Println("signin jjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjj")

	// Put your logic here

	// UserController_Unsecure: end_implement
	res := &app.UserMedia{}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *UserController) Update(ctx *app.UpdateUserContext) error {
	// UserController_Update: start_implement

	// Put your logic here

	// UserController_Update: end_implement
	return nil
}
