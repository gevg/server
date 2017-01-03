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

// CreateObservationPayload is the Observation create action payload.
type CreateObservationPayload struct {
	Observation *Observation `form:"observation,omitempty" json:"observation,omitempty" xml:"observation,omitempty"`
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
		scheme = "http"
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
		scheme = "http"
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
		scheme = "http"
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
		scheme = "http"
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
		scheme = "http"
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
		scheme = "http"
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
