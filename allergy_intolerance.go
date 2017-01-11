package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"net/http"

	"os"
	"time"

	"net"
	"net/url"
	"strings"

	"github.com/goa-fhir/server/app"
	"github.com/goadesign/goa"
)

// AllergyIntoleranceController implements the AllergyIntolerance resource.
type AllergyIntoleranceController struct {
	*goa.Controller
}

var cc = &http.Client{Timeout: time.Second}

func getJson(url string, target interface{}) error {
	r, err := cc.Get(url)
	if err != nil && hasTimedOut(err) {
		fmt.Println("A timeout error occured")
		os.Exit(3)
	}
	defer r.Body.Close()

	//io.Copy(os.Stdout, r.Body)
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("%s", body)

	//b, _ := json.Marshal(body)
	return json.Unmarshal(body, &target)
	//return json.NewDecoder(r.Body).Decode(target)
}

// NewAllergyIntoleranceController creates a AllergyIntolerance controller.
func NewAllergyIntoleranceController(service *goa.Service) *AllergyIntoleranceController {
	return &AllergyIntoleranceController{Controller: service.NewController("AllergyIntoleranceController")}
}

// Read runs the read action.
func (c *AllergyIntoleranceController) Read(ctx *app.ReadAllergyIntoleranceContext) error {
	// AllergyIntoleranceController_Read: start_implement
	//cc := &http.Client{Timeout: time.Second}

	//response, err := http.Get("https://fhir-open-api.smarthealthit.org/AllergyIntolerance/59")
	// r, err := cc.Get("http://nprogram.azurewebsites.net/Patient/1?_format=json")
	// if err != nil && hasTimedOut(err) {
	// 	fmt.Println("A timeout error occured")
	// 	os.Exit(3)
	// }

	// fmt.Println("AllergyIntolerance read: ", response.StatusCode, " ", response.Status)

	//--A simple http GET
	// body, _ := ioutil.ReadAll(response.Body)
	// response.Body.Close()
	// fmt.Printf("%s", body)

	//--A simple custom JSON-parsing example
	foo1 := new(app.Patient)
	getJson("http://nprogram.azurewebsites.net/Patient/1?_format=json", foo1)
	fmt.Println("Name:", foo1.Address)
	//fmt.Println("Length:", len(foo1.Address))

	//fmt.Println("Name:", *foo1.Address[1].PostalCode)

	for f := range foo1.Name {
		fmt.Println("Use:", foo1.Name[f].Use)
		fmt.Println("Family:", foo1.Name[f].Family)
		fmt.Println("Given:", foo1.Name[f].Given)
		fmt.Println("Prefix:", foo1.Name[f].Prefix)
	}

	//fmt.Println("Name:", *foo1.Address[1].PostalCode)

	for f := range foo1.Address {
		fmt.Println("Use:", foo1.Address[f].Use)
		fmt.Println("Line:", foo1.Address[f].Line)
		fmt.Println("City:", foo1.Address[f].City)
		fmt.Println("PostalCode:", foo1.Address[f].PostalCode)
	}

	for f := range foo1.Identifier {
		fmt.Println("Use:", foo1.Identifier[f].Use)
		fmt.Println("Type:", foo1.Identifier[f].Type.Coding)
		fmt.Println("System:", *foo1.Identifier[f].System)
		fmt.Println("Value:", *foo1.Identifier[f].Value)

	}
	// json.Marshal(foo1)
	// fmt.Println("Marshal:", foo1)
	//fmt.Printf("%v\n", *foo1.Gender)

	//foo2 := app.PatientMedia{}
	var foo2 interface{}
	getJson("http://nprogram.azurewebsites.net/Patient/1?_format=json", &foo2)
	fmt.Println("Active:", foo2)
	fmt.Printf("%T\n", foo2)

	//fooEndcode :=

	// var p app.AllergyIntoleranceMedia
	// err = json.Unmarshal([]byte(response.Body), &p)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Fprintln(p)

	//--Parse into an interface
	// var ks = []byte(response)
	// var f interface{}
	// err := json.Unmarshal(ks, &f)
	// if err != nil {
	// 	fmt.Fprintln(err)
	// 	os.Exit(1)
	// }
	// fmt.Fprintln(f)

	res := &app.AllergyIntoleranceMedia{}
	return ctx.OK(res)

	// Put your logic here
	//AllergyIntoleranceController_Read: end_implement

}

func IsJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

func hasTimedOut(err error) bool {
	switch err := err.(type) {
	case *url.Error:
		if err, ok := err.Err.(net.Error); ok && err.Timeout() {
			return true
		}
	case net.Error:
		if err.Timeout() {
			return true
		}
	case *net.OpError:
		if err.Timeout() {
			return true
		}
	}
	errTxt := "use the closed network connection"
	if err != nil && strings.Contains(err.Error(), errTxt) {
		return true
	}
	return false
}
