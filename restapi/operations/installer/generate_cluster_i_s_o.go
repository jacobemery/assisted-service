// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GenerateClusterISOHandlerFunc turns a function with the right signature into a generate cluster i s o handler
type GenerateClusterISOHandlerFunc func(GenerateClusterISOParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GenerateClusterISOHandlerFunc) Handle(params GenerateClusterISOParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GenerateClusterISOHandler interface for that can handle valid generate cluster i s o params
type GenerateClusterISOHandler interface {
	Handle(GenerateClusterISOParams, interface{}) middleware.Responder
}

// NewGenerateClusterISO creates a new http.Handler for the generate cluster i s o operation
func NewGenerateClusterISO(ctx *middleware.Context, handler GenerateClusterISOHandler) *GenerateClusterISO {
	return &GenerateClusterISO{Context: ctx, Handler: handler}
}

/*GenerateClusterISO swagger:route POST /v1/clusters/{cluster_id}/downloads/image installer generateClusterISO

Creates a new OpenShift per-cluster Discovery ISO.

*/
type GenerateClusterISO struct {
	Context *middleware.Context
	Handler GenerateClusterISOHandler
}

func (o *GenerateClusterISO) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGenerateClusterISOParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}