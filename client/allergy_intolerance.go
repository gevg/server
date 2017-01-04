package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/net/websocket"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// CreateAllergyIntolerancePayload is the AllergyIntolerance create action payload.
type CreateAllergyIntolerancePayload struct {
	// Category of the identified Substance. See http://hl7.org/fhir/ValueSet/allergy-intolerance-category
	Category *string `form:"category,omitempty" json:"category,omitempty" xml:"category,omitempty"`
	// Estimate of the potential clinical harm, or seriousness, of the reaction to the identified Substance. See http://hl7.org/fhir/ValueSet/allergy-intolerance-criticality
	Criticality *string `form:"criticality,omitempty" json:"criticality,omitempty" xml:"criticality,omitempty"`
	// This records identifiers associated with this allergy/intolerance concern that are defined by business processes and/or
	// 		used to refer to it when a direct URL reference to the resource itself is not appropriate (e.g. in CDA documents, or in written / printed documentation).
	Identifier []*Identifier `form:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// Represents the date and/or time of the last known occurrence of a reaction event.
	LastOccurence *time.Time `form:"lastOccurence,omitempty" json:"lastOccurence,omitempty" xml:"lastOccurence,omitempty"`
	// Additional narrative about the propensity for the Adverse Reaction, not captured in other fields..
	Note *Annotation `form:"note,omitempty" json:"note,omitempty" xml:"note,omitempty"`
	// Record of the date and/or time of the onset of the Allergy or Intolerance.
	Onset *time.Time `form:"onset,omitempty" json:"onset,omitempty" xml:"onset,omitempty"`
	// The patient who has the allergy or intolerance.
	Patient *HL7Reference `form:"patient,omitempty" json:"patient,omitempty" xml:"patient,omitempty"`
	// Details about each adverse reaction event linked to exposure to the identified Substance.
	Reaction *Reaction `form:"reaction,omitempty" json:"reaction,omitempty" xml:"reaction,omitempty"`
	// Date when the sensitivity was recorded.
	RecordedDate *time.Time `form:"recordedDate,omitempty" json:"recordedDate,omitempty" xml:"recordedDate,omitempty"`
	// Individual who recorded the record and takes responsibility for its conten.
	Recorder *HL7Reference `form:"recorder,omitempty" json:"recorder,omitempty" xml:"recorder,omitempty"`
	// The source of the information about the allergy that is recorded.
	Reporter *HL7Reference `form:"reporter,omitempty" json:"reporter,omitempty" xml:"reporter,omitempty"`
	// Assertion about certainty associated with a propensity, or potential risk, of a reaction to the identified Substance. See http://hl7.org/fhir/ValueSet/allergy-intolerance-status
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// Identification of the underlying physiological mechanism for the reaction risk. See http://hl7.org/fhir/ValueSet/allergy-intolerance-type
	Type *string `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
}

// CreateAllergyIntolerancePath computes a request path to the create action of AllergyIntolerance.
func CreateAllergyIntolerancePath(patientID int) string {
	return fmt.Sprintf("/nosh/patients/%v/allergy.intolerance", patientID)
}

// Record new vital
func (c *Client) CreateAllergyIntolerance(ctx context.Context, path string, payload *CreateAllergyIntolerancePayload) (*http.Response, error) {
	req, err := c.NewCreateAllergyIntoleranceRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateAllergyIntoleranceRequest create the request corresponding to the create action endpoint of the AllergyIntolerance resource.
func (c *Client) NewCreateAllergyIntoleranceRequest(ctx context.Context, path string, payload *CreateAllergyIntolerancePayload) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*")
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// DeleteAllergyIntolerancePath computes a request path to the delete action of AllergyIntolerance.
func DeleteAllergyIntolerancePath(patientID int, allergyIntoleranceID int) string {
	return fmt.Sprintf("/nosh/patients/%v/allergy.intolerance/%v", patientID, allergyIntoleranceID)
}

// DeleteAllergyIntolerance makes a request to the delete action endpoint of the AllergyIntolerance resource
func (c *Client) DeleteAllergyIntolerance(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteAllergyIntoleranceRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteAllergyIntoleranceRequest create the request corresponding to the delete action endpoint of the AllergyIntolerance resource.
func (c *Client) NewDeleteAllergyIntoleranceRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ListAllergyIntolerancePath computes a request path to the list action of AllergyIntolerance.
func ListAllergyIntolerancePath(patientID int) string {
	return fmt.Sprintf("/nosh/patients/%v/allergy.intolerance", patientID)
}

// List all allergy_intolerances in patient optionally filtering by year
func (c *Client) ListAllergyIntolerance(ctx context.Context, path string, years []int) (*http.Response, error) {
	req, err := c.NewListAllergyIntoleranceRequest(ctx, path, years)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListAllergyIntoleranceRequest create the request corresponding to the list action endpoint of the AllergyIntolerance resource.
func (c *Client) NewListAllergyIntoleranceRequest(ctx context.Context, path string, years []int) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	for _, p := range years {
		tmp43 := strconv.Itoa(p)
		values.Add("years", tmp43)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// RateAllergyIntolerancePayload is the AllergyIntolerance rate action payload.
type RateAllergyIntolerancePayload struct {
	Rating string `form:"rating" json:"rating" xml:"rating"`
}

// RateAllergyIntolerancePath computes a request path to the rate action of AllergyIntolerance.
func RateAllergyIntolerancePath(patientID int, allergyIntoleranceID int) string {
	return fmt.Sprintf("/nosh/patients/%v/allergy.intolerance/%v/actions/rate", patientID, allergyIntoleranceID)
}

// RateAllergyIntolerance makes a request to the rate action endpoint of the AllergyIntolerance resource
func (c *Client) RateAllergyIntolerance(ctx context.Context, path string, payload *RateAllergyIntolerancePayload) (*http.Response, error) {
	req, err := c.NewRateAllergyIntoleranceRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewRateAllergyIntoleranceRequest create the request corresponding to the rate action endpoint of the AllergyIntolerance resource.
func (c *Client) NewRateAllergyIntoleranceRequest(ctx context.Context, path string, payload *RateAllergyIntolerancePayload) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*")
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("PUT", u.String(), &body)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ShowAllergyIntolerancePath computes a request path to the show action of AllergyIntolerance.
func ShowAllergyIntolerancePath(patientID int, allergyIntoleranceID int) string {
	return fmt.Sprintf("/nosh/patients/%v/allergy.intolerance/%v", patientID, allergyIntoleranceID)
}

// Retrieve vital with given id
func (c *Client) ShowAllergyIntolerance(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowAllergyIntoleranceRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowAllergyIntoleranceRequest create the request corresponding to the show action endpoint of the AllergyIntolerance resource.
func (c *Client) NewShowAllergyIntoleranceRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// UpdateAllergyIntolerancePath computes a request path to the update action of AllergyIntolerance.
func UpdateAllergyIntolerancePath(patientID int, allergyIntoleranceID int) string {
	return fmt.Sprintf("/nosh/patients/%v/allergy.intolerance/%v", patientID, allergyIntoleranceID)
}

// UpdateAllergyIntolerance makes a request to the update action endpoint of the AllergyIntolerance resource
func (c *Client) UpdateAllergyIntolerance(ctx context.Context, path string, payload *AllergyIntolerancePayload) (*http.Response, error) {
	req, err := c.NewUpdateAllergyIntoleranceRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateAllergyIntoleranceRequest create the request corresponding to the update action endpoint of the AllergyIntolerance resource.
func (c *Client) NewUpdateAllergyIntoleranceRequest(ctx context.Context, path string, payload *AllergyIntolerancePayload) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*")
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("PATCH", u.String(), &body)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// WatchAllergyIntolerancePath computes a request path to the watch action of AllergyIntolerance.
func WatchAllergyIntolerancePath(patientID int, allergyIntoleranceID int) string {
	return fmt.Sprintf("/nosh/patients/%v/allergy.intolerance/%v/watch", patientID, allergyIntoleranceID)
}

// Retrieve vital with given id
func (c *Client) WatchAllergyIntolerance(ctx context.Context, path string) (*websocket.Conn, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "ws"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	url_ := u.String()
	cfg, err := websocket.NewConfig(url_, url_)
	if err != nil {
		return nil, err
	}
	return websocket.DialConfig(cfg)
}
