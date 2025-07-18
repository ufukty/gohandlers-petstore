// Code generated by gohandlers v0.37.0. DO NOT EDIT.

package pets

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/ufukty/gohandlers/pkg/gohandlers"
)

func (pe *Pets) ListHandlers() map[string]gohandlers.HandlerInfo {
	return map[string]gohandlers.HandlerInfo{
		"Create": {Method: "POST", Path: "/create", Ref: pe.Create},
		"Delete": {Method: "DELETE", Path: "/pets/{id}", Ref: pe.Delete},
		"Get":    {Method: "GET", Path: "/pets/{id}", Ref: pe.Get},
		"List":   {Method: "GET", Path: "/pets", Ref: pe.List},
	}
}

func join(segments ...string) string {
	url := ""
	for i, segment := range segments {
		if i != 0 && !strings.HasPrefix(segment, "/") {
			url += "/"
		}
		url += segment
	}
	return url
}

func (bq CreateRequest) Build(host string) (*http.Request, error) {
	uri := "/create"
	body := bytes.NewBuffer([]byte{})
	if err := json.NewEncoder(body).Encode(bq); err != nil {
		return nil, fmt.Errorf("json.Encoder.Encode: %w", err)
	}
	r, err := http.NewRequest("POST", join(host, uri), body)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest: %w", err)
	}
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Content-Length", fmt.Sprintf("%d", body.Len()))
	return r, nil
}

func (bq *CreateRequest) Parse(rq *http.Request) error {
	if !strings.HasPrefix(rq.Header.Get("Content-Type"), "application/json") {
		return fmt.Errorf("invalid content type for request: %s", rq.Header.Get("Content-Type"))
	}
	if err := json.NewDecoder(rq.Body).Decode(bq); err != nil {
		return fmt.Errorf("decoding body: %w", err)
	}
	return nil
}

func (bq CreateRequest) Validate() (issues map[string]any) {
	issues = map[string]any{}
	if issue := bq.Name.Validate(); issue != nil {
		issues["name"] = issue
	}
	if issue := bq.Tag.Validate(); issue != nil {
		issues["tag"] = issue
	}
	return
}

func (bs CreateResponse) Write(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(bs); err != nil {
		return fmt.Errorf("encoding the body: %w", err)
	}
	return nil
}

func (bs *CreateResponse) Parse(rs *http.Response) error {
	if !strings.HasPrefix(rs.Header.Get("Content-Type"), "application/json") {
		return fmt.Errorf("invalid content type for request: %s", rs.Header.Get("Content-Type"))
	}
	if err := json.NewDecoder(rs.Body).Decode(bs); err != nil {
		return fmt.Errorf("decoding the body: %w", err)
	}
	return nil
}

func (bq DeleteRequest) Build(host string) (*http.Request, error) {
	uri := "/pets/{id}"
	encoded, err := bq.ID.ToRoute()
	if err != nil {
		return nil, fmt.Errorf("DeleteRequest.ID.ToRoute: %w", err)
	}
	uri = strings.Replace(uri, "{id}", encoded, 1)
	r, err := http.NewRequest("DELETE", join(host, uri), nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest: %w", err)
	}
	return r, nil
}

func (bq *DeleteRequest) Parse(rq *http.Request) error {
	if err := bq.ID.FromRoute(rq.PathValue("id")); err != nil {
		return fmt.Errorf("DeleteRequest.ID.FromRoute: %w", err)
	}
	return nil
}

func (bq DeleteRequest) Validate() (issues map[string]any) {
	issues = map[string]any{}
	if issue := bq.ID.Validate(); issue != nil {
		issues["id"] = issue
	}
	return
}

func (bq GetRequest) Build(host string) (*http.Request, error) {
	uri := "/pets/{id}"
	encoded, err := bq.ID.ToRoute()
	if err != nil {
		return nil, fmt.Errorf("GetRequest.ID.ToRoute: %w", err)
	}
	uri = strings.Replace(uri, "{id}", encoded, 1)
	r, err := http.NewRequest("GET", join(host, uri), nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest: %w", err)
	}
	return r, nil
}

func (bq *GetRequest) Parse(rq *http.Request) error {
	if err := bq.ID.FromRoute(rq.PathValue("id")); err != nil {
		return fmt.Errorf("GetRequest.ID.FromRoute: %w", err)
	}
	return nil
}

func (bq GetRequest) Validate() (issues map[string]any) {
	issues = map[string]any{}
	if issue := bq.ID.Validate(); issue != nil {
		issues["id"] = issue
	}
	return
}

func (bs GetResponse) Write(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(bs); err != nil {
		return fmt.Errorf("encoding the body: %w", err)
	}
	return nil
}

func (bs *GetResponse) Parse(rs *http.Response) error {
	if !strings.HasPrefix(rs.Header.Get("Content-Type"), "application/json") {
		return fmt.Errorf("invalid content type for request: %s", rs.Header.Get("Content-Type"))
	}
	if err := json.NewDecoder(rs.Body).Decode(bs); err != nil {
		return fmt.Errorf("decoding the body: %w", err)
	}
	return nil
}

func (bq ListRequest) Build(host string) (*http.Request, error) {
	uri := "/pets"
	q := []string{}
	encoded, ok, err := bq.Limit.ToQuery()
	if err != nil {
		return nil, fmt.Errorf("ListRequest.Limit.ToQuery: %w", err)
	}
	if ok {
		q = append(q, fmt.Sprintf("limit=%s", encoded))
	}
	if len(q) > 0 {
		uri = fmt.Sprintf("%s?%s", uri, strings.Join(q, "&"))
	}
	r, err := http.NewRequest("GET", join(host, uri), nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest: %w", err)
	}
	return r, nil
}

func (bq *ListRequest) Parse(rq *http.Request) error {
	q := rq.URL.Query()
	if q.Has("limit") {
		err := bq.Limit.FromQuery(q.Get("limit"))
		if err != nil {
			return fmt.Errorf("ListRequest.Limit.FromQuery: %w", err)
		}
	}
	return nil
}

func (bq ListRequest) Validate() (issues map[string]any) {
	issues = map[string]any{}
	if issue := bq.Limit.Validate(); issue != nil {
		issues["limit"] = issue
	}
	return
}

func (bs ListResponse) Write(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(bs); err != nil {
		return fmt.Errorf("encoding the body: %w", err)
	}
	return nil
}

func (bs *ListResponse) Parse(rs *http.Response) error {
	if !strings.HasPrefix(rs.Header.Get("Content-Type"), "application/json") {
		return fmt.Errorf("invalid content type for request: %s", rs.Header.Get("Content-Type"))
	}
	if err := json.NewDecoder(rs.Body).Decode(bs); err != nil {
		return fmt.Errorf("decoding the body: %w", err)
	}
	return nil
}
