//go:generate goagen bootstrap -d github.com/goa-fhir/server/design

package main

import (
	"fmt"
	"os"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/inconshreveable/log15"
	"github.com/goa-fhir/server/app"
	"github.com/goa-fhir/server/models"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var (
	// ErrUnauthorized is the error returned for unauthorized requests.
	ErrUnauthorized = goa.NewErrorClass("unauthorized", 401)
	db              *gorm.DB
	logger          log15.Logger
	patient_db      *models.PatientDB
	vitals_db       *models.VitalDB
	user_db         *models.UserDB
)

const (
	host     = "localhost" //set to ip address of postgres container
	port     = 5432
	user     = "postgres"
	password = "iceman22"
	dbname   = "nc"
)

func main() {

	psqlInfo := fmt.Sprintf(" host =% s port =% d user =% s "+"password =% s dbname =% s sslmode = disable", host, port, user, password, dbname)

	//db, err := sql.Open(" postgres", psqlInfo)
	db, err := gorm.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	db.LogMode(true)

	db.DropTable(&models.User{}, &models.Patient{})
	db.AutoMigrate(&models.User{}, &models.Patient{})

	user_db = models.NewUserDB(db)
	patient_db = models.NewPatientDB(db)
	db.DB().SetMaxOpenConns(50)

	// Create service
	service := goa.New("Secure")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount security middlewares
	jwtMiddleware, err := NewJWTMiddleware()
	exitOnFailure(err)
	app.UseJWTMiddleware(service, jwtMiddleware)

	// Security middleware used to secure the creation of JWT tokens.
	app.UseSigninBasicAuthMiddleware(service, NewBasicAuthMiddleware())

	// Mount "js" controller
	// c4 := NewJsController(service)
	// app.MountJsController(service, c4)
	// // Mount "JWT" controller
	s1, err := NewJWTController(service)
	exitOnFailure(err)
	app.MountJWTController(service, s1)
	// Mount "public" controller
	// c6 := NewPublicController(service)
	// app.MountPublicController(service, c6)
	// Mount "swagger" controller
	//c7 := NewSwaggerController(service)
	//app.MountSwaggerController(service, c7)

	// Mount "NutritionRequest" controller
	c := NewNutritionRequestController(service)
	app.MountNutritionRequestController(service, c)
	// Mount "basic" controller
	c2 := NewBasicController(service)
	app.MountBasicController(service, c2)
	// Mount "health" controller
	c3 := NewHealthController(service)
	app.MountHealthController(service, c3)
	// Mount "jwt" controller
	// c4 := NewJWTController(service)
	// app.MountJWTController(service, c4)
	// Mount "patient" controller
	c5 := NewPatientController(service)
	app.MountPatientController(service, c5)
	// Mount "user" controller
	c6 := NewUserController(service)
	app.MountUserController(service, c6)
	// Mount "vital" controller
	c7 := NewVitalController(service)
	app.MountVitalController(service, c7)
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
