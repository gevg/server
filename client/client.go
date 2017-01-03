package client

import (
	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
)

// Client is the Secure service client.
type Client struct {
	*goaclient.Client
	BasicAuthSigner             goaclient.Signer
	JWTSigner                   goaclient.Signer
	SigninBasicAuthSigner       goaclient.Signer
	Oauth2ClientBasicAuthSigner goaclient.Signer
	Encoder                     *goa.HTTPEncoder
	Decoder                     *goa.HTTPDecoder
}

// New instantiates the client.
func New(c goaclient.Doer) *Client {
	client := &Client{
		Client:  goaclient.New(c),
		Encoder: goa.NewHTTPEncoder(),
		Decoder: goa.NewHTTPDecoder(),
	}

	// Setup encoders and decoders
	client.Encoder.Register(goa.NewJSONEncoder, "application/json")
	client.Decoder.Register(goa.NewJSONDecoder, "application/json")

	// Setup default encoder and decoder
	client.Encoder.Register(goa.NewJSONEncoder, "*/*")
	client.Decoder.Register(goa.NewJSONDecoder, "*/*")

	return client
}

// SetBasicAuthSigner sets the request signer for the basic_auth security scheme.
func (c *Client) SetBasicAuthSigner(signer goaclient.Signer) {
	c.BasicAuthSigner = signer
}

// SetJWTSigner sets the request signer for the jwt security scheme.
func (c *Client) SetJWTSigner(signer goaclient.Signer) {
	c.JWTSigner = signer
}

// SetSigninBasicAuthSigner sets the request signer for the SigninBasicAuth security scheme.
func (c *Client) SetSigninBasicAuthSigner(signer goaclient.Signer) {
	c.SigninBasicAuthSigner = signer
}

// SetOauth2ClientBasicAuthSigner sets the request signer for the oauth2_client_basic_auth security scheme.
func (c *Client) SetOauth2ClientBasicAuthSigner(signer goaclient.Signer) {
	c.Oauth2ClientBasicAuthSigner = signer
}
