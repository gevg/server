// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/jamesallain/goa-fhir/design
// --out=$(GOPATH)\src\github.com\jamesallain\nosh\server
// --version=v1.1.0-dirty
//
// API "Secure": Model Helpers
//
// The content of this file is auto-generated, DO NOT MODIFY

package models

import (
	"github.com/goadesign/goa"
	"github.com/jamesallain/goa-fhir/app"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
	"time"
)

// MediaType Retrieval Functions

// ListVital returns an array of view: default.
func (m *VitalDB) ListVital(ctx context.Context, patientID int) []*app.Vital {
	defer goa.MeasureSince([]string{"goa", "db", "vital", "listvital"}, time.Now())

	var native []*Vital
	var objs []*app.Vital
	err := m.Db.Scopes(VitalFilterByPatient(patientID, m.Db)).Table(m.TableName()).Preload("Patient").Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Vital", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.VitalToVital())
	}

	return objs
}

// VitalToVital loads a Vital and builds the default view of media type Vital.
func (m *Vital) VitalToVital() *app.Vital {
	vital := &app.Vital{}
	tmp1 := m.Patient.PatientToPatientLink()
	vital.Links = &app.VitalLinks{Patient: tmp1}
	vital.ID = m.ID
	vital.Name = m.Name
	tmp2 := &m.Patient
	vital.Patient = tmp2.PatientToPatientTiny() // %!s(MISSING)
	vital.Varietal = m.Varietal
	vital.Vineyard = m.Vineyard
	vital.Vintage = m.Vintage

	return vital
}

// OneVital loads a Vital and builds the default view of media type Vital.
func (m *VitalDB) OneVital(ctx context.Context, id int, patientID int) (*app.Vital, error) {
	defer goa.MeasureSince([]string{"goa", "db", "vital", "onevital"}, time.Now())

	var native Vital
	err := m.Db.Scopes(VitalFilterByPatient(patientID, m.Db)).Table(m.TableName()).Preload("Patient").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Vital", "error", err.Error())
		return nil, err
	}

	view := *native.VitalToVital()
	return &view, err
}

// MediaType Retrieval Functions

// ListVitalFull returns an array of view: full.
func (m *VitalDB) ListVitalFull(ctx context.Context, patientID int) []*app.VitalFull {
	defer goa.MeasureSince([]string{"goa", "db", "vital", "listvitalfull"}, time.Now())

	var native []*Vital
	var objs []*app.VitalFull
	err := m.Db.Scopes(VitalFilterByPatient(patientID, m.Db)).Table(m.TableName()).Preload("Patient").Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Vital", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.VitalToVitalFull())
	}

	return objs
}

// VitalToVitalFull loads a Vital and builds the full view of media type Vital.
func (m *Vital) VitalToVitalFull() *app.VitalFull {
	vital := &app.VitalFull{}
	tmp1 := m.Patient.PatientToPatientLink()
	vital.Links = &app.VitalLinks{Patient: tmp1}
	vital.Color = m.Color
	vital.Country = m.Country
	vital.CreatedAt = m.CreatedAt
	vital.ID = m.ID
	vital.Name = m.Name
	tmp2 := &m.Patient
	vital.Patient = tmp2.PatientToPatient() // %!s(MISSING)
	vital.Region = m.Region
	vital.Review = m.Review
	vital.Sweetness = m.Sweetness
	vital.UpdatedAt = &m.UpdatedAt
	vital.Varietal = m.Varietal
	vital.Vineyard = m.Vineyard
	vital.Vintage = m.Vintage

	return vital
}

// OneVitalFull loads a Vital and builds the full view of media type Vital.
func (m *VitalDB) OneVitalFull(ctx context.Context, id int, patientID int) (*app.VitalFull, error) {
	defer goa.MeasureSince([]string{"goa", "db", "vital", "onevitalfull"}, time.Now())

	var native Vital
	err := m.Db.Scopes(VitalFilterByPatient(patientID, m.Db)).Table(m.TableName()).Preload("Patient").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Vital", "error", err.Error())
		return nil, err
	}

	view := *native.VitalToVitalFull()
	return &view, err
}

// MediaType Retrieval Functions

// ListVitalTiny returns an array of view: tiny.
func (m *VitalDB) ListVitalTiny(ctx context.Context, patientID int) []*app.VitalTiny {
	defer goa.MeasureSince([]string{"goa", "db", "vital", "listvitaltiny"}, time.Now())

	var native []*Vital
	var objs []*app.VitalTiny
	err := m.Db.Scopes(VitalFilterByPatient(patientID, m.Db)).Table(m.TableName()).Preload("Patient").Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Vital", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.VitalToVitalTiny())
	}

	return objs
}

// VitalToVitalTiny loads a Vital and builds the tiny view of media type Vital.
func (m *Vital) VitalToVitalTiny() *app.VitalTiny {
	vital := &app.VitalTiny{}
	tmp1 := m.Patient.PatientToPatientLink()
	vital.Links = &app.VitalLinks{Patient: tmp1}
	vital.ID = m.ID
	vital.Name = m.Name

	return vital
}

// OneVitalTiny loads a Vital and builds the tiny view of media type Vital.
func (m *VitalDB) OneVitalTiny(ctx context.Context, id int, patientID int) (*app.VitalTiny, error) {
	defer goa.MeasureSince([]string{"goa", "db", "vital", "onevitaltiny"}, time.Now())

	var native Vital
	err := m.Db.Scopes(VitalFilterByPatient(patientID, m.Db)).Table(m.TableName()).Preload("Patient").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Vital", "error", err.Error())
		return nil, err
	}

	view := *native.VitalToVitalTiny()
	return &view, err
}
