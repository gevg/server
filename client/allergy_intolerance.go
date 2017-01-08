package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// ReadAllergyIntolerancePath computes a request path to the read action of AllergyIntolerance.
func ReadAllergyIntolerancePath(patientID int, allergy string) string {
	return fmt.Sprintf("/nosh/patients/%v/allergy.intolerance/%v.intoleranceID", patientID, allergy)
}

// Retrieve vital with given id
func (c *Client) ReadAllergyIntolerance(ctx context.Context, path string, contained *string, containedType *string, count *int, elements *string, has *string, id *int, include *string, lastUpdate *time.Time, list *string, profile *string, query *string, revinclude *string, security *string, sort *string, summary *string, tag *string, text *string, type_ *string, allergyIntoleranceID *int) (*http.Response, error) {
	req, err := c.NewReadAllergyIntoleranceRequest(ctx, path, contained, containedType, count, elements, has, id, include, lastUpdate, list, profile, query, revinclude, security, sort, summary, tag, text, type_, allergyIntoleranceID)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewReadAllergyIntoleranceRequest create the request corresponding to the read action endpoint of the AllergyIntolerance resource.
func (c *Client) NewReadAllergyIntoleranceRequest(ctx context.Context, path string, contained *string, containedType *string, count *int, elements *string, has *string, id *int, include *string, lastUpdate *time.Time, list *string, profile *string, query *string, revinclude *string, security *string, sort *string, summary *string, tag *string, text *string, type_ *string, allergyIntoleranceID *int) (*http.Request, error) {
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
		tmp46 := strconv.Itoa(*count)
		values.Set("_count", tmp46)
	}
	if elements != nil {
		values.Set("_elements", *elements)
	}
	if has != nil {
		values.Set("_has", *has)
	}
	if id != nil {
		tmp47 := strconv.Itoa(*id)
		values.Set("_id", tmp47)
	}
	if include != nil {
		values.Set("_include", *include)
	}
	if lastUpdate != nil {
		tmp48 := lastUpdate.Format(time.RFC3339)
		values.Set("_lastUpdate", tmp48)
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
	if allergyIntoleranceID != nil {
		tmp49 := strconv.Itoa(*allergyIntoleranceID)
		values.Set("allergy.intoleranceID", tmp49)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
