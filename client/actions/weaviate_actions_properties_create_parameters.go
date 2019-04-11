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

package actions

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

	models "github.com/creativesoftwarefdn/weaviate/models"
)

// NewWeaviateActionsPropertiesCreateParams creates a new WeaviateActionsPropertiesCreateParams object
// with the default values initialized.
func NewWeaviateActionsPropertiesCreateParams() *WeaviateActionsPropertiesCreateParams {
	var ()
	return &WeaviateActionsPropertiesCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWeaviateActionsPropertiesCreateParamsWithTimeout creates a new WeaviateActionsPropertiesCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWeaviateActionsPropertiesCreateParamsWithTimeout(timeout time.Duration) *WeaviateActionsPropertiesCreateParams {
	var ()
	return &WeaviateActionsPropertiesCreateParams{

		timeout: timeout,
	}
}

// NewWeaviateActionsPropertiesCreateParamsWithContext creates a new WeaviateActionsPropertiesCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewWeaviateActionsPropertiesCreateParamsWithContext(ctx context.Context) *WeaviateActionsPropertiesCreateParams {
	var ()
	return &WeaviateActionsPropertiesCreateParams{

		Context: ctx,
	}
}

// NewWeaviateActionsPropertiesCreateParamsWithHTTPClient creates a new WeaviateActionsPropertiesCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWeaviateActionsPropertiesCreateParamsWithHTTPClient(client *http.Client) *WeaviateActionsPropertiesCreateParams {
	var ()
	return &WeaviateActionsPropertiesCreateParams{
		HTTPClient: client,
	}
}

/*WeaviateActionsPropertiesCreateParams contains all the parameters to send to the API endpoint
for the weaviate actions properties create operation typically these are written to a http.Request
*/
type WeaviateActionsPropertiesCreateParams struct {

	/*ActionID
	  Unique ID of the Action.

	*/
	ActionID strfmt.UUID
	/*Body*/
	Body *models.SingleRef
	/*PropertyName
	  Unique name of the property related to the Action.

	*/
	PropertyName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the weaviate actions properties create params
func (o *WeaviateActionsPropertiesCreateParams) WithTimeout(timeout time.Duration) *WeaviateActionsPropertiesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the weaviate actions properties create params
func (o *WeaviateActionsPropertiesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the weaviate actions properties create params
func (o *WeaviateActionsPropertiesCreateParams) WithContext(ctx context.Context) *WeaviateActionsPropertiesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the weaviate actions properties create params
func (o *WeaviateActionsPropertiesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the weaviate actions properties create params
func (o *WeaviateActionsPropertiesCreateParams) WithHTTPClient(client *http.Client) *WeaviateActionsPropertiesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the weaviate actions properties create params
func (o *WeaviateActionsPropertiesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActionID adds the actionID to the weaviate actions properties create params
func (o *WeaviateActionsPropertiesCreateParams) WithActionID(actionID strfmt.UUID) *WeaviateActionsPropertiesCreateParams {
	o.SetActionID(actionID)
	return o
}

// SetActionID adds the actionId to the weaviate actions properties create params
func (o *WeaviateActionsPropertiesCreateParams) SetActionID(actionID strfmt.UUID) {
	o.ActionID = actionID
}

// WithBody adds the body to the weaviate actions properties create params
func (o *WeaviateActionsPropertiesCreateParams) WithBody(body *models.SingleRef) *WeaviateActionsPropertiesCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the weaviate actions properties create params
func (o *WeaviateActionsPropertiesCreateParams) SetBody(body *models.SingleRef) {
	o.Body = body
}

// WithPropertyName adds the propertyName to the weaviate actions properties create params
func (o *WeaviateActionsPropertiesCreateParams) WithPropertyName(propertyName string) *WeaviateActionsPropertiesCreateParams {
	o.SetPropertyName(propertyName)
	return o
}

// SetPropertyName adds the propertyName to the weaviate actions properties create params
func (o *WeaviateActionsPropertiesCreateParams) SetPropertyName(propertyName string) {
	o.PropertyName = propertyName
}

// WriteToRequest writes these params to a swagger request
func (o *WeaviateActionsPropertiesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param actionId
	if err := r.SetPathParam("actionId", o.ActionID.String()); err != nil {
		return err
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param propertyName
	if err := r.SetPathParam("propertyName", o.PropertyName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}