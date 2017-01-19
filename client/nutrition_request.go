// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/goa-fhir/server/design
// --out=$(GOPATH)\src\github.com\goa-fhir\server
// --version=v1.1.0-dirty
//
// API "goa-FHIR": NutritionRequest Resource Client
//
// The content of this file is auto-generated, DO NOT MODIFY

package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// ReadNutritionRequestPath computes a request path to the read action of NutritionRequest.
func ReadNutritionRequestPath(patientID int, nutritionRequestID int) string {
	param0 := strconv.Itoa(patientID)
	param1 := strconv.Itoa(nutritionRequestID)

	return fmt.Sprintf("/nosh/patients/%s/nutrition.requests/%s", param0, param1)
}

// Retrieve vital with given id
func (c *Client) ReadNutritionRequest(ctx context.Context, path string, contained *string, containedType *string, count *int, elements *string, has *string, id *int, include *string, lastUpdate *time.Time, list *string, profile *string, query *string, revinclude *string, security *string, sort *string, summary *string, tag *string, text *string, type_ *string) (*http.Response, error) {
	req, err := c.NewReadNutritionRequestRequest(ctx, path, contained, containedType, count, elements, has, id, include, lastUpdate, list, profile, query, revinclude, security, sort, summary, tag, text, type_)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewReadNutritionRequestRequest create the request corresponding to the read action endpoint of the NutritionRequest resource.
func (c *Client) NewReadNutritionRequestRequest(ctx context.Context, path string, contained *string, containedType *string, count *int, elements *string, has *string, id *int, include *string, lastUpdate *time.Time, list *string, profile *string, query *string, revinclude *string, security *string, sort *string, summary *string, tag *string, text *string, type_ *string) (*http.Request, error) {
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
		tmp48 := strconv.Itoa(*count)
		values.Set("_count", tmp48)
	}
	if elements != nil {
		values.Set("_elements", *elements)
	}
	if has != nil {
		values.Set("_has", *has)
	}
	if id != nil {
		tmp49 := strconv.Itoa(*id)
		values.Set("_id", tmp49)
	}
	if include != nil {
		values.Set("_include", *include)
	}
	if lastUpdate != nil {
		tmp50 := lastUpdate.Format(time.RFC3339)
		values.Set("_lastUpdate", tmp50)
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
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
