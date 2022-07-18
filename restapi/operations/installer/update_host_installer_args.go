// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdateHostInstallerArgsHandlerFunc turns a function with the right signature into a update host installer args handler
type UpdateHostInstallerArgsHandlerFunc func(UpdateHostInstallerArgsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateHostInstallerArgsHandlerFunc) Handle(params UpdateHostInstallerArgsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// UpdateHostInstallerArgsHandler interface for that can handle valid update host installer args params
type UpdateHostInstallerArgsHandler interface {
	Handle(UpdateHostInstallerArgsParams, interface{}) middleware.Responder
}

// NewUpdateHostInstallerArgs creates a new http.Handler for the update host installer args operation
func NewUpdateHostInstallerArgs(ctx *middleware.Context, handler UpdateHostInstallerArgsHandler) *UpdateHostInstallerArgs {
	return &UpdateHostInstallerArgs{Context: ctx, Handler: handler}
}

/*UpdateHostInstallerArgs swagger:route PATCH /v1/clusters/{cluster_id}/hosts/{host_id}/installer-args installer updateHostInstallerArgs

Updates a host's installer arguments.

*/
type UpdateHostInstallerArgs struct {
	Context *middleware.Context
	Handler UpdateHostInstallerArgsHandler
}

func (o *UpdateHostInstallerArgs) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateHostInstallerArgsParams()

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