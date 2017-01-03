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

// CreateNutritionRequestPayload is the NutritionRequest create action payload.
type CreateNutritionRequestPayload struct {
	// A link to a record of allergies or intolerances  which should be included in the nutrition order.
	AllergyIntolerance []*Reference `form:"allergyIntolerance,omitempty" json:"allergyIntolerance,omitempty" xml:"allergyIntolerance,omitempty"`
	// Date of creation
	CreatedAt *time.Time `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	// The date and time that this nutrition order was requested.
	DateTime *time.Time `form:"dateTime,omitempty" json:"dateTime,omitempty" xml:"dateTime,omitempty"`
	// An encounter that provides additional information about the healthcare context in which this request is made.
	Encounter *Reference `form:"encounter,omitempty" json:"encounter,omitempty" xml:"encounter,omitempty"`
	// Feeding provided through the gastrointestinal tract via a tube, catheter, or stoma that delivers nutrition distal to the oral cavity.
	EnteralFormula EnteralFormulaCollection `form:"enteralFormula,omitempty" json:"enteralFormula,omitempty" xml:"enteralFormula,omitempty"`
	// This modifier is used to convey order-specific modifiers about the type of food that should NOT be given. These can be derived from
	// 		patient allergies, intolerances, or preferences such as No Red Meat, No Soy or No Wheat or  Gluten-Free.  While it should not be necessary to repeat allergy or intolerance
	// 		information captured in the referenced allergyIntolerance resource in the excludeFoodModifier, this element may be used to convey additional specificity related to foods that should be
	// 		eliminated from the patient’s diet for any reason.  This modifier applies to the entire nutrition order inclusive of the oral diet, nutritional supplements and enteral formula feedings.
	// 		See http://hl7.org/fhir/ValueSet/food-type
	ExcludeFoodModifier []*CodeableConcept `form:"excludeFoodModifier,omitempty" json:"excludeFoodModifier,omitempty" xml:"excludeFoodModifier,omitempty"`
	// This modifier is used to convey order-specific modifiers about the type of food that should be given. These can be derived
	// 		from patient allergies, intolerances, or preferences such as Halal, Vegan or Kosher. This modifier applies to the entire nutrition order inclusive of the oral diet, nutritional
	// 		supplements and enteral formula feedings. See http://hl7.org/fhir/ValueSet/encounter-diet
	FoodPreferenceModifier []*CodeableConcept `form:"foodPreferenceModifier,omitempty" json:"foodPreferenceModifier,omitempty" xml:"foodPreferenceModifier,omitempty"`
	// API href of nutrition request
	Href string `form:"href" json:"href" xml:"href"`
	// ID of nutrition request
	ID int `form:"id" json:"id" xml:"id"`
	// Identifiers assigned to this order by the order sender or by the order receiver.
	Identifier []*Identifier `form:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	// Diet given orally in contrast to enteral (tube) feeding.
	OralDiet OralDietCollection `form:"oralDiet,omitempty" json:"oralDiet,omitempty" xml:"oralDiet,omitempty"`
	// The practitioner that holds legal responsibility for ordering the diet, nutritional supplement, or formula feedings.
	Orderer *Reference `form:"orderer,omitempty" json:"orderer,omitempty" xml:"orderer,omitempty"`
	// The person (patient) who needs the nutrition order for an oral diet, nutritional supplement and/or enteral or formula feeding.
	Patient *Reference `form:"patient,omitempty" json:"patient,omitempty" xml:"patient,omitempty"`
	// The workflow status of the nutrition order/request. See http://hl7.org/fhir/nutrition-request-status
	Status string `form:"status" json:"status" xml:"status"`
	// Oral nutritional products given in order to add further nutritional value to the patient's diet.
	Supplement []*Supplement `form:"supplement,omitempty" json:"supplement,omitempty" xml:"supplement,omitempty"`
	// Date of last update
	UpdatedAt *time.Time `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

// CreateNutritionRequestPath computes a request path to the create action of NutritionRequest.
func CreateNutritionRequestPath(patientID int) string {
	return fmt.Sprintf("/nosh/patients/%v/nutrition_requests", patientID)
}

// Record new vital
func (c *Client) CreateNutritionRequest(ctx context.Context, path string, payload *CreateNutritionRequestPayload) (*http.Response, error) {
	req, err := c.NewCreateNutritionRequestRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateNutritionRequestRequest create the request corresponding to the create action endpoint of the NutritionRequest resource.
func (c *Client) NewCreateNutritionRequestRequest(ctx context.Context, path string, payload *CreateNutritionRequestPayload) (*http.Request, error) {
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

// DeleteNutritionRequestPath computes a request path to the delete action of NutritionRequest.
func DeleteNutritionRequestPath(patientID int, nutritionRequestID int) string {
	return fmt.Sprintf("/nosh/patients/%v/nutrition_requests/%v", patientID, nutritionRequestID)
}

// DeleteNutritionRequest makes a request to the delete action endpoint of the NutritionRequest resource
func (c *Client) DeleteNutritionRequest(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteNutritionRequestRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteNutritionRequestRequest create the request corresponding to the delete action endpoint of the NutritionRequest resource.
func (c *Client) NewDeleteNutritionRequestRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListNutritionRequestPath computes a request path to the list action of NutritionRequest.
func ListNutritionRequestPath(patientID int) string {
	return fmt.Sprintf("/nosh/patients/%v/nutrition_requests", patientID)
}

// List all nutrition_requests in patient optionally filtering by year
func (c *Client) ListNutritionRequest(ctx context.Context, path string, years []int) (*http.Response, error) {
	req, err := c.NewListNutritionRequestRequest(ctx, path, years)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListNutritionRequestRequest create the request corresponding to the list action endpoint of the NutritionRequest resource.
func (c *Client) NewListNutritionRequestRequest(ctx context.Context, path string, years []int) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	for _, p := range years {
		tmp44 := strconv.Itoa(p)
		values.Add("years", tmp44)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// RateNutritionRequestPayload is the NutritionRequest rate action payload.
type RateNutritionRequestPayload struct {
	Rating string `form:"rating" json:"rating" xml:"rating"`
}

// RateNutritionRequestPath computes a request path to the rate action of NutritionRequest.
func RateNutritionRequestPath(patientID int, nutritionRequestID int) string {
	return fmt.Sprintf("/nosh/patients/%v/nutrition_requests/%v/actions/rate", patientID, nutritionRequestID)
}

// RateNutritionRequest makes a request to the rate action endpoint of the NutritionRequest resource
func (c *Client) RateNutritionRequest(ctx context.Context, path string, payload *RateNutritionRequestPayload) (*http.Response, error) {
	req, err := c.NewRateNutritionRequestRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewRateNutritionRequestRequest create the request corresponding to the rate action endpoint of the NutritionRequest resource.
func (c *Client) NewRateNutritionRequestRequest(ctx context.Context, path string, payload *RateNutritionRequestPayload) (*http.Request, error) {
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

// ShowNutritionRequestPath computes a request path to the show action of NutritionRequest.
func ShowNutritionRequestPath(patientID int, nutritionRequestID int) string {
	return fmt.Sprintf("/nosh/patients/%v/nutrition_requests/%v", patientID, nutritionRequestID)
}

// Retrieve vital with given id
func (c *Client) ShowNutritionRequest(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowNutritionRequestRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowNutritionRequestRequest create the request corresponding to the show action endpoint of the NutritionRequest resource.
func (c *Client) NewShowNutritionRequestRequest(ctx context.Context, path string) (*http.Request, error) {
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

// UpdateNutritionRequestPath computes a request path to the update action of NutritionRequest.
func UpdateNutritionRequestPath(patientID int, nutritionRequestID int) string {
	return fmt.Sprintf("/nosh/patients/%v/nutrition_requests/%v", patientID, nutritionRequestID)
}

// UpdateNutritionRequest makes a request to the update action endpoint of the NutritionRequest resource
func (c *Client) UpdateNutritionRequest(ctx context.Context, path string, payload *NutritionRequestPayload) (*http.Response, error) {
	req, err := c.NewUpdateNutritionRequestRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateNutritionRequestRequest create the request corresponding to the update action endpoint of the NutritionRequest resource.
func (c *Client) NewUpdateNutritionRequestRequest(ctx context.Context, path string, payload *NutritionRequestPayload) (*http.Request, error) {
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

// WatchNutritionRequestPath computes a request path to the watch action of NutritionRequest.
func WatchNutritionRequestPath(patientID int, nutritionRequestID int) string {
	return fmt.Sprintf("/nosh/patients/%v/nutrition_requests/%v/watch", patientID, nutritionRequestID)
}

// Retrieve vital with given id
func (c *Client) WatchNutritionRequest(ctx context.Context, path string) (*websocket.Conn, error) {
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
