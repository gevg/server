package main

import (
	"fmt"
	"log"

	"net/http"

	"io/ioutil"

	"github.com/goa-fhir/server/app"
	"github.com/goadesign/goa"
)

// AllergyIntoleranceController implements the AllergyIntolerance resource.
type AllergyIntoleranceController struct {
	*goa.Controller
}

// NewAllergyIntoleranceController creates a AllergyIntolerance controller.
func NewAllergyIntoleranceController(service *goa.Service) *AllergyIntoleranceController {
	return &AllergyIntoleranceController{Controller: service.NewController("AllergyIntoleranceController")}
}

// Read runs the read action.
func (c *AllergyIntoleranceController) Read(ctx *app.ReadAllergyIntoleranceContext) error {
	// AllergyIntoleranceController_Read: start_implement

	response, err := http.Get("https://fhir-open-api.smarthealthit.org/AllergyIntolerance/59")
	fmt.Println("AllergyIntolerance read")
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		//_, err := io.Copy(os.Stdout, response.Body)
		if err != nil {
			log.Fatal(err)
		}
	}
	//fmt.Println(response.Body)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fileType := http.DetectContentType(body)

	fmt.Println("FileType: ", fileType)
	//fmt.Println(body)

	// Put your logic here
	//AllergyIntoleranceController_Read: end_implement
	res := &app.AllergyIntoleranceMedia{}
	return ctx.OK(res)

}
