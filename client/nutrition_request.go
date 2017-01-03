package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/net/websocket"
	"net/http"
	"net/url"
	"strconv"
)

// CreateNutritionRequestPayload is the NutritionRequest create action payload.
type CreateNutritionRequestPayload struct {
	Address                *string `form:"address,omitempty" json:"address,omitempty" xml:"address,omitempty"`
	AddressPostalCode      *string `form:"address_postal_code,omitempty" json:"address_postal_code,omitempty" xml:"address_postal_code,omitempty"`
	EnteralFormula         *string `form:"enteralFormula,omitempty" json:"enteralFormula,omitempty" xml:"enteralFormula,omitempty"`
	FoodPreferenceModifier *struct {
		Coding *string `form:"coding,omitempty" json:"coding,omitempty" xml:"coding,omitempty"`
		ID     *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
		Text   *string `form:"text,omitempty" json:"text,omitempty" xml:"text,omitempty"`
	} `form:"foodPreferenceModifier,omitempty" json:"foodPreferenceModifier,omitempty" xml:"foodPreferenceModifier,omitempty"`
	Href             *string           `form:"href,omitempty" json:"href,omitempty" xml:"href,omitempty"`
	ID               *string           `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Identifier       *string           `form:"identifier,omitempty" json:"identifier,omitempty" xml:"identifier,omitempty"`
	Links            *string           `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
	NutritionRequest *NutritionRequest `form:"nutritionRequest,omitempty" json:"nutritionRequest,omitempty" xml:"nutritionRequest,omitempty"`
	OralDiet         *string           `form:"oralDiet,omitempty" json:"oralDiet,omitempty" xml:"oralDiet,omitempty"`
	Patient          *struct {
		ID        *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
		Reference *string `form:"reference,omitempty" json:"reference,omitempty" xml:"reference,omitempty"`
	} `form:"patient,omitempty" json:"patient,omitempty" xml:"patient,omitempty"`
	Supplement *string `form:"supplement,omitempty" json:"supplement,omitempty" xml:"supplement,omitempty"`
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
		scheme = "http"
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
		scheme = "http"
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
		scheme = "http"
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
		scheme = "http"
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
		scheme = "http"
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
		scheme = "http"
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
