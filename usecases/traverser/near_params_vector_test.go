//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2021 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

package traverser

import (
	"context"
	"reflect"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/semi-technologies/weaviate/entities/additional"
	"github.com/semi-technologies/weaviate/entities/near"
	"github.com/semi-technologies/weaviate/entities/search"
)

func Test_nearParamsVector_validateNearParams(t *testing.T) {
	type args struct {
		nearVector   *near.NearVectorParams
		nearObject   *near.NearObjectParams
		moduleParams map[string]interface{}
		className    []string
	}
	tests := []struct {
		name       string
		args       args
		wantErr    bool
		errMessage string
	}{
		{
			name: "Should be OK, when all near params are nil",
			args: args{
				nearVector:   nil,
				nearObject:   nil,
				moduleParams: nil,
				className:    nil,
			},
			wantErr: false,
		},
		{
			name: "Should be OK, when nearVector param is set",
			args: args{
				nearVector:   &near.NearVectorParams{},
				nearObject:   nil,
				moduleParams: nil,
				className:    nil,
			},
			wantErr: false,
		},
		{
			name: "Should be OK, when nearObject param is set",
			args: args{
				nearVector:   nil,
				nearObject:   &near.NearObjectParams{},
				moduleParams: nil,
				className:    nil,
			},
			wantErr: false,
		},
		{
			name: "Should be OK, when moduleParams param is set",
			args: args{
				nearVector: nil,
				nearObject: nil,
				moduleParams: map[string]interface{}{
					"nearCustomText": &nearCustomTextParams{},
				},
				className: nil,
			},
			wantErr: false,
		},
		{
			name: "Should throw error, when nearVector and nearObject is set",
			args: args{
				nearVector:   &near.NearVectorParams{},
				nearObject:   &near.NearObjectParams{},
				moduleParams: nil,
				className:    nil,
			},
			wantErr:    true,
			errMessage: "found both 'nearVector' and 'nearObject' parameters which are conflicting, choose one instead",
		},
		{
			name: "Should throw error, when nearVector and moduleParams is set",
			args: args{
				nearVector: &near.NearVectorParams{},
				nearObject: nil,
				moduleParams: map[string]interface{}{
					"nearCustomText": &nearCustomTextParams{},
				},
				className: nil,
			},
			wantErr:    true,
			errMessage: "found both 'nearText' and 'nearVector' parameters which are conflicting, choose one instead",
		},
		{
			name: "Should throw error, when nearObject and moduleParams is set",
			args: args{
				nearVector: nil,
				nearObject: &near.NearObjectParams{},
				moduleParams: map[string]interface{}{
					"nearCustomText": &nearCustomTextParams{},
				},
				className: nil,
			},
			wantErr:    true,
			errMessage: "found both 'nearText' and 'nearObject' parameters which are conflicting, choose one instead",
		},
		{
			name: "Should throw error, when nearVector and nearObject and moduleParams is set",
			args: args{
				nearVector: &near.NearVectorParams{},
				nearObject: &near.NearObjectParams{},
				moduleParams: map[string]interface{}{
					"nearCustomText": &nearCustomTextParams{},
				},
				className: nil,
			},
			wantErr:    true,
			errMessage: "found 'nearText' and 'nearVector' and 'nearObject' parameters which are conflicting, choose one instead",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &nearParamsVector{
				modulesProvider: &fakeModulesProvider{},
				search:          &fakeNearParamsSearcher{},
			}
			err := e.validateNearParams(tt.args.nearVector, tt.args.nearObject, tt.args.moduleParams, tt.args.className...)
			if (err != nil) != tt.wantErr {
				t.Errorf("nearParamsVector.validateNearParams() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil && tt.errMessage != err.Error() {
				t.Errorf("nearParamsVector.validateNearParams() error = %v, errMessage = %v", err, tt.errMessage)
			}
		})
	}
}

func Test_nearParamsVector_vectorFromParams(t *testing.T) {
	type args struct {
		ctx          context.Context
		nearVector   *near.NearVectorParams
		nearObject   *near.NearObjectParams
		moduleParams map[string]interface{}
		className    string
	}
	tests := []struct {
		name    string
		args    args
		want    []float32
		wantErr bool
	}{
		{
			name: "Should get vector from nearVector",
			args: args{
				nearVector: &near.NearVectorParams{
					Vector: []float32{1.1, 1.0, 0.1},
				},
			},
			want:    []float32{1.1, 1.0, 0.1},
			wantErr: false,
		},
		{
			name: "Should get vector from nearObject",
			args: args{
				nearObject: &near.NearObjectParams{
					ID: "uuid",
				},
			},
			want:    []float32{1.0, 1.0, 1.0},
			wantErr: false,
		},
		{
			name: "Should get vector from nearText",
			args: args{
				moduleParams: map[string]interface{}{
					"nearCustomText": &nearCustomTextParams{
						Values: []string{"a"},
					},
				},
			},
			want:    []float32{1, 2, 3},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &nearParamsVector{
				modulesProvider: &fakeModulesProvider{},
				search:          &fakeNearParamsSearcher{},
			}
			got, err := e.vectorFromParams(tt.args.ctx, tt.args.nearVector, tt.args.nearObject, tt.args.moduleParams, tt.args.className)
			if (err != nil) != tt.wantErr {
				t.Errorf("nearParamsVector.vectorFromParams() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nearParamsVector.vectorFromParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nearParamsVector_extractCertaintyFromParams(t *testing.T) {
	type args struct {
		nearVector   *near.NearVectorParams
		nearObject   *near.NearObjectParams
		moduleParams map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Should extract certainty from nearVector",
			args: args{
				nearVector: &near.NearVectorParams{
					Certainty: 0.88,
				},
			},
			want: 0.88,
		},
		{
			name: "Should extract certainty from nearObject",
			args: args{
				nearObject: &near.NearObjectParams{
					Certainty: 0.99,
				},
			},
			want: 0.99,
		},
		{
			name: "Should extract certainty from nearText",
			args: args{
				moduleParams: map[string]interface{}{
					"nearCustomText": &nearCustomTextParams{
						Certainty: 0.77,
					},
				},
			},
			want: 0.77,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &nearParamsVector{
				modulesProvider: &fakeModulesProvider{},
				search:          &fakeNearParamsSearcher{},
			}
			if got := e.extractCertaintyFromParams(tt.args.nearVector, tt.args.nearObject, tt.args.moduleParams); got != tt.want {
				t.Errorf("nearParamsVector.extractCertaintyFromParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

type fakeNearParamsSearcher struct{}

func (f *fakeNearParamsSearcher) ObjectByID(ctx context.Context, id strfmt.UUID,
	props search.SelectProperties, additional additional.Properties) (*search.Result, error) {
	return &search.Result{
		Vector: []float32{1.0, 1.0, 1.0},
	}, nil
}