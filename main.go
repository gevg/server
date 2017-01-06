//go:generate goagen bootstrap -d github.com/goa-fhir/server/design

package main

import (
	"github.com/goa-fhir/server/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	// Create service
	service := goa.New("Secure")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "AllergyIntolerance" controller
	c := NewAllergyIntoleranceController(service)
	app.MountAllergyIntoleranceController(service, c)
	// Mount "NutritionRequest" controller
	c2 := NewNutritionRequestController(service)
	app.MountNutritionRequestController(service, c2)
	// Mount "Observation" controller
	c3 := NewObservationController(service)
	app.MountObservationController(service, c3)
	// Mount "basic" controller
	c4 := NewBasicController(service)
	app.MountBasicController(service, c4)
	// Mount "health" controller
	c5 := NewHealthController(service)
	app.MountHealthController(service, c5)
	// Mount "jwt" controller
	c6 := NewJWTController(service)
	app.MountJWTController(service, c6)
	// Mount "patient" controller
	c7 := NewPatientController(service)
	app.MountPatientController(service, c7)
	// Mount "user" controller
	c8 := NewUserController(service)
	app.MountUserController(service, c8)

	// Start service
	if err := service.ListenAndServe(":8088"); err != nil {
		service.LogError("startup", "err", err)
	}
}
