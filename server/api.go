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

// Generated by: OpenAPI Generator (https://openapi-generator.tech)

package server

import (
	"net/http"

	"github.com/coinbase/rosetta-sdk-go/types"
)

// AccountAPIRouter defines the required methods for binding the api requests to a responses for the
// AccountAPI
// The AccountAPIRouter implementation should parse necessary information from the http request,
// pass the data to a AccountAPIServicer to perform the required actions, then write the service
// results to the http response.
type AccountAPIRouter interface {
	AccountBalance(http.ResponseWriter, *http.Request)
}

// BlockAPIRouter defines the required methods for binding the api requests to a responses for the
// BlockAPI
// The BlockAPIRouter implementation should parse necessary information from the http request,
// pass the data to a BlockAPIServicer to perform the required actions, then write the service
// results to the http response.
type BlockAPIRouter interface {
	Block(http.ResponseWriter, *http.Request)
	BlockTransaction(http.ResponseWriter, *http.Request)
}

// ConstructionAPIRouter defines the required methods for binding the api requests to a responses
// for the ConstructionAPI
// The ConstructionAPIRouter implementation should parse necessary information from the http
// request,
// pass the data to a ConstructionAPIServicer to perform the required actions, then write the
// service results to the http response.
type ConstructionAPIRouter interface {
	TransactionConstruction(http.ResponseWriter, *http.Request)
	TransactionSubmit(http.ResponseWriter, *http.Request)
}

// MempoolAPIRouter defines the required methods for binding the api requests to a responses for the
// MempoolAPI
// The MempoolAPIRouter implementation should parse necessary information from the http request,
// pass the data to a MempoolAPIServicer to perform the required actions, then write the service
// results to the http response.
type MempoolAPIRouter interface {
	Mempool(http.ResponseWriter, *http.Request)
	MempoolTransaction(http.ResponseWriter, *http.Request)
}

// NetworkAPIRouter defines the required methods for binding the api requests to a responses for the
// NetworkAPI
// The NetworkAPIRouter implementation should parse necessary information from the http request,
// pass the data to a NetworkAPIServicer to perform the required actions, then write the service
// results to the http response.
type NetworkAPIRouter interface {
	NetworkStatus(http.ResponseWriter, *http.Request)
}

// AccountAPIServicer defines the api actions for the AccountAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type AccountAPIServicer interface {
	AccountBalance(*types.AccountBalanceRequest) (*types.AccountBalanceResponse, *types.Error)
}

// BlockAPIServicer defines the api actions for the BlockAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type BlockAPIServicer interface {
	Block(*types.BlockRequest) (*types.BlockResponse, *types.Error)
	BlockTransaction(*types.BlockTransactionRequest) (*types.BlockTransactionResponse, *types.Error)
}

// ConstructionAPIServicer defines the api actions for the ConstructionAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type ConstructionAPIServicer interface {
	TransactionConstruction(
		*types.TransactionConstructionRequest,
	) (*types.TransactionConstructionResponse, *types.Error)
	TransactionSubmit(
		*types.TransactionSubmitRequest,
	) (*types.TransactionSubmitResponse, *types.Error)
}

// MempoolAPIServicer defines the api actions for the MempoolAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type MempoolAPIServicer interface {
	Mempool(*types.MempoolRequest) (*types.MempoolResponse, *types.Error)
	MempoolTransaction(
		*types.MempoolTransactionRequest,
	) (*types.MempoolTransactionResponse, *types.Error)
}

// NetworkAPIServicer defines the api actions for the NetworkAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type NetworkAPIServicer interface {
	NetworkStatus(*types.NetworkStatusRequest) (*types.NetworkStatusResponse, *types.Error)
}