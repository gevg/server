package main

import (
	"github.com/goa-fhir/server/app"
	"github.com/goadesign/goa"
)

// Oauth2ProviderController implements the oauth2_provider resource.
type Oauth2ProviderController struct {
	*goa.Controller
}

// NewOauth2ProviderController creates a oauth2_provider controller.
func NewOauth2ProviderController(service *goa.Service) *Oauth2ProviderController {
	return &Oauth2ProviderController{Controller: service.NewController("Oauth2ProviderController")}
}

// Authorize runs the authorize action.
func (c *Oauth2ProviderController) Authorize(ctx *app.AuthorizeOauth2ProviderContext) error {
	// Oauth2ProviderController_Authorize: start_implement

	// Put your logic here

	// Oauth2ProviderController_Authorize: end_implement
	return nil
}

// GetToken runs the get_token action.
func (c *Oauth2ProviderController) GetToken(ctx *app.GetTokenOauth2ProviderContext) error {
	// Oauth2ProviderController_GetToken: start_implement

	// Put your logic here

	// Oauth2ProviderController_GetToken: end_implement
	// res := &app.TokenMedia{}
	// return ctx.OK(res)
	return nil

}
