/*
OKX v5 API

OpenAPI specification for Okx exchange - Rest API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rest

import (
	"fmt"
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// EthStakingAPIService EthStakingAPI service
type EthStakingAPIService service

type ApiCreateFinanceStakingDefiEthPurchaseV5Request struct {
	ctx context.Context
	ApiService *EthStakingAPIService
	createFinanceStakingDefiEthPurchaseV5Req *CreateFinanceStakingDefiEthPurchaseV5Req
}

// The request body for CreateFinanceStakingDefiEthPurchaseV5
func (r ApiCreateFinanceStakingDefiEthPurchaseV5Request) CreateFinanceStakingDefiEthPurchaseV5Req(createFinanceStakingDefiEthPurchaseV5Req CreateFinanceStakingDefiEthPurchaseV5Req) ApiCreateFinanceStakingDefiEthPurchaseV5Request {
	r.createFinanceStakingDefiEthPurchaseV5Req = &createFinanceStakingDefiEthPurchaseV5Req
	return r
}

func (r ApiCreateFinanceStakingDefiEthPurchaseV5Request) Execute() (*CreateFinanceStakingDefiEthPurchaseV5Resp, *http.Response, error) {
	return r.ApiService.CreateFinanceStakingDefiEthPurchaseV5Execute(r)
}

/*
CreateFinanceStakingDefiEthPurchaseV5 POST / Purchase

Staking ETH for BETH


Only the assets in the funding account can be used.



#### Rate Limit: 2 requests per second 

#### Rate limit rule: User ID 

#### Permission: Trade 



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateFinanceStakingDefiEthPurchaseV5Request
*/
func (a *EthStakingAPIService) CreateFinanceStakingDefiEthPurchaseV5(ctx context.Context) ApiCreateFinanceStakingDefiEthPurchaseV5Request {
	return ApiCreateFinanceStakingDefiEthPurchaseV5Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CreateFinanceStakingDefiEthPurchaseV5Resp
func (a *EthStakingAPIService) CreateFinanceStakingDefiEthPurchaseV5Execute(r ApiCreateFinanceStakingDefiEthPurchaseV5Request) (*CreateFinanceStakingDefiEthPurchaseV5Resp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateFinanceStakingDefiEthPurchaseV5Resp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EthStakingAPIService.CreateFinanceStakingDefiEthPurchaseV5")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v5/finance/staking-defi/eth/purchase"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createFinanceStakingDefiEthPurchaseV5Req == nil {
		return localVarReturnValue, nil, reportError("createFinanceStakingDefiEthPurchaseV5Req is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createFinanceStakingDefiEthPurchaseV5Req
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextOKXAuth).(Auth); ok {
			localVarHeaderParams["OK-ACCESS-KEY"] = auth.APIKey
			localVarHeaderParams["OK-ACCESS-PASSPHRASE"] = auth.Passphrase
		}
	}
	
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode >= 400 && localVarHTTPResponse.StatusCode < 500 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = strings.TrimSpace(fmt.Sprintf("%s %s", localVarHTTPResponse.Status, *v.Msg))
			newErr.model = &v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode >= 500 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = strings.TrimSpace(fmt.Sprintf("%s %s", localVarHTTPResponse.Status, *v.Msg))
			newErr.model = &v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	if *localVarReturnValue.Code != "0" {
		var v *APIError = &APIError{
			Code: localVarReturnValue.Code,
			Msg: localVarReturnValue.Msg,
		}
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: *localVarReturnValue.Msg,
			model: v,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiCreateFinanceStakingDefiEthRedeemV5Request struct {
	ctx context.Context
	ApiService *EthStakingAPIService
	createFinanceStakingDefiEthRedeemV5Req *CreateFinanceStakingDefiEthRedeemV5Req
}

// The request body for CreateFinanceStakingDefiEthRedeemV5
func (r ApiCreateFinanceStakingDefiEthRedeemV5Request) CreateFinanceStakingDefiEthRedeemV5Req(createFinanceStakingDefiEthRedeemV5Req CreateFinanceStakingDefiEthRedeemV5Req) ApiCreateFinanceStakingDefiEthRedeemV5Request {
	r.createFinanceStakingDefiEthRedeemV5Req = &createFinanceStakingDefiEthRedeemV5Req
	return r
}

func (r ApiCreateFinanceStakingDefiEthRedeemV5Request) Execute() (*CreateFinanceStakingDefiEthRedeemV5Resp, *http.Response, error) {
	return r.ApiService.CreateFinanceStakingDefiEthRedeemV5Execute(r)
}

/*
CreateFinanceStakingDefiEthRedeemV5 POST / Redeem

Only the assets in the funding account can be used. If your BETH is in your trading account, you can make funding transfer first.



#### Rate Limit: 2 requests per second 

#### Rate limit rule: User ID 

#### Permission: Trade 



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateFinanceStakingDefiEthRedeemV5Request
*/
func (a *EthStakingAPIService) CreateFinanceStakingDefiEthRedeemV5(ctx context.Context) ApiCreateFinanceStakingDefiEthRedeemV5Request {
	return ApiCreateFinanceStakingDefiEthRedeemV5Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CreateFinanceStakingDefiEthRedeemV5Resp
func (a *EthStakingAPIService) CreateFinanceStakingDefiEthRedeemV5Execute(r ApiCreateFinanceStakingDefiEthRedeemV5Request) (*CreateFinanceStakingDefiEthRedeemV5Resp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateFinanceStakingDefiEthRedeemV5Resp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EthStakingAPIService.CreateFinanceStakingDefiEthRedeemV5")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v5/finance/staking-defi/eth/redeem"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createFinanceStakingDefiEthRedeemV5Req == nil {
		return localVarReturnValue, nil, reportError("createFinanceStakingDefiEthRedeemV5Req is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createFinanceStakingDefiEthRedeemV5Req
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextOKXAuth).(Auth); ok {
			localVarHeaderParams["OK-ACCESS-KEY"] = auth.APIKey
			localVarHeaderParams["OK-ACCESS-PASSPHRASE"] = auth.Passphrase
		}
	}
	
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode >= 400 && localVarHTTPResponse.StatusCode < 500 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = strings.TrimSpace(fmt.Sprintf("%s %s", localVarHTTPResponse.Status, *v.Msg))
			newErr.model = &v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode >= 500 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = strings.TrimSpace(fmt.Sprintf("%s %s", localVarHTTPResponse.Status, *v.Msg))
			newErr.model = &v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	if *localVarReturnValue.Code != "0" {
		var v *APIError = &APIError{
			Code: localVarReturnValue.Code,
			Msg: localVarReturnValue.Msg,
		}
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: *localVarReturnValue.Msg,
			model: v,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetFinanceStakingDefiEthApyHistoryV5Request struct {
	ctx context.Context
	ApiService *EthStakingAPIService
	days *string
}

// Get the days of APY(Annual percentage yield) history record in the past  No more than 365 days
func (r ApiGetFinanceStakingDefiEthApyHistoryV5Request) Days(days string) ApiGetFinanceStakingDefiEthApyHistoryV5Request {
	r.days = &days
	return r
}

func (r ApiGetFinanceStakingDefiEthApyHistoryV5Request) Execute() (*GetFinanceStakingDefiEthApyHistoryV5Resp, *http.Response, error) {
	return r.ApiService.GetFinanceStakingDefiEthApyHistoryV5Execute(r)
}

/*
GetFinanceStakingDefiEthApyHistoryV5 GET / APY history (Public)

Public endpoints don't need authorization.

#### Rate Limit: 6 requests per second 

#### Rate limit rule: IP 



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetFinanceStakingDefiEthApyHistoryV5Request
*/
func (a *EthStakingAPIService) GetFinanceStakingDefiEthApyHistoryV5(ctx context.Context) ApiGetFinanceStakingDefiEthApyHistoryV5Request {
	return ApiGetFinanceStakingDefiEthApyHistoryV5Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetFinanceStakingDefiEthApyHistoryV5Resp
func (a *EthStakingAPIService) GetFinanceStakingDefiEthApyHistoryV5Execute(r ApiGetFinanceStakingDefiEthApyHistoryV5Request) (*GetFinanceStakingDefiEthApyHistoryV5Resp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetFinanceStakingDefiEthApyHistoryV5Resp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EthStakingAPIService.GetFinanceStakingDefiEthApyHistoryV5")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v5/finance/staking-defi/eth/apy-history"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.days == nil {
		return localVarReturnValue, nil, reportError("days is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "days", r.days, "form", "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode >= 400 && localVarHTTPResponse.StatusCode < 500 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = strings.TrimSpace(fmt.Sprintf("%s %s", localVarHTTPResponse.Status, *v.Msg))
			newErr.model = &v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode >= 500 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = strings.TrimSpace(fmt.Sprintf("%s %s", localVarHTTPResponse.Status, *v.Msg))
			newErr.model = &v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	if *localVarReturnValue.Code != "0" {
		var v *APIError = &APIError{
			Code: localVarReturnValue.Code,
			Msg: localVarReturnValue.Msg,
		}
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: *localVarReturnValue.Msg,
			model: v,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetFinanceStakingDefiEthBalanceV5Request struct {
	ctx context.Context
	ApiService *EthStakingAPIService
}

func (r ApiGetFinanceStakingDefiEthBalanceV5Request) Execute() (*GetFinanceStakingDefiEthBalanceV5Resp, *http.Response, error) {
	return r.ApiService.GetFinanceStakingDefiEthBalanceV5Execute(r)
}

/*
GetFinanceStakingDefiEthBalanceV5 GET / Balance

The balance is a snapshot summarized all BETH assets (including assets in redeeming) in account.

#### Rate Limit: 6 requests per second 

#### Rate limit rule: User ID 

#### Permission: Read 



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetFinanceStakingDefiEthBalanceV5Request
*/
func (a *EthStakingAPIService) GetFinanceStakingDefiEthBalanceV5(ctx context.Context) ApiGetFinanceStakingDefiEthBalanceV5Request {
	return ApiGetFinanceStakingDefiEthBalanceV5Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetFinanceStakingDefiEthBalanceV5Resp
func (a *EthStakingAPIService) GetFinanceStakingDefiEthBalanceV5Execute(r ApiGetFinanceStakingDefiEthBalanceV5Request) (*GetFinanceStakingDefiEthBalanceV5Resp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetFinanceStakingDefiEthBalanceV5Resp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EthStakingAPIService.GetFinanceStakingDefiEthBalanceV5")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v5/finance/staking-defi/eth/balance"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextOKXAuth).(Auth); ok {
			localVarHeaderParams["OK-ACCESS-KEY"] = auth.APIKey
			localVarHeaderParams["OK-ACCESS-PASSPHRASE"] = auth.Passphrase
		}
	}
	
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode >= 400 && localVarHTTPResponse.StatusCode < 500 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = strings.TrimSpace(fmt.Sprintf("%s %s", localVarHTTPResponse.Status, *v.Msg))
			newErr.model = &v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode >= 500 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = strings.TrimSpace(fmt.Sprintf("%s %s", localVarHTTPResponse.Status, *v.Msg))
			newErr.model = &v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	if *localVarReturnValue.Code != "0" {
		var v *APIError = &APIError{
			Code: localVarReturnValue.Code,
			Msg: localVarReturnValue.Msg,
		}
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: *localVarReturnValue.Msg,
			model: v,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetFinanceStakingDefiEthProductInfoV5Request struct {
	ctx context.Context
	ApiService *EthStakingAPIService
}

func (r ApiGetFinanceStakingDefiEthProductInfoV5Request) Execute() (*GetFinanceStakingDefiEthProductInfoV5Resp, *http.Response, error) {
	return r.ApiService.GetFinanceStakingDefiEthProductInfoV5Execute(r)
}

/*
GetFinanceStakingDefiEthProductInfoV5 GET / Product info

#### Rate Limit: 3 requests per second 

#### Rate limit rule: User ID 

#### Permission: Read 



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetFinanceStakingDefiEthProductInfoV5Request
*/
func (a *EthStakingAPIService) GetFinanceStakingDefiEthProductInfoV5(ctx context.Context) ApiGetFinanceStakingDefiEthProductInfoV5Request {
	return ApiGetFinanceStakingDefiEthProductInfoV5Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetFinanceStakingDefiEthProductInfoV5Resp
func (a *EthStakingAPIService) GetFinanceStakingDefiEthProductInfoV5Execute(r ApiGetFinanceStakingDefiEthProductInfoV5Request) (*GetFinanceStakingDefiEthProductInfoV5Resp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetFinanceStakingDefiEthProductInfoV5Resp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EthStakingAPIService.GetFinanceStakingDefiEthProductInfoV5")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v5/finance/staking-defi/eth/product-info"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextOKXAuth).(Auth); ok {
			localVarHeaderParams["OK-ACCESS-KEY"] = auth.APIKey
			localVarHeaderParams["OK-ACCESS-PASSPHRASE"] = auth.Passphrase
		}
	}
	
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode >= 400 && localVarHTTPResponse.StatusCode < 500 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = strings.TrimSpace(fmt.Sprintf("%s %s", localVarHTTPResponse.Status, *v.Msg))
			newErr.model = &v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode >= 500 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = strings.TrimSpace(fmt.Sprintf("%s %s", localVarHTTPResponse.Status, *v.Msg))
			newErr.model = &v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	if *localVarReturnValue.Code != "0" {
		var v *APIError = &APIError{
			Code: localVarReturnValue.Code,
			Msg: localVarReturnValue.Msg,
		}
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: *localVarReturnValue.Msg,
			model: v,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetFinanceStakingDefiEthPurchaseRedeemHistoryV5Request struct {
	ctx context.Context
	ApiService *EthStakingAPIService
	type_ *string
	status *string
	after *string
	before *string
	limit *string
}

// Type  &#x60;purchase&#x60;  &#x60;redeem&#x60;
func (r ApiGetFinanceStakingDefiEthPurchaseRedeemHistoryV5Request) Type_(type_ string) ApiGetFinanceStakingDefiEthPurchaseRedeemHistoryV5Request {
	r.type_ = &type_
	return r
}

// Status  &#x60;pending&#x60;  &#x60;success&#x60;  &#x60;failed&#x60;
func (r ApiGetFinanceStakingDefiEthPurchaseRedeemHistoryV5Request) Status(status string) ApiGetFinanceStakingDefiEthPurchaseRedeemHistoryV5Request {
	r.status = &status
	return r
}

// Pagination of data to return records earlier than the &#x60;requestTime&#x60;. The value passed is the corresponding &#x60;timestamp&#x60;
func (r ApiGetFinanceStakingDefiEthPurchaseRedeemHistoryV5Request) After(after string) ApiGetFinanceStakingDefiEthPurchaseRedeemHistoryV5Request {
	r.after = &after
	return r
}

// Pagination of data to return records newer than the &#x60;requestTime&#x60;. The value passed is the corresponding &#x60;timestamp&#x60;
func (r ApiGetFinanceStakingDefiEthPurchaseRedeemHistoryV5Request) Before(before string) ApiGetFinanceStakingDefiEthPurchaseRedeemHistoryV5Request {
	r.before = &before
	return r
}

// Number of results per request. The default is &#x60;100&#x60;. The maximum is &#x60;100&#x60;.
func (r ApiGetFinanceStakingDefiEthPurchaseRedeemHistoryV5Request) Limit(limit string) ApiGetFinanceStakingDefiEthPurchaseRedeemHistoryV5Request {
	r.limit = &limit
	return r
}

func (r ApiGetFinanceStakingDefiEthPurchaseRedeemHistoryV5Request) Execute() (*GetFinanceStakingDefiEthPurchaseRedeemHistoryV5Resp, *http.Response, error) {
	return r.ApiService.GetFinanceStakingDefiEthPurchaseRedeemHistoryV5Execute(r)
}

/*
GetFinanceStakingDefiEthPurchaseRedeemHistoryV5 GET / Purchase&Redeem history

#### Rate Limit: 6 requests per second 

#### Rate limit rule: User ID 

#### Permission: Read 



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetFinanceStakingDefiEthPurchaseRedeemHistoryV5Request
*/
func (a *EthStakingAPIService) GetFinanceStakingDefiEthPurchaseRedeemHistoryV5(ctx context.Context) ApiGetFinanceStakingDefiEthPurchaseRedeemHistoryV5Request {
	return ApiGetFinanceStakingDefiEthPurchaseRedeemHistoryV5Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetFinanceStakingDefiEthPurchaseRedeemHistoryV5Resp
func (a *EthStakingAPIService) GetFinanceStakingDefiEthPurchaseRedeemHistoryV5Execute(r ApiGetFinanceStakingDefiEthPurchaseRedeemHistoryV5Request) (*GetFinanceStakingDefiEthPurchaseRedeemHistoryV5Resp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetFinanceStakingDefiEthPurchaseRedeemHistoryV5Resp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EthStakingAPIService.GetFinanceStakingDefiEthPurchaseRedeemHistoryV5")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v5/finance/staking-defi/eth/purchase-redeem-history"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")
	} else {
		var defaultValue string = ""
		r.type_ = &defaultValue
	}
	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "form", "")
	} else {
		var defaultValue string = ""
		r.status = &defaultValue
	}
	if r.after != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "after", r.after, "form", "")
	} else {
		var defaultValue string = ""
		r.after = &defaultValue
	}
	if r.before != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "before", r.before, "form", "")
	} else {
		var defaultValue string = ""
		r.before = &defaultValue
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	} else {
		var defaultValue string = ""
		r.limit = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextOKXAuth).(Auth); ok {
			localVarHeaderParams["OK-ACCESS-KEY"] = auth.APIKey
			localVarHeaderParams["OK-ACCESS-PASSPHRASE"] = auth.Passphrase
		}
	}
	
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode >= 400 && localVarHTTPResponse.StatusCode < 500 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = strings.TrimSpace(fmt.Sprintf("%s %s", localVarHTTPResponse.Status, *v.Msg))
			newErr.model = &v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode >= 500 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = strings.TrimSpace(fmt.Sprintf("%s %s", localVarHTTPResponse.Status, *v.Msg))
			newErr.model = &v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	if *localVarReturnValue.Code != "0" {
		var v *APIError = &APIError{
			Code: localVarReturnValue.Code,
			Msg: localVarReturnValue.Msg,
		}
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: *localVarReturnValue.Msg,
			model: v,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
