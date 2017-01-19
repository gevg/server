package main

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	"os"

	"net"
	"net/url"
	"strings"

	"github.com/goa-fhir/server/app"
	"github.com/goadesign/goa"
)

//"github.com/pquerna/ffjson""

// Decode Gob file
func Load(path string, object interface{}) error {
	file, err := os.Open(path)
	if err == nil {
		decoder := json.NewDecoder(file)
		err = decoder.Decode(object)
	}
	file.Close()
	return err
}

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
	//var sqlParams []interface{}
	x := json.RawMessage{}

	absPath, _ := filepath.Abs("bundle.json")
	fmt.Println(absPath)
	file, err := os.Open(absPath)
	if err == nil {
		_ = json.NewDecoder(file).Decode(&x)
	}
	file.Close()

	//getJson("http://nprogram.azurewebsites.net/Patient/1?_format=json", patient)
	//getJson("http://localhost:3001", patient)
	// //getJson("https://open-ic.epic.com/FHIR/api/FHIR/DSTU2/Patient/Tbt3KuCY0B5PSrJvCu2j-PlK.aiHsu2xUjUM8bWpetXoB", patient)
	//_, err = json.NewDecoder(file).Decode(patient)
	//, err := cc.Get("http://fhirtest.uhn.ca/baseDstu2/Patient/EXexample")
	//r, err := cc.Get("http://nprogram.azurewebsites.net/Patient/1?_format=json")

	// if err != nil {
	// 	fmt.Println("A timeout error occured")
	// 	os.Exit(3)
	// }
	// defer r.Body.Close()

	// body, _ := ioutil.ReadAll(r.Body)

	// isJSON := IsJSON(string(body))
	// if isJSON == true {
	// 	fmt.Println("Yesssssssssssssssssssssssssssssssssssssssssssssssss")
	// 	_ = json.Unmarshal(body, &x)
	// } else {
	// 	fmt.Println("NOoooooooooooooooooooooooooooooooooooooooooooo")
	// 	_ = xml.Unmarshal(body, &x)
	// }

	//err = json.Unmarshal([]byte(r), &f)

	// //b, _ := json.Marshal(body)
	// return json.Unmarshal(body, &target)
	//fmt.Printf("%s", x)

	fmt.Println()
	//main_rt := gjson.Get(string(x), "resourceType")
	//sub_rt := gjson.Get(string(x), "entry.#.resource.resourceType")
	// m, ok := gjson.Parse(string(x)).Value().(map[string]interface{})
	// if !ok {
	// 	// not a map
	// }

	//value1 := gjson.Get(string(x), "*")
	//value2 := gjson.Get(string(x), "entry.#.resource.code.coding.#.display")
	//value2 := gjson.Get(string(x), "entry.#.resource.component.#.valueQuantity.value")
	//value2 := gjson.Get(string(x), "entry.#.resource.component.#.code.coding.#.display")

	//value3 := gjson.Get(string(x), "identifier")
	//fmt.Println("value *:", value1.String())

	//fmt.Println("resource:", value11.String())
	//fmt.Println("entry:", value2.String())
	//fmt.Println("value:", value3.String())
	//result := gjson.Get(string(x), "entry.#.resource")

	// //results := gjson.GetMany(string(x), "entry.#.resource.code.coding.#.display", "entry.#.resource.component.#.code.coding.#.display", "entry.#.resource.component.#.valueQuantity.Value")
	// results := gjson.GetMany(string(x), "entry.#.resource.component.#.code.coding.#.display", "entry.#.resource.component.#.valueQuantity.value")

	//_ = json.Unmarshal(x, &patient)
	//_, _ = json.Marshal(string(x))

	//_ = json.NewDecoder(r.Body).Decode(&x)
	//fmt.Println(results)
	//fmt.Printf("Indentifer:", x.Matches[0].Ad, "\n")
	//fmt.Println("Address:", x.Address)
	//fmt.Println("Telecom:", x.Telecom)
	//fmt.Println("CareProvider:", x.CareProvider)
	//fmt.Println("Name:", x.Name)
	//fmt.Printf("Contact:", x.Contact)

	//fmt.Println("x is equal to:", x.Address[0].Text)

	// res := &app.Patient{}
	// res.Active = patient.Active
	// res.BirthDate = patient.BirthDate
	// res.Gender = patient.Gender
	// //res.Telecom = patient.Telecom
	//res.Address=patient.Address
	//res.Address=

	//json.Marshal(patient)

	// 	type int64array []int64

	// func (a int64array) Value() (driver.Value, error) {
	//     // Format a in PostgreSQL's array input format {1,2,3} and return it as as string or []byte.
	// }
	//_, _ = json.Marshal(x)

	//g, _ := json.Marshal(x)
	//fmt.Println(string(g))
	// _, err = stmt.Exec(uuid.NewV4().String(), string(g))
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	fmt.Println("Error with db")
	// }

	//json.Marshal(x)

	// t := app.Observation{}
	// s := reflect.ValueOf(&t).Elem()
	// typeOfT := s.Type()

	// for i := 0; i < s.NumField(); i++ {
	// 	//f := s.Field(i)
	// 	//fmt.Printf("%d: %s %s = %v\n", i,
	// 	//typeOfT.Field(i).Name, f.Type(), f.Interface())
	// 	//test2 := fmt.Sprintf("%T", s.Field(i))
	// 	test2 := fmt.Sprintf("%T", s.Field(i))

	// 	fmt.Println(reflect.TypeOf(test2).Kind())
	// 	//fmt.Println(*test2)

	// 	test := fmt.Sprintf("%s", strings.ToLower(typeOfT.Field(i).Name))
	// 	fmt.Println(test)

	// 	value := gjson.Get(string(x), test)
	// 	fmt.Println()
	// 	fmt.Println("value *:", value.String())

	// }

	//res := &app.AllergyIntoleranceMedia{}
	return nil
	//return ctx.OK(res)

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
