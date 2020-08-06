// Copyright 2020 Coinbase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by mockery v1.0.0. DO NOT EDIT.

package syncer

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	types "github.com/coinbase/rosetta-sdk-go/types"
)

// Helper is an autogenerated mock type for the Helper type
type Helper struct {
	mock.Mock
}

// Block provides a mock function with given fields: _a0, _a1, _a2
func (_m *Helper) Block(_a0 context.Context, _a1 *types.NetworkIdentifier, _a2 *types.PartialBlockIdentifier) (*types.Block, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *types.Block
	if rf, ok := ret.Get(0).(func(context.Context, *types.NetworkIdentifier, *types.PartialBlockIdentifier) *types.Block); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Block)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.NetworkIdentifier, *types.PartialBlockIdentifier) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NetworkStatus provides a mock function with given fields: _a0, _a1
func (_m *Helper) NetworkStatus(_a0 context.Context, _a1 *types.NetworkIdentifier) (*types.NetworkStatusResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.NetworkStatusResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.NetworkIdentifier) *types.NetworkStatusResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.NetworkStatusResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.NetworkIdentifier) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
