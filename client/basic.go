package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// SecureBasicPath computes a request path to the secure action of basic.
func SecureBasicPath() string {
	return fmt.Sprintf("/nosh/basic")
}

// This action is secure with the basic_auth scheme
func (c *Client) SecureBasic(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewSecureBasicRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewSecureBasicRequest create the request corresponding to the secure action endpoint of the basic resource.
func (c *Client) NewSecureBasicRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.BasicAuthSigner != nil {
		c.BasicAuthSigner.Sign(req)
	}
	return req, nil
}

// UnsecureBasicPath computes a request path to the unsecure action of basic.
func UnsecureBasicPath() string {
	return fmt.Sprintf("/nosh/basic/unsecure")
}

// This action does not require auth
func (c *Client) UnsecureBasic(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewUnsecureBasicRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUnsecureBasicRequest create the request corresponding to the unsecure action endpoint of the basic resource.
func (c *Client) NewUnsecureBasicRequest(ctx context.Context, path string) (*http.Request, error) {
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
