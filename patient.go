package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/goa-fhir/server/app"
	"github.com/goadesign/goa"
	"github.com/satori/go.uuid"
	"github.com/tidwall/gjson"
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
	fmt.Println("Patient create")

	// Put your logic here

	// PatientController_Create: end_implement
	return nil
}

// Delete runs the delete action.
func (c *PatientController) Delete(ctx *app.DeletePatientContext) error {
	// PatientController_Delete: start_implement
	fmt.Println("Patient delete")

	// Put your logic here

	// PatientController_Delete: end_implement
	return nil
}

// Read runs the read action.
func (c *PatientController) Read(ctx *app.ReadPatientContext) error {
	var sqlParams []interface{}
	x := json.RawMessage{}

	//getJson("https://open-ic.epic.com/FHIR/api/FHIR/DSTU2/Patient/Tbt3KuCY0B5PSrJvCu2j-PlK.aiHsu2xUjUM8bWpetXoB", &x)
	getJson("http://fhirtest.uhn.ca/baseDstu2/Patient/EXexample", x)
	//, err := cc.Get("http://fhirtest.uhn.ca/baseDstu2/Patient/EXexample")

	fmt.Println("start", x)
	//Setting up database ---------------------------------------------------------------------------------
	psqlInfo := fmt.Sprintf("host =% s port =% d user =% s "+"password =% s dbname =% s sslmode = disable", POSTGRES_HOST, POSTGRES_PORT, POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_DB)
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
	var query = `INSERT INTO nc.hl7.patient (
	patient_id,
	meta,
	identifier,
	active,
	telecom,
	gender,
	birth_date,
	deceased_boolean,
	deceased_date_time,
	address,
	marital_status,
	photo,
	contact,
	animal,
	communication,
	care_provider,
	managing_organization,
	link
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18);`

	sqlParams = append(sqlParams, uuid.NewV4().String())
	//sqlParams = append(sqlParams, patient.patientJSON)
	//sqlParams = append(sqlParams, *patient.BirthDate)
	//sqlParams = append(sqlParams, patient.Gender)
	//sqlParams = append(sqlParams, patient.Telecom)

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

	fmt.Println(string(x))
	ResourceType := gjson.Get(string(x), "resourceType")
	fmt.Println("resourceType", ResourceType)

	result := gjson.Get(string(x), "entry.#.resource")
	result.ForEach(func(key, value gjson.Result) bool {
		// o := app.Observation{}
		o := new(app.Observation)

		//fmt.Println(key.Value, value.String())
		//sub := gjson.Get(value.String(), "component")
		// for _, name := range sub.Array() {
		// 	//println(name.String())
		// 	fmt.Println(name.Get("code.coding.#.valueQuantity.value"))
		// }
		//fmt.Println(gjson.Parse(value.String()).Get("display"))
		json.Unmarshal([]byte(value.String()), &o)
		// for _,i range len(o.Component){
		// 	//fmt.Println(o.Component.)
		// }
		// record := app.Observation{}
		// record.Component = o.Component

		_, err := json.Marshal(o.Component)
		if err != nil {
			fmt.Println(err.Error())
		}

		json.Marshal(o)
		//test := &o
		fmt.Printf("%s", &o.Component)

		return true // keep iterating
	})

	_, err = stmt.Exec(uuid.NewV4().String(), x)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Error with db")
	}

	res := &app.PatientMedia{}
	return ctx.OK(res)
}

// Search runs the search action.
func (c *PatientController) Search(ctx *app.SearchPatientContext) error {
	// PatientController_Search: start_implement
	fmt.Println("Patient search")

	// Put your logic here

	// PatientController_Search: end_implement
	res := app.PatientMediaCollection{}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *PatientController) Update(ctx *app.UpdatePatientContext) error {
	// PatientController_Update: start_implement
	fmt.Println("Patient update")

	// Put your logic here

	// PatientController_Update: end_implement
	return nil
}

// Vread runs the vread action.
func (c *PatientController) Vread(ctx *app.VreadPatientContext) error {
	// PatientController_Vread: start_implement
	fmt.Println("Patient vread")

	// Put your logic here

	// PatientController_Vread: end_implement
	res := &app.PatientMedia{}
	return ctx.OK(res)
}

var cc = &http.Client{Timeout: time.Second}

func getJson(url string, target json.RawMessage) error {
	r, err := cc.Get(url)
	if err != nil && hasTimedOut(err) {
		fmt.Println("A timeout error occured")
		os.Exit(3)
	}
	defer r.Body.Close()

	//io.Copy(os.Stdout, r.Body)
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(body))
	//return json.Marshal(body)
	_ = json.Unmarshal(body, &target)
	json.Marshal(&target)
	return nil
	//return json.NewDecoder(r.Body).Decode(target)
}
