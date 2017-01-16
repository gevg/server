package main

import (
	"database/sql"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"reflect"

	"net/http"

	"os"
	"time"

	"net"
	"net/url"
	"strings"

	"github.com/goa-fhir/server/app"
	"github.com/goadesign/goa"
	"github.com/satori/go.uuid"
	"github.com/tidwall/gjson"
)

//"github.com/pquerna/ffjson""

type Sections map[string]json.RawMessage

type patient struct {
	patientJSON json.RawMessage
	app.Patient
}

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

type Results struct {
	Matches []app.Patient
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
	fmt.Println("body")
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
	//var x app.Patient
	var sqlParams []interface{}

	var x json.RawMessage
	//x := &app.Patient{}

	//getJson("http://nprogram.azurewebsites.net/Patient/1?_format=json", patient)
	//getJson("http://localhost:3001", patient)

	//curl -i -H "Accept: application/json+fhir" "https://fhir-open.sandboxcerner.com/dstu2/d075cf8b-3261-481d-97e5-ba6c48d3b41f/MedicationOrder?patient=2744010&status=active"
	// //--A simple custom JSON-parsing example
	// //getJson("https://open-ic.epic.com/FHIR/api/FHIR/DSTU2/Patient/Tbt3KuCY0B5PSrJvCu2j-PlK.aiHsu2xUjUM8bWpetXoB", patient)
	// //getJson("https://open-ic.epic.com/FHIR/api/FHIR/DSTU2/Patient/Tbt3KuCY0B5PSrJvCu2j-PlK.aiHsu2xUjUM8bWpetXoB", patient)
	// //getJson("https://fhir-open.sandboxcerner.com/dstu2/d075cf8b-3261-481d-97e5-ba6c48d3b41f/MedicationOrder?patient=2744010&status=active", patient)
	// //getJson("https://fhir-open.sandboxcerner.com/dstu2/d075cf8b-3261-481d-97e5-ba6c48d3b41f/MedicationOrder?_format=json&patient=2744010&status=active", patient)
	//_, err = json.NewDecoder(file).Decode(patient)
	//, err := cc.Get("http://fhirtest.uhn.ca/baseDstu2/Patient/EXexample")
	r, err := cc.Get("http://fhirtest.uhn.ca/baseDstu2/Observation/2927")
	//r, err := cc.Get("http://nprogram.azurewebsites.net/Patient/1?_format=json")
	//r, err := cc.Get("https://fhir-open-api.smarthealthit.org/Observation/691-bp")

	//r, err := cc.Get("https://fhir-open-api.smarthealthit.org/AllergyIntolerance/59")

	if err != nil {
		fmt.Println("A timeout error occured")
		os.Exit(3)
	}
	defer r.Body.Close()

	body, _ := ioutil.ReadAll(r.Body)

	isJSON := IsJSON(string(body))
	if isJSON == true {
		fmt.Println("Yesssssssssssssssssssssssssssssssssssssssssssssssss")
		_ = json.Unmarshal(body, &x)
	} else {
		fmt.Println("NOoooooooooooooooooooooooooooooooooooooooooooo")
		_ = xml.Unmarshal(body, &x)
	}

	// fmt.Println("body")

	//err = json.Unmarshal([]byte(r), &f)

	// //b, _ := json.Marshal(body)
	// return json.Unmarshal(body, &target)
	fmt.Printf("%s", body)

	fmt.Println()
	value11 := gjson.Get(string(x), "resourceType")

	value1 := gjson.Get(string(body), "*")
	value2 := gjson.Get(string(body), "subject")
	value3 := gjson.Get(string(body), "identifier")
	fmt.Println()
	fmt.Println("value *:", value1.String())

	fmt.Println("resource:", value11.String())
	fmt.Println("Subject:", value2.String())
	fmt.Println("value:", value3.String())

	//_ = json.Unmarshal(x, &patient)
	//_, _ = json.Marshal(string(x))

	//_ = json.NewDecoder(r.Body).Decode(&x)
	fmt.Println()
	//fmt.Printf("Indentifer:", x.Matches[0].Ad, "\n")
	//fmt.Println("Address:", x.Address)
	//fmt.Println("Telecom:", x.Telecom)
	//fmt.Println("CareProvider:", x.CareProvider)
	//fmt.Println("Name:", x.Name)
	//fmt.Printf("Contact:", x.Contact)
	fmt.Println()

	//fmt.Println("x is equal to:", x.Address[0].Text)

	psqlInfo := fmt.Sprintf(" host =% s port =% d user =% s "+"password =% s dbname =% s sslmode = disable", POSTGRES_HOST, POSTGRES_PORT, POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_DB)

	//db, err := sql.Open(" postgres", psqlInfo)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Could not connect to db!")
	}

	tx, err := db.Begin()

	if err != nil {
		fmt.Println("Error beginning transaction")
		return err
	}

	//fmt.Println("unmarshal json", x.Identifier)

	//var query = `INSERT INTO nc.hl7.patient (id, address) VALUES ($1, $2);`
	var query = `INSERT INTO nc.hl7.patient (id, address) VALUES ($1, $2);`
	//var query = `INSERT INTO nc.hl7.patient (id, address) VALUES ($1, json_each($2));`

	//var query = `INSERT INTO nc.hl7.patient (id, address) VALUES ($1, $2, $3, '{"telecom":$4}');`
	//var query = `INSERT INTO nc.hl7.patient (id, active, birth_date, telecom) VALUES ($1, $2, $3, $4);`

	sqlParams = append(sqlParams, uuid.NewV4().String())
	//sqlParams = append(sqlParams, patient.patientJSON)
	//sqlParams = append(sqlParams, *patient.BirthDate)
	//sqlParams = append(sqlParams, patient.Gender)
	//sqlParams = append(sqlParams, patient.Telecom)
	fmt.Println()
	fmt.Println(sqlParams)
	fmt.Println()
	//fmt.Println(patient.Telecom, patient.patientJSON)

	stmt, err := tx.Prepare(query)
	if err != nil {
		fmt.Println("Error beginning transaction")
		fmt.Println(err.Error())
		return err
	}

	defer func() {
		if err == nil {
			log.Println("Commit Item")
			tx.Commit()
		} else {
			log.Println("Rollback Item")
			tx.Rollback()
			fmt.Println(err.Error())

			//return ctx.NotFound()
		}
		stmt.Close()
	}()

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

	g, _ := json.Marshal(x)
	fmt.Println(string(g))
	_, err = stmt.Exec(uuid.NewV4().String(), string(g))
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Error with db")
	}

	//json.Marshal(x)

	type T struct {
		A int
		B string
	}

	t := app.Observation{}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()

	for i := 0; i < s.NumField(); i++ {
		//f := s.Field(i)
		//fmt.Printf("%d: %s %s = %v\n", i,
		//typeOfT.Field(i).Name, f.Type(), f.Interface())
		//test2 := fmt.Sprintf("%T", s.Field(i))
		test2 := fmt.Sprintf("%T", s.Field(i))

		fmt.Println(reflect.TypeOf(test2).Kind())
		//fmt.Println(*test2)

		test := fmt.Sprintf("%s", strings.ToLower(typeOfT.Field(i).Name))
		fmt.Println(test)

		value := gjson.Get(string(x), test)
		fmt.Println()
		fmt.Println("value *:", value.String())

	}

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

func OpenFile(object interface{}) (j json.RawMessage) {
	absPath, _ := filepath.Abs("patient.json")
	//absPath, _ := filepath.Abs("../mypackage/data/file.txt")
	fmt.Println(absPath)
	file, err := os.Open(absPath)
	if err == nil {
		decoder := json.NewDecoder(file)
		err = decoder.Decode(object)
	}
	file.Close()

	//return file
	return nil
}
