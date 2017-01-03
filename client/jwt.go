package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
	"strconv"
)

// SecureJWTPath computes a request path to the secure action of jwt.
func SecureJWTPath() string {
	return fmt.Sprintf("/nosh/jwt")
}

// This action is secured with the jwt scheme
func (c *Client) SecureJWT(ctx context.Context, path string, fail *bool) (*http.Response, error) {
	req, err := c.NewSecureJWTRequest(ctx, path, fail)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewSecureJWTRequest create the request corresponding to the secure action endpoint of the jwt resource.
func (c *Client) NewSecureJWTRequest(ctx context.Context, path string, fail *bool) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if fail != nil {
		tmp46 := strconv.FormatBool(*fail)
		values.Set("fail", tmp46)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		c.JWTSigner.Sign(req)
	}
	return req, nil
}

// SigninJWTPayload is the jwt signin action payload.
type SigninJWTPayload struct {
	Password string `form:"password" json:"password" xml:"password"`
	Username string `form:"username" json:"username" xml:"username"`
}

// SigninJWTPath computes a request path to the signin action of jwt.
func SigninJWTPath() string {
	return fmt.Sprintf("/nosh/jwt/signin")
}

// Creates a valid JWT
func (c *Client) SigninJWT(ctx context.Context, path string, payload *SigninJWTPayload) (*http.Response, error) {
	req, err := c.NewSigninJWTRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewSigninJWTRequest create the request corresponding to the signin action endpoint of the jwt resource.
func (c *Client) NewSigninJWTRequest(ctx context.Context, path string, payload *SigninJWTPayload) (*http.Request, error) {
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
	if c.SigninBasicAuthSigner != nil {
		c.SigninBasicAuthSigner.Sign(req)
	}
	return req, nil
}

// SignupJWTPayload is the jwt signup action payload.
type SignupJWTPayload struct {
	AddressCity       string `form:"address_city" json:"address_city" xml:"address_city"`
	AddressLine       string `form:"address_line" json:"address_line" xml:"address_line"`
	AddressPostalCode string `form:"address_postal_code" json:"address_postal_code" xml:"address_postal_code"`
	AddressState      string `form:"address_state" json:"address_state" xml:"address_state"`
	Email             string `form:"email" json:"email" xml:"email"`
	FirstName         string `form:"first_name" json:"first_name" xml:"first_name"`
	LastName          string `form:"last_name" json:"last_name" xml:"last_name"`
	Password          string `form:"password" json:"password" xml:"password"`
	Username          string `form:"username" json:"username" xml:"username"`
}

// SignupJWTPath computes a request path to the signup action of jwt.
func SignupJWTPath() string {
	return fmt.Sprintf("/nosh/jwt/signup")
}

// Signup user
func (c *Client) SignupJWT(ctx context.Context, path string, payload *SignupJWTPayload) (*http.Response, error) {
	req, err := c.NewSignupJWTRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewSignupJWTRequest create the request corresponding to the signup action endpoint of the jwt resource.
func (c *Client) NewSignupJWTRequest(ctx context.Context, path string, payload *SignupJWTPayload) (*http.Request, error) {
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

// UnsecureJWTPath computes a request path to the unsecure action of jwt.
func UnsecureJWTPath() string {
	return fmt.Sprintf("/nosh/jwt/unsecure")
}

// This action does not require auth
func (c *Client) UnsecureJWT(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewUnsecureJWTRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUnsecureJWTRequest create the request corresponding to the unsecure action endpoint of the jwt resource.
func (c *Client) NewUnsecureJWTRequest(ctx context.Context, path string) (*http.Request, error) {
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
