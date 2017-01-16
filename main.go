//go:generate goagen bootstrap -d github.com/goa-fhir/server/design

package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/goa-fhir/server/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	_ "github.com/lib/pq"
)

//	"github.com/jinzhu/gorm"

// var (
// 	// ErrUnauthorized is the error returned for unauthorized requests.
// 	ErrUnauthorized = goa.NewErrorClass("unauthorized", 401)
// 	db              *gorm.DB
// 	logger          log15.Logger
// 	patient_db      *models.PatientDB
// 	//vitals_db       *models.VitalDB
// 	user_db *models.UserDB
// )

const (
	POSTGRES_PASSWORD = "iceman22"
	POSTGRES_USER     = "postgres"
	POSTGRES_DB       = "nc"
	POSTGRES_HOST     = "localhost"
	POSTGRES_PORT     = 5432
)

func main() {

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

	// db.LogMode(true)

	// db.DropTable(&models.User{}, &models.Patient{})
	// db.AutoMigrate(&models.User{}, &models.Patient{})

	// user_db = models.NewUserDB(db)
	// patient_db = models.NewPatientDB(db)
	// db.DB().SetMaxOpenConns(50)

	// Create service
	service := goa.New("Secure")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// // Mount security middlewares
	// jwtMiddleware, err := NewJWTMiddleware()
	// exitOnFailure(err)
	// app.UseJWTMiddleware(service, jwtMiddleware)

	// // Security middleware used to secure the creation of JWT tokens.
	// app.UseSigninBasicAuthMiddleware(service, NewBasicAuthMiddleware())

	// Mount "js" controller
	// c4 := NewJsController(service)
	// app.MountJsController(service, c4)
	// // Mount "JWT" controller
	// s1, err := NewJWTController(service)
	// exitOnFailure(err)
	// app.MountJWTController(service, s1)

	// Mount "AllergyIntolerance" controller
	c := NewAllergyIntoleranceController(service)
	app.MountAllergyIntoleranceController(service, c)
	// Mount "NutritionRequest" controller
	c2 := NewNutritionRequestController(service)
	app.MountNutritionRequestController(service, c2)
	// Mount "Observation" controller
	c3 := NewObservationController(service)
	app.MountObservationController(service, c3)
	// Mount "health" controller
	c4 := NewHealthController(service)
	app.MountHealthController(service, c4)
	// Mount "js" controller
	// c5 := NewJsController(service)
	// app.MountJsController(service, c5)
	// Mount "oauth2_provider" controller
	c6 := NewOauth2ProviderController(service)
	app.MountOauth2ProviderController(service, c6)
	// Mount "patient" controller
	c7 := NewPatientController(service)
	app.MountPatientController(service, c7)
	// Mount "public" controller
	c8 := NewPublicController(service)
	app.MountPublicController(service, c8)
	// Mount "swagger" controller
	c9 := NewSwaggerController(service)
	app.MountSwaggerController(service, c9)
	// Mount "user" controller
	c10 := NewUserController(service)
	app.MountUserController(service, c10)
	// Start service
	if err := service.ListenAndServe(":8088"); err != nil {
		service.LogError("startup", "err", err)
	}
}

// exitOnFailure prints a fatal error message and exits the process with status 1.
func exitOnFailure(err error) {
	if err == nil {
		return
	}
	fmt.Fprintf(os.Stderr, "[CRIT] %s", err.Error())
	os.Exit(1)
}
