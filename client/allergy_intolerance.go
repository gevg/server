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

// CreateAllergyIntolerancePayload is the AllergyIntolerance create action payload.
type CreateAllergyIntolerancePayload struct {
	Active  *bool `form:"active,omitempty" json:"active,omitempty" xml:"active,omitempty"`
	Address *struct {
		// City
		City   string `form:"city" json:"city" xml:"city"`
		Number *struct {
			// Street name
			Other *string `form:"other,omitempty" json:"other,omitempty" xml:"other,omitempty"`
		} `form:"number,omitempty" json:"number,omitempty" xml:"number,omitempty"`
		// Street name
		Street *string `form:"street,omitempty" json:"street,omitempty" xml:"street,omitempty"`
	} `form:"address,omitempty" json:"address,omitempty" xml:"address,omitempty"`
	BirthDate     *string `form:"birthDate,omitempty" json:"birthDate,omitempty" xml:"birthDate,omitempty"`
	Country       *string `form:"country,omitempty" json:"country,omitempty" xml:"country,omitempty"`
	Deceased      *bool   `form:"deceased,omitempty" json:"deceased,omitempty" xml:"deceased,omitempty"`
	Gender        *string `form:"gender,omitempty" json:"gender,omitempty" xml:"gender,omitempty"`
	MultipleBirth *bool   `form:"multiple_birth,omitempty" json:"multiple_birth,omitempty" xml:"multiple_birth,omitempty"`
	Region        *string `form:"region,omitempty" json:"region,omitempty" xml:"region,omitempty"`
	Review        *string `form:"review,omitempty" json:"review,omitempty" xml:"review,omitempty"`
	Sweetness     *int    `form:"sweetness,omitempty" json:"sweetness,omitempty" xml:"sweetness,omitempty"`
}

// CreateAllergyIntolerancePath computes a request path to the create action of AllergyIntolerance.
func CreateAllergyIntolerancePath(patientID int) string {
	return fmt.Sprintf("/nosh/patients/%v/allergy_intolerance", patientID)
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
		scheme = "http"
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
	return fmt.Sprintf("/nosh/patients/%v/allergy_intolerance/%v", patientID, allergyIntoleranceID)
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
		scheme = "http"
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
	return fmt.Sprintf("/nosh/patients/%v/allergy_intolerance", patientID)
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
		scheme = "http"
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
	return fmt.Sprintf("/nosh/patients/%v/allergy_intolerance/%v/actions/rate", patientID, allergyIntoleranceID)
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
		scheme = "http"
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
	return fmt.Sprintf("/nosh/patients/%v/allergy_intolerance/%v", patientID, allergyIntoleranceID)
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
		scheme = "http"
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
	return fmt.Sprintf("/nosh/patients/%v/allergy_intolerance/%v", patientID, allergyIntoleranceID)
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
		scheme = "http"
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
	return fmt.Sprintf("/nosh/patients/%v/allergy_intolerance/%v/watch", patientID, allergyIntoleranceID)
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
