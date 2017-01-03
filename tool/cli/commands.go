package cli

import (
	"encoding/json"
	"fmt"
	"github.com/goa-fhir/server/client"
	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
	uuid "github.com/goadesign/goa/uuid"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type (
	// CreateAllergyIntoleranceCommand is the command line data structure for the create action of AllergyIntolerance
	CreateAllergyIntoleranceCommand struct {
		Payload     string
		ContentType string
		// Patient ID
		PatientID   int
		PrettyPrint bool
	}

	// DeleteAllergyIntoleranceCommand is the command line data structure for the delete action of AllergyIntolerance
	DeleteAllergyIntoleranceCommand struct {
		AllergyIntoleranceID int
		// Patient ID
		PatientID   int
		PrettyPrint bool
	}

	// ListAllergyIntoleranceCommand is the command line data structure for the list action of AllergyIntolerance
	ListAllergyIntoleranceCommand struct {
		// Patient ID
		PatientID int
		// Filter by years
		Years       []int
		PrettyPrint bool
	}

	// RateAllergyIntoleranceCommand is the command line data structure for the rate action of AllergyIntolerance
	RateAllergyIntoleranceCommand struct {
		Payload              string
		ContentType          string
		AllergyIntoleranceID int
		// Patient ID
		PatientID   int
		PrettyPrint bool
	}

	// ShowAllergyIntoleranceCommand is the command line data structure for the show action of AllergyIntolerance
	ShowAllergyIntoleranceCommand struct {
		AllergyIntoleranceID int
		// Patient ID
		PatientID   int
		PrettyPrint bool
	}

	// UpdateAllergyIntoleranceCommand is the command line data structure for the update action of AllergyIntolerance
	UpdateAllergyIntoleranceCommand struct {
		Payload              string
		ContentType          string
		AllergyIntoleranceID int
		// Patient ID
		PatientID   int
		PrettyPrint bool
	}

	// WatchAllergyIntoleranceCommand is the command line data structure for the watch action of AllergyIntolerance
	WatchAllergyIntoleranceCommand struct {
		AllergyIntoleranceID int
		// Patient ID
		PatientID   int
		PrettyPrint bool
	}

	// CreateNutritionRequestCommand is the command line data structure for the create action of NutritionRequest
	CreateNutritionRequestCommand struct {
		Payload     string
		ContentType string
		// Patient ID
		PatientID   int
		PrettyPrint bool
	}

	// DeleteNutritionRequestCommand is the command line data structure for the delete action of NutritionRequest
	DeleteNutritionRequestCommand struct {
		NutritionRequestID int
		// Patient ID
		PatientID   int
		PrettyPrint bool
	}

	// ListNutritionRequestCommand is the command line data structure for the list action of NutritionRequest
	ListNutritionRequestCommand struct {
		// Patient ID
		PatientID int
		// Filter by years
		Years       []int
		PrettyPrint bool
	}

	// RateNutritionRequestCommand is the command line data structure for the rate action of NutritionRequest
	RateNutritionRequestCommand struct {
		Payload            string
		ContentType        string
		NutritionRequestID int
		// Patient ID
		PatientID   int
		PrettyPrint bool
	}

	// ShowNutritionRequestCommand is the command line data structure for the show action of NutritionRequest
	ShowNutritionRequestCommand struct {
		NutritionRequestID int
		// Patient ID
		PatientID   int
		PrettyPrint bool
	}

	// UpdateNutritionRequestCommand is the command line data structure for the update action of NutritionRequest
	UpdateNutritionRequestCommand struct {
		Payload            string
		ContentType        string
		NutritionRequestID int
		// Patient ID
		PatientID   int
		PrettyPrint bool
	}

	// WatchNutritionRequestCommand is the command line data structure for the watch action of NutritionRequest
	WatchNutritionRequestCommand struct {
		NutritionRequestID int
		// Patient ID
		PatientID   int
		PrettyPrint bool
	}

	// CreateObservationCommand is the command line data structure for the create action of Observation
	CreateObservationCommand struct {
		Payload     string
		ContentType string
		// Patient ID
		PatientID   int
		PrettyPrint bool
	}

	// DeleteObservationCommand is the command line data structure for the delete action of Observation
	DeleteObservationCommand struct {
		ObservationID int
		// Patient ID
		PatientID   int
		PrettyPrint bool
	}

	// ListObservationCommand is the command line data structure for the list action of Observation
	ListObservationCommand struct {
		// Patient ID
		PatientID int
		// Filter by years
		Years       []int
		PrettyPrint bool
	}

	// RateObservationCommand is the command line data structure for the rate action of Observation
	RateObservationCommand struct {
		ObservationID int
		// Patient ID
		PatientID   int
		PrettyPrint bool
	}

	// ShowObservationCommand is the command line data structure for the show action of Observation
	ShowObservationCommand struct {
		ObservationID int
		// Patient ID
		PatientID   int
		PrettyPrint bool
	}

	// UpdateObservationCommand is the command line data structure for the update action of Observation
	UpdateObservationCommand struct {
		Payload       string
		ContentType   string
		ObservationID int
		// Patient ID
		PatientID   int
		PrettyPrint bool
	}

	// WatchObservationCommand is the command line data structure for the watch action of Observation
	WatchObservationCommand struct {
		ObservationID int
		// Patient ID
		PatientID   int
		PrettyPrint bool
	}

	// SecureBasicCommand is the command line data structure for the secure action of basic
	SecureBasicCommand struct {
		PrettyPrint bool
	}

	// UnsecureBasicCommand is the command line data structure for the unsecure action of basic
	UnsecureBasicCommand struct {
		PrettyPrint bool
	}

	// HealthHealthCommand is the command line data structure for the health action of health
	HealthHealthCommand struct {
		PrettyPrint bool
	}

	// SecureJWTCommand is the command line data structure for the secure action of jwt
	SecureJWTCommand struct {
		// Force auth failure via JWT validation middleware
		Fail        string
		PrettyPrint bool
	}

	// SigninJWTCommand is the command line data structure for the signin action of jwt
	SigninJWTCommand struct {
		Payload     string
		ContentType string
		PrettyPrint bool
	}

	// SignupJWTCommand is the command line data structure for the signup action of jwt
	SignupJWTCommand struct {
		Payload     string
		ContentType string
		PrettyPrint bool
	}

	// UnsecureJWTCommand is the command line data structure for the unsecure action of jwt
	UnsecureJWTCommand struct {
		PrettyPrint bool
	}

	// CreatePatientCommand is the command line data structure for the create action of patient
	CreatePatientCommand struct {
		Payload     string
		ContentType string
		PrettyPrint bool
	}

	// DeletePatientCommand is the command line data structure for the delete action of patient
	DeletePatientCommand struct {
		// Patient ID
		PatientID   int
		PrettyPrint bool
	}

	// ListPatientCommand is the command line data structure for the list action of patient
	ListPatientCommand struct {
		// Force auth failure via JWT validation middleware
		Fail        string
		PrettyPrint bool
	}

	// ShowPatientCommand is the command line data structure for the show action of patient
	ShowPatientCommand struct {
		// Patient ID
		PatientID   int
		PrettyPrint bool
	}

	// UpdatePatientCommand is the command line data structure for the update action of patient
	UpdatePatientCommand struct {
		Payload     string
		ContentType string
		// Patient ID
		PatientID   int
		PrettyPrint bool
	}

	// CreateUserCommand is the command line data structure for the create action of user
	CreateUserCommand struct {
		Payload     string
		ContentType string
		PrettyPrint bool
	}

	// DeleteUserCommand is the command line data structure for the delete action of user
	DeleteUserCommand struct {
		// User ID
		UserID      int
		PrettyPrint bool
	}

	// ListUserCommand is the command line data structure for the list action of user
	ListUserCommand struct {
		PrettyPrint bool
	}

	// ShowUserCommand is the command line data structure for the show action of user
	ShowUserCommand struct {
		// User ID
		UserID      int
		PrettyPrint bool
	}

	// SigninUserCommand is the command line data structure for the signin action of user
	SigninUserCommand struct {
		Payload     string
		ContentType string
		PrettyPrint bool
	}

	// SignupUserCommand is the command line data structure for the signup action of user
	SignupUserCommand struct {
		Payload     string
		ContentType string
		PrettyPrint bool
	}

	// UpdateUserCommand is the command line data structure for the update action of user
	UpdateUserCommand struct {
		Payload     string
		ContentType string
		// User ID
		UserID      int
		PrettyPrint bool
	}
)

// RegisterCommands registers the resource action CLI commands.
func RegisterCommands(app *cobra.Command, c *client.Client) {
	var command, sub *cobra.Command
	command = &cobra.Command{
		Use:   "create",
		Short: `create action`,
	}
	tmp1 := new(CreateAllergyIntoleranceCommand)
	sub = &cobra.Command{
		Use:   `AllergyIntolerance ["/nosh/patients/PATIENTID/allergy_intolerance"]`,
		Short: ``,
		Long: `

Payload example:

{
   "active": true,
   "address": {
      "city": "Autem molestiae asperiores consequatur animi.",
      "number": "hi",
      "street": "Quod dolor distinctio suscipit ut perspiciatis."
   },
   "birthDate": "Merlot",
   "country": "USA",
   "deceased": false,
   "gender": "male",
   "multiple_birth": false,
   "region": "Napa Valley",
   "review": "Great and inexpensive",
   "sweetness": 4
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp1.Run(c, args) },
	}
	tmp1.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp1.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp2 := new(CreateNutritionRequestCommand)
	sub = &cobra.Command{
		Use:   `NutritionRequest ["/nosh/patients/PATIENTID/nutrition_requests"]`,
		Short: ``,
		Long: `

Payload example:

{
   "address": "Suscipit in in reiciendis officiis saepe voluptate.",
   "address_postal_code": 46032,
   "enteralFormula": "Dolores aperiam et voluptas dolorem deleniti maxime.",
   "foodPreferenceModifier": {
      "coding": "Qui et.",
      "id": "Eaque sequi quaerat et.",
      "text": "Molestiae quae eaque."
   },
   "href": "Veritatis sit nam dolorem.",
   "id": "Tenetur non molestiae.",
   "identifier": "Iure repellat numquam.",
   "links": "Deserunt consequatur aliquam animi.",
   "nutritionRequest": {
      "allergyIntolerance": [
         {
            "display": "Dolores aut reprehenderit adipisci.",
            "reference": "Voluptate et."
         }
      ],
      "dateTime": "2006-11-15T18:31:46-05:00",
      "enteralFormula": [
         {
            "additiveProductName": "Cupiditate necessitatibus libero rem esse eos quibusdam.",
            "additiveType": {
               "coding": [
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  },
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  }
               ],
               "text": "Vitae perferendis odit corporis error consequatur."
            },
            "administrativeInstruction": "Cumque aliquam.",
            "baseFormulaType": {
               "coding": [
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  },
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  }
               ],
               "text": "Vitae perferendis odit corporis error consequatur."
            },
            "baseFormulatProdcutName": "Repudiandae deserunt quia cumque dicta est.",
            "caloricDensity": {
               "code": "Ex voluptatem qui nihil quis.",
               "comparator": "\u003e",
               "system": "http://walter.biz/manuel",
               "unit": "Molestiae eos.",
               "value": 0.3374809955821533
            },
            "maxVolumeToDeliver": {
               "code": "Ex voluptatem qui nihil quis.",
               "comparator": "\u003e",
               "system": "http://walter.biz/manuel",
               "unit": "Molestiae eos.",
               "value": 0.3374809955821533
            },
            "routeofAdministration": {
               "coding": [
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  },
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  }
               ],
               "text": "Vitae perferendis odit corporis error consequatur."
            }
         }
      ],
      "excludeFoodModifier": [
         {
            "coding": [
               {
                  "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                  "display": "Provident quam porro maiores.",
                  "system": "Quisquam voluptas ut aut explicabo et.",
                  "userSelected": true,
                  "version": "Qui eaque et veritatis dolorum voluptatibus."
               },
               {
                  "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                  "display": "Provident quam porro maiores.",
                  "system": "Quisquam voluptas ut aut explicabo et.",
                  "userSelected": true,
                  "version": "Qui eaque et veritatis dolorum voluptatibus."
               }
            ],
            "text": "Vitae perferendis odit corporis error consequatur."
         }
      ],
      "foodPreferenceModifier": [
         {
            "coding": [
               {
                  "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                  "display": "Provident quam porro maiores.",
                  "system": "Quisquam voluptas ut aut explicabo et.",
                  "userSelected": true,
                  "version": "Qui eaque et veritatis dolorum voluptatibus."
               },
               {
                  "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                  "display": "Provident quam porro maiores.",
                  "system": "Quisquam voluptas ut aut explicabo et.",
                  "userSelected": true,
                  "version": "Qui eaque et veritatis dolorum voluptatibus."
               }
            ],
            "text": "Vitae perferendis odit corporis error consequatur."
         },
         {
            "coding": [
               {
                  "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                  "display": "Provident quam porro maiores.",
                  "system": "Quisquam voluptas ut aut explicabo et.",
                  "userSelected": true,
                  "version": "Qui eaque et veritatis dolorum voluptatibus."
               },
               {
                  "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                  "display": "Provident quam porro maiores.",
                  "system": "Quisquam voluptas ut aut explicabo et.",
                  "userSelected": true,
                  "version": "Qui eaque et veritatis dolorum voluptatibus."
               }
            ],
            "text": "Vitae perferendis odit corporis error consequatur."
         },
         {
            "coding": [
               {
                  "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                  "display": "Provident quam porro maiores.",
                  "system": "Quisquam voluptas ut aut explicabo et.",
                  "userSelected": true,
                  "version": "Qui eaque et veritatis dolorum voluptatibus."
               },
               {
                  "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                  "display": "Provident quam porro maiores.",
                  "system": "Quisquam voluptas ut aut explicabo et.",
                  "userSelected": true,
                  "version": "Qui eaque et veritatis dolorum voluptatibus."
               }
            ],
            "text": "Vitae perferendis odit corporis error consequatur."
         }
      ],
      "href": "/patients/1/nutrition_requests/1",
      "id": 1,
      "identifier": [
         {
            "CodeableConcept": {
               "coding": [
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  },
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  }
               ],
               "text": "Vitae perferendis odit corporis error consequatur."
            },
            "assigner": {
               "display": "Dolores aut reprehenderit adipisci.",
               "reference": "Voluptate et."
            },
            "period": {
               "end": "2001-10-30T19:00:22-04:00",
               "start": "1970-01-13T19:32:51-05:00"
            }
         },
         {
            "CodeableConcept": {
               "coding": [
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  },
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  }
               ],
               "text": "Vitae perferendis odit corporis error consequatur."
            },
            "assigner": {
               "display": "Dolores aut reprehenderit adipisci.",
               "reference": "Voluptate et."
            },
            "period": {
               "end": "2001-10-30T19:00:22-04:00",
               "start": "1970-01-13T19:32:51-05:00"
            }
         }
      ],
      "links": {
         "enteralFormula": [
            {
               "additiveProductName": "Cupiditate necessitatibus libero rem esse eos quibusdam.",
               "additiveType": {
                  "coding": [
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     },
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     }
                  ],
                  "text": "Vitae perferendis odit corporis error consequatur."
               },
               "administrativeInstruction": "Cumque aliquam.",
               "baseFormulaType": {
                  "coding": [
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     },
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     }
                  ],
                  "text": "Vitae perferendis odit corporis error consequatur."
               },
               "baseFormulatProdcutName": "Repudiandae deserunt quia cumque dicta est.",
               "caloricDensity": {
                  "code": "Ex voluptatem qui nihil quis.",
                  "comparator": "\u003e",
                  "system": "http://walter.biz/manuel",
                  "unit": "Molestiae eos.",
                  "value": 0.3374809955821533
               },
               "maxVolumeToDeliver": {
                  "code": "Ex voluptatem qui nihil quis.",
                  "comparator": "\u003e",
                  "system": "http://walter.biz/manuel",
                  "unit": "Molestiae eos.",
                  "value": 0.3374809955821533
               },
               "routeofAdministration": {
                  "coding": [
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     },
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     }
                  ],
                  "text": "Vitae perferendis odit corporis error consequatur."
               }
            },
            {
               "additiveProductName": "Cupiditate necessitatibus libero rem esse eos quibusdam.",
               "additiveType": {
                  "coding": [
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     },
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     }
                  ],
                  "text": "Vitae perferendis odit corporis error consequatur."
               },
               "administrativeInstruction": "Cumque aliquam.",
               "baseFormulaType": {
                  "coding": [
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     },
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     }
                  ],
                  "text": "Vitae perferendis odit corporis error consequatur."
               },
               "baseFormulatProdcutName": "Repudiandae deserunt quia cumque dicta est.",
               "caloricDensity": {
                  "code": "Ex voluptatem qui nihil quis.",
                  "comparator": "\u003e",
                  "system": "http://walter.biz/manuel",
                  "unit": "Molestiae eos.",
                  "value": 0.3374809955821533
               },
               "maxVolumeToDeliver": {
                  "code": "Ex voluptatem qui nihil quis.",
                  "comparator": "\u003e",
                  "system": "http://walter.biz/manuel",
                  "unit": "Molestiae eos.",
                  "value": 0.3374809955821533
               },
               "routeofAdministration": {
                  "coding": [
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     },
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     }
                  ],
                  "text": "Vitae perferendis odit corporis error consequatur."
               }
            }
         ],
         "oralDiet": [
            {
               "fluidConsistencyType": {
                  "coding": [
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     },
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     }
                  ],
                  "text": "Vitae perferendis odit corporis error consequatur."
               },
               "instruction": "Qui fugit error.",
               "nutrient": {
                  "amount": {
                     "code": "Ex voluptatem qui nihil quis.",
                     "comparator": "\u003e",
                     "system": "http://walter.biz/manuel",
                     "unit": "Molestiae eos.",
                     "value": 0.3374809955821533
                  },
                  "modifier": {
                     "coding": [
                        {
                           "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                           "display": "Provident quam porro maiores.",
                           "system": "Quisquam voluptas ut aut explicabo et.",
                           "userSelected": true,
                           "version": "Qui eaque et veritatis dolorum voluptatibus."
                        },
                        {
                           "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                           "display": "Provident quam porro maiores.",
                           "system": "Quisquam voluptas ut aut explicabo et.",
                           "userSelected": true,
                           "version": "Qui eaque et veritatis dolorum voluptatibus."
                        }
                     ],
                     "text": "Vitae perferendis odit corporis error consequatur."
                  }
               },
               "schedule": {
                  "code": "AM",
                  "event": "1978-12-17T03:44:52-05:00",
                  "repeat": {
                     "extension": {
                        "ValueAddress": {
                           "city": "Consequatur omnis omnis tempore et.",
                           "country": "Suscipit ut quia sequi atque numquam.",
                           "distinct": "Quia fuga.",
                           "line": "Vero dicta voluptatibus consectetur ut minus.",
                           "postalCode": 8606407676718086019,
                           "state": "Neque quod similique est quia.",
                           "type": "postal",
                           "use": "work"
                        },
                        "ValueAnnotation": {
                           "authorReference": {
                              "display": "Dolores aut reprehenderit adipisci.",
                              "reference": "Voluptate et."
                           },
                           "authorString": "Quasi cupiditate quia debitis vel.",
                           "time": "1992-01-31T12:26:43-05:00"
                        },
                        "ValueAttachment": {
                           "contentType": "In et sunt iure quo tempora consequuntur.",
                           "creation": "2010-04-13T22:20:06-04:00",
                           "data": "Iusto aperiam.",
                           "hash": "Unde consequuntur esse qui.",
                           "size": 7170819194107951440,
                           "title": "Voluptatem id eum cupiditate."
                        },
                        "ValueBase64Binary": "Est distinctio voluptate tenetur.",
                        "ValueBoolean": true,
                        "ValueCode": "Placeat vero temporibus et aliquid ut incidunt.",
                        "ValueCodeableConcept": {
                           "coding": [
                              {
                                 "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                                 "display": "Provident quam porro maiores.",
                                 "system": "Quisquam voluptas ut aut explicabo et.",
                                 "userSelected": true,
                                 "version": "Qui eaque et veritatis dolorum voluptatibus."
                              },
                              {
                                 "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                                 "display": "Provident quam porro maiores.",
                                 "system": "Quisquam voluptas ut aut explicabo et.",
                                 "userSelected": true,
                                 "version": "Qui eaque et veritatis dolorum voluptatibus."
                              }
                           ],
                           "text": "Vitae perferendis odit corporis error consequatur."
                        },
                        "ValueCoding": {
                           "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                           "display": "Provident quam porro maiores.",
                           "system": "Quisquam voluptas ut aut explicabo et.",
                           "userSelected": true,
                           "version": "Qui eaque et veritatis dolorum voluptatibus."
                        },
                        "ValueContactPoint": "Dolor quaerat delectus et.",
                        "ValueDate": "1980-05-15T19:46:08-04:00",
                        "ValueDateTime": "1981-04-19T10:31:33-04:00",
                        "ValueDecimal": 0.7932081018594244,
                        "ValueHumanName": {
                           "family": [
                              "Ea culpa qui a enim.",
                              "Ea culpa qui a enim.",
                              "Ea culpa qui a enim."
                           ],
                           "given": "Reiciendis odit.",
                           "period": {
                              "end": "2001-10-30T19:00:22-04:00",
                              "start": "1970-01-13T19:32:51-05:00"
                           },
                           "prefix": "Natus inventore optio eaque.",
                           "suffix": "Id qui delectus quia qui.",
                           "use": "temp"
                        },
                        "ValueId": "Molestiae praesentium quidem corrupti nihil eum.",
                        "ValueIdentifier": {
                           "CodeableConcept": {
                              "coding": [
                                 {
                                    "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                                    "display": "Provident quam porro maiores.",
                                    "system": "Quisquam voluptas ut aut explicabo et.",
                                    "userSelected": true,
                                    "version": "Qui eaque et veritatis dolorum voluptatibus."
                                 },
                                 {
                                    "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                                    "display": "Provident quam porro maiores.",
                                    "system": "Quisquam voluptas ut aut explicabo et.",
                                    "userSelected": true,
                                    "version": "Qui eaque et veritatis dolorum voluptatibus."
                                 }
                              ],
                              "text": "Vitae perferendis odit corporis error consequatur."
                           },
                           "assigner": {
                              "display": "Dolores aut reprehenderit adipisci.",
                              "reference": "Voluptate et."
                           },
                           "period": {
                              "end": "2001-10-30T19:00:22-04:00",
                              "start": "1970-01-13T19:32:51-05:00"
                           }
                        },
                        "ValueInstant": "1977-11-08T04:06:45-05:00",
                        "ValueInteger": 8012361700862165525,
                        "ValueMarkdown": "Sed neque reiciendis.",
                        "ValueMeta": {
                           "lastUpdated": "1975-12-19T03:26:58-05:00",
                           "profile": "http://green.info/marilyne",
                           "security": {
                              "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                              "display": "Provident quam porro maiores.",
                              "system": "Quisquam voluptas ut aut explicabo et.",
                              "userSelected": true,
                              "version": "Qui eaque et veritatis dolorum voluptatibus."
                           },
                           "tag": {
                              "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                              "display": "Provident quam porro maiores.",
                              "system": "Quisquam voluptas ut aut explicabo et.",
                              "userSelected": true,
                              "version": "Qui eaque et veritatis dolorum voluptatibus."
                           },
                           "versionId": "Corrupti nemo dolor et non vel ea."
                        },
                        "ValueOid": "Laudantium porro.",
                        "ValuePeriod": {
                           "end": "2001-10-30T19:00:22-04:00",
                           "start": "1970-01-13T19:32:51-05:00"
                        },
                        "ValuePositiveInt": 0.6018552052933437,
                        "ValueQuantity": {
                           "code": "Ex voluptatem qui nihil quis.",
                           "comparator": "\u003e",
                           "system": "http://walter.biz/manuel",
                           "unit": "Molestiae eos.",
                           "value": 0.3374809955821533
                        },
                        "ValueRange": "Similique quia quis.",
                        "ValueRatio": 7314611204127242704,
                        "ValueReference": {
                           "display": "Dolores aut reprehenderit adipisci.",
                           "reference": "Voluptate et."
                        },
                        "ValueSampledData": "Unde odio aut suscipit atque.",
                        "ValueSignature": "Libero sequi ut.",
                        "ValueString": "Molestiae in.",
                        "ValueTime": "2013-02-01T03:56:35-05:00",
                        "ValueTiming": "1993-07-31T06:07:28-04:00",
                        "ValueUnsignedInt": 0.8739176042943996,
                        "ValueUri": "A expedita vel aut voluptatum veritatis.",
                        "url": "Nam earum repellendus officia in rerum aut."
                     },
                     "id": "Dolores labore laudantium."
                  }
               },
               "texture": {
                  "foodType": {
                     "coding": [
                        {
                           "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                           "display": "Provident quam porro maiores.",
                           "system": "Quisquam voluptas ut aut explicabo et.",
                           "userSelected": true,
                           "version": "Qui eaque et veritatis dolorum voluptatibus."
                        },
                        {
                           "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                           "display": "Provident quam porro maiores.",
                           "system": "Quisquam voluptas ut aut explicabo et.",
                           "userSelected": true,
                           "version": "Qui eaque et veritatis dolorum voluptatibus."
                        }
                     ],
                     "text": "Vitae perferendis odit corporis error consequatur."
                  },
                  "modifier": {
                     "coding": [
                        {
                           "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                           "display": "Provident quam porro maiores.",
                           "system": "Quisquam voluptas ut aut explicabo et.",
                           "userSelected": true,
                           "version": "Qui eaque et veritatis dolorum voluptatibus."
                        },
                        {
                           "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                           "display": "Provident quam porro maiores.",
                           "system": "Quisquam voluptas ut aut explicabo et.",
                           "userSelected": true,
                           "version": "Qui eaque et veritatis dolorum voluptatibus."
                        }
                     ],
                     "text": "Vitae perferendis odit corporis error consequatur."
                  }
               },
               "type": {
                  "coding": [
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     },
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     }
                  ],
                  "text": "Vitae perferendis odit corporis error consequatur."
               }
            }
         ]
      },
      "oralDiet": [
         {
            "fluidConsistencyType": {
               "coding": [
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  },
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  }
               ],
               "text": "Vitae perferendis odit corporis error consequatur."
            },
            "instruction": "Qui fugit error.",
            "nutrient": {
               "amount": {
                  "code": "Ex voluptatem qui nihil quis.",
                  "comparator": "\u003e",
                  "system": "http://walter.biz/manuel",
                  "unit": "Molestiae eos.",
                  "value": 0.3374809955821533
               },
               "modifier": {
                  "coding": [
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     },
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     }
                  ],
                  "text": "Vitae perferendis odit corporis error consequatur."
               }
            },
            "schedule": {
               "code": "AM",
               "event": "1978-12-17T03:44:52-05:00",
               "repeat": {
                  "extension": {
                     "ValueAddress": {
                        "city": "Consequatur omnis omnis tempore et.",
                        "country": "Suscipit ut quia sequi atque numquam.",
                        "distinct": "Quia fuga.",
                        "line": "Vero dicta voluptatibus consectetur ut minus.",
                        "postalCode": 8606407676718086019,
                        "state": "Neque quod similique est quia.",
                        "type": "postal",
                        "use": "work"
                     },
                     "ValueAnnotation": {
                        "authorReference": {
                           "display": "Dolores aut reprehenderit adipisci.",
                           "reference": "Voluptate et."
                        },
                        "authorString": "Quasi cupiditate quia debitis vel.",
                        "time": "1992-01-31T12:26:43-05:00"
                     },
                     "ValueAttachment": {
                        "contentType": "In et sunt iure quo tempora consequuntur.",
                        "creation": "2010-04-13T22:20:06-04:00",
                        "data": "Iusto aperiam.",
                        "hash": "Unde consequuntur esse qui.",
                        "size": 7170819194107951440,
                        "title": "Voluptatem id eum cupiditate."
                     },
                     "ValueBase64Binary": "Est distinctio voluptate tenetur.",
                     "ValueBoolean": true,
                     "ValueCode": "Placeat vero temporibus et aliquid ut incidunt.",
                     "ValueCodeableConcept": {
                        "coding": [
                           {
                              "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                              "display": "Provident quam porro maiores.",
                              "system": "Quisquam voluptas ut aut explicabo et.",
                              "userSelected": true,
                              "version": "Qui eaque et veritatis dolorum voluptatibus."
                           },
                           {
                              "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                              "display": "Provident quam porro maiores.",
                              "system": "Quisquam voluptas ut aut explicabo et.",
                              "userSelected": true,
                              "version": "Qui eaque et veritatis dolorum voluptatibus."
                           }
                        ],
                        "text": "Vitae perferendis odit corporis error consequatur."
                     },
                     "ValueCoding": {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     },
                     "ValueContactPoint": "Dolor quaerat delectus et.",
                     "ValueDate": "1980-05-15T19:46:08-04:00",
                     "ValueDateTime": "1981-04-19T10:31:33-04:00",
                     "ValueDecimal": 0.7932081018594244,
                     "ValueHumanName": {
                        "family": [
                           "Ea culpa qui a enim.",
                           "Ea culpa qui a enim.",
                           "Ea culpa qui a enim."
                        ],
                        "given": "Reiciendis odit.",
                        "period": {
                           "end": "2001-10-30T19:00:22-04:00",
                           "start": "1970-01-13T19:32:51-05:00"
                        },
                        "prefix": "Natus inventore optio eaque.",
                        "suffix": "Id qui delectus quia qui.",
                        "use": "temp"
                     },
                     "ValueId": "Molestiae praesentium quidem corrupti nihil eum.",
                     "ValueIdentifier": {
                        "CodeableConcept": {
                           "coding": [
                              {
                                 "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                                 "display": "Provident quam porro maiores.",
                                 "system": "Quisquam voluptas ut aut explicabo et.",
                                 "userSelected": true,
                                 "version": "Qui eaque et veritatis dolorum voluptatibus."
                              },
                              {
                                 "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                                 "display": "Provident quam porro maiores.",
                                 "system": "Quisquam voluptas ut aut explicabo et.",
                                 "userSelected": true,
                                 "version": "Qui eaque et veritatis dolorum voluptatibus."
                              }
                           ],
                           "text": "Vitae perferendis odit corporis error consequatur."
                        },
                        "assigner": {
                           "display": "Dolores aut reprehenderit adipisci.",
                           "reference": "Voluptate et."
                        },
                        "period": {
                           "end": "2001-10-30T19:00:22-04:00",
                           "start": "1970-01-13T19:32:51-05:00"
                        }
                     },
                     "ValueInstant": "1977-11-08T04:06:45-05:00",
                     "ValueInteger": 8012361700862165525,
                     "ValueMarkdown": "Sed neque reiciendis.",
                     "ValueMeta": {
                        "lastUpdated": "1975-12-19T03:26:58-05:00",
                        "profile": "http://green.info/marilyne",
                        "security": {
                           "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                           "display": "Provident quam porro maiores.",
                           "system": "Quisquam voluptas ut aut explicabo et.",
                           "userSelected": true,
                           "version": "Qui eaque et veritatis dolorum voluptatibus."
                        },
                        "tag": {
                           "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                           "display": "Provident quam porro maiores.",
                           "system": "Quisquam voluptas ut aut explicabo et.",
                           "userSelected": true,
                           "version": "Qui eaque et veritatis dolorum voluptatibus."
                        },
                        "versionId": "Corrupti nemo dolor et non vel ea."
                     },
                     "ValueOid": "Laudantium porro.",
                     "ValuePeriod": {
                        "end": "2001-10-30T19:00:22-04:00",
                        "start": "1970-01-13T19:32:51-05:00"
                     },
                     "ValuePositiveInt": 0.6018552052933437,
                     "ValueQuantity": {
                        "code": "Ex voluptatem qui nihil quis.",
                        "comparator": "\u003e",
                        "system": "http://walter.biz/manuel",
                        "unit": "Molestiae eos.",
                        "value": 0.3374809955821533
                     },
                     "ValueRange": "Similique quia quis.",
                     "ValueRatio": 7314611204127242704,
                     "ValueReference": {
                        "display": "Dolores aut reprehenderit adipisci.",
                        "reference": "Voluptate et."
                     },
                     "ValueSampledData": "Unde odio aut suscipit atque.",
                     "ValueSignature": "Libero sequi ut.",
                     "ValueString": "Molestiae in.",
                     "ValueTime": "2013-02-01T03:56:35-05:00",
                     "ValueTiming": "1993-07-31T06:07:28-04:00",
                     "ValueUnsignedInt": 0.8739176042943996,
                     "ValueUri": "A expedita vel aut voluptatum veritatis.",
                     "url": "Nam earum repellendus officia in rerum aut."
                  },
                  "id": "Dolores labore laudantium."
               }
            },
            "texture": {
               "foodType": {
                  "coding": [
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     },
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     }
                  ],
                  "text": "Vitae perferendis odit corporis error consequatur."
               },
               "modifier": {
                  "coding": [
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     },
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     }
                  ],
                  "text": "Vitae perferendis odit corporis error consequatur."
               }
            },
            "type": {
               "coding": [
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  },
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  }
               ],
               "text": "Vitae perferendis odit corporis error consequatur."
            }
         },
         {
            "fluidConsistencyType": {
               "coding": [
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  },
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  }
               ],
               "text": "Vitae perferendis odit corporis error consequatur."
            },
            "instruction": "Qui fugit error.",
            "nutrient": {
               "amount": {
                  "code": "Ex voluptatem qui nihil quis.",
                  "comparator": "\u003e",
                  "system": "http://walter.biz/manuel",
                  "unit": "Molestiae eos.",
                  "value": 0.3374809955821533
               },
               "modifier": {
                  "coding": [
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     },
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     }
                  ],
                  "text": "Vitae perferendis odit corporis error consequatur."
               }
            },
            "schedule": {
               "code": "AM",
               "event": "1978-12-17T03:44:52-05:00",
               "repeat": {
                  "extension": {
                     "ValueAddress": {
                        "city": "Consequatur omnis omnis tempore et.",
                        "country": "Suscipit ut quia sequi atque numquam.",
                        "distinct": "Quia fuga.",
                        "line": "Vero dicta voluptatibus consectetur ut minus.",
                        "postalCode": 8606407676718086019,
                        "state": "Neque quod similique est quia.",
                        "type": "postal",
                        "use": "work"
                     },
                     "ValueAnnotation": {
                        "authorReference": {
                           "display": "Dolores aut reprehenderit adipisci.",
                           "reference": "Voluptate et."
                        },
                        "authorString": "Quasi cupiditate quia debitis vel.",
                        "time": "1992-01-31T12:26:43-05:00"
                     },
                     "ValueAttachment": {
                        "contentType": "In et sunt iure quo tempora consequuntur.",
                        "creation": "2010-04-13T22:20:06-04:00",
                        "data": "Iusto aperiam.",
                        "hash": "Unde consequuntur esse qui.",
                        "size": 7170819194107951440,
                        "title": "Voluptatem id eum cupiditate."
                     },
                     "ValueBase64Binary": "Est distinctio voluptate tenetur.",
                     "ValueBoolean": true,
                     "ValueCode": "Placeat vero temporibus et aliquid ut incidunt.",
                     "ValueCodeableConcept": {
                        "coding": [
                           {
                              "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                              "display": "Provident quam porro maiores.",
                              "system": "Quisquam voluptas ut aut explicabo et.",
                              "userSelected": true,
                              "version": "Qui eaque et veritatis dolorum voluptatibus."
                           },
                           {
                              "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                              "display": "Provident quam porro maiores.",
                              "system": "Quisquam voluptas ut aut explicabo et.",
                              "userSelected": true,
                              "version": "Qui eaque et veritatis dolorum voluptatibus."
                           }
                        ],
                        "text": "Vitae perferendis odit corporis error consequatur."
                     },
                     "ValueCoding": {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     },
                     "ValueContactPoint": "Dolor quaerat delectus et.",
                     "ValueDate": "1980-05-15T19:46:08-04:00",
                     "ValueDateTime": "1981-04-19T10:31:33-04:00",
                     "ValueDecimal": 0.7932081018594244,
                     "ValueHumanName": {
                        "family": [
                           "Ea culpa qui a enim.",
                           "Ea culpa qui a enim.",
                           "Ea culpa qui a enim."
                        ],
                        "given": "Reiciendis odit.",
                        "period": {
                           "end": "2001-10-30T19:00:22-04:00",
                           "start": "1970-01-13T19:32:51-05:00"
                        },
                        "prefix": "Natus inventore optio eaque.",
                        "suffix": "Id qui delectus quia qui.",
                        "use": "temp"
                     },
                     "ValueId": "Molestiae praesentium quidem corrupti nihil eum.",
                     "ValueIdentifier": {
                        "CodeableConcept": {
                           "coding": [
                              {
                                 "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                                 "display": "Provident quam porro maiores.",
                                 "system": "Quisquam voluptas ut aut explicabo et.",
                                 "userSelected": true,
                                 "version": "Qui eaque et veritatis dolorum voluptatibus."
                              },
                              {
                                 "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                                 "display": "Provident quam porro maiores.",
                                 "system": "Quisquam voluptas ut aut explicabo et.",
                                 "userSelected": true,
                                 "version": "Qui eaque et veritatis dolorum voluptatibus."
                              }
                           ],
                           "text": "Vitae perferendis odit corporis error consequatur."
                        },
                        "assigner": {
                           "display": "Dolores aut reprehenderit adipisci.",
                           "reference": "Voluptate et."
                        },
                        "period": {
                           "end": "2001-10-30T19:00:22-04:00",
                           "start": "1970-01-13T19:32:51-05:00"
                        }
                     },
                     "ValueInstant": "1977-11-08T04:06:45-05:00",
                     "ValueInteger": 8012361700862165525,
                     "ValueMarkdown": "Sed neque reiciendis.",
                     "ValueMeta": {
                        "lastUpdated": "1975-12-19T03:26:58-05:00",
                        "profile": "http://green.info/marilyne",
                        "security": {
                           "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                           "display": "Provident quam porro maiores.",
                           "system": "Quisquam voluptas ut aut explicabo et.",
                           "userSelected": true,
                           "version": "Qui eaque et veritatis dolorum voluptatibus."
                        },
                        "tag": {
                           "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                           "display": "Provident quam porro maiores.",
                           "system": "Quisquam voluptas ut aut explicabo et.",
                           "userSelected": true,
                           "version": "Qui eaque et veritatis dolorum voluptatibus."
                        },
                        "versionId": "Corrupti nemo dolor et non vel ea."
                     },
                     "ValueOid": "Laudantium porro.",
                     "ValuePeriod": {
                        "end": "2001-10-30T19:00:22-04:00",
                        "start": "1970-01-13T19:32:51-05:00"
                     },
                     "ValuePositiveInt": 0.6018552052933437,
                     "ValueQuantity": {
                        "code": "Ex voluptatem qui nihil quis.",
                        "comparator": "\u003e",
                        "system": "http://walter.biz/manuel",
                        "unit": "Molestiae eos.",
                        "value": 0.3374809955821533
                     },
                     "ValueRange": "Similique quia quis.",
                     "ValueRatio": 7314611204127242704,
                     "ValueReference": {
                        "display": "Dolores aut reprehenderit adipisci.",
                        "reference": "Voluptate et."
                     },
                     "ValueSampledData": "Unde odio aut suscipit atque.",
                     "ValueSignature": "Libero sequi ut.",
                     "ValueString": "Molestiae in.",
                     "ValueTime": "2013-02-01T03:56:35-05:00",
                     "ValueTiming": "1993-07-31T06:07:28-04:00",
                     "ValueUnsignedInt": 0.8739176042943996,
                     "ValueUri": "A expedita vel aut voluptatum veritatis.",
                     "url": "Nam earum repellendus officia in rerum aut."
                  },
                  "id": "Dolores labore laudantium."
               }
            },
            "texture": {
               "foodType": {
                  "coding": [
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     },
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     }
                  ],
                  "text": "Vitae perferendis odit corporis error consequatur."
               },
               "modifier": {
                  "coding": [
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     },
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     }
                  ],
                  "text": "Vitae perferendis odit corporis error consequatur."
               }
            },
            "type": {
               "coding": [
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  },
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  }
               ],
               "text": "Vitae perferendis odit corporis error consequatur."
            }
         }
      ],
      "orderer": {
         "display": "Dolores aut reprehenderit adipisci.",
         "reference": "Voluptate et."
      },
      "patient": {
         "display": "Dolores aut reprehenderit adipisci.",
         "reference": "Voluptate et."
      },
      "status": "completed",
      "supplement": [
         {
            "instruction": "Labore nihil sint quibusdam et.",
            "productName": "Ut impedit.",
            "quantity": {
               "code": "Ex voluptatem qui nihil quis.",
               "comparator": "\u003e",
               "system": "http://walter.biz/manuel",
               "unit": "Molestiae eos.",
               "value": 0.3374809955821533
            },
            "type": {
               "coding": [
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  },
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  }
               ],
               "text": "Vitae perferendis odit corporis error consequatur."
            }
         },
         {
            "instruction": "Labore nihil sint quibusdam et.",
            "productName": "Ut impedit.",
            "quantity": {
               "code": "Ex voluptatem qui nihil quis.",
               "comparator": "\u003e",
               "system": "http://walter.biz/manuel",
               "unit": "Molestiae eos.",
               "value": 0.3374809955821533
            },
            "type": {
               "coding": [
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  },
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  }
               ],
               "text": "Vitae perferendis odit corporis error consequatur."
            }
         }
      ]
   },
   "oralDiet": "At ex.",
   "patient": {
      "id": "Rerum recusandae voluptatum fugiat.",
      "reference": "Enim reiciendis et sed fuga."
   },
   "supplement": "Est autem."
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp2.Run(c, args) },
	}
	tmp2.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp2.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp3 := new(CreateObservationCommand)
	sub = &cobra.Command{
		Use:   `Observation ["/nosh/patients/PATIENTID/observation"]`,
		Short: ``,
		Long: `

Payload example:

{
   "observation": {
      "bodySite": {
         "coding": [
            {
               "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
               "display": "Provident quam porro maiores.",
               "system": "Quisquam voluptas ut aut explicabo et.",
               "userSelected": true,
               "version": "Qui eaque et veritatis dolorum voluptatibus."
            },
            {
               "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
               "display": "Provident quam porro maiores.",
               "system": "Quisquam voluptas ut aut explicabo et.",
               "userSelected": true,
               "version": "Qui eaque et veritatis dolorum voluptatibus."
            }
         ],
         "text": "Vitae perferendis odit corporis error consequatur."
      },
      "category": {
         "coding": [
            {
               "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
               "display": "Provident quam porro maiores.",
               "system": "Quisquam voluptas ut aut explicabo et.",
               "userSelected": true,
               "version": "Qui eaque et veritatis dolorum voluptatibus."
            },
            {
               "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
               "display": "Provident quam porro maiores.",
               "system": "Quisquam voluptas ut aut explicabo et.",
               "userSelected": true,
               "version": "Qui eaque et veritatis dolorum voluptatibus."
            }
         ],
         "text": "Vitae perferendis odit corporis error consequatur."
      },
      "code": {
         "coding": [
            {
               "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
               "display": "Provident quam porro maiores.",
               "system": "Quisquam voluptas ut aut explicabo et.",
               "userSelected": true,
               "version": "Qui eaque et veritatis dolorum voluptatibus."
            },
            {
               "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
               "display": "Provident quam porro maiores.",
               "system": "Quisquam voluptas ut aut explicabo et.",
               "userSelected": true,
               "version": "Qui eaque et veritatis dolorum voluptatibus."
            }
         ],
         "text": "Vitae perferendis odit corporis error consequatur."
      },
      "comments": "Voluptate sint officiis qui nulla sapiente minus.",
      "component": [
         {
            "code": {
               "coding": [
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  },
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  }
               ],
               "text": "Vitae perferendis odit corporis error consequatur."
            },
            "dataAbsentReason": {
               "coding": [
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  },
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  }
               ],
               "text": "Vitae perferendis odit corporis error consequatur."
            },
            "referenceRange": {
               "age": {
                  "high": {
                     "code": "Ex voluptatem qui nihil quis.",
                     "comparator": "\u003e",
                     "system": "http://walter.biz/manuel",
                     "unit": "Molestiae eos.",
                     "value": 0.3374809955821533
                  },
                  "low": {
                     "code": "Ex voluptatem qui nihil quis.",
                     "comparator": "\u003e",
                     "system": "http://walter.biz/manuel",
                     "unit": "Molestiae eos.",
                     "value": 0.3374809955821533
                  }
               },
               "high": {
                  "code": "Ex voluptatem qui nihil quis.",
                  "comparator": "\u003e",
                  "system": "http://walter.biz/manuel",
                  "unit": "Molestiae eos.",
                  "value": 0.3374809955821533
               },
               "low": {
                  "code": "Ex voluptatem qui nihil quis.",
                  "comparator": "\u003e",
                  "system": "http://walter.biz/manuel",
                  "unit": "Molestiae eos.",
                  "value": 0.3374809955821533
               },
               "meaning": {
                  "coding": [
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     },
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     }
                  ],
                  "text": "Vitae perferendis odit corporis error consequatur."
               },
               "text": "Est suscipit modi dignissimos."
            },
            "valueAttachment": {
               "contentType": "In et sunt iure quo tempora consequuntur.",
               "creation": "2010-04-13T22:20:06-04:00",
               "data": "Iusto aperiam.",
               "hash": "Unde consequuntur esse qui.",
               "size": 7170819194107951440,
               "title": "Voluptatem id eum cupiditate."
            },
            "valueCodeableConcept": {
               "coding": [
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  },
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  }
               ],
               "text": "Vitae perferendis odit corporis error consequatur."
            },
            "valueDatTime": "2006-06-03T15:46:06-04:00",
            "valuePeriod": {
               "end": "2001-10-30T19:00:22-04:00",
               "start": "1970-01-13T19:32:51-05:00"
            },
            "valueQuantity": {
               "code": "Ex voluptatem qui nihil quis.",
               "comparator": "\u003e",
               "system": "http://walter.biz/manuel",
               "unit": "Molestiae eos.",
               "value": 0.3374809955821533
            },
            "valueRange": {
               "high": {
                  "code": "Ex voluptatem qui nihil quis.",
                  "comparator": "\u003e",
                  "system": "http://walter.biz/manuel",
                  "unit": "Molestiae eos.",
                  "value": 0.3374809955821533
               },
               "low": {
                  "code": "Ex voluptatem qui nihil quis.",
                  "comparator": "\u003e",
                  "system": "http://walter.biz/manuel",
                  "unit": "Molestiae eos.",
                  "value": 0.3374809955821533
               }
            },
            "valueSampledData": {
               "data": "Quis eos earum.",
               "dimensions": 9133536065128410703,
               "factor": 0.2928210072171706,
               "lowerLimit": 0.5963619438089043,
               "origin": {
                  "code": "Ex voluptatem qui nihil quis.",
                  "comparator": "\u003e",
                  "system": "http://walter.biz/manuel",
                  "unit": "Molestiae eos.",
                  "value": 0.3374809955821533
               },
               "period": 0.47228831209882643,
               "upperLimit": 0.6070474669537046
            },
            "valueTime": "1980-07-14T00:10:15-04:00"
         },
         {
            "code": {
               "coding": [
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  },
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  }
               ],
               "text": "Vitae perferendis odit corporis error consequatur."
            },
            "dataAbsentReason": {
               "coding": [
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  },
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  }
               ],
               "text": "Vitae perferendis odit corporis error consequatur."
            },
            "referenceRange": {
               "age": {
                  "high": {
                     "code": "Ex voluptatem qui nihil quis.",
                     "comparator": "\u003e",
                     "system": "http://walter.biz/manuel",
                     "unit": "Molestiae eos.",
                     "value": 0.3374809955821533
                  },
                  "low": {
                     "code": "Ex voluptatem qui nihil quis.",
                     "comparator": "\u003e",
                     "system": "http://walter.biz/manuel",
                     "unit": "Molestiae eos.",
                     "value": 0.3374809955821533
                  }
               },
               "high": {
                  "code": "Ex voluptatem qui nihil quis.",
                  "comparator": "\u003e",
                  "system": "http://walter.biz/manuel",
                  "unit": "Molestiae eos.",
                  "value": 0.3374809955821533
               },
               "low": {
                  "code": "Ex voluptatem qui nihil quis.",
                  "comparator": "\u003e",
                  "system": "http://walter.biz/manuel",
                  "unit": "Molestiae eos.",
                  "value": 0.3374809955821533
               },
               "meaning": {
                  "coding": [
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     },
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     }
                  ],
                  "text": "Vitae perferendis odit corporis error consequatur."
               },
               "text": "Est suscipit modi dignissimos."
            },
            "valueAttachment": {
               "contentType": "In et sunt iure quo tempora consequuntur.",
               "creation": "2010-04-13T22:20:06-04:00",
               "data": "Iusto aperiam.",
               "hash": "Unde consequuntur esse qui.",
               "size": 7170819194107951440,
               "title": "Voluptatem id eum cupiditate."
            },
            "valueCodeableConcept": {
               "coding": [
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  },
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  }
               ],
               "text": "Vitae perferendis odit corporis error consequatur."
            },
            "valueDatTime": "2006-06-03T15:46:06-04:00",
            "valuePeriod": {
               "end": "2001-10-30T19:00:22-04:00",
               "start": "1970-01-13T19:32:51-05:00"
            },
            "valueQuantity": {
               "code": "Ex voluptatem qui nihil quis.",
               "comparator": "\u003e",
               "system": "http://walter.biz/manuel",
               "unit": "Molestiae eos.",
               "value": 0.3374809955821533
            },
            "valueRange": {
               "high": {
                  "code": "Ex voluptatem qui nihil quis.",
                  "comparator": "\u003e",
                  "system": "http://walter.biz/manuel",
                  "unit": "Molestiae eos.",
                  "value": 0.3374809955821533
               },
               "low": {
                  "code": "Ex voluptatem qui nihil quis.",
                  "comparator": "\u003e",
                  "system": "http://walter.biz/manuel",
                  "unit": "Molestiae eos.",
                  "value": 0.3374809955821533
               }
            },
            "valueSampledData": {
               "data": "Quis eos earum.",
               "dimensions": 9133536065128410703,
               "factor": 0.2928210072171706,
               "lowerLimit": 0.5963619438089043,
               "origin": {
                  "code": "Ex voluptatem qui nihil quis.",
                  "comparator": "\u003e",
                  "system": "http://walter.biz/manuel",
                  "unit": "Molestiae eos.",
                  "value": 0.3374809955821533
               },
               "period": 0.47228831209882643,
               "upperLimit": 0.6070474669537046
            },
            "valueTime": "1980-07-14T00:10:15-04:00"
         },
         {
            "code": {
               "coding": [
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  },
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  }
               ],
               "text": "Vitae perferendis odit corporis error consequatur."
            },
            "dataAbsentReason": {
               "coding": [
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  },
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  }
               ],
               "text": "Vitae perferendis odit corporis error consequatur."
            },
            "referenceRange": {
               "age": {
                  "high": {
                     "code": "Ex voluptatem qui nihil quis.",
                     "comparator": "\u003e",
                     "system": "http://walter.biz/manuel",
                     "unit": "Molestiae eos.",
                     "value": 0.3374809955821533
                  },
                  "low": {
                     "code": "Ex voluptatem qui nihil quis.",
                     "comparator": "\u003e",
                     "system": "http://walter.biz/manuel",
                     "unit": "Molestiae eos.",
                     "value": 0.3374809955821533
                  }
               },
               "high": {
                  "code": "Ex voluptatem qui nihil quis.",
                  "comparator": "\u003e",
                  "system": "http://walter.biz/manuel",
                  "unit": "Molestiae eos.",
                  "value": 0.3374809955821533
               },
               "low": {
                  "code": "Ex voluptatem qui nihil quis.",
                  "comparator": "\u003e",
                  "system": "http://walter.biz/manuel",
                  "unit": "Molestiae eos.",
                  "value": 0.3374809955821533
               },
               "meaning": {
                  "coding": [
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     },
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     }
                  ],
                  "text": "Vitae perferendis odit corporis error consequatur."
               },
               "text": "Est suscipit modi dignissimos."
            },
            "valueAttachment": {
               "contentType": "In et sunt iure quo tempora consequuntur.",
               "creation": "2010-04-13T22:20:06-04:00",
               "data": "Iusto aperiam.",
               "hash": "Unde consequuntur esse qui.",
               "size": 7170819194107951440,
               "title": "Voluptatem id eum cupiditate."
            },
            "valueCodeableConcept": {
               "coding": [
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  },
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  }
               ],
               "text": "Vitae perferendis odit corporis error consequatur."
            },
            "valueDatTime": "2006-06-03T15:46:06-04:00",
            "valuePeriod": {
               "end": "2001-10-30T19:00:22-04:00",
               "start": "1970-01-13T19:32:51-05:00"
            },
            "valueQuantity": {
               "code": "Ex voluptatem qui nihil quis.",
               "comparator": "\u003e",
               "system": "http://walter.biz/manuel",
               "unit": "Molestiae eos.",
               "value": 0.3374809955821533
            },
            "valueRange": {
               "high": {
                  "code": "Ex voluptatem qui nihil quis.",
                  "comparator": "\u003e",
                  "system": "http://walter.biz/manuel",
                  "unit": "Molestiae eos.",
                  "value": 0.3374809955821533
               },
               "low": {
                  "code": "Ex voluptatem qui nihil quis.",
                  "comparator": "\u003e",
                  "system": "http://walter.biz/manuel",
                  "unit": "Molestiae eos.",
                  "value": 0.3374809955821533
               }
            },
            "valueSampledData": {
               "data": "Quis eos earum.",
               "dimensions": 9133536065128410703,
               "factor": 0.2928210072171706,
               "lowerLimit": 0.5963619438089043,
               "origin": {
                  "code": "Ex voluptatem qui nihil quis.",
                  "comparator": "\u003e",
                  "system": "http://walter.biz/manuel",
                  "unit": "Molestiae eos.",
                  "value": 0.3374809955821533
               },
               "period": 0.47228831209882643,
               "upperLimit": 0.6070474669537046
            },
            "valueTime": "1980-07-14T00:10:15-04:00"
         }
      ],
      "dateAbsentReason": {
         "coding": [
            {
               "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
               "display": "Provident quam porro maiores.",
               "system": "Quisquam voluptas ut aut explicabo et.",
               "userSelected": true,
               "version": "Qui eaque et veritatis dolorum voluptatibus."
            },
            {
               "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
               "display": "Provident quam porro maiores.",
               "system": "Quisquam voluptas ut aut explicabo et.",
               "userSelected": true,
               "version": "Qui eaque et veritatis dolorum voluptatibus."
            }
         ],
         "text": "Vitae perferendis odit corporis error consequatur."
      },
      "device": {
         "display": "Dolores aut reprehenderit adipisci.",
         "reference": "Voluptate et."
      },
      "effectiveDateTime": "1989-07-31T17:52:24-04:00",
      "effectivePeriod": {
         "end": "2001-10-30T19:00:22-04:00",
         "start": "1970-01-13T19:32:51-05:00"
      },
      "encounter": {
         "display": "Dolores aut reprehenderit adipisci.",
         "reference": "Voluptate et."
      },
      "identifier": [
         {
            "CodeableConcept": {
               "coding": [
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  },
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  }
               ],
               "text": "Vitae perferendis odit corporis error consequatur."
            },
            "assigner": {
               "display": "Dolores aut reprehenderit adipisci.",
               "reference": "Voluptate et."
            },
            "period": {
               "end": "2001-10-30T19:00:22-04:00",
               "start": "1970-01-13T19:32:51-05:00"
            }
         },
         {
            "CodeableConcept": {
               "coding": [
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  },
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  }
               ],
               "text": "Vitae perferendis odit corporis error consequatur."
            },
            "assigner": {
               "display": "Dolores aut reprehenderit adipisci.",
               "reference": "Voluptate et."
            },
            "period": {
               "end": "2001-10-30T19:00:22-04:00",
               "start": "1970-01-13T19:32:51-05:00"
            }
         }
      ],
      "interpretation": {
         "coding": [
            {
               "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
               "display": "Provident quam porro maiores.",
               "system": "Quisquam voluptas ut aut explicabo et.",
               "userSelected": true,
               "version": "Qui eaque et veritatis dolorum voluptatibus."
            },
            {
               "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
               "display": "Provident quam porro maiores.",
               "system": "Quisquam voluptas ut aut explicabo et.",
               "userSelected": true,
               "version": "Qui eaque et veritatis dolorum voluptatibus."
            }
         ],
         "text": "Vitae perferendis odit corporis error consequatur."
      },
      "issued": "1980-02-17T14:56:04-05:00",
      "method": {
         "coding": [
            {
               "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
               "display": "Provident quam porro maiores.",
               "system": "Quisquam voluptas ut aut explicabo et.",
               "userSelected": true,
               "version": "Qui eaque et veritatis dolorum voluptatibus."
            },
            {
               "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
               "display": "Provident quam porro maiores.",
               "system": "Quisquam voluptas ut aut explicabo et.",
               "userSelected": true,
               "version": "Qui eaque et veritatis dolorum voluptatibus."
            }
         ],
         "text": "Vitae perferendis odit corporis error consequatur."
      },
      "performer": [
         {
            "display": "Dolores aut reprehenderit adipisci.",
            "reference": "Voluptate et."
         },
         {
            "display": "Dolores aut reprehenderit adipisci.",
            "reference": "Voluptate et."
         },
         {
            "display": "Dolores aut reprehenderit adipisci.",
            "reference": "Voluptate et."
         }
      ],
      "referenceRange": [
         {
            "age": {
               "high": {
                  "code": "Ex voluptatem qui nihil quis.",
                  "comparator": "\u003e",
                  "system": "http://walter.biz/manuel",
                  "unit": "Molestiae eos.",
                  "value": 0.3374809955821533
               },
               "low": {
                  "code": "Ex voluptatem qui nihil quis.",
                  "comparator": "\u003e",
                  "system": "http://walter.biz/manuel",
                  "unit": "Molestiae eos.",
                  "value": 0.3374809955821533
               }
            },
            "high": {
               "code": "Ex voluptatem qui nihil quis.",
               "comparator": "\u003e",
               "system": "http://walter.biz/manuel",
               "unit": "Molestiae eos.",
               "value": 0.3374809955821533
            },
            "low": {
               "code": "Ex voluptatem qui nihil quis.",
               "comparator": "\u003e",
               "system": "http://walter.biz/manuel",
               "unit": "Molestiae eos.",
               "value": 0.3374809955821533
            },
            "meaning": {
               "coding": [
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  },
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  }
               ],
               "text": "Vitae perferendis odit corporis error consequatur."
            },
            "text": "Est suscipit modi dignissimos."
         },
         {
            "age": {
               "high": {
                  "code": "Ex voluptatem qui nihil quis.",
                  "comparator": "\u003e",
                  "system": "http://walter.biz/manuel",
                  "unit": "Molestiae eos.",
                  "value": 0.3374809955821533
               },
               "low": {
                  "code": "Ex voluptatem qui nihil quis.",
                  "comparator": "\u003e",
                  "system": "http://walter.biz/manuel",
                  "unit": "Molestiae eos.",
                  "value": 0.3374809955821533
               }
            },
            "high": {
               "code": "Ex voluptatem qui nihil quis.",
               "comparator": "\u003e",
               "system": "http://walter.biz/manuel",
               "unit": "Molestiae eos.",
               "value": 0.3374809955821533
            },
            "low": {
               "code": "Ex voluptatem qui nihil quis.",
               "comparator": "\u003e",
               "system": "http://walter.biz/manuel",
               "unit": "Molestiae eos.",
               "value": 0.3374809955821533
            },
            "meaning": {
               "coding": [
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  },
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  }
               ],
               "text": "Vitae perferendis odit corporis error consequatur."
            },
            "text": "Est suscipit modi dignissimos."
         },
         {
            "age": {
               "high": {
                  "code": "Ex voluptatem qui nihil quis.",
                  "comparator": "\u003e",
                  "system": "http://walter.biz/manuel",
                  "unit": "Molestiae eos.",
                  "value": 0.3374809955821533
               },
               "low": {
                  "code": "Ex voluptatem qui nihil quis.",
                  "comparator": "\u003e",
                  "system": "http://walter.biz/manuel",
                  "unit": "Molestiae eos.",
                  "value": 0.3374809955821533
               }
            },
            "high": {
               "code": "Ex voluptatem qui nihil quis.",
               "comparator": "\u003e",
               "system": "http://walter.biz/manuel",
               "unit": "Molestiae eos.",
               "value": 0.3374809955821533
            },
            "low": {
               "code": "Ex voluptatem qui nihil quis.",
               "comparator": "\u003e",
               "system": "http://walter.biz/manuel",
               "unit": "Molestiae eos.",
               "value": 0.3374809955821533
            },
            "meaning": {
               "coding": [
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  },
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  }
               ],
               "text": "Vitae perferendis odit corporis error consequatur."
            },
            "text": "Est suscipit modi dignissimos."
         }
      ],
      "related": [
         {
            "target": {
               "coding": [
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  },
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  }
               ],
               "text": "Vitae perferendis odit corporis error consequatur."
            },
            "type": "Consequatur itaque."
         }
      ],
      "specimen": {
         "display": "Dolores aut reprehenderit adipisci.",
         "reference": "Voluptate et."
      },
      "status": "preliminary",
      "valueAttachment": {
         "contentType": "In et sunt iure quo tempora consequuntur.",
         "creation": "2010-04-13T22:20:06-04:00",
         "data": "Iusto aperiam.",
         "hash": "Unde consequuntur esse qui.",
         "size": 7170819194107951440,
         "title": "Voluptatem id eum cupiditate."
      },
      "valueCodeableConcept": {
         "coding": [
            {
               "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
               "display": "Provident quam porro maiores.",
               "system": "Quisquam voluptas ut aut explicabo et.",
               "userSelected": true,
               "version": "Qui eaque et veritatis dolorum voluptatibus."
            },
            {
               "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
               "display": "Provident quam porro maiores.",
               "system": "Quisquam voluptas ut aut explicabo et.",
               "userSelected": true,
               "version": "Qui eaque et veritatis dolorum voluptatibus."
            }
         ],
         "text": "Vitae perferendis odit corporis error consequatur."
      },
      "valueDatTime": "1976-02-02T04:43:55-05:00",
      "valuePeriod": {
         "end": "2001-10-30T19:00:22-04:00",
         "start": "1970-01-13T19:32:51-05:00"
      },
      "valueQuantity": {
         "code": "Ex voluptatem qui nihil quis.",
         "comparator": "\u003e",
         "system": "http://walter.biz/manuel",
         "unit": "Molestiae eos.",
         "value": 0.3374809955821533
      },
      "valueRange": {
         "high": {
            "code": "Ex voluptatem qui nihil quis.",
            "comparator": "\u003e",
            "system": "http://walter.biz/manuel",
            "unit": "Molestiae eos.",
            "value": 0.3374809955821533
         },
         "low": {
            "code": "Ex voluptatem qui nihil quis.",
            "comparator": "\u003e",
            "system": "http://walter.biz/manuel",
            "unit": "Molestiae eos.",
            "value": 0.3374809955821533
         }
      },
      "valueSampledData": {
         "data": "Quis eos earum.",
         "dimensions": 9133536065128410703,
         "factor": 0.2928210072171706,
         "lowerLimit": 0.5963619438089043,
         "origin": {
            "code": "Ex voluptatem qui nihil quis.",
            "comparator": "\u003e",
            "system": "http://walter.biz/manuel",
            "unit": "Molestiae eos.",
            "value": 0.3374809955821533
         },
         "period": 0.47228831209882643,
         "upperLimit": 0.6070474669537046
      },
      "valueString": "Nihil fugit est.",
      "valueTime": "2014-09-30T03:57:06-04:00"
   }
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp3.Run(c, args) },
	}
	tmp3.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp3.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp4 := new(CreatePatientCommand)
	sub = &cobra.Command{
		Use:   `patient ["/nosh/patients"]`,
		Short: `This resource uses JWT to secure its endpoints`,
		Long: `This resource uses JWT to secure its endpoints

Payload example:

{
   "name": "Nisi velit assumenda in."
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp4.Run(c, args) },
	}
	tmp4.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp4.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp5 := new(CreateUserCommand)
	sub = &cobra.Command{
		Use:   `user ["/nosh/users"]`,
		Short: `This resource uses JWT to secure its endpoints`,
		Long: `This resource uses JWT to secure its endpoints

Payload example:

{
   "email": "jim.smith@gmail.com",
   "password": "password"
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp5.Run(c, args) },
	}
	tmp5.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp5.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "delete",
		Short: `delete action`,
	}
	tmp6 := new(DeleteAllergyIntoleranceCommand)
	sub = &cobra.Command{
		Use:   `AllergyIntolerance ["/nosh/patients/PATIENTID/allergy_intolerance/ALLERGY_INTOLERANCEID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp6.Run(c, args) },
	}
	tmp6.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp6.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp7 := new(DeleteNutritionRequestCommand)
	sub = &cobra.Command{
		Use:   `NutritionRequest ["/nosh/patients/PATIENTID/nutrition_requests/NUTRITION_REQUESTID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp7.Run(c, args) },
	}
	tmp7.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp7.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp8 := new(DeleteObservationCommand)
	sub = &cobra.Command{
		Use:   `Observation ["/nosh/patients/PATIENTID/observation/OBSERVATIONID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp8.Run(c, args) },
	}
	tmp8.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp8.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp9 := new(DeletePatientCommand)
	sub = &cobra.Command{
		Use:   `patient ["/nosh/patients/PATIENTID"]`,
		Short: `This resource uses JWT to secure its endpoints`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp9.Run(c, args) },
	}
	tmp9.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp9.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp10 := new(DeleteUserCommand)
	sub = &cobra.Command{
		Use:   `user ["/nosh/users/USERID"]`,
		Short: `This resource uses JWT to secure its endpoints`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp10.Run(c, args) },
	}
	tmp10.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp10.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "health",
		Short: `Perform health check.`,
	}
	tmp11 := new(HealthHealthCommand)
	sub = &cobra.Command{
		Use:   `health ["/nosh/_ah/health"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp11.Run(c, args) },
	}
	tmp11.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp11.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "list",
		Short: `list action`,
	}
	tmp12 := new(ListAllergyIntoleranceCommand)
	sub = &cobra.Command{
		Use:   `AllergyIntolerance ["/nosh/patients/PATIENTID/allergy_intolerance"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp12.Run(c, args) },
	}
	tmp12.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp12.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp13 := new(ListNutritionRequestCommand)
	sub = &cobra.Command{
		Use:   `NutritionRequest ["/nosh/patients/PATIENTID/nutrition_requests"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp13.Run(c, args) },
	}
	tmp13.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp13.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp14 := new(ListObservationCommand)
	sub = &cobra.Command{
		Use:   `Observation ["/nosh/patients/PATIENTID/observation"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp14.Run(c, args) },
	}
	tmp14.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp14.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp15 := new(ListPatientCommand)
	sub = &cobra.Command{
		Use:   `patient [("/nosh/patients/jwt"|"/nosh/patients")]`,
		Short: `This resource uses JWT to secure its endpoints`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp15.Run(c, args) },
	}
	tmp15.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp15.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp16 := new(ListUserCommand)
	sub = &cobra.Command{
		Use:   `user ["/nosh/users"]`,
		Short: `This resource uses JWT to secure its endpoints`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp16.Run(c, args) },
	}
	tmp16.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp16.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "rate",
		Short: `rate action`,
	}
	tmp17 := new(RateAllergyIntoleranceCommand)
	sub = &cobra.Command{
		Use:   `AllergyIntolerance ["/nosh/patients/PATIENTID/allergy_intolerance/ALLERGY_INTOLERANCEID/actions/rate"]`,
		Short: ``,
		Long: `

Payload example:

{
   "rating": "Quaerat quas debitis illo delectus."
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp17.Run(c, args) },
	}
	tmp17.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp17.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp18 := new(RateNutritionRequestCommand)
	sub = &cobra.Command{
		Use:   `NutritionRequest ["/nosh/patients/PATIENTID/nutrition_requests/NUTRITION_REQUESTID/actions/rate"]`,
		Short: ``,
		Long: `

Payload example:

{
   "rating": "Enim architecto quos."
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp18.Run(c, args) },
	}
	tmp18.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp18.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp19 := new(RateObservationCommand)
	sub = &cobra.Command{
		Use:   `Observation ["/nosh/patients/PATIENTID/observation/OBSERVATIONID/actions/rate"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp19.Run(c, args) },
	}
	tmp19.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp19.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "secure",
		Short: `secure action`,
	}
	tmp20 := new(SecureBasicCommand)
	sub = &cobra.Command{
		Use:   `basic ["/nosh/basic"]`,
		Short: `This resource uses basic auth to secure its endpoints`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp20.Run(c, args) },
	}
	tmp20.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp20.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp21 := new(SecureJWTCommand)
	sub = &cobra.Command{
		Use:   `jwt ["/nosh/jwt"]`,
		Short: `This resource uses JWT to secure its endpoints`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp21.Run(c, args) },
	}
	tmp21.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp21.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "show",
		Short: `show action`,
	}
	tmp22 := new(ShowAllergyIntoleranceCommand)
	sub = &cobra.Command{
		Use:   `AllergyIntolerance ["/nosh/patients/PATIENTID/allergy_intolerance/ALLERGY_INTOLERANCEID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp22.Run(c, args) },
	}
	tmp22.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp22.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp23 := new(ShowNutritionRequestCommand)
	sub = &cobra.Command{
		Use:   `NutritionRequest ["/nosh/patients/PATIENTID/nutrition_requests/NUTRITION_REQUESTID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp23.Run(c, args) },
	}
	tmp23.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp23.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp24 := new(ShowObservationCommand)
	sub = &cobra.Command{
		Use:   `Observation ["/nosh/patients/PATIENTID/observation/OBSERVATIONID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp24.Run(c, args) },
	}
	tmp24.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp24.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp25 := new(ShowPatientCommand)
	sub = &cobra.Command{
		Use:   `patient ["/nosh/patients/PATIENTID"]`,
		Short: `This resource uses JWT to secure its endpoints`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp25.Run(c, args) },
	}
	tmp25.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp25.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp26 := new(ShowUserCommand)
	sub = &cobra.Command{
		Use:   `user ["/nosh/users/USERID"]`,
		Short: `This resource uses JWT to secure its endpoints`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp26.Run(c, args) },
	}
	tmp26.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp26.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "signin",
		Short: `signin action`,
	}
	tmp27 := new(SigninJWTCommand)
	sub = &cobra.Command{
		Use:   `jwt ["/nosh/jwt/signin"]`,
		Short: `This resource uses JWT to secure its endpoints`,
		Long: `This resource uses JWT to secure its endpoints

Payload example:

{
   "password": "Tenetur ratione nulla et.",
   "username": "Labore quos autem non quos."
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp27.Run(c, args) },
	}
	tmp27.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp27.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp28 := new(SigninUserCommand)
	sub = &cobra.Command{
		Use:   `user ["/nosh/users/jwt/signin"]`,
		Short: `This resource uses JWT to secure its endpoints`,
		Long: `This resource uses JWT to secure its endpoints

Payload example:

{
   "password": "password",
   "username": "jim"
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp28.Run(c, args) },
	}
	tmp28.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp28.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "signup",
		Short: `signup action`,
	}
	tmp29 := new(SignupJWTCommand)
	sub = &cobra.Command{
		Use:   `jwt ["/nosh/jwt/signup"]`,
		Short: `This resource uses JWT to secure its endpoints`,
		Long: `This resource uses JWT to secure its endpoints

Payload example:

{
   "address_city": "Velit corporis atque impedit et iste inventore.",
   "address_line": "Impedit omnis quibusdam non.",
   "address_postal_code": "Nihil laborum numquam debitis alias non illo.",
   "address_state": "Molestiae a similique saepe.",
   "email": "Assumenda delectus et qui a quibusdam magni.",
   "first_name": "Quo non labore.",
   "last_name": "Aperiam quidem et sunt architecto repudiandae facilis.",
   "password": "Magni dolor laborum molestiae hic atque qui.",
   "username": "Provident veniam est."
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp29.Run(c, args) },
	}
	tmp29.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp29.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp30 := new(SignupUserCommand)
	sub = &cobra.Command{
		Use:   `user ["/nosh/users/jwt/signup"]`,
		Short: `This resource uses JWT to secure its endpoints`,
		Long: `This resource uses JWT to secure its endpoints

Payload example:

{
   "address_city": "Carmel",
   "address_line": "533 Worth Ct",
   "address_postal_code": 46032,
   "address_state": "IN",
   "email": "jim.smith@gmail.com",
   "first_name": "Jim",
   "last_name": "Smith",
   "password": "password",
   "username": "jim"
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp30.Run(c, args) },
	}
	tmp30.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp30.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "unsecure",
		Short: `unsecure action`,
	}
	tmp31 := new(UnsecureBasicCommand)
	sub = &cobra.Command{
		Use:   `basic ["/nosh/basic/unsecure"]`,
		Short: `This resource uses basic auth to secure its endpoints`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp31.Run(c, args) },
	}
	tmp31.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp31.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp32 := new(UnsecureJWTCommand)
	sub = &cobra.Command{
		Use:   `jwt ["/nosh/jwt/unsecure"]`,
		Short: `This resource uses JWT to secure its endpoints`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp32.Run(c, args) },
	}
	tmp32.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp32.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "update",
		Short: `update action`,
	}
	tmp33 := new(UpdateAllergyIntoleranceCommand)
	sub = &cobra.Command{
		Use:   `AllergyIntolerance ["/nosh/patients/PATIENTID/allergy_intolerance/ALLERGY_INTOLERANCEID"]`,
		Short: ``,
		Long: `

Payload example:

{
   "active": true,
   "address": {
      "city": "Autem molestiae asperiores consequatur animi.",
      "number": "hi",
      "street": "Quod dolor distinctio suscipit ut perspiciatis."
   },
   "birthDate": "Merlot",
   "country": "USA",
   "deceased": false,
   "gender": "male",
   "multiple_birth": false,
   "region": "Napa Valley",
   "review": "Great and inexpensive",
   "sweetness": 4
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp33.Run(c, args) },
	}
	tmp33.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp33.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp34 := new(UpdateNutritionRequestCommand)
	sub = &cobra.Command{
		Use:   `NutritionRequest ["/nosh/patients/PATIENTID/nutrition_requests/NUTRITION_REQUESTID"]`,
		Short: ``,
		Long: `

Payload example:

{
   "address": "Suscipit in in reiciendis officiis saepe voluptate.",
   "address_postal_code": 46032,
   "enteralFormula": "Dolores aperiam et voluptas dolorem deleniti maxime.",
   "foodPreferenceModifier": {
      "coding": "Qui et.",
      "id": "Eaque sequi quaerat et.",
      "text": "Molestiae quae eaque."
   },
   "href": "Veritatis sit nam dolorem.",
   "id": "Tenetur non molestiae.",
   "identifier": "Iure repellat numquam.",
   "links": "Deserunt consequatur aliquam animi.",
   "nutritionRequest": {
      "allergyIntolerance": [
         {
            "display": "Dolores aut reprehenderit adipisci.",
            "reference": "Voluptate et."
         }
      ],
      "dateTime": "2006-11-15T18:31:46-05:00",
      "enteralFormula": [
         {
            "additiveProductName": "Cupiditate necessitatibus libero rem esse eos quibusdam.",
            "additiveType": {
               "coding": [
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  },
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  }
               ],
               "text": "Vitae perferendis odit corporis error consequatur."
            },
            "administrativeInstruction": "Cumque aliquam.",
            "baseFormulaType": {
               "coding": [
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  },
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  }
               ],
               "text": "Vitae perferendis odit corporis error consequatur."
            },
            "baseFormulatProdcutName": "Repudiandae deserunt quia cumque dicta est.",
            "caloricDensity": {
               "code": "Ex voluptatem qui nihil quis.",
               "comparator": "\u003e",
               "system": "http://walter.biz/manuel",
               "unit": "Molestiae eos.",
               "value": 0.3374809955821533
            },
            "maxVolumeToDeliver": {
               "code": "Ex voluptatem qui nihil quis.",
               "comparator": "\u003e",
               "system": "http://walter.biz/manuel",
               "unit": "Molestiae eos.",
               "value": 0.3374809955821533
            },
            "routeofAdministration": {
               "coding": [
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  },
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  }
               ],
               "text": "Vitae perferendis odit corporis error consequatur."
            }
         }
      ],
      "excludeFoodModifier": [
         {
            "coding": [
               {
                  "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                  "display": "Provident quam porro maiores.",
                  "system": "Quisquam voluptas ut aut explicabo et.",
                  "userSelected": true,
                  "version": "Qui eaque et veritatis dolorum voluptatibus."
               },
               {
                  "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                  "display": "Provident quam porro maiores.",
                  "system": "Quisquam voluptas ut aut explicabo et.",
                  "userSelected": true,
                  "version": "Qui eaque et veritatis dolorum voluptatibus."
               }
            ],
            "text": "Vitae perferendis odit corporis error consequatur."
         }
      ],
      "foodPreferenceModifier": [
         {
            "coding": [
               {
                  "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                  "display": "Provident quam porro maiores.",
                  "system": "Quisquam voluptas ut aut explicabo et.",
                  "userSelected": true,
                  "version": "Qui eaque et veritatis dolorum voluptatibus."
               },
               {
                  "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                  "display": "Provident quam porro maiores.",
                  "system": "Quisquam voluptas ut aut explicabo et.",
                  "userSelected": true,
                  "version": "Qui eaque et veritatis dolorum voluptatibus."
               }
            ],
            "text": "Vitae perferendis odit corporis error consequatur."
         },
         {
            "coding": [
               {
                  "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                  "display": "Provident quam porro maiores.",
                  "system": "Quisquam voluptas ut aut explicabo et.",
                  "userSelected": true,
                  "version": "Qui eaque et veritatis dolorum voluptatibus."
               },
               {
                  "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                  "display": "Provident quam porro maiores.",
                  "system": "Quisquam voluptas ut aut explicabo et.",
                  "userSelected": true,
                  "version": "Qui eaque et veritatis dolorum voluptatibus."
               }
            ],
            "text": "Vitae perferendis odit corporis error consequatur."
         },
         {
            "coding": [
               {
                  "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                  "display": "Provident quam porro maiores.",
                  "system": "Quisquam voluptas ut aut explicabo et.",
                  "userSelected": true,
                  "version": "Qui eaque et veritatis dolorum voluptatibus."
               },
               {
                  "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                  "display": "Provident quam porro maiores.",
                  "system": "Quisquam voluptas ut aut explicabo et.",
                  "userSelected": true,
                  "version": "Qui eaque et veritatis dolorum voluptatibus."
               }
            ],
            "text": "Vitae perferendis odit corporis error consequatur."
         }
      ],
      "href": "/patients/1/nutrition_requests/1",
      "id": 1,
      "identifier": [
         {
            "CodeableConcept": {
               "coding": [
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  },
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  }
               ],
               "text": "Vitae perferendis odit corporis error consequatur."
            },
            "assigner": {
               "display": "Dolores aut reprehenderit adipisci.",
               "reference": "Voluptate et."
            },
            "period": {
               "end": "2001-10-30T19:00:22-04:00",
               "start": "1970-01-13T19:32:51-05:00"
            }
         },
         {
            "CodeableConcept": {
               "coding": [
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  },
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  }
               ],
               "text": "Vitae perferendis odit corporis error consequatur."
            },
            "assigner": {
               "display": "Dolores aut reprehenderit adipisci.",
               "reference": "Voluptate et."
            },
            "period": {
               "end": "2001-10-30T19:00:22-04:00",
               "start": "1970-01-13T19:32:51-05:00"
            }
         }
      ],
      "links": {
         "enteralFormula": [
            {
               "additiveProductName": "Cupiditate necessitatibus libero rem esse eos quibusdam.",
               "additiveType": {
                  "coding": [
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     },
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     }
                  ],
                  "text": "Vitae perferendis odit corporis error consequatur."
               },
               "administrativeInstruction": "Cumque aliquam.",
               "baseFormulaType": {
                  "coding": [
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     },
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     }
                  ],
                  "text": "Vitae perferendis odit corporis error consequatur."
               },
               "baseFormulatProdcutName": "Repudiandae deserunt quia cumque dicta est.",
               "caloricDensity": {
                  "code": "Ex voluptatem qui nihil quis.",
                  "comparator": "\u003e",
                  "system": "http://walter.biz/manuel",
                  "unit": "Molestiae eos.",
                  "value": 0.3374809955821533
               },
               "maxVolumeToDeliver": {
                  "code": "Ex voluptatem qui nihil quis.",
                  "comparator": "\u003e",
                  "system": "http://walter.biz/manuel",
                  "unit": "Molestiae eos.",
                  "value": 0.3374809955821533
               },
               "routeofAdministration": {
                  "coding": [
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     },
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     }
                  ],
                  "text": "Vitae perferendis odit corporis error consequatur."
               }
            },
            {
               "additiveProductName": "Cupiditate necessitatibus libero rem esse eos quibusdam.",
               "additiveType": {
                  "coding": [
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     },
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     }
                  ],
                  "text": "Vitae perferendis odit corporis error consequatur."
               },
               "administrativeInstruction": "Cumque aliquam.",
               "baseFormulaType": {
                  "coding": [
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     },
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     }
                  ],
                  "text": "Vitae perferendis odit corporis error consequatur."
               },
               "baseFormulatProdcutName": "Repudiandae deserunt quia cumque dicta est.",
               "caloricDensity": {
                  "code": "Ex voluptatem qui nihil quis.",
                  "comparator": "\u003e",
                  "system": "http://walter.biz/manuel",
                  "unit": "Molestiae eos.",
                  "value": 0.3374809955821533
               },
               "maxVolumeToDeliver": {
                  "code": "Ex voluptatem qui nihil quis.",
                  "comparator": "\u003e",
                  "system": "http://walter.biz/manuel",
                  "unit": "Molestiae eos.",
                  "value": 0.3374809955821533
               },
               "routeofAdministration": {
                  "coding": [
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     },
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     }
                  ],
                  "text": "Vitae perferendis odit corporis error consequatur."
               }
            }
         ],
         "oralDiet": [
            {
               "fluidConsistencyType": {
                  "coding": [
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     },
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     }
                  ],
                  "text": "Vitae perferendis odit corporis error consequatur."
               },
               "instruction": "Qui fugit error.",
               "nutrient": {
                  "amount": {
                     "code": "Ex voluptatem qui nihil quis.",
                     "comparator": "\u003e",
                     "system": "http://walter.biz/manuel",
                     "unit": "Molestiae eos.",
                     "value": 0.3374809955821533
                  },
                  "modifier": {
                     "coding": [
                        {
                           "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                           "display": "Provident quam porro maiores.",
                           "system": "Quisquam voluptas ut aut explicabo et.",
                           "userSelected": true,
                           "version": "Qui eaque et veritatis dolorum voluptatibus."
                        },
                        {
                           "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                           "display": "Provident quam porro maiores.",
                           "system": "Quisquam voluptas ut aut explicabo et.",
                           "userSelected": true,
                           "version": "Qui eaque et veritatis dolorum voluptatibus."
                        }
                     ],
                     "text": "Vitae perferendis odit corporis error consequatur."
                  }
               },
               "schedule": {
                  "code": "AM",
                  "event": "1978-12-17T03:44:52-05:00",
                  "repeat": {
                     "extension": {
                        "ValueAddress": {
                           "city": "Consequatur omnis omnis tempore et.",
                           "country": "Suscipit ut quia sequi atque numquam.",
                           "distinct": "Quia fuga.",
                           "line": "Vero dicta voluptatibus consectetur ut minus.",
                           "postalCode": 8606407676718086019,
                           "state": "Neque quod similique est quia.",
                           "type": "postal",
                           "use": "work"
                        },
                        "ValueAnnotation": {
                           "authorReference": {
                              "display": "Dolores aut reprehenderit adipisci.",
                              "reference": "Voluptate et."
                           },
                           "authorString": "Quasi cupiditate quia debitis vel.",
                           "time": "1992-01-31T12:26:43-05:00"
                        },
                        "ValueAttachment": {
                           "contentType": "In et sunt iure quo tempora consequuntur.",
                           "creation": "2010-04-13T22:20:06-04:00",
                           "data": "Iusto aperiam.",
                           "hash": "Unde consequuntur esse qui.",
                           "size": 7170819194107951440,
                           "title": "Voluptatem id eum cupiditate."
                        },
                        "ValueBase64Binary": "Est distinctio voluptate tenetur.",
                        "ValueBoolean": true,
                        "ValueCode": "Placeat vero temporibus et aliquid ut incidunt.",
                        "ValueCodeableConcept": {
                           "coding": [
                              {
                                 "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                                 "display": "Provident quam porro maiores.",
                                 "system": "Quisquam voluptas ut aut explicabo et.",
                                 "userSelected": true,
                                 "version": "Qui eaque et veritatis dolorum voluptatibus."
                              },
                              {
                                 "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                                 "display": "Provident quam porro maiores.",
                                 "system": "Quisquam voluptas ut aut explicabo et.",
                                 "userSelected": true,
                                 "version": "Qui eaque et veritatis dolorum voluptatibus."
                              }
                           ],
                           "text": "Vitae perferendis odit corporis error consequatur."
                        },
                        "ValueCoding": {
                           "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                           "display": "Provident quam porro maiores.",
                           "system": "Quisquam voluptas ut aut explicabo et.",
                           "userSelected": true,
                           "version": "Qui eaque et veritatis dolorum voluptatibus."
                        },
                        "ValueContactPoint": "Dolor quaerat delectus et.",
                        "ValueDate": "1980-05-15T19:46:08-04:00",
                        "ValueDateTime": "1981-04-19T10:31:33-04:00",
                        "ValueDecimal": 0.7932081018594244,
                        "ValueHumanName": {
                           "family": [
                              "Ea culpa qui a enim.",
                              "Ea culpa qui a enim.",
                              "Ea culpa qui a enim."
                           ],
                           "given": "Reiciendis odit.",
                           "period": {
                              "end": "2001-10-30T19:00:22-04:00",
                              "start": "1970-01-13T19:32:51-05:00"
                           },
                           "prefix": "Natus inventore optio eaque.",
                           "suffix": "Id qui delectus quia qui.",
                           "use": "temp"
                        },
                        "ValueId": "Molestiae praesentium quidem corrupti nihil eum.",
                        "ValueIdentifier": {
                           "CodeableConcept": {
                              "coding": [
                                 {
                                    "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                                    "display": "Provident quam porro maiores.",
                                    "system": "Quisquam voluptas ut aut explicabo et.",
                                    "userSelected": true,
                                    "version": "Qui eaque et veritatis dolorum voluptatibus."
                                 },
                                 {
                                    "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                                    "display": "Provident quam porro maiores.",
                                    "system": "Quisquam voluptas ut aut explicabo et.",
                                    "userSelected": true,
                                    "version": "Qui eaque et veritatis dolorum voluptatibus."
                                 }
                              ],
                              "text": "Vitae perferendis odit corporis error consequatur."
                           },
                           "assigner": {
                              "display": "Dolores aut reprehenderit adipisci.",
                              "reference": "Voluptate et."
                           },
                           "period": {
                              "end": "2001-10-30T19:00:22-04:00",
                              "start": "1970-01-13T19:32:51-05:00"
                           }
                        },
                        "ValueInstant": "1977-11-08T04:06:45-05:00",
                        "ValueInteger": 8012361700862165525,
                        "ValueMarkdown": "Sed neque reiciendis.",
                        "ValueMeta": {
                           "lastUpdated": "1975-12-19T03:26:58-05:00",
                           "profile": "http://green.info/marilyne",
                           "security": {
                              "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                              "display": "Provident quam porro maiores.",
                              "system": "Quisquam voluptas ut aut explicabo et.",
                              "userSelected": true,
                              "version": "Qui eaque et veritatis dolorum voluptatibus."
                           },
                           "tag": {
                              "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                              "display": "Provident quam porro maiores.",
                              "system": "Quisquam voluptas ut aut explicabo et.",
                              "userSelected": true,
                              "version": "Qui eaque et veritatis dolorum voluptatibus."
                           },
                           "versionId": "Corrupti nemo dolor et non vel ea."
                        },
                        "ValueOid": "Laudantium porro.",
                        "ValuePeriod": {
                           "end": "2001-10-30T19:00:22-04:00",
                           "start": "1970-01-13T19:32:51-05:00"
                        },
                        "ValuePositiveInt": 0.6018552052933437,
                        "ValueQuantity": {
                           "code": "Ex voluptatem qui nihil quis.",
                           "comparator": "\u003e",
                           "system": "http://walter.biz/manuel",
                           "unit": "Molestiae eos.",
                           "value": 0.3374809955821533
                        },
                        "ValueRange": "Similique quia quis.",
                        "ValueRatio": 7314611204127242704,
                        "ValueReference": {
                           "display": "Dolores aut reprehenderit adipisci.",
                           "reference": "Voluptate et."
                        },
                        "ValueSampledData": "Unde odio aut suscipit atque.",
                        "ValueSignature": "Libero sequi ut.",
                        "ValueString": "Molestiae in.",
                        "ValueTime": "2013-02-01T03:56:35-05:00",
                        "ValueTiming": "1993-07-31T06:07:28-04:00",
                        "ValueUnsignedInt": 0.8739176042943996,
                        "ValueUri": "A expedita vel aut voluptatum veritatis.",
                        "url": "Nam earum repellendus officia in rerum aut."
                     },
                     "id": "Dolores labore laudantium."
                  }
               },
               "texture": {
                  "foodType": {
                     "coding": [
                        {
                           "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                           "display": "Provident quam porro maiores.",
                           "system": "Quisquam voluptas ut aut explicabo et.",
                           "userSelected": true,
                           "version": "Qui eaque et veritatis dolorum voluptatibus."
                        },
                        {
                           "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                           "display": "Provident quam porro maiores.",
                           "system": "Quisquam voluptas ut aut explicabo et.",
                           "userSelected": true,
                           "version": "Qui eaque et veritatis dolorum voluptatibus."
                        }
                     ],
                     "text": "Vitae perferendis odit corporis error consequatur."
                  },
                  "modifier": {
                     "coding": [
                        {
                           "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                           "display": "Provident quam porro maiores.",
                           "system": "Quisquam voluptas ut aut explicabo et.",
                           "userSelected": true,
                           "version": "Qui eaque et veritatis dolorum voluptatibus."
                        },
                        {
                           "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                           "display": "Provident quam porro maiores.",
                           "system": "Quisquam voluptas ut aut explicabo et.",
                           "userSelected": true,
                           "version": "Qui eaque et veritatis dolorum voluptatibus."
                        }
                     ],
                     "text": "Vitae perferendis odit corporis error consequatur."
                  }
               },
               "type": {
                  "coding": [
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     },
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     }
                  ],
                  "text": "Vitae perferendis odit corporis error consequatur."
               }
            }
         ]
      },
      "oralDiet": [
         {
            "fluidConsistencyType": {
               "coding": [
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  },
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  }
               ],
               "text": "Vitae perferendis odit corporis error consequatur."
            },
            "instruction": "Qui fugit error.",
            "nutrient": {
               "amount": {
                  "code": "Ex voluptatem qui nihil quis.",
                  "comparator": "\u003e",
                  "system": "http://walter.biz/manuel",
                  "unit": "Molestiae eos.",
                  "value": 0.3374809955821533
               },
               "modifier": {
                  "coding": [
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     },
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     }
                  ],
                  "text": "Vitae perferendis odit corporis error consequatur."
               }
            },
            "schedule": {
               "code": "AM",
               "event": "1978-12-17T03:44:52-05:00",
               "repeat": {
                  "extension": {
                     "ValueAddress": {
                        "city": "Consequatur omnis omnis tempore et.",
                        "country": "Suscipit ut quia sequi atque numquam.",
                        "distinct": "Quia fuga.",
                        "line": "Vero dicta voluptatibus consectetur ut minus.",
                        "postalCode": 8606407676718086019,
                        "state": "Neque quod similique est quia.",
                        "type": "postal",
                        "use": "work"
                     },
                     "ValueAnnotation": {
                        "authorReference": {
                           "display": "Dolores aut reprehenderit adipisci.",
                           "reference": "Voluptate et."
                        },
                        "authorString": "Quasi cupiditate quia debitis vel.",
                        "time": "1992-01-31T12:26:43-05:00"
                     },
                     "ValueAttachment": {
                        "contentType": "In et sunt iure quo tempora consequuntur.",
                        "creation": "2010-04-13T22:20:06-04:00",
                        "data": "Iusto aperiam.",
                        "hash": "Unde consequuntur esse qui.",
                        "size": 7170819194107951440,
                        "title": "Voluptatem id eum cupiditate."
                     },
                     "ValueBase64Binary": "Est distinctio voluptate tenetur.",
                     "ValueBoolean": true,
                     "ValueCode": "Placeat vero temporibus et aliquid ut incidunt.",
                     "ValueCodeableConcept": {
                        "coding": [
                           {
                              "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                              "display": "Provident quam porro maiores.",
                              "system": "Quisquam voluptas ut aut explicabo et.",
                              "userSelected": true,
                              "version": "Qui eaque et veritatis dolorum voluptatibus."
                           },
                           {
                              "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                              "display": "Provident quam porro maiores.",
                              "system": "Quisquam voluptas ut aut explicabo et.",
                              "userSelected": true,
                              "version": "Qui eaque et veritatis dolorum voluptatibus."
                           }
                        ],
                        "text": "Vitae perferendis odit corporis error consequatur."
                     },
                     "ValueCoding": {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     },
                     "ValueContactPoint": "Dolor quaerat delectus et.",
                     "ValueDate": "1980-05-15T19:46:08-04:00",
                     "ValueDateTime": "1981-04-19T10:31:33-04:00",
                     "ValueDecimal": 0.7932081018594244,
                     "ValueHumanName": {
                        "family": [
                           "Ea culpa qui a enim.",
                           "Ea culpa qui a enim.",
                           "Ea culpa qui a enim."
                        ],
                        "given": "Reiciendis odit.",
                        "period": {
                           "end": "2001-10-30T19:00:22-04:00",
                           "start": "1970-01-13T19:32:51-05:00"
                        },
                        "prefix": "Natus inventore optio eaque.",
                        "suffix": "Id qui delectus quia qui.",
                        "use": "temp"
                     },
                     "ValueId": "Molestiae praesentium quidem corrupti nihil eum.",
                     "ValueIdentifier": {
                        "CodeableConcept": {
                           "coding": [
                              {
                                 "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                                 "display": "Provident quam porro maiores.",
                                 "system": "Quisquam voluptas ut aut explicabo et.",
                                 "userSelected": true,
                                 "version": "Qui eaque et veritatis dolorum voluptatibus."
                              },
                              {
                                 "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                                 "display": "Provident quam porro maiores.",
                                 "system": "Quisquam voluptas ut aut explicabo et.",
                                 "userSelected": true,
                                 "version": "Qui eaque et veritatis dolorum voluptatibus."
                              }
                           ],
                           "text": "Vitae perferendis odit corporis error consequatur."
                        },
                        "assigner": {
                           "display": "Dolores aut reprehenderit adipisci.",
                           "reference": "Voluptate et."
                        },
                        "period": {
                           "end": "2001-10-30T19:00:22-04:00",
                           "start": "1970-01-13T19:32:51-05:00"
                        }
                     },
                     "ValueInstant": "1977-11-08T04:06:45-05:00",
                     "ValueInteger": 8012361700862165525,
                     "ValueMarkdown": "Sed neque reiciendis.",
                     "ValueMeta": {
                        "lastUpdated": "1975-12-19T03:26:58-05:00",
                        "profile": "http://green.info/marilyne",
                        "security": {
                           "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                           "display": "Provident quam porro maiores.",
                           "system": "Quisquam voluptas ut aut explicabo et.",
                           "userSelected": true,
                           "version": "Qui eaque et veritatis dolorum voluptatibus."
                        },
                        "tag": {
                           "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                           "display": "Provident quam porro maiores.",
                           "system": "Quisquam voluptas ut aut explicabo et.",
                           "userSelected": true,
                           "version": "Qui eaque et veritatis dolorum voluptatibus."
                        },
                        "versionId": "Corrupti nemo dolor et non vel ea."
                     },
                     "ValueOid": "Laudantium porro.",
                     "ValuePeriod": {
                        "end": "2001-10-30T19:00:22-04:00",
                        "start": "1970-01-13T19:32:51-05:00"
                     },
                     "ValuePositiveInt": 0.6018552052933437,
                     "ValueQuantity": {
                        "code": "Ex voluptatem qui nihil quis.",
                        "comparator": "\u003e",
                        "system": "http://walter.biz/manuel",
                        "unit": "Molestiae eos.",
                        "value": 0.3374809955821533
                     },
                     "ValueRange": "Similique quia quis.",
                     "ValueRatio": 7314611204127242704,
                     "ValueReference": {
                        "display": "Dolores aut reprehenderit adipisci.",
                        "reference": "Voluptate et."
                     },
                     "ValueSampledData": "Unde odio aut suscipit atque.",
                     "ValueSignature": "Libero sequi ut.",
                     "ValueString": "Molestiae in.",
                     "ValueTime": "2013-02-01T03:56:35-05:00",
                     "ValueTiming": "1993-07-31T06:07:28-04:00",
                     "ValueUnsignedInt": 0.8739176042943996,
                     "ValueUri": "A expedita vel aut voluptatum veritatis.",
                     "url": "Nam earum repellendus officia in rerum aut."
                  },
                  "id": "Dolores labore laudantium."
               }
            },
            "texture": {
               "foodType": {
                  "coding": [
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     },
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     }
                  ],
                  "text": "Vitae perferendis odit corporis error consequatur."
               },
               "modifier": {
                  "coding": [
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     },
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     }
                  ],
                  "text": "Vitae perferendis odit corporis error consequatur."
               }
            },
            "type": {
               "coding": [
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  },
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  }
               ],
               "text": "Vitae perferendis odit corporis error consequatur."
            }
         },
         {
            "fluidConsistencyType": {
               "coding": [
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  },
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  }
               ],
               "text": "Vitae perferendis odit corporis error consequatur."
            },
            "instruction": "Qui fugit error.",
            "nutrient": {
               "amount": {
                  "code": "Ex voluptatem qui nihil quis.",
                  "comparator": "\u003e",
                  "system": "http://walter.biz/manuel",
                  "unit": "Molestiae eos.",
                  "value": 0.3374809955821533
               },
               "modifier": {
                  "coding": [
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     },
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     }
                  ],
                  "text": "Vitae perferendis odit corporis error consequatur."
               }
            },
            "schedule": {
               "code": "AM",
               "event": "1978-12-17T03:44:52-05:00",
               "repeat": {
                  "extension": {
                     "ValueAddress": {
                        "city": "Consequatur omnis omnis tempore et.",
                        "country": "Suscipit ut quia sequi atque numquam.",
                        "distinct": "Quia fuga.",
                        "line": "Vero dicta voluptatibus consectetur ut minus.",
                        "postalCode": 8606407676718086019,
                        "state": "Neque quod similique est quia.",
                        "type": "postal",
                        "use": "work"
                     },
                     "ValueAnnotation": {
                        "authorReference": {
                           "display": "Dolores aut reprehenderit adipisci.",
                           "reference": "Voluptate et."
                        },
                        "authorString": "Quasi cupiditate quia debitis vel.",
                        "time": "1992-01-31T12:26:43-05:00"
                     },
                     "ValueAttachment": {
                        "contentType": "In et sunt iure quo tempora consequuntur.",
                        "creation": "2010-04-13T22:20:06-04:00",
                        "data": "Iusto aperiam.",
                        "hash": "Unde consequuntur esse qui.",
                        "size": 7170819194107951440,
                        "title": "Voluptatem id eum cupiditate."
                     },
                     "ValueBase64Binary": "Est distinctio voluptate tenetur.",
                     "ValueBoolean": true,
                     "ValueCode": "Placeat vero temporibus et aliquid ut incidunt.",
                     "ValueCodeableConcept": {
                        "coding": [
                           {
                              "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                              "display": "Provident quam porro maiores.",
                              "system": "Quisquam voluptas ut aut explicabo et.",
                              "userSelected": true,
                              "version": "Qui eaque et veritatis dolorum voluptatibus."
                           },
                           {
                              "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                              "display": "Provident quam porro maiores.",
                              "system": "Quisquam voluptas ut aut explicabo et.",
                              "userSelected": true,
                              "version": "Qui eaque et veritatis dolorum voluptatibus."
                           }
                        ],
                        "text": "Vitae perferendis odit corporis error consequatur."
                     },
                     "ValueCoding": {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     },
                     "ValueContactPoint": "Dolor quaerat delectus et.",
                     "ValueDate": "1980-05-15T19:46:08-04:00",
                     "ValueDateTime": "1981-04-19T10:31:33-04:00",
                     "ValueDecimal": 0.7932081018594244,
                     "ValueHumanName": {
                        "family": [
                           "Ea culpa qui a enim.",
                           "Ea culpa qui a enim.",
                           "Ea culpa qui a enim."
                        ],
                        "given": "Reiciendis odit.",
                        "period": {
                           "end": "2001-10-30T19:00:22-04:00",
                           "start": "1970-01-13T19:32:51-05:00"
                        },
                        "prefix": "Natus inventore optio eaque.",
                        "suffix": "Id qui delectus quia qui.",
                        "use": "temp"
                     },
                     "ValueId": "Molestiae praesentium quidem corrupti nihil eum.",
                     "ValueIdentifier": {
                        "CodeableConcept": {
                           "coding": [
                              {
                                 "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                                 "display": "Provident quam porro maiores.",
                                 "system": "Quisquam voluptas ut aut explicabo et.",
                                 "userSelected": true,
                                 "version": "Qui eaque et veritatis dolorum voluptatibus."
                              },
                              {
                                 "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                                 "display": "Provident quam porro maiores.",
                                 "system": "Quisquam voluptas ut aut explicabo et.",
                                 "userSelected": true,
                                 "version": "Qui eaque et veritatis dolorum voluptatibus."
                              }
                           ],
                           "text": "Vitae perferendis odit corporis error consequatur."
                        },
                        "assigner": {
                           "display": "Dolores aut reprehenderit adipisci.",
                           "reference": "Voluptate et."
                        },
                        "period": {
                           "end": "2001-10-30T19:00:22-04:00",
                           "start": "1970-01-13T19:32:51-05:00"
                        }
                     },
                     "ValueInstant": "1977-11-08T04:06:45-05:00",
                     "ValueInteger": 8012361700862165525,
                     "ValueMarkdown": "Sed neque reiciendis.",
                     "ValueMeta": {
                        "lastUpdated": "1975-12-19T03:26:58-05:00",
                        "profile": "http://green.info/marilyne",
                        "security": {
                           "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                           "display": "Provident quam porro maiores.",
                           "system": "Quisquam voluptas ut aut explicabo et.",
                           "userSelected": true,
                           "version": "Qui eaque et veritatis dolorum voluptatibus."
                        },
                        "tag": {
                           "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                           "display": "Provident quam porro maiores.",
                           "system": "Quisquam voluptas ut aut explicabo et.",
                           "userSelected": true,
                           "version": "Qui eaque et veritatis dolorum voluptatibus."
                        },
                        "versionId": "Corrupti nemo dolor et non vel ea."
                     },
                     "ValueOid": "Laudantium porro.",
                     "ValuePeriod": {
                        "end": "2001-10-30T19:00:22-04:00",
                        "start": "1970-01-13T19:32:51-05:00"
                     },
                     "ValuePositiveInt": 0.6018552052933437,
                     "ValueQuantity": {
                        "code": "Ex voluptatem qui nihil quis.",
                        "comparator": "\u003e",
                        "system": "http://walter.biz/manuel",
                        "unit": "Molestiae eos.",
                        "value": 0.3374809955821533
                     },
                     "ValueRange": "Similique quia quis.",
                     "ValueRatio": 7314611204127242704,
                     "ValueReference": {
                        "display": "Dolores aut reprehenderit adipisci.",
                        "reference": "Voluptate et."
                     },
                     "ValueSampledData": "Unde odio aut suscipit atque.",
                     "ValueSignature": "Libero sequi ut.",
                     "ValueString": "Molestiae in.",
                     "ValueTime": "2013-02-01T03:56:35-05:00",
                     "ValueTiming": "1993-07-31T06:07:28-04:00",
                     "ValueUnsignedInt": 0.8739176042943996,
                     "ValueUri": "A expedita vel aut voluptatum veritatis.",
                     "url": "Nam earum repellendus officia in rerum aut."
                  },
                  "id": "Dolores labore laudantium."
               }
            },
            "texture": {
               "foodType": {
                  "coding": [
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     },
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     }
                  ],
                  "text": "Vitae perferendis odit corporis error consequatur."
               },
               "modifier": {
                  "coding": [
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     },
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     }
                  ],
                  "text": "Vitae perferendis odit corporis error consequatur."
               }
            },
            "type": {
               "coding": [
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  },
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  }
               ],
               "text": "Vitae perferendis odit corporis error consequatur."
            }
         }
      ],
      "orderer": {
         "display": "Dolores aut reprehenderit adipisci.",
         "reference": "Voluptate et."
      },
      "patient": {
         "display": "Dolores aut reprehenderit adipisci.",
         "reference": "Voluptate et."
      },
      "status": "completed",
      "supplement": [
         {
            "instruction": "Labore nihil sint quibusdam et.",
            "productName": "Ut impedit.",
            "quantity": {
               "code": "Ex voluptatem qui nihil quis.",
               "comparator": "\u003e",
               "system": "http://walter.biz/manuel",
               "unit": "Molestiae eos.",
               "value": 0.3374809955821533
            },
            "type": {
               "coding": [
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  },
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  }
               ],
               "text": "Vitae perferendis odit corporis error consequatur."
            }
         },
         {
            "instruction": "Labore nihil sint quibusdam et.",
            "productName": "Ut impedit.",
            "quantity": {
               "code": "Ex voluptatem qui nihil quis.",
               "comparator": "\u003e",
               "system": "http://walter.biz/manuel",
               "unit": "Molestiae eos.",
               "value": 0.3374809955821533
            },
            "type": {
               "coding": [
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  },
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  }
               ],
               "text": "Vitae perferendis odit corporis error consequatur."
            }
         }
      ]
   },
   "oralDiet": "At ex.",
   "patient": {
      "id": "Rerum recusandae voluptatum fugiat.",
      "reference": "Enim reiciendis et sed fuga."
   },
   "supplement": "Est autem."
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp34.Run(c, args) },
	}
	tmp34.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp34.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp35 := new(UpdateObservationCommand)
	sub = &cobra.Command{
		Use:   `Observation ["/nosh/patients/PATIENTID/observation/OBSERVATIONID"]`,
		Short: ``,
		Long: `

Payload example:

{
   "observation": {
      "bodySite": {
         "coding": [
            {
               "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
               "display": "Provident quam porro maiores.",
               "system": "Quisquam voluptas ut aut explicabo et.",
               "userSelected": true,
               "version": "Qui eaque et veritatis dolorum voluptatibus."
            },
            {
               "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
               "display": "Provident quam porro maiores.",
               "system": "Quisquam voluptas ut aut explicabo et.",
               "userSelected": true,
               "version": "Qui eaque et veritatis dolorum voluptatibus."
            }
         ],
         "text": "Vitae perferendis odit corporis error consequatur."
      },
      "category": {
         "coding": [
            {
               "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
               "display": "Provident quam porro maiores.",
               "system": "Quisquam voluptas ut aut explicabo et.",
               "userSelected": true,
               "version": "Qui eaque et veritatis dolorum voluptatibus."
            },
            {
               "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
               "display": "Provident quam porro maiores.",
               "system": "Quisquam voluptas ut aut explicabo et.",
               "userSelected": true,
               "version": "Qui eaque et veritatis dolorum voluptatibus."
            }
         ],
         "text": "Vitae perferendis odit corporis error consequatur."
      },
      "code": {
         "coding": [
            {
               "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
               "display": "Provident quam porro maiores.",
               "system": "Quisquam voluptas ut aut explicabo et.",
               "userSelected": true,
               "version": "Qui eaque et veritatis dolorum voluptatibus."
            },
            {
               "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
               "display": "Provident quam porro maiores.",
               "system": "Quisquam voluptas ut aut explicabo et.",
               "userSelected": true,
               "version": "Qui eaque et veritatis dolorum voluptatibus."
            }
         ],
         "text": "Vitae perferendis odit corporis error consequatur."
      },
      "comments": "Voluptate sint officiis qui nulla sapiente minus.",
      "component": [
         {
            "code": {
               "coding": [
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  },
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  }
               ],
               "text": "Vitae perferendis odit corporis error consequatur."
            },
            "dataAbsentReason": {
               "coding": [
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  },
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  }
               ],
               "text": "Vitae perferendis odit corporis error consequatur."
            },
            "referenceRange": {
               "age": {
                  "high": {
                     "code": "Ex voluptatem qui nihil quis.",
                     "comparator": "\u003e",
                     "system": "http://walter.biz/manuel",
                     "unit": "Molestiae eos.",
                     "value": 0.3374809955821533
                  },
                  "low": {
                     "code": "Ex voluptatem qui nihil quis.",
                     "comparator": "\u003e",
                     "system": "http://walter.biz/manuel",
                     "unit": "Molestiae eos.",
                     "value": 0.3374809955821533
                  }
               },
               "high": {
                  "code": "Ex voluptatem qui nihil quis.",
                  "comparator": "\u003e",
                  "system": "http://walter.biz/manuel",
                  "unit": "Molestiae eos.",
                  "value": 0.3374809955821533
               },
               "low": {
                  "code": "Ex voluptatem qui nihil quis.",
                  "comparator": "\u003e",
                  "system": "http://walter.biz/manuel",
                  "unit": "Molestiae eos.",
                  "value": 0.3374809955821533
               },
               "meaning": {
                  "coding": [
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     },
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     }
                  ],
                  "text": "Vitae perferendis odit corporis error consequatur."
               },
               "text": "Est suscipit modi dignissimos."
            },
            "valueAttachment": {
               "contentType": "In et sunt iure quo tempora consequuntur.",
               "creation": "2010-04-13T22:20:06-04:00",
               "data": "Iusto aperiam.",
               "hash": "Unde consequuntur esse qui.",
               "size": 7170819194107951440,
               "title": "Voluptatem id eum cupiditate."
            },
            "valueCodeableConcept": {
               "coding": [
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  },
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  }
               ],
               "text": "Vitae perferendis odit corporis error consequatur."
            },
            "valueDatTime": "2006-06-03T15:46:06-04:00",
            "valuePeriod": {
               "end": "2001-10-30T19:00:22-04:00",
               "start": "1970-01-13T19:32:51-05:00"
            },
            "valueQuantity": {
               "code": "Ex voluptatem qui nihil quis.",
               "comparator": "\u003e",
               "system": "http://walter.biz/manuel",
               "unit": "Molestiae eos.",
               "value": 0.3374809955821533
            },
            "valueRange": {
               "high": {
                  "code": "Ex voluptatem qui nihil quis.",
                  "comparator": "\u003e",
                  "system": "http://walter.biz/manuel",
                  "unit": "Molestiae eos.",
                  "value": 0.3374809955821533
               },
               "low": {
                  "code": "Ex voluptatem qui nihil quis.",
                  "comparator": "\u003e",
                  "system": "http://walter.biz/manuel",
                  "unit": "Molestiae eos.",
                  "value": 0.3374809955821533
               }
            },
            "valueSampledData": {
               "data": "Quis eos earum.",
               "dimensions": 9133536065128410703,
               "factor": 0.2928210072171706,
               "lowerLimit": 0.5963619438089043,
               "origin": {
                  "code": "Ex voluptatem qui nihil quis.",
                  "comparator": "\u003e",
                  "system": "http://walter.biz/manuel",
                  "unit": "Molestiae eos.",
                  "value": 0.3374809955821533
               },
               "period": 0.47228831209882643,
               "upperLimit": 0.6070474669537046
            },
            "valueTime": "1980-07-14T00:10:15-04:00"
         },
         {
            "code": {
               "coding": [
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  },
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  }
               ],
               "text": "Vitae perferendis odit corporis error consequatur."
            },
            "dataAbsentReason": {
               "coding": [
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  },
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  }
               ],
               "text": "Vitae perferendis odit corporis error consequatur."
            },
            "referenceRange": {
               "age": {
                  "high": {
                     "code": "Ex voluptatem qui nihil quis.",
                     "comparator": "\u003e",
                     "system": "http://walter.biz/manuel",
                     "unit": "Molestiae eos.",
                     "value": 0.3374809955821533
                  },
                  "low": {
                     "code": "Ex voluptatem qui nihil quis.",
                     "comparator": "\u003e",
                     "system": "http://walter.biz/manuel",
                     "unit": "Molestiae eos.",
                     "value": 0.3374809955821533
                  }
               },
               "high": {
                  "code": "Ex voluptatem qui nihil quis.",
                  "comparator": "\u003e",
                  "system": "http://walter.biz/manuel",
                  "unit": "Molestiae eos.",
                  "value": 0.3374809955821533
               },
               "low": {
                  "code": "Ex voluptatem qui nihil quis.",
                  "comparator": "\u003e",
                  "system": "http://walter.biz/manuel",
                  "unit": "Molestiae eos.",
                  "value": 0.3374809955821533
               },
               "meaning": {
                  "coding": [
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     },
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     }
                  ],
                  "text": "Vitae perferendis odit corporis error consequatur."
               },
               "text": "Est suscipit modi dignissimos."
            },
            "valueAttachment": {
               "contentType": "In et sunt iure quo tempora consequuntur.",
               "creation": "2010-04-13T22:20:06-04:00",
               "data": "Iusto aperiam.",
               "hash": "Unde consequuntur esse qui.",
               "size": 7170819194107951440,
               "title": "Voluptatem id eum cupiditate."
            },
            "valueCodeableConcept": {
               "coding": [
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  },
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  }
               ],
               "text": "Vitae perferendis odit corporis error consequatur."
            },
            "valueDatTime": "2006-06-03T15:46:06-04:00",
            "valuePeriod": {
               "end": "2001-10-30T19:00:22-04:00",
               "start": "1970-01-13T19:32:51-05:00"
            },
            "valueQuantity": {
               "code": "Ex voluptatem qui nihil quis.",
               "comparator": "\u003e",
               "system": "http://walter.biz/manuel",
               "unit": "Molestiae eos.",
               "value": 0.3374809955821533
            },
            "valueRange": {
               "high": {
                  "code": "Ex voluptatem qui nihil quis.",
                  "comparator": "\u003e",
                  "system": "http://walter.biz/manuel",
                  "unit": "Molestiae eos.",
                  "value": 0.3374809955821533
               },
               "low": {
                  "code": "Ex voluptatem qui nihil quis.",
                  "comparator": "\u003e",
                  "system": "http://walter.biz/manuel",
                  "unit": "Molestiae eos.",
                  "value": 0.3374809955821533
               }
            },
            "valueSampledData": {
               "data": "Quis eos earum.",
               "dimensions": 9133536065128410703,
               "factor": 0.2928210072171706,
               "lowerLimit": 0.5963619438089043,
               "origin": {
                  "code": "Ex voluptatem qui nihil quis.",
                  "comparator": "\u003e",
                  "system": "http://walter.biz/manuel",
                  "unit": "Molestiae eos.",
                  "value": 0.3374809955821533
               },
               "period": 0.47228831209882643,
               "upperLimit": 0.6070474669537046
            },
            "valueTime": "1980-07-14T00:10:15-04:00"
         },
         {
            "code": {
               "coding": [
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  },
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  }
               ],
               "text": "Vitae perferendis odit corporis error consequatur."
            },
            "dataAbsentReason": {
               "coding": [
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  },
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  }
               ],
               "text": "Vitae perferendis odit corporis error consequatur."
            },
            "referenceRange": {
               "age": {
                  "high": {
                     "code": "Ex voluptatem qui nihil quis.",
                     "comparator": "\u003e",
                     "system": "http://walter.biz/manuel",
                     "unit": "Molestiae eos.",
                     "value": 0.3374809955821533
                  },
                  "low": {
                     "code": "Ex voluptatem qui nihil quis.",
                     "comparator": "\u003e",
                     "system": "http://walter.biz/manuel",
                     "unit": "Molestiae eos.",
                     "value": 0.3374809955821533
                  }
               },
               "high": {
                  "code": "Ex voluptatem qui nihil quis.",
                  "comparator": "\u003e",
                  "system": "http://walter.biz/manuel",
                  "unit": "Molestiae eos.",
                  "value": 0.3374809955821533
               },
               "low": {
                  "code": "Ex voluptatem qui nihil quis.",
                  "comparator": "\u003e",
                  "system": "http://walter.biz/manuel",
                  "unit": "Molestiae eos.",
                  "value": 0.3374809955821533
               },
               "meaning": {
                  "coding": [
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     },
                     {
                        "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                        "display": "Provident quam porro maiores.",
                        "system": "Quisquam voluptas ut aut explicabo et.",
                        "userSelected": true,
                        "version": "Qui eaque et veritatis dolorum voluptatibus."
                     }
                  ],
                  "text": "Vitae perferendis odit corporis error consequatur."
               },
               "text": "Est suscipit modi dignissimos."
            },
            "valueAttachment": {
               "contentType": "In et sunt iure quo tempora consequuntur.",
               "creation": "2010-04-13T22:20:06-04:00",
               "data": "Iusto aperiam.",
               "hash": "Unde consequuntur esse qui.",
               "size": 7170819194107951440,
               "title": "Voluptatem id eum cupiditate."
            },
            "valueCodeableConcept": {
               "coding": [
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  },
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  }
               ],
               "text": "Vitae perferendis odit corporis error consequatur."
            },
            "valueDatTime": "2006-06-03T15:46:06-04:00",
            "valuePeriod": {
               "end": "2001-10-30T19:00:22-04:00",
               "start": "1970-01-13T19:32:51-05:00"
            },
            "valueQuantity": {
               "code": "Ex voluptatem qui nihil quis.",
               "comparator": "\u003e",
               "system": "http://walter.biz/manuel",
               "unit": "Molestiae eos.",
               "value": 0.3374809955821533
            },
            "valueRange": {
               "high": {
                  "code": "Ex voluptatem qui nihil quis.",
                  "comparator": "\u003e",
                  "system": "http://walter.biz/manuel",
                  "unit": "Molestiae eos.",
                  "value": 0.3374809955821533
               },
               "low": {
                  "code": "Ex voluptatem qui nihil quis.",
                  "comparator": "\u003e",
                  "system": "http://walter.biz/manuel",
                  "unit": "Molestiae eos.",
                  "value": 0.3374809955821533
               }
            },
            "valueSampledData": {
               "data": "Quis eos earum.",
               "dimensions": 9133536065128410703,
               "factor": 0.2928210072171706,
               "lowerLimit": 0.5963619438089043,
               "origin": {
                  "code": "Ex voluptatem qui nihil quis.",
                  "comparator": "\u003e",
                  "system": "http://walter.biz/manuel",
                  "unit": "Molestiae eos.",
                  "value": 0.3374809955821533
               },
               "period": 0.47228831209882643,
               "upperLimit": 0.6070474669537046
            },
            "valueTime": "1980-07-14T00:10:15-04:00"
         }
      ],
      "dateAbsentReason": {
         "coding": [
            {
               "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
               "display": "Provident quam porro maiores.",
               "system": "Quisquam voluptas ut aut explicabo et.",
               "userSelected": true,
               "version": "Qui eaque et veritatis dolorum voluptatibus."
            },
            {
               "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
               "display": "Provident quam porro maiores.",
               "system": "Quisquam voluptas ut aut explicabo et.",
               "userSelected": true,
               "version": "Qui eaque et veritatis dolorum voluptatibus."
            }
         ],
         "text": "Vitae perferendis odit corporis error consequatur."
      },
      "device": {
         "display": "Dolores aut reprehenderit adipisci.",
         "reference": "Voluptate et."
      },
      "effectiveDateTime": "1989-07-31T17:52:24-04:00",
      "effectivePeriod": {
         "end": "2001-10-30T19:00:22-04:00",
         "start": "1970-01-13T19:32:51-05:00"
      },
      "encounter": {
         "display": "Dolores aut reprehenderit adipisci.",
         "reference": "Voluptate et."
      },
      "identifier": [
         {
            "CodeableConcept": {
               "coding": [
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  },
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  }
               ],
               "text": "Vitae perferendis odit corporis error consequatur."
            },
            "assigner": {
               "display": "Dolores aut reprehenderit adipisci.",
               "reference": "Voluptate et."
            },
            "period": {
               "end": "2001-10-30T19:00:22-04:00",
               "start": "1970-01-13T19:32:51-05:00"
            }
         },
         {
            "CodeableConcept": {
               "coding": [
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  },
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  }
               ],
               "text": "Vitae perferendis odit corporis error consequatur."
            },
            "assigner": {
               "display": "Dolores aut reprehenderit adipisci.",
               "reference": "Voluptate et."
            },
            "period": {
               "end": "2001-10-30T19:00:22-04:00",
               "start": "1970-01-13T19:32:51-05:00"
            }
         }
      ],
      "interpretation": {
         "coding": [
            {
               "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
               "display": "Provident quam porro maiores.",
               "system": "Quisquam voluptas ut aut explicabo et.",
               "userSelected": true,
               "version": "Qui eaque et veritatis dolorum voluptatibus."
            },
            {
               "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
               "display": "Provident quam porro maiores.",
               "system": "Quisquam voluptas ut aut explicabo et.",
               "userSelected": true,
               "version": "Qui eaque et veritatis dolorum voluptatibus."
            }
         ],
         "text": "Vitae perferendis odit corporis error consequatur."
      },
      "issued": "1980-02-17T14:56:04-05:00",
      "method": {
         "coding": [
            {
               "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
               "display": "Provident quam porro maiores.",
               "system": "Quisquam voluptas ut aut explicabo et.",
               "userSelected": true,
               "version": "Qui eaque et veritatis dolorum voluptatibus."
            },
            {
               "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
               "display": "Provident quam porro maiores.",
               "system": "Quisquam voluptas ut aut explicabo et.",
               "userSelected": true,
               "version": "Qui eaque et veritatis dolorum voluptatibus."
            }
         ],
         "text": "Vitae perferendis odit corporis error consequatur."
      },
      "performer": [
         {
            "display": "Dolores aut reprehenderit adipisci.",
            "reference": "Voluptate et."
         },
         {
            "display": "Dolores aut reprehenderit adipisci.",
            "reference": "Voluptate et."
         },
         {
            "display": "Dolores aut reprehenderit adipisci.",
            "reference": "Voluptate et."
         }
      ],
      "referenceRange": [
         {
            "age": {
               "high": {
                  "code": "Ex voluptatem qui nihil quis.",
                  "comparator": "\u003e",
                  "system": "http://walter.biz/manuel",
                  "unit": "Molestiae eos.",
                  "value": 0.3374809955821533
               },
               "low": {
                  "code": "Ex voluptatem qui nihil quis.",
                  "comparator": "\u003e",
                  "system": "http://walter.biz/manuel",
                  "unit": "Molestiae eos.",
                  "value": 0.3374809955821533
               }
            },
            "high": {
               "code": "Ex voluptatem qui nihil quis.",
               "comparator": "\u003e",
               "system": "http://walter.biz/manuel",
               "unit": "Molestiae eos.",
               "value": 0.3374809955821533
            },
            "low": {
               "code": "Ex voluptatem qui nihil quis.",
               "comparator": "\u003e",
               "system": "http://walter.biz/manuel",
               "unit": "Molestiae eos.",
               "value": 0.3374809955821533
            },
            "meaning": {
               "coding": [
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  },
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  }
               ],
               "text": "Vitae perferendis odit corporis error consequatur."
            },
            "text": "Est suscipit modi dignissimos."
         },
         {
            "age": {
               "high": {
                  "code": "Ex voluptatem qui nihil quis.",
                  "comparator": "\u003e",
                  "system": "http://walter.biz/manuel",
                  "unit": "Molestiae eos.",
                  "value": 0.3374809955821533
               },
               "low": {
                  "code": "Ex voluptatem qui nihil quis.",
                  "comparator": "\u003e",
                  "system": "http://walter.biz/manuel",
                  "unit": "Molestiae eos.",
                  "value": 0.3374809955821533
               }
            },
            "high": {
               "code": "Ex voluptatem qui nihil quis.",
               "comparator": "\u003e",
               "system": "http://walter.biz/manuel",
               "unit": "Molestiae eos.",
               "value": 0.3374809955821533
            },
            "low": {
               "code": "Ex voluptatem qui nihil quis.",
               "comparator": "\u003e",
               "system": "http://walter.biz/manuel",
               "unit": "Molestiae eos.",
               "value": 0.3374809955821533
            },
            "meaning": {
               "coding": [
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  },
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  }
               ],
               "text": "Vitae perferendis odit corporis error consequatur."
            },
            "text": "Est suscipit modi dignissimos."
         },
         {
            "age": {
               "high": {
                  "code": "Ex voluptatem qui nihil quis.",
                  "comparator": "\u003e",
                  "system": "http://walter.biz/manuel",
                  "unit": "Molestiae eos.",
                  "value": 0.3374809955821533
               },
               "low": {
                  "code": "Ex voluptatem qui nihil quis.",
                  "comparator": "\u003e",
                  "system": "http://walter.biz/manuel",
                  "unit": "Molestiae eos.",
                  "value": 0.3374809955821533
               }
            },
            "high": {
               "code": "Ex voluptatem qui nihil quis.",
               "comparator": "\u003e",
               "system": "http://walter.biz/manuel",
               "unit": "Molestiae eos.",
               "value": 0.3374809955821533
            },
            "low": {
               "code": "Ex voluptatem qui nihil quis.",
               "comparator": "\u003e",
               "system": "http://walter.biz/manuel",
               "unit": "Molestiae eos.",
               "value": 0.3374809955821533
            },
            "meaning": {
               "coding": [
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  },
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  }
               ],
               "text": "Vitae perferendis odit corporis error consequatur."
            },
            "text": "Est suscipit modi dignissimos."
         }
      ],
      "related": [
         {
            "target": {
               "coding": [
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  },
                  {
                     "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
                     "display": "Provident quam porro maiores.",
                     "system": "Quisquam voluptas ut aut explicabo et.",
                     "userSelected": true,
                     "version": "Qui eaque et veritatis dolorum voluptatibus."
                  }
               ],
               "text": "Vitae perferendis odit corporis error consequatur."
            },
            "type": "Consequatur itaque."
         }
      ],
      "specimen": {
         "display": "Dolores aut reprehenderit adipisci.",
         "reference": "Voluptate et."
      },
      "status": "preliminary",
      "valueAttachment": {
         "contentType": "In et sunt iure quo tempora consequuntur.",
         "creation": "2010-04-13T22:20:06-04:00",
         "data": "Iusto aperiam.",
         "hash": "Unde consequuntur esse qui.",
         "size": 7170819194107951440,
         "title": "Voluptatem id eum cupiditate."
      },
      "valueCodeableConcept": {
         "coding": [
            {
               "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
               "display": "Provident quam porro maiores.",
               "system": "Quisquam voluptas ut aut explicabo et.",
               "userSelected": true,
               "version": "Qui eaque et veritatis dolorum voluptatibus."
            },
            {
               "code": "At eligendi ullam officiis reprehenderit perspiciatis.",
               "display": "Provident quam porro maiores.",
               "system": "Quisquam voluptas ut aut explicabo et.",
               "userSelected": true,
               "version": "Qui eaque et veritatis dolorum voluptatibus."
            }
         ],
         "text": "Vitae perferendis odit corporis error consequatur."
      },
      "valueDatTime": "1976-02-02T04:43:55-05:00",
      "valuePeriod": {
         "end": "2001-10-30T19:00:22-04:00",
         "start": "1970-01-13T19:32:51-05:00"
      },
      "valueQuantity": {
         "code": "Ex voluptatem qui nihil quis.",
         "comparator": "\u003e",
         "system": "http://walter.biz/manuel",
         "unit": "Molestiae eos.",
         "value": 0.3374809955821533
      },
      "valueRange": {
         "high": {
            "code": "Ex voluptatem qui nihil quis.",
            "comparator": "\u003e",
            "system": "http://walter.biz/manuel",
            "unit": "Molestiae eos.",
            "value": 0.3374809955821533
         },
         "low": {
            "code": "Ex voluptatem qui nihil quis.",
            "comparator": "\u003e",
            "system": "http://walter.biz/manuel",
            "unit": "Molestiae eos.",
            "value": 0.3374809955821533
         }
      },
      "valueSampledData": {
         "data": "Quis eos earum.",
         "dimensions": 9133536065128410703,
         "factor": 0.2928210072171706,
         "lowerLimit": 0.5963619438089043,
         "origin": {
            "code": "Ex voluptatem qui nihil quis.",
            "comparator": "\u003e",
            "system": "http://walter.biz/manuel",
            "unit": "Molestiae eos.",
            "value": 0.3374809955821533
         },
         "period": 0.47228831209882643,
         "upperLimit": 0.6070474669537046
      },
      "valueString": "Nihil fugit est.",
      "valueTime": "2014-09-30T03:57:06-04:00"
   }
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp35.Run(c, args) },
	}
	tmp35.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp35.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp36 := new(UpdatePatientCommand)
	sub = &cobra.Command{
		Use:   `patient ["/nosh/patients/PATIENTID"]`,
		Short: `This resource uses JWT to secure its endpoints`,
		Long: `This resource uses JWT to secure its endpoints

Payload example:

{
   "name": "Excepturi sit cumque sunt."
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp36.Run(c, args) },
	}
	tmp36.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp36.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp37 := new(UpdateUserCommand)
	sub = &cobra.Command{
		Use:   `user ["/nosh/users/USERID"]`,
		Short: `This resource uses JWT to secure its endpoints`,
		Long: `This resource uses JWT to secure its endpoints

Payload example:

{
   "email": "jim.smith@gmail.com",
   "password": "password"
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp37.Run(c, args) },
	}
	tmp37.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp37.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "watch",
		Short: `watch action`,
	}
	tmp38 := new(WatchAllergyIntoleranceCommand)
	sub = &cobra.Command{
		Use:   `AllergyIntolerance ["/nosh/patients/PATIENTID/allergy_intolerance/ALLERGY_INTOLERANCEID/watch"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp38.Run(c, args) },
	}
	tmp38.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp38.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp39 := new(WatchNutritionRequestCommand)
	sub = &cobra.Command{
		Use:   `NutritionRequest ["/nosh/patients/PATIENTID/nutrition_requests/NUTRITION_REQUESTID/watch"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp39.Run(c, args) },
	}
	tmp39.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp39.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp40 := new(WatchObservationCommand)
	sub = &cobra.Command{
		Use:   `Observation ["/nosh/patients/PATIENTID/observation/OBSERVATIONID/watch"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp40.Run(c, args) },
	}
	tmp40.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp40.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
}

func intFlagVal(name string, parsed int) *int {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func float64FlagVal(name string, parsed float64) *float64 {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func boolFlagVal(name string, parsed bool) *bool {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func stringFlagVal(name string, parsed string) *string {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func hasFlag(name string) bool {
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "--"+name) {
			return true
		}
	}
	return false
}

func jsonVal(val string) (*interface{}, error) {
	var t interface{}
	err := json.Unmarshal([]byte(val), &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func jsonArray(ins []string) ([]interface{}, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []interface{}
	for _, id := range ins {
		val, err := jsonVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, val)
	}
	return vals, nil
}

func timeVal(val string) (*time.Time, error) {
	t, err := time.Parse(time.RFC3339, val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func timeArray(ins []string) ([]time.Time, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []time.Time
	for _, id := range ins {
		val, err := timeVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func uuidVal(val string) (*uuid.UUID, error) {
	t, err := uuid.FromString(val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func uuidArray(ins []string) ([]uuid.UUID, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []uuid.UUID
	for _, id := range ins {
		val, err := uuidVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func float64Val(val string) (*float64, error) {
	t, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func float64Array(ins []string) ([]float64, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []float64
	for _, id := range ins {
		val, err := float64Val(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func boolVal(val string) (*bool, error) {
	t, err := strconv.ParseBool(val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func boolArray(ins []string) ([]bool, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []bool
	for _, id := range ins {
		val, err := boolVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

// Run makes the HTTP request corresponding to the CreateAllergyIntoleranceCommand command.
func (cmd *CreateAllergyIntoleranceCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/nosh/patients/%v/allergy_intolerance", cmd.PatientID)
	}
	var payload client.CreateAllergyIntolerancePayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CreateAllergyIntolerance(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateAllergyIntoleranceCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
	var patientID int
	cc.Flags().IntVar(&cmd.PatientID, "patientID", patientID, `Patient ID`)
}

// Run makes the HTTP request corresponding to the DeleteAllergyIntoleranceCommand command.
func (cmd *DeleteAllergyIntoleranceCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/nosh/patients/%v/allergy_intolerance/%v", cmd.PatientID, cmd.AllergyIntoleranceID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.DeleteAllergyIntolerance(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *DeleteAllergyIntoleranceCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var allergyIntoleranceID int
	cc.Flags().IntVar(&cmd.AllergyIntoleranceID, "allergy_intoleranceID", allergyIntoleranceID, ``)
	var patientID int
	cc.Flags().IntVar(&cmd.PatientID, "patientID", patientID, `Patient ID`)
}

// Run makes the HTTP request corresponding to the ListAllergyIntoleranceCommand command.
func (cmd *ListAllergyIntoleranceCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/nosh/patients/%v/allergy_intolerance", cmd.PatientID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ListAllergyIntolerance(ctx, path, cmd.Years)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListAllergyIntoleranceCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var patientID int
	cc.Flags().IntVar(&cmd.PatientID, "patientID", patientID, `Patient ID`)
	var years []int
	cc.Flags().IntSliceVar(&cmd.Years, "years", years, `Filter by years`)
}

// Run makes the HTTP request corresponding to the RateAllergyIntoleranceCommand command.
func (cmd *RateAllergyIntoleranceCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/nosh/patients/%v/allergy_intolerance/%v/actions/rate", cmd.PatientID, cmd.AllergyIntoleranceID)
	}
	var payload client.RateAllergyIntolerancePayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.RateAllergyIntolerance(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *RateAllergyIntoleranceCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
	var allergyIntoleranceID int
	cc.Flags().IntVar(&cmd.AllergyIntoleranceID, "allergy_intoleranceID", allergyIntoleranceID, ``)
	var patientID int
	cc.Flags().IntVar(&cmd.PatientID, "patientID", patientID, `Patient ID`)
}

// Run makes the HTTP request corresponding to the ShowAllergyIntoleranceCommand command.
func (cmd *ShowAllergyIntoleranceCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/nosh/patients/%v/allergy_intolerance/%v", cmd.PatientID, cmd.AllergyIntoleranceID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ShowAllergyIntolerance(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowAllergyIntoleranceCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var allergyIntoleranceID int
	cc.Flags().IntVar(&cmd.AllergyIntoleranceID, "allergy_intoleranceID", allergyIntoleranceID, ``)
	var patientID int
	cc.Flags().IntVar(&cmd.PatientID, "patientID", patientID, `Patient ID`)
}

// Run makes the HTTP request corresponding to the UpdateAllergyIntoleranceCommand command.
func (cmd *UpdateAllergyIntoleranceCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/nosh/patients/%v/allergy_intolerance/%v", cmd.PatientID, cmd.AllergyIntoleranceID)
	}
	var payload client.AllergyIntolerancePayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.UpdateAllergyIntolerance(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UpdateAllergyIntoleranceCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
	var allergyIntoleranceID int
	cc.Flags().IntVar(&cmd.AllergyIntoleranceID, "allergy_intoleranceID", allergyIntoleranceID, ``)
	var patientID int
	cc.Flags().IntVar(&cmd.PatientID, "patientID", patientID, `Patient ID`)
}

// Run establishes a websocket connection for the WatchAllergyIntoleranceCommand command.
func (cmd *WatchAllergyIntoleranceCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/nosh/patients/%v/allergy_intolerance/%v/watch", cmd.PatientID, cmd.AllergyIntoleranceID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	ws, err := c.WatchAllergyIntolerance(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}
	go goaclient.WSWrite(ws)
	goaclient.WSRead(ws)

	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *WatchAllergyIntoleranceCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var allergyIntoleranceID int
	cc.Flags().IntVar(&cmd.AllergyIntoleranceID, "allergy_intoleranceID", allergyIntoleranceID, ``)
	var patientID int
	cc.Flags().IntVar(&cmd.PatientID, "patientID", patientID, `Patient ID`)
}

// Run makes the HTTP request corresponding to the CreateNutritionRequestCommand command.
func (cmd *CreateNutritionRequestCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/nosh/patients/%v/nutrition_requests", cmd.PatientID)
	}
	var payload client.CreateNutritionRequestPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CreateNutritionRequest(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateNutritionRequestCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
	var patientID int
	cc.Flags().IntVar(&cmd.PatientID, "patientID", patientID, `Patient ID`)
}

// Run makes the HTTP request corresponding to the DeleteNutritionRequestCommand command.
func (cmd *DeleteNutritionRequestCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/nosh/patients/%v/nutrition_requests/%v", cmd.PatientID, cmd.NutritionRequestID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.DeleteNutritionRequest(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *DeleteNutritionRequestCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var nutritionRequestID int
	cc.Flags().IntVar(&cmd.NutritionRequestID, "nutrition_requestID", nutritionRequestID, ``)
	var patientID int
	cc.Flags().IntVar(&cmd.PatientID, "patientID", patientID, `Patient ID`)
}

// Run makes the HTTP request corresponding to the ListNutritionRequestCommand command.
func (cmd *ListNutritionRequestCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/nosh/patients/%v/nutrition_requests", cmd.PatientID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ListNutritionRequest(ctx, path, cmd.Years)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListNutritionRequestCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var patientID int
	cc.Flags().IntVar(&cmd.PatientID, "patientID", patientID, `Patient ID`)
	var years []int
	cc.Flags().IntSliceVar(&cmd.Years, "years", years, `Filter by years`)
}

// Run makes the HTTP request corresponding to the RateNutritionRequestCommand command.
func (cmd *RateNutritionRequestCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/nosh/patients/%v/nutrition_requests/%v/actions/rate", cmd.PatientID, cmd.NutritionRequestID)
	}
	var payload client.RateNutritionRequestPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.RateNutritionRequest(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *RateNutritionRequestCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
	var nutritionRequestID int
	cc.Flags().IntVar(&cmd.NutritionRequestID, "nutrition_requestID", nutritionRequestID, ``)
	var patientID int
	cc.Flags().IntVar(&cmd.PatientID, "patientID", patientID, `Patient ID`)
}

// Run makes the HTTP request corresponding to the ShowNutritionRequestCommand command.
func (cmd *ShowNutritionRequestCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/nosh/patients/%v/nutrition_requests/%v", cmd.PatientID, cmd.NutritionRequestID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ShowNutritionRequest(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowNutritionRequestCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var nutritionRequestID int
	cc.Flags().IntVar(&cmd.NutritionRequestID, "nutrition_requestID", nutritionRequestID, ``)
	var patientID int
	cc.Flags().IntVar(&cmd.PatientID, "patientID", patientID, `Patient ID`)
}

// Run makes the HTTP request corresponding to the UpdateNutritionRequestCommand command.
func (cmd *UpdateNutritionRequestCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/nosh/patients/%v/nutrition_requests/%v", cmd.PatientID, cmd.NutritionRequestID)
	}
	var payload client.NutritionRequestPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.UpdateNutritionRequest(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UpdateNutritionRequestCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
	var nutritionRequestID int
	cc.Flags().IntVar(&cmd.NutritionRequestID, "nutrition_requestID", nutritionRequestID, ``)
	var patientID int
	cc.Flags().IntVar(&cmd.PatientID, "patientID", patientID, `Patient ID`)
}

// Run establishes a websocket connection for the WatchNutritionRequestCommand command.
func (cmd *WatchNutritionRequestCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/nosh/patients/%v/nutrition_requests/%v/watch", cmd.PatientID, cmd.NutritionRequestID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	ws, err := c.WatchNutritionRequest(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}
	go goaclient.WSWrite(ws)
	goaclient.WSRead(ws)

	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *WatchNutritionRequestCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var nutritionRequestID int
	cc.Flags().IntVar(&cmd.NutritionRequestID, "nutrition_requestID", nutritionRequestID, ``)
	var patientID int
	cc.Flags().IntVar(&cmd.PatientID, "patientID", patientID, `Patient ID`)
}

// Run makes the HTTP request corresponding to the CreateObservationCommand command.
func (cmd *CreateObservationCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/nosh/patients/%v/observation", cmd.PatientID)
	}
	var payload client.CreateObservationPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CreateObservation(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateObservationCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
	var patientID int
	cc.Flags().IntVar(&cmd.PatientID, "patientID", patientID, `Patient ID`)
}

// Run makes the HTTP request corresponding to the DeleteObservationCommand command.
func (cmd *DeleteObservationCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/nosh/patients/%v/observation/%v", cmd.PatientID, cmd.ObservationID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.DeleteObservation(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *DeleteObservationCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var observationID int
	cc.Flags().IntVar(&cmd.ObservationID, "observationID", observationID, ``)
	var patientID int
	cc.Flags().IntVar(&cmd.PatientID, "patientID", patientID, `Patient ID`)
}

// Run makes the HTTP request corresponding to the ListObservationCommand command.
func (cmd *ListObservationCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/nosh/patients/%v/observation", cmd.PatientID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ListObservation(ctx, path, cmd.Years)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListObservationCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var patientID int
	cc.Flags().IntVar(&cmd.PatientID, "patientID", patientID, `Patient ID`)
	var years []int
	cc.Flags().IntSliceVar(&cmd.Years, "years", years, `Filter by years`)
}

// Run makes the HTTP request corresponding to the RateObservationCommand command.
func (cmd *RateObservationCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/nosh/patients/%v/observation/%v/actions/rate", cmd.PatientID, cmd.ObservationID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.RateObservation(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *RateObservationCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var observationID int
	cc.Flags().IntVar(&cmd.ObservationID, "observationID", observationID, ``)
	var patientID int
	cc.Flags().IntVar(&cmd.PatientID, "patientID", patientID, `Patient ID`)
}

// Run makes the HTTP request corresponding to the ShowObservationCommand command.
func (cmd *ShowObservationCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/nosh/patients/%v/observation/%v", cmd.PatientID, cmd.ObservationID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ShowObservation(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowObservationCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var observationID int
	cc.Flags().IntVar(&cmd.ObservationID, "observationID", observationID, ``)
	var patientID int
	cc.Flags().IntVar(&cmd.PatientID, "patientID", patientID, `Patient ID`)
}

// Run makes the HTTP request corresponding to the UpdateObservationCommand command.
func (cmd *UpdateObservationCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/nosh/patients/%v/observation/%v", cmd.PatientID, cmd.ObservationID)
	}
	var payload client.ObservationPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.UpdateObservation(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UpdateObservationCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
	var observationID int
	cc.Flags().IntVar(&cmd.ObservationID, "observationID", observationID, ``)
	var patientID int
	cc.Flags().IntVar(&cmd.PatientID, "patientID", patientID, `Patient ID`)
}

// Run establishes a websocket connection for the WatchObservationCommand command.
func (cmd *WatchObservationCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/nosh/patients/%v/observation/%v/watch", cmd.PatientID, cmd.ObservationID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	ws, err := c.WatchObservation(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}
	go goaclient.WSWrite(ws)
	goaclient.WSRead(ws)

	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *WatchObservationCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var observationID int
	cc.Flags().IntVar(&cmd.ObservationID, "observationID", observationID, ``)
	var patientID int
	cc.Flags().IntVar(&cmd.PatientID, "patientID", patientID, `Patient ID`)
}

// Run makes the HTTP request corresponding to the SecureBasicCommand command.
func (cmd *SecureBasicCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/nosh/basic"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.SecureBasic(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *SecureBasicCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the UnsecureBasicCommand command.
func (cmd *UnsecureBasicCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/nosh/basic/unsecure"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.UnsecureBasic(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UnsecureBasicCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the HealthHealthCommand command.
func (cmd *HealthHealthCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/nosh/_ah/health"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.HealthHealth(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *HealthHealthCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the SecureJWTCommand command.
func (cmd *SecureJWTCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/nosh/jwt"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	var tmp41 *bool
	if cmd.Fail != "" {
		var err error
		tmp41, err = boolVal(cmd.Fail)
		if err != nil {
			goa.LogError(ctx, "failed to parse flag into *bool value", "flag", "--fail", "err", err)
			return err
		}
	}
	resp, err := c.SecureJWT(ctx, path, tmp41)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *SecureJWTCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var fail string
	cc.Flags().StringVar(&cmd.Fail, "fail", fail, `Force auth failure via JWT validation middleware`)
}

// Run makes the HTTP request corresponding to the SigninJWTCommand command.
func (cmd *SigninJWTCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/nosh/jwt/signin"
	}
	var payload client.SigninJWTPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.SigninJWT(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *SigninJWTCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
}

// Run makes the HTTP request corresponding to the SignupJWTCommand command.
func (cmd *SignupJWTCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/nosh/jwt/signup"
	}
	var payload client.SignupJWTPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.SignupJWT(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *SignupJWTCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
}

// Run makes the HTTP request corresponding to the UnsecureJWTCommand command.
func (cmd *UnsecureJWTCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/nosh/jwt/unsecure"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.UnsecureJWT(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UnsecureJWTCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the CreatePatientCommand command.
func (cmd *CreatePatientCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/nosh/patients"
	}
	var payload client.CreatePatientPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CreatePatient(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreatePatientCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
}

// Run makes the HTTP request corresponding to the DeletePatientCommand command.
func (cmd *DeletePatientCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/nosh/patients/%v", cmd.PatientID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.DeletePatient(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *DeletePatientCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var patientID int
	cc.Flags().IntVar(&cmd.PatientID, "patientID", patientID, `Patient ID`)
}

// Run makes the HTTP request corresponding to the ListPatientCommand command.
func (cmd *ListPatientCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/nosh/patients/jwt"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	var tmp42 *bool
	if cmd.Fail != "" {
		var err error
		tmp42, err = boolVal(cmd.Fail)
		if err != nil {
			goa.LogError(ctx, "failed to parse flag into *bool value", "flag", "--fail", "err", err)
			return err
		}
	}
	resp, err := c.ListPatient(ctx, path, tmp42)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListPatientCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var fail string
	cc.Flags().StringVar(&cmd.Fail, "fail", fail, `Force auth failure via JWT validation middleware`)
}

// Run makes the HTTP request corresponding to the ShowPatientCommand command.
func (cmd *ShowPatientCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/nosh/patients/%v", cmd.PatientID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ShowPatient(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowPatientCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var patientID int
	cc.Flags().IntVar(&cmd.PatientID, "patientID", patientID, `Patient ID`)
}

// Run makes the HTTP request corresponding to the UpdatePatientCommand command.
func (cmd *UpdatePatientCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/nosh/patients/%v", cmd.PatientID)
	}
	var payload client.UpdatePatientPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.UpdatePatient(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UpdatePatientCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
	var patientID int
	cc.Flags().IntVar(&cmd.PatientID, "patientID", patientID, `Patient ID`)
}

// Run makes the HTTP request corresponding to the CreateUserCommand command.
func (cmd *CreateUserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/nosh/users"
	}
	var payload client.CreateUserPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CreateUser(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateUserCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
}

// Run makes the HTTP request corresponding to the DeleteUserCommand command.
func (cmd *DeleteUserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/nosh/users/%v", cmd.UserID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.DeleteUser(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *DeleteUserCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var userID int
	cc.Flags().IntVar(&cmd.UserID, "userID", userID, `User ID`)
}

// Run makes the HTTP request corresponding to the ListUserCommand command.
func (cmd *ListUserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/nosh/users"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ListUser(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListUserCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the ShowUserCommand command.
func (cmd *ShowUserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/nosh/users/%v", cmd.UserID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ShowUser(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowUserCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var userID int
	cc.Flags().IntVar(&cmd.UserID, "userID", userID, `User ID`)
}

// Run makes the HTTP request corresponding to the SigninUserCommand command.
func (cmd *SigninUserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/nosh/users/jwt/signin"
	}
	var payload client.SigninUserPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.SigninUser(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *SigninUserCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
}

// Run makes the HTTP request corresponding to the SignupUserCommand command.
func (cmd *SignupUserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/nosh/users/jwt/signup"
	}
	var payload client.SignupUserPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.SignupUser(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *SignupUserCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
}

// Run makes the HTTP request corresponding to the UpdateUserCommand command.
func (cmd *UpdateUserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/nosh/users/%v", cmd.UserID)
	}
	var payload client.UpdateUserPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.UpdateUser(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UpdateUserCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
	var userID int
	cc.Flags().IntVar(&cmd.UserID, "userID", userID, `User ID`)
}
