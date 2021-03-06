// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/goa-fhir/server/design
// --out=$(GOPATH)\src\github.com\goa-fhir\server
// --version=v1.1.0-dirty
//
// API "goa-FHIR": Application Resource Href Factories
//
// The content of this file is auto-generated, DO NOT MODIFY

package app

import (
	"fmt"
	"strings"
)

// ObservationHref returns the resource href.
func ObservationHref(patientID, observationID interface{}) string {
	parampatientID := strings.TrimLeftFunc(fmt.Sprintf("%v", patientID), func(r rune) bool { return r == '/' })
	paramobservationID := strings.TrimLeftFunc(fmt.Sprintf("%v", observationID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/nosh/patients/%v/observation/%v", parampatientID, paramobservationID)
}

// PatientHref returns the resource href.
func PatientHref(patientID interface{}) string {
	parampatientID := strings.TrimLeftFunc(fmt.Sprintf("%v", patientID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/nosh/patients/%v", parampatientID)
}

// UserHref returns the resource href.
func UserHref(userID interface{}) string {
	paramuserID := strings.TrimLeftFunc(fmt.Sprintf("%v", userID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/nosh/users/%v", paramuserID)
}
