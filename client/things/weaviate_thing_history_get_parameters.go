/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@creativesoftwarefdn.org
 */ // Code generated by go-swagger; DO NOT EDIT.

package things

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewWeaviateThingHistoryGetParams creates a new WeaviateThingHistoryGetParams object
// with the default values initialized.
func NewWeaviateThingHistoryGetParams() *WeaviateThingHistoryGetParams {
	var ()
	return &WeaviateThingHistoryGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWeaviateThingHistoryGetParamsWithTimeout creates a new WeaviateThingHistoryGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWeaviateThingHistoryGetParamsWithTimeout(timeout time.Duration) *WeaviateThingHistoryGetParams {
	var ()
	return &WeaviateThingHistoryGetParams{

		timeout: timeout,
	}
}

// NewWeaviateThingHistoryGetParamsWithContext creates a new WeaviateThingHistoryGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewWeaviateThingHistoryGetParamsWithContext(ctx context.Context) *WeaviateThingHistoryGetParams {
	var ()
	return &WeaviateThingHistoryGetParams{

		Context: ctx,
	}
}

// NewWeaviateThingHistoryGetParamsWithHTTPClient creates a new WeaviateThingHistoryGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWeaviateThingHistoryGetParamsWithHTTPClient(client *http.Client) *WeaviateThingHistoryGetParams {
	var ()
	return &WeaviateThingHistoryGetParams{
		HTTPClient: client,
	}
}

/*WeaviateThingHistoryGetParams contains all the parameters to send to the API endpoint
for the weaviate thing history get operation typically these are written to a http.Request
*/
type WeaviateThingHistoryGetParams struct {

	/*ThingID
	  Unique ID of the Thing.

	*/
	ThingID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the weaviate thing history get params
func (o *WeaviateThingHistoryGetParams) WithTimeout(timeout time.Duration) *WeaviateThingHistoryGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the weaviate thing history get params
func (o *WeaviateThingHistoryGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the weaviate thing history get params
func (o *WeaviateThingHistoryGetParams) WithContext(ctx context.Context) *WeaviateThingHistoryGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the weaviate thing history get params
func (o *WeaviateThingHistoryGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the weaviate thing history get params
func (o *WeaviateThingHistoryGetParams) WithHTTPClient(client *http.Client) *WeaviateThingHistoryGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the weaviate thing history get params
func (o *WeaviateThingHistoryGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithThingID adds the thingID to the weaviate thing history get params
func (o *WeaviateThingHistoryGetParams) WithThingID(thingID strfmt.UUID) *WeaviateThingHistoryGetParams {
	o.SetThingID(thingID)
	return o
}

// SetThingID adds the thingId to the weaviate thing history get params
func (o *WeaviateThingHistoryGetParams) SetThingID(thingID strfmt.UUID) {
	o.ThingID = thingID
}

// WriteToRequest writes these params to a swagger request
func (o *WeaviateThingHistoryGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param thingId
	if err := r.SetPathParam("thingId", o.ThingID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}