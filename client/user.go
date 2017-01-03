package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// CreateUserPayload is the user create action payload.
type CreateUserPayload struct {
	// Email of user
	Email string `form:"email" json:"email" xml:"email"`
	// Password of user
	Password string `form:"password" json:"password" xml:"password"`
}

// CreateUserPath computes a request path to the create action of user.
func CreateUserPath() string {
	return fmt.Sprintf("/nosh/users")
}

// Create new user
func (c *Client) CreateUser(ctx context.Context, path string, payload *CreateUserPayload) (*http.Response, error) {
	req, err := c.NewCreateUserRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateUserRequest create the request corresponding to the create action endpoint of the user resource.
func (c *Client) NewCreateUserRequest(ctx context.Context, path string, payload *CreateUserPayload) (*http.Request, error) {
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

// DeleteUserPath computes a request path to the delete action of user.
func DeleteUserPath(userID int) string {
	return fmt.Sprintf("/nosh/users/%v", userID)
}

// DeleteUser makes a request to the delete action endpoint of the user resource
func (c *Client) DeleteUser(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteUserRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteUserRequest create the request corresponding to the delete action endpoint of the user resource.
func (c *Client) NewDeleteUserRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListUserPath computes a request path to the list action of user.
func ListUserPath() string {
	return fmt.Sprintf("/nosh/users")
}

// Retrieve all users.
func (c *Client) ListUser(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListUserRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListUserRequest create the request corresponding to the list action endpoint of the user resource.
func (c *Client) NewListUserRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ShowUserPath computes a request path to the show action of user.
func ShowUserPath(userID int) string {
	return fmt.Sprintf("/nosh/users/%v", userID)
}

// Retrieve user with given id. IDs 1 and 2 pre-exist in the system.
func (c *Client) ShowUser(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowUserRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowUserRequest create the request corresponding to the show action endpoint of the user resource.
func (c *Client) NewShowUserRequest(ctx context.Context, path string) (*http.Request, error) {
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

// SigninUserPayload is the user signin action payload.
type SigninUserPayload struct {
	// Password of user
	Password string `form:"password" json:"password" xml:"password"`
	// Username of user
	Username string `form:"username" json:"username" xml:"username"`
}

// SigninUserPath computes a request path to the signin action of user.
func SigninUserPath() string {
	return fmt.Sprintf("/nosh/users/jwt/signin")
}

// Creates a valid JWT
func (c *Client) SigninUser(ctx context.Context, path string, payload *SigninUserPayload) (*http.Response, error) {
	req, err := c.NewSigninUserRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewSigninUserRequest create the request corresponding to the signin action endpoint of the user resource.
func (c *Client) NewSigninUserRequest(ctx context.Context, path string, payload *SigninUserPayload) (*http.Request, error) {
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

// SignupUserPayload is the user signup action payload.
type SignupUserPayload struct {
	// Name of city, town etc.
	AddressCity string `form:"address_city" json:"address_city" xml:"address_city"`
	// Street name, number, direction & P.O. Box etc.
	AddressLine string `form:"address_line" json:"address_line" xml:"address_line"`
	// Postal code for area
	AddressPostalCode int `form:"address_postal_code" json:"address_postal_code" xml:"address_postal_code"`
	// Sub-unit of country (abbreviations ok)
	AddressState string `form:"address_state" json:"address_state" xml:"address_state"`
	// Email of user
	Email string `form:"email" json:"email" xml:"email"`
	// First name of user
	FirstName string `form:"first_name" json:"first_name" xml:"first_name"`
	// Last name of user
	LastName string `form:"last_name" json:"last_name" xml:"last_name"`
	// Password of user
	Password string `form:"password" json:"password" xml:"password"`
	// Username of user
	Username string `form:"username" json:"username" xml:"username"`
}

// SignupUserPath computes a request path to the signup action of user.
func SignupUserPath() string {
	return fmt.Sprintf("/nosh/users/jwt/signup")
}

// Signup user
func (c *Client) SignupUser(ctx context.Context, path string, payload *SignupUserPayload) (*http.Response, error) {
	req, err := c.NewSignupUserRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewSignupUserRequest create the request corresponding to the signup action endpoint of the user resource.
func (c *Client) NewSignupUserRequest(ctx context.Context, path string, payload *SignupUserPayload) (*http.Request, error) {
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

// UpdateUserPayload is the user update action payload.
type UpdateUserPayload struct {
	// Email of user
	Email string `form:"email" json:"email" xml:"email"`
	// Password of user
	Password string `form:"password" json:"password" xml:"password"`
}

// UpdateUserPath computes a request path to the update action of user.
func UpdateUserPath(userID int) string {
	return fmt.Sprintf("/nosh/users/%v", userID)
}

// Change user name
func (c *Client) UpdateUser(ctx context.Context, path string, payload *UpdateUserPayload) (*http.Response, error) {
	req, err := c.NewUpdateUserRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateUserRequest create the request corresponding to the update action endpoint of the user resource.
func (c *Client) NewUpdateUserRequest(ctx context.Context, path string, payload *UpdateUserPayload) (*http.Request, error) {
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
