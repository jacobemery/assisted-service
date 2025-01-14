// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V2ListEventsURL generates an URL for the v2 list events operation
type V2ListEventsURL struct {
	Categories   []string
	ClusterID    *strfmt.UUID
	ClusterLevel *bool
	DeletedHosts *bool
	HostID       *strfmt.UUID
	HostIds      []strfmt.UUID
	InfraEnvID   *strfmt.UUID
	Limit        *int64
	Message      *string
	Offset       *int64
	Severities   []string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *V2ListEventsURL) WithBasePath(bp string) *V2ListEventsURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *V2ListEventsURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *V2ListEventsURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/v2/events"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/api/assisted-install"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var categoriesIR []string
	for _, categoriesI := range o.Categories {
		categoriesIS := categoriesI
		if categoriesIS != "" {
			categoriesIR = append(categoriesIR, categoriesIS)
		}
	}

	categories := swag.JoinByFormat(categoriesIR, "")

	if len(categories) > 0 {
		qsv := categories[0]
		if qsv != "" {
			qs.Set("categories", qsv)
		}
	}

	var clusterIDQ string
	if o.ClusterID != nil {
		clusterIDQ = o.ClusterID.String()
	}
	if clusterIDQ != "" {
		qs.Set("cluster_id", clusterIDQ)
	}

	var clusterLevelQ string
	if o.ClusterLevel != nil {
		clusterLevelQ = swag.FormatBool(*o.ClusterLevel)
	}
	if clusterLevelQ != "" {
		qs.Set("cluster_level", clusterLevelQ)
	}

	var deletedHostsQ string
	if o.DeletedHosts != nil {
		deletedHostsQ = swag.FormatBool(*o.DeletedHosts)
	}
	if deletedHostsQ != "" {
		qs.Set("deleted_hosts", deletedHostsQ)
	}

	var hostIDQ string
	if o.HostID != nil {
		hostIDQ = o.HostID.String()
	}
	if hostIDQ != "" {
		qs.Set("host_id", hostIDQ)
	}

	var hostIdsIR []string
	for _, hostIdsI := range o.HostIds {
		hostIdsIS := hostIdsI.String()
		if hostIdsIS != "" {
			hostIdsIR = append(hostIdsIR, hostIdsIS)
		}
	}

	hostIds := swag.JoinByFormat(hostIdsIR, "")

	if len(hostIds) > 0 {
		qsv := hostIds[0]
		if qsv != "" {
			qs.Set("host_ids", qsv)
		}
	}

	var infraEnvIDQ string
	if o.InfraEnvID != nil {
		infraEnvIDQ = o.InfraEnvID.String()
	}
	if infraEnvIDQ != "" {
		qs.Set("infra_env_id", infraEnvIDQ)
	}

	var limitQ string
	if o.Limit != nil {
		limitQ = swag.FormatInt64(*o.Limit)
	}
	if limitQ != "" {
		qs.Set("limit", limitQ)
	}

	var messageQ string
	if o.Message != nil {
		messageQ = *o.Message
	}
	if messageQ != "" {
		qs.Set("message", messageQ)
	}

	var offsetQ string
	if o.Offset != nil {
		offsetQ = swag.FormatInt64(*o.Offset)
	}
	if offsetQ != "" {
		qs.Set("offset", offsetQ)
	}

	var severitiesIR []string
	for _, severitiesI := range o.Severities {
		severitiesIS := severitiesI
		if severitiesIS != "" {
			severitiesIR = append(severitiesIR, severitiesIS)
		}
	}

	severities := swag.JoinByFormat(severitiesIR, "")

	if len(severities) > 0 {
		qsv := severities[0]
		if qsv != "" {
			qs.Set("severities", qsv)
		}
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *V2ListEventsURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *V2ListEventsURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *V2ListEventsURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on V2ListEventsURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on V2ListEventsURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *V2ListEventsURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
