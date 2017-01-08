package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// AuthorizeOauth2ProviderPath computes a request path to the authorize action of oauth2_provider.
func AuthorizeOauth2ProviderPath() string {
	return fmt.Sprintf("/nosh/oauth2/auth")
}

// Authorize OAuth2 client
func (c *Client) AuthorizeOauth2Provider(ctx context.Context, path string, clientID string, responseType string, contained *string, containedType *string, count *int, elements *string, has *string, id *int, include *string, lastUpdate *time.Time, list *string, profile *string, query *string, revinclude *string, security *string, sort *string, summary *string, tag *string, text *string, type_ *string, redirectURI *string, scope *string, state *string) (*http.Response, error) {
	req, err := c.NewAuthorizeOauth2ProviderRequest(ctx, path, clientID, responseType, contained, containedType, count, elements, has, id, include, lastUpdate, list, profile, query, revinclude, security, sort, summary, tag, text, type_, redirectURI, scope, state)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewAuthorizeOauth2ProviderRequest create the request corresponding to the authorize action endpoint of the oauth2_provider resource.
func (c *Client) NewAuthorizeOauth2ProviderRequest(ctx context.Context, path string, clientID string, responseType string, contained *string, containedType *string, count *int, elements *string, has *string, id *int, include *string, lastUpdate *time.Time, list *string, profile *string, query *string, revinclude *string, security *string, sort *string, summary *string, tag *string, text *string, type_ *string, redirectURI *string, scope *string, state *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	values.Set("client_id", clientID)
	values.Set("response_type", responseType)
	if contained != nil {
		values.Set("_contained", *contained)
	}
	if containedType != nil {
		values.Set("_containedType", *containedType)
	}
	if count != nil {
		tmp59 := strconv.Itoa(*count)
		values.Set("_count", tmp59)
	}
	if elements != nil {
		values.Set("_elements", *elements)
	}
	if has != nil {
		values.Set("_has", *has)
	}
	if id != nil {
		tmp60 := strconv.Itoa(*id)
		values.Set("_id", tmp60)
	}
	if include != nil {
		values.Set("_include", *include)
	}
	if lastUpdate != nil {
		tmp61 := lastUpdate.Format(time.RFC3339)
		values.Set("_lastUpdate", tmp61)
	}
	if list != nil {
		values.Set("_list", *list)
	}
	if profile != nil {
		values.Set("_profile", *profile)
	}
	if query != nil {
		values.Set("_query", *query)
	}
	if revinclude != nil {
		values.Set("_revinclude", *revinclude)
	}
	if security != nil {
		values.Set("_security", *security)
	}
	if sort != nil {
		values.Set("_sort", *sort)
	}
	if summary != nil {
		values.Set("_summary", *summary)
	}
	if tag != nil {
		values.Set("_tag", *tag)
	}
	if text != nil {
		values.Set("_text", *text)
	}
	if type_ != nil {
		values.Set("_type", *type_)
	}
	if redirectURI != nil {
		values.Set("redirect_uri", *redirectURI)
	}
	if scope != nil {
		values.Set("scope", *scope)
	}
	if state != nil {
		values.Set("state", *state)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// GetTokenOauth2ProviderPath computes a request path to the get_token action of oauth2_provider.
func GetTokenOauth2ProviderPath() string {
	return fmt.Sprintf("/nosh/oauth2/token")
}

// Get access token from authorization code or refresh token
func (c *Client) GetTokenOauth2Provider(ctx context.Context, path string, payload *TokenPayload, contained *string, containedType *string, count *int, elements *string, has *string, id *int, include *string, lastUpdate *time.Time, list *string, profile *string, query *string, revinclude *string, security *string, sort *string, summary *string, tag *string, text *string, type_ *string) (*http.Response, error) {
	req, err := c.NewGetTokenOauth2ProviderRequest(ctx, path, payload, contained, containedType, count, elements, has, id, include, lastUpdate, list, profile, query, revinclude, security, sort, summary, tag, text, type_)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetTokenOauth2ProviderRequest create the request corresponding to the get_token action endpoint of the oauth2_provider resource.
func (c *Client) NewGetTokenOauth2ProviderRequest(ctx context.Context, path string, payload *TokenPayload, contained *string, containedType *string, count *int, elements *string, has *string, id *int, include *string, lastUpdate *time.Time, list *string, profile *string, query *string, revinclude *string, security *string, sort *string, summary *string, tag *string, text *string, type_ *string) (*http.Request, error) {
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
	values := u.Query()
	if contained != nil {
		values.Set("_contained", *contained)
	}
	if containedType != nil {
		values.Set("_containedType", *containedType)
	}
	if count != nil {
		tmp62 := strconv.Itoa(*count)
		values.Set("_count", tmp62)
	}
	if elements != nil {
		values.Set("_elements", *elements)
	}
	if has != nil {
		values.Set("_has", *has)
	}
	if id != nil {
		tmp63 := strconv.Itoa(*id)
		values.Set("_id", tmp63)
	}
	if include != nil {
		values.Set("_include", *include)
	}
	if lastUpdate != nil {
		tmp64 := lastUpdate.Format(time.RFC3339)
		values.Set("_lastUpdate", tmp64)
	}
	if list != nil {
		values.Set("_list", *list)
	}
	if profile != nil {
		values.Set("_profile", *profile)
	}
	if query != nil {
		values.Set("_query", *query)
	}
	if revinclude != nil {
		values.Set("_revinclude", *revinclude)
	}
	if security != nil {
		values.Set("_security", *security)
	}
	if sort != nil {
		values.Set("_sort", *sort)
	}
	if summary != nil {
		values.Set("_summary", *summary)
	}
	if tag != nil {
		values.Set("_tag", *tag)
	}
	if text != nil {
		values.Set("_text", *text)
	}
	if type_ != nil {
		values.Set("_type", *type_)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	if c.Oauth2ClientBasicAuthSigner != nil {
		c.Oauth2ClientBasicAuthSigner.Sign(req)
	}
	return req, nil
}