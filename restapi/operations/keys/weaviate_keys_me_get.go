/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2018 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * AUTHOR: Bob van Luijt (bob@kub.design)
 * See www.creativesoftwarefdn.org for details
 * Contact: @CreativeSofwFdn / bob@kub.design
 */
// Code generated by go-swagger; DO NOT EDIT.

package keys

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeaviateKeysMeGetHandlerFunc turns a function with the right signature into a weaviate keys me get handler
type WeaviateKeysMeGetHandlerFunc func(WeaviateKeysMeGetParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaviateKeysMeGetHandlerFunc) Handle(params WeaviateKeysMeGetParams) middleware.Responder {
	return fn(params)
}

// WeaviateKeysMeGetHandler interface for that can handle valid weaviate keys me get params
type WeaviateKeysMeGetHandler interface {
	Handle(WeaviateKeysMeGetParams) middleware.Responder
}

// NewWeaviateKeysMeGet creates a new http.Handler for the weaviate keys me get operation
func NewWeaviateKeysMeGet(ctx *middleware.Context, handler WeaviateKeysMeGetHandler) *WeaviateKeysMeGet {
	return &WeaviateKeysMeGet{Context: ctx, Handler: handler}
}

/*WeaviateKeysMeGet swagger:route GET /keys/me keys weaviateKeysMeGet

Get a key based on the key used to do the request.

Get the key-information of the key used.

*/
type WeaviateKeysMeGet struct {
	Context *middleware.Context
	Handler WeaviateKeysMeGetHandler
}

func (o *WeaviateKeysMeGet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWeaviateKeysMeGetParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
