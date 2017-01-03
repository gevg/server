package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
	"strconv"
)

// CreatePatientPayload is the patient create action payload.
type CreatePatientPayload struct {
	Name string `form:"name" json:"name" xml:"name"`
}

// CreatePatientPath computes a request path to the create action of patient.
func CreatePatientPath() string {
	return fmt.Sprintf("/nosh/patients")
}

// Create new patient
func (c *Client) CreatePatient(ctx context.Context, path string, payload *CreatePatientPayload) (*http.Response, error) {
	req, err := c.NewCreatePatientRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreatePatientRequest create the request corresponding to the create action endpoint of the patient resource.
func (c *Client) NewCreatePatientRequest(ctx context.Context, path string, payload *CreatePatientPayload) (*http.Request, error) {
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

// DeletePatientPath computes a request path to the delete action of patient.
func DeletePatientPath(patientID int) string {
	return fmt.Sprintf("/nosh/patients/%v", patientID)
}

// DeletePatient makes a request to the delete action endpoint of the patient resource
func (c *Client) DeletePatient(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeletePatientRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeletePatientRequest create the request corresponding to the delete action endpoint of the patient resource.
func (c *Client) NewDeletePatientRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListPatientPath computes a request path to the list action of patient.
func ListPatientPath() string {
	return fmt.Sprintf("/nosh/patients/jwt")
}

// ListPatientPath2 computes a request path to the list action of patient.
func ListPatientPath2() string {
	return fmt.Sprintf("/nosh/patients")
}

// Retrieve all patients.
func (c *Client) ListPatient(ctx context.Context, path string, fail *bool) (*http.Response, error) {
	req, err := c.NewListPatientRequest(ctx, path, fail)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListPatientRequest create the request corresponding to the list action endpoint of the patient resource.
func (c *Client) NewListPatientRequest(ctx context.Context, path string, fail *bool) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if fail != nil {
		tmp47 := strconv.FormatBool(*fail)
		values.Set("fail", tmp47)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ShowPatientPath computes a request path to the show action of patient.
func ShowPatientPath(patientID int) string {
	return fmt.Sprintf("/nosh/patients/%v", patientID)
}

// Retrieve patient with given id. IDs 1 and 2 pre-exist in the system.
func (c *Client) ShowPatient(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowPatientRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowPatientRequest create the request corresponding to the show action endpoint of the patient resource.
func (c *Client) NewShowPatientRequest(ctx context.Context, path string) (*http.Request, error) {
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

// UpdatePatientPayload is the patient update action payload.
type UpdatePatientPayload struct {
	Name string `form:"name" json:"name" xml:"name"`
}

// UpdatePatientPath computes a request path to the update action of patient.
func UpdatePatientPath(patientID int) string {
	return fmt.Sprintf("/nosh/patients/%v", patientID)
}

// Change patient name
func (c *Client) UpdatePatient(ctx context.Context, path string, payload *UpdatePatientPayload) (*http.Response, error) {
	req, err := c.NewUpdatePatientRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdatePatientRequest create the request corresponding to the update action endpoint of the patient resource.
func (c *Client) NewUpdatePatientRequest(ctx context.Context, path string, payload *UpdatePatientPayload) (*http.Request, error) {
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
