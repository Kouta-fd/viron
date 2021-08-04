package helpers

import (
	"net/http"
	"strings"

	"github.com/cam-inc/viron/packages/golang/constant"

	"github.com/getkin/kin-openapi/openapi3"
)

func ref(r string, org string, rep string) string {
	if strings.Contains(r, org) {
		return strings.Replace(r, org, rep, -1)
	}
	return r
}

func UpperCamelToLowerCamel(operationID string) string {
	if operationID == "" {
		return operationID
	}
	first := operationID[:1]
	lower := strings.ToLower(first)
	return strings.Replace(operationID, first, lower, 1)
}

func opeRef(ope *openapi3.Operation, org string, rep string) error {
	if ope.OperationID != "" {
		ope.OperationID = UpperCamelToLowerCamel(ope.OperationID)
	}
	if ope.RequestBody != nil {
		ope.RequestBody.Ref = ref(ope.RequestBody.Ref, org, rep)
	}
	for _, r := range ope.Responses {
		r.Ref = ref(r.Ref, org, rep)
	}
	for _, p := range ope.Parameters {
		p.Ref = ref(p.Ref, org, rep)
	}
	return nil
}

func Ref(docRoot *openapi3.T, org string, rep string) error {

	// paths
	for _, pathItem := range docRoot.Paths {
		pathItem.Ref = ref(pathItem.Ref, org, rep)
		if pathItem.Get != nil {
			if err := opeRef(pathItem.Get, org, rep); err != nil {
				return err
			}
		}
		if pathItem.Post != nil {
			if err := opeRef(pathItem.Post, org, rep); err != nil {
				return err
			}
		}
		if pathItem.Put != nil {
			if err := opeRef(pathItem.Put, org, rep); err != nil {
				return err
			}
		}
		if pathItem.Delete != nil {
			if err := opeRef(pathItem.Delete, org, rep); err != nil {
				return err
			}
		}
	}

	// components
	for _, param := range docRoot.Components.Parameters {
		param.Ref = ref(param.Ref, org, rep)
	}
	for _, schema := range docRoot.Components.Schemas {
		schema.Ref = ref(schema.Ref, org, rep)
		if schema.Value != nil {

			for _, oo := range schema.Value.OneOf {
				oo.Ref = ref(oo.Ref, org, rep)
			}
			for _, ao := range schema.Value.AllOf {
				ao.Ref = ref(ao.Ref, org, rep)
			}
			for _, any := range schema.Value.AnyOf {
				any.Ref = ref(any.Ref, org, rep)
			}
		}
	}
	for _, response := range docRoot.Components.Responses {
		response.Ref = ref(response.Ref, org, rep)
	}
	for _, link := range docRoot.Components.Links {
		link.Ref = ref(link.Ref, org, rep)
	}
	return nil
}

func MethodNameLower(method string) string {
	switch strings.ToLower(method) {
	case constant.API_METHOD_GET:
		return constant.API_METHOD_GET
	case constant.API_METHOD_POST:
		return constant.API_METHOD_POST
	case constant.API_METHOD_PUT:
		return constant.API_METHOD_PUT
	case constant.API_METHOD_DELETE:
		return constant.API_METHOD_DELETE
	}
	return ""
}

func MethodNameUpper(method string) string {
	switch strings.ToUpper(method) {
	case http.MethodGet:
		return http.MethodGet
	case http.MethodPost:
		return http.MethodPost
	case http.MethodPut:
		return http.MethodPut
	case http.MethodDelete:
		return http.MethodDelete
	}
	return ""

}

/*
func HasJWT(docRoot *openapi3.T) bool {
	for _, path := range docRoot.Paths {
		path.Get
	}
}
*/
