// Package adminroles provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.1 DO NOT EDIT.
package adminroles

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	externalRef0 "github.com/cam-inc/viron/packages/golang/routes/components"
	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
)

const (
	JwtScopes = "jwt.Scopes"
)

// Defines values for VironAdminRolePermissionPermission.
const (
	VironAdminRolePermissionPermissionDeny VironAdminRolePermissionPermission = "deny"

	VironAdminRolePermissionPermissionRead VironAdminRolePermissionPermission = "read"

	VironAdminRolePermissionPermissionWrite VironAdminRolePermissionPermission = "write"
)

// VironAdminRole defines model for VironAdminRole.
type VironAdminRole struct {

	// ロールID
	Id string `json:"id"`

	// 権限
	Permissions []VironAdminRolePermission `json:"permissions"`
}

// VironAdminRoleCreatePayload defines model for VironAdminRoleCreatePayload.
type VironAdminRoleCreatePayload struct {

	// ロールID
	Id string `json:"id"`

	// 権限
	Permissions []VironAdminRolePermission `json:"permissions"`
}

// VironAdminRoleList defines model for VironAdminRoleList.
type VironAdminRoleList []VironAdminRole

// VironAdminRoleListWithPager defines model for VironAdminRoleListWithPager.
type VironAdminRoleListWithPager struct {
	// Embedded struct due to allOf(./components.yaml#/components/schemas/VironPager)
	externalRef0.VironPager `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	List VironAdminRoleList `json:"list"`
}

// VironAdminRolePermission defines model for VironAdminRolePermission.
type VironAdminRolePermission struct {
	Permission VironAdminRolePermissionPermission `json:"permission"`
	ResourceId string                             `json:"resourceId"`
}

// VironAdminRolePermissionPermission defines model for VironAdminRolePermission.Permission.
type VironAdminRolePermissionPermission string

// VironAdminRoleUpdatePayload defines model for VironAdminRoleUpdatePayload.
type VironAdminRoleUpdatePayload struct {

	// 権限
	Permissions []VironAdminRolePermission `json:"permissions"`
}

// CreateVironAdminRoleJSONBody defines parameters for CreateVironAdminRole.
type CreateVironAdminRoleJSONBody VironAdminRoleCreatePayload

// UpdateVironAdminRoleJSONBody defines parameters for UpdateVironAdminRole.
type UpdateVironAdminRoleJSONBody VironAdminRoleUpdatePayload

// CreateVironAdminRoleJSONRequestBody defines body for CreateVironAdminRole for application/json ContentType.
type CreateVironAdminRoleJSONRequestBody CreateVironAdminRoleJSONBody

// UpdateVironAdminRoleJSONRequestBody defines body for UpdateVironAdminRole for application/json ContentType.
type UpdateVironAdminRoleJSONRequestBody UpdateVironAdminRoleJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// list admin roles
	// (GET /viron/adminroles)
	ListVironAdminRoles(w http.ResponseWriter, r *http.Request)
	// create an admin role
	// (POST /viron/adminroles)
	CreateVironAdminRole(w http.ResponseWriter, r *http.Request)
	// delete an admin role
	// (DELETE /viron/adminroles/{id})
	RemoveVironAdminRole(w http.ResponseWriter, r *http.Request, id externalRef0.VironIdPathParam)
	// update an admin role
	// (PUT /viron/adminroles/{id})
	UpdateVironAdminRole(w http.ResponseWriter, r *http.Request, id externalRef0.VironIdPathParam)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
}

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

// ListVironAdminRoles operation middleware
func (siw *ServerInterfaceWrapper) ListVironAdminRoles(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, JwtScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ListVironAdminRoles(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// CreateVironAdminRole operation middleware
func (siw *ServerInterfaceWrapper) CreateVironAdminRole(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, JwtScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateVironAdminRole(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// RemoveVironAdminRole operation middleware
func (siw *ServerInterfaceWrapper) RemoveVironAdminRole(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id externalRef0.VironIdPathParam

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter id: %s", err), http.StatusBadRequest)
		return
	}

	ctx = context.WithValue(ctx, JwtScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.RemoveVironAdminRole(w, r, id)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// UpdateVironAdminRole operation middleware
func (siw *ServerInterfaceWrapper) UpdateVironAdminRole(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id externalRef0.VironIdPathParam

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter id: %s", err), http.StatusBadRequest)
		return
	}

	ctx = context.WithValue(ctx, JwtScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateVironAdminRole(w, r, id)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL     string
	BaseRouter  chi.Router
	Middlewares []MiddlewareFunc
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/viron/adminroles", wrapper.ListVironAdminRoles)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/viron/adminroles", wrapper.CreateVironAdminRole)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/viron/adminroles/{id}", wrapper.RemoveVironAdminRole)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/viron/adminroles/{id}", wrapper.UpdateVironAdminRole)
	})

	return r
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+RVXWsbRxT9K+K2j1ut7PZpn/r1YmqwMNR9EHoY717JY3ZnpjOzchexUNVQ2kLBlGIw",
	"SR4SQpwQEkLIQyAff0aJHP2LMDNC2g/FljEOgbxIK83Oveeec+69Qwh5IjhDphUEQ1DhHibEPu5Qydl3",
	"UULZNo/R/CMkFyg1RXtOI/MZoQolFZpyBgGMDx+ND1+ODx9u/Age6EwgBKC0pKwPuQcCZUKVopyp+t3J",
	"/QfTkyPwgGpM7PmXEnsQwBf+AqM/A+iX0bXngU2aWV4iJckgzz2Q+GtKJUYQdAzsMpDu/ALf3cdQmwjl",
	"6D9IJBrbJIs5iT5nIjap0rbiywOrw1kW/Beq99qkj9IEJ3G81YOgM0/TLKRpZiSJP5zXBcm9qlbxrILV",
	"gduaq9zZMHW+urWiCnrUfCNKZ8jSxISWSIwwB5JqBA8iZFkh08JCEhVPZYgb1n2V4wrcwrtFyVdQ/GcR",
	"nWf9T8HH51vYvE1Zj9fhEZOrIXmMqkEEbSiBIe3RkGgbygNNtZl78O3AQPNjutso3AEPBiidetBqtppr",
	"4MFvXwnSN9R0zGBl2o3VzhAMaTawkcvap1yvCSdInzLi0GmZYlljGJT7aV6pJrumvbpepb6zx7fPjv6c",
	"D6I3z39/d+90/Md/0+On5mF0Oh79Ox69Ho9OwIO+5KmYXzr7/8X01h3f/TAC1vMXCbKlNCr5XDdwgYwI",
	"CgF83Ww1122Ves+awXe0Wkodo8EQ+qjrSi2tBLwKqZtLSZWoBGfKRV9vtczXTBk7Y4SIZ5L7+8p1onPk",
	"5afEYnhZ15Vr2PrJGldhmEqqM2uK/QMNQadrpFNpkhCZzbxR8ZkmfeOiqgMMv4KrFQh7dXPy11GNMLfW",
	"dqq2Mr2FSn/Po+yayCrv07zc0Mb5eU23tWuCciWpQltHg7CCXufJlXt10/tDGuVOwRg1Xqjl27//mZ7c",
	"rWm5jQkf1LUURJIENUq1+h5d3HFcbURtYowtSWKnTEWZb+qQVybQ1bw6gR6I9GK7T248mxw/qVHkVtlH",
	"oui6W6i8l1dqoasIldp0l3B6Pj8cAiMJLtkfeTd/HwAA//8qHVu9fgwAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	pathPrefix := path.Dir(pathToFile)

	for rawPath, rawFunc := range externalRef0.PathToRawSpec(path.Join(pathPrefix, "./components.yaml")) {
		if _, ok := res[rawPath]; ok {
			// it is not possible to compare functions in golang, so always overwrite the old value
		}
		res[rawPath] = rawFunc
	}
	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
