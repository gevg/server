package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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
	fmt.Printf("%s\n %s\n", "hi", ctx.Payload)
	// Put your logic here

	//phone := "14158586273"
	// QueryEscape escapes the phone string so
	// it can be safely placed inside a URL query
	//safePhone := url.QueryEscape(phone)
	//http://sqlonfhir-stu3.azurewebsites.net/fhir/
	//url := fmt.Sprintf("http://apilayer.net/api/validate?access_key=YOUR_ACCESS_KEY&number=%s", safePhone)
	//url := fmt.Sprintf("http://sqlonfhir-stu3.azurewebsites.net/fhir/patient")
	url := fmt.Sprintf("http://demo.oridashi.com.au:8291/patient")

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	// Fill the record with the data from the JSON
	var record Numverify

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	fmt.Println("Gender = ", record.gender)
	fmt.Println("Birth Date   = ", record.birthDate)
	// fmt.Println("Location  = ", record.Location)
	// fmt.Println("Carrier   = ", record.Carrier)
	// fmt.Println("LineType  = ", record.LineType)

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
