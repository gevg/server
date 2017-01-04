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

// CreateObservationPayload is the Observation create action payload.
type CreateObservationPayload struct {
	// Indicates the site on the subject's body where the observation was made (i.e. the target site). See http://hl7.org/fhir/ValueSet/body-site
	BodySite *CodeableConcept `form:"bodySite,omitempty" json:"bodySite,omitempty" xml:"bodySite,omitempty"`
	// A code that classifies the general type of observation being made.  This is used  for searching, sorting and display purposes. See http://hl7.org/fhir/ValueSet/observation-category
	Category *CodeableConcept `form:"category,omitempty" json:"category,omitempty" xml:"category,omitempty"`
	// Describes what was observed. Sometimes this is called the observation 'name'. See http://hl7.org/fhir/ValueSet/observation-codes
	Code *CodeableConcept `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// May include statements about significant, unexpected or unreliable values, or information about the source of the value where this may be relevant to the interpretation of the result.
	Comments *string `form:"comments,omitempty" json:"comments,omitempty" xml:"comments,omitempty"`
	// Some observations have multiple component observations.  These component observations are expressed as separate code
	// 		value pairs that share the same attributes.  Examples include systolic and diastolic component observations for blood pressure measurement and multiple
	// 		component observations for genetics observations.
	Component ComponentMediaCollection `form:"component,omitempty" json:"component,omitempty" xml:"component,omitempty"`
	// Provides a reason why the expected value in the element Observation.value[x] is missing. See http://hl7.org/fhir/ValueSet/observation-valueabsentreason
	DateAbsentReason *CodeableConcept `form:"dateAbsentReason,omitempty" json:"dateAbsentReason,omitempty" xml:"dateAbsentReason,omitempty"`
	// The device used to generate the observation data.
	Device *HL7Reference `form:"device,omitempty" json:"device,omitempty" xml:"device,omitempty"`
	// The time or time-period the observed value is asserted as being true. For biological subjects - e.g. human patients -
	// 		this is usually called the "physiologically relevant time". This is usually either the time of the procedure or of specimen collection,
	// 		but very often the source of the date/time is not known, only the date/time itself.
	EffectiveDateTime *time.Time `form:"effectiveDateTime,omitempty" json:"effectiveDateTime,omitempty" xml:"effectiveDateTime,omitempty"`
	// The time or time-period the observed value is asserted as being true. For biological subjects - e.g. human patients -
	// 		this is usually called the "physiologically relevant time". This is usually either the time of the procedure or of specimen collection,
	// 		but very often the source of the date/time is not known, only the date/time itself.
	EffectivePeriod *Period `form:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty" xml:"effectivePeriod,omitempty"`
	// The healthcare event  (e.g. a patient and healthcare provider interaction) during which this observation is made.
	Encounter *HL7Reference `form:"encounter,omitempty" json:"encounter,omitempty" xml:"encounter,omitempty"`
	// A unique identifier for the simple observation instance.
	Identifier []*Identifier `form:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The assessment made based on the result of the observation.  Intended as a simple compact code often
	// 		placed adjacent to the result value in reports and flow sheets to signal the meaning/normalcy status of the result. Otherwise known as abnormal flag.
	// 		See http://hl7.org/fhir/ValueSet/observation-interpretation
	Interpretation *CodeableConcept `form:"interpretation,omitempty" json:"interpretation,omitempty" xml:"interpretation,omitempty"`
	// The date and time this observation was made available to providers, typically after the results have been reviewed and verified.
	Issued *time.Time `form:"issued,omitempty" json:"issued,omitempty" xml:"issued,omitempty"`
	// Indicates the mechanism used to perform the observation. See http://hl7.org/fhir/ValueSet/observation-methods
	Method *CodeableConcept `form:"method,omitempty" json:"method,omitempty" xml:"method,omitempty"`
	// Who was responsible for asserting the observed value as 'true'.
	Performer []*HL7Reference `form:"performer,omitempty" json:"performer,omitempty" xml:"performer,omitempty"`
	// Guidance on how to interpret the value by comparison to a normal or recommended range.
	ReferenceRange ReferenceRangeMediaCollection `form:"referenceRange,omitempty" json:"referenceRange,omitempty" xml:"referenceRange,omitempty"`
	// A  reference to another resource (usually another Observation but could  also be a QuestionnaireAnswer) whose relationship is defined by the relationship type code.
	Related RelatedMediaCollection `form:"related,omitempty" json:"related,omitempty" xml:"related,omitempty"`
	// The specimen that was used when this observation was made.
	Specimen *HL7Reference `form:"specimen,omitempty" json:"specimen,omitempty" xml:"specimen,omitempty"`
	// The status of the result value. See http://hl7.org/fhir/ValueSet/observation-status
	Status               *string          `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	ValueAttachment      *Attachment      `form:"valueAttachment,omitempty" json:"valueAttachment,omitempty" xml:"valueAttachment,omitempty"`
	ValueCodeableConcept *CodeableConcept `form:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty" xml:"valueCodeableConcept,omitempty"`
	ValueDatTime         *time.Time       `form:"valueDatTime,omitempty" json:"valueDatTime,omitempty" xml:"valueDatTime,omitempty"`
	ValuePeriod          *Period          `form:"valuePeriod,omitempty" json:"valuePeriod,omitempty" xml:"valuePeriod,omitempty"`
	ValueQuantity        *Quantity        `form:"valueQuantity,omitempty" json:"valueQuantity,omitempty" xml:"valueQuantity,omitempty"`
	ValueRange           *Range           `form:"valueRange,omitempty" json:"valueRange,omitempty" xml:"valueRange,omitempty"`
	ValueSampledData     *SampleData      `form:"valueSampledData,omitempty" json:"valueSampledData,omitempty" xml:"valueSampledData,omitempty"`
	ValueString          *string          `form:"valueString,omitempty" json:"valueString,omitempty" xml:"valueString,omitempty"`
	ValueTime            *time.Time       `form:"valueTime,omitempty" json:"valueTime,omitempty" xml:"valueTime,omitempty"`
}

// CreateObservationPath computes a request path to the create action of Observation.
func CreateObservationPath(patientID int) string {
	return fmt.Sprintf("/nosh/patients/%v/observation", patientID)
}

// Record new vital
func (c *Client) CreateObservation(ctx context.Context, path string, payload *CreateObservationPayload) (*http.Response, error) {
	req, err := c.NewCreateObservationRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateObservationRequest create the request corresponding to the create action endpoint of the Observation resource.
func (c *Client) NewCreateObservationRequest(ctx context.Context, path string, payload *CreateObservationPayload) (*http.Request, error) {
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

// DeleteObservationPath computes a request path to the delete action of Observation.
func DeleteObservationPath(patientID int, observationID int) string {
	return fmt.Sprintf("/nosh/patients/%v/observation/%v", patientID, observationID)
}

// DeleteObservation makes a request to the delete action endpoint of the Observation resource
func (c *Client) DeleteObservation(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteObservationRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteObservationRequest create the request corresponding to the delete action endpoint of the Observation resource.
func (c *Client) NewDeleteObservationRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListObservationPath computes a request path to the list action of Observation.
func ListObservationPath(patientID int) string {
	return fmt.Sprintf("/nosh/patients/%v/observation", patientID)
}

// List all observations in patient optionally filtering by year
func (c *Client) ListObservation(ctx context.Context, path string, years []int) (*http.Response, error) {
	req, err := c.NewListObservationRequest(ctx, path, years)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListObservationRequest create the request corresponding to the list action endpoint of the Observation resource.
func (c *Client) NewListObservationRequest(ctx context.Context, path string, years []int) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	for _, p := range years {
		tmp45 := strconv.Itoa(p)
		values.Add("years", tmp45)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// RateObservationPath computes a request path to the rate action of Observation.
func RateObservationPath(patientID int, observationID int) string {
	return fmt.Sprintf("/nosh/patients/%v/observation/%v/actions/rate", patientID, observationID)
}

// RateObservation makes a request to the rate action endpoint of the Observation resource
func (c *Client) RateObservation(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewRateObservationRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewRateObservationRequest create the request corresponding to the rate action endpoint of the Observation resource.
func (c *Client) NewRateObservationRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ShowObservationPath computes a request path to the show action of Observation.
func ShowObservationPath(patientID int, observationID int) string {
	return fmt.Sprintf("/nosh/patients/%v/observation/%v", patientID, observationID)
}

// Retrieve vital with given id
func (c *Client) ShowObservation(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowObservationRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowObservationRequest create the request corresponding to the show action endpoint of the Observation resource.
func (c *Client) NewShowObservationRequest(ctx context.Context, path string) (*http.Request, error) {
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

// UpdateObservationPath computes a request path to the update action of Observation.
func UpdateObservationPath(patientID int, observationID int) string {
	return fmt.Sprintf("/nosh/patients/%v/observation/%v", patientID, observationID)
}

// UpdateObservation makes a request to the update action endpoint of the Observation resource
func (c *Client) UpdateObservation(ctx context.Context, path string, payload *ObservationPayload) (*http.Response, error) {
	req, err := c.NewUpdateObservationRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateObservationRequest create the request corresponding to the update action endpoint of the Observation resource.
func (c *Client) NewUpdateObservationRequest(ctx context.Context, path string, payload *ObservationPayload) (*http.Request, error) {
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

// WatchObservationPath computes a request path to the watch action of Observation.
func WatchObservationPath(patientID int, observationID int) string {
	return fmt.Sprintf("/nosh/patients/%v/observation/%v/watch", patientID, observationID)
}

// Retrieve vital with given id
func (c *Client) WatchObservation(ctx context.Context, path string) (*websocket.Conn, error) {
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
