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


// OnChainEarnAPIService OnChainEarnAPI service
type OnChainEarnAPIService service

type ApiCreateFinanceStakingDefiCancelV5Request struct {
	ctx context.Context
	ApiService *OnChainEarnAPIService
	createFinanceStakingDefiCancelV5Req *CreateFinanceStakingDefiCancelV5Req
}

// The request body for CreateFinanceStakingDefiCancelV5
func (r ApiCreateFinanceStakingDefiCancelV5Request) CreateFinanceStakingDefiCancelV5Req(createFinanceStakingDefiCancelV5Req CreateFinanceStakingDefiCancelV5Req) ApiCreateFinanceStakingDefiCancelV5Request {
	r.createFinanceStakingDefiCancelV5Req = &createFinanceStakingDefiCancelV5Req
	return r
}

func (r ApiCreateFinanceStakingDefiCancelV5Request) Execute() (*CreateFinanceStakingDefiCancelV5Resp, *http.Response, error) {
	return r.ApiService.CreateFinanceStakingDefiCancelV5Execute(r)
}

/*
CreateFinanceStakingDefiCancelV5 POST / Cancel purchases/redemptions

**_After cancelling, returning funds will go to the funding account._**

#### Rate Limit: 2 requests per second 

#### Rate limit rule: User ID 

#### Permission: Trade 



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateFinanceStakingDefiCancelV5Request
*/
func (a *OnChainEarnAPIService) CreateFinanceStakingDefiCancelV5(ctx context.Context) ApiCreateFinanceStakingDefiCancelV5Request {
	return ApiCreateFinanceStakingDefiCancelV5Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CreateFinanceStakingDefiCancelV5Resp
func (a *OnChainEarnAPIService) CreateFinanceStakingDefiCancelV5Execute(r ApiCreateFinanceStakingDefiCancelV5Request) (*CreateFinanceStakingDefiCancelV5Resp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateFinanceStakingDefiCancelV5Resp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OnChainEarnAPIService.CreateFinanceStakingDefiCancelV5")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v5/finance/staking-defi/cancel"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createFinanceStakingDefiCancelV5Req == nil {
		return localVarReturnValue, nil, reportError("createFinanceStakingDefiCancelV5Req is required and must be specified")
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
	localVarPostBody = r.createFinanceStakingDefiCancelV5Req
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

type ApiCreateFinanceStakingDefiPurchaseV5Request struct {
	ctx context.Context
	ApiService *OnChainEarnAPIService
	createFinanceStakingDefiPurchaseV5Req *CreateFinanceStakingDefiPurchaseV5Req
}

// The request body for CreateFinanceStakingDefiPurchaseV5
func (r ApiCreateFinanceStakingDefiPurchaseV5Request) CreateFinanceStakingDefiPurchaseV5Req(createFinanceStakingDefiPurchaseV5Req CreateFinanceStakingDefiPurchaseV5Req) ApiCreateFinanceStakingDefiPurchaseV5Request {
	r.createFinanceStakingDefiPurchaseV5Req = &createFinanceStakingDefiPurchaseV5Req
	return r
}

func (r ApiCreateFinanceStakingDefiPurchaseV5Request) Execute() (*CreateFinanceStakingDefiPurchaseV5Resp, *http.Response, error) {
	return r.ApiService.CreateFinanceStakingDefiPurchaseV5Execute(r)
}

/*
CreateFinanceStakingDefiPurchaseV5 POST / Purchase

#### Rate Limit: 2 requests per second 

#### Rate limit rule: User ID 

#### Permission: Trade 



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateFinanceStakingDefiPurchaseV5Request
*/
func (a *OnChainEarnAPIService) CreateFinanceStakingDefiPurchaseV5(ctx context.Context) ApiCreateFinanceStakingDefiPurchaseV5Request {
	return ApiCreateFinanceStakingDefiPurchaseV5Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CreateFinanceStakingDefiPurchaseV5Resp
func (a *OnChainEarnAPIService) CreateFinanceStakingDefiPurchaseV5Execute(r ApiCreateFinanceStakingDefiPurchaseV5Request) (*CreateFinanceStakingDefiPurchaseV5Resp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateFinanceStakingDefiPurchaseV5Resp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OnChainEarnAPIService.CreateFinanceStakingDefiPurchaseV5")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v5/finance/staking-defi/purchase"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createFinanceStakingDefiPurchaseV5Req == nil {
		return localVarReturnValue, nil, reportError("createFinanceStakingDefiPurchaseV5Req is required and must be specified")
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
	localVarPostBody = r.createFinanceStakingDefiPurchaseV5Req
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

type ApiCreateFinanceStakingDefiRedeemV5Request struct {
	ctx context.Context
	ApiService *OnChainEarnAPIService
	createFinanceStakingDefiRedeemV5Req *CreateFinanceStakingDefiRedeemV5Req
}

// The request body for CreateFinanceStakingDefiRedeemV5
func (r ApiCreateFinanceStakingDefiRedeemV5Request) CreateFinanceStakingDefiRedeemV5Req(createFinanceStakingDefiRedeemV5Req CreateFinanceStakingDefiRedeemV5Req) ApiCreateFinanceStakingDefiRedeemV5Request {
	r.createFinanceStakingDefiRedeemV5Req = &createFinanceStakingDefiRedeemV5Req
	return r
}

func (r ApiCreateFinanceStakingDefiRedeemV5Request) Execute() (*CreateFinanceStakingDefiRedeemV5Resp, *http.Response, error) {
	return r.ApiService.CreateFinanceStakingDefiRedeemV5Execute(r)
}

/*
CreateFinanceStakingDefiRedeemV5 POST / Redeem

#### Rate Limit: 2 requests per second 

#### Rate limit rule: User ID 

#### Permission: Trade 



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateFinanceStakingDefiRedeemV5Request
*/
func (a *OnChainEarnAPIService) CreateFinanceStakingDefiRedeemV5(ctx context.Context) ApiCreateFinanceStakingDefiRedeemV5Request {
	return ApiCreateFinanceStakingDefiRedeemV5Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CreateFinanceStakingDefiRedeemV5Resp
func (a *OnChainEarnAPIService) CreateFinanceStakingDefiRedeemV5Execute(r ApiCreateFinanceStakingDefiRedeemV5Request) (*CreateFinanceStakingDefiRedeemV5Resp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateFinanceStakingDefiRedeemV5Resp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OnChainEarnAPIService.CreateFinanceStakingDefiRedeemV5")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v5/finance/staking-defi/redeem"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createFinanceStakingDefiRedeemV5Req == nil {
		return localVarReturnValue, nil, reportError("createFinanceStakingDefiRedeemV5Req is required and must be specified")
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
	localVarPostBody = r.createFinanceStakingDefiRedeemV5Req
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

type ApiGetFinanceStakingDefiOffersV5Request struct {
	ctx context.Context
	ApiService *OnChainEarnAPIService
	productId *string
	protocolType *string
	ccy *string
}

// Product ID
func (r ApiGetFinanceStakingDefiOffersV5Request) ProductId(productId string) ApiGetFinanceStakingDefiOffersV5Request {
	r.productId = &productId
	return r
}

// Protocol type  &#x60;defi&#x60;: on-chain earn
func (r ApiGetFinanceStakingDefiOffersV5Request) ProtocolType(protocolType string) ApiGetFinanceStakingDefiOffersV5Request {
	r.protocolType = &protocolType
	return r
}

// Investment currency, e.g. &#x60;BTC&#x60;
func (r ApiGetFinanceStakingDefiOffersV5Request) Ccy(ccy string) ApiGetFinanceStakingDefiOffersV5Request {
	r.ccy = &ccy
	return r
}

func (r ApiGetFinanceStakingDefiOffersV5Request) Execute() (*GetFinanceStakingDefiOffersV5Resp, *http.Response, error) {
	return r.ApiService.GetFinanceStakingDefiOffersV5Execute(r)
}

/*
GetFinanceStakingDefiOffersV5 GET / Offers

#### Rate Limit: 3 requests per second 

#### Rate limit rule: User ID 

#### Permission: Read 



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetFinanceStakingDefiOffersV5Request
*/
func (a *OnChainEarnAPIService) GetFinanceStakingDefiOffersV5(ctx context.Context) ApiGetFinanceStakingDefiOffersV5Request {
	return ApiGetFinanceStakingDefiOffersV5Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetFinanceStakingDefiOffersV5Resp
func (a *OnChainEarnAPIService) GetFinanceStakingDefiOffersV5Execute(r ApiGetFinanceStakingDefiOffersV5Request) (*GetFinanceStakingDefiOffersV5Resp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetFinanceStakingDefiOffersV5Resp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OnChainEarnAPIService.GetFinanceStakingDefiOffersV5")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v5/finance/staking-defi/offers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.productId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "productId", r.productId, "form", "")
	} else {
		var defaultValue string = ""
		r.productId = &defaultValue
	}
	if r.protocolType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "protocolType", r.protocolType, "form", "")
	} else {
		var defaultValue string = ""
		r.protocolType = &defaultValue
	}
	if r.ccy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ccy", r.ccy, "form", "")
	} else {
		var defaultValue string = ""
		r.ccy = &defaultValue
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

type ApiGetFinanceStakingDefiOrdersActiveV5Request struct {
	ctx context.Context
	ApiService *OnChainEarnAPIService
	productId *string
	protocolType *string
	ccy *string
	state *string
}

// Product ID
func (r ApiGetFinanceStakingDefiOrdersActiveV5Request) ProductId(productId string) ApiGetFinanceStakingDefiOrdersActiveV5Request {
	r.productId = &productId
	return r
}

// Protocol type  &#x60;defi&#x60;: on-chain earn
func (r ApiGetFinanceStakingDefiOrdersActiveV5Request) ProtocolType(protocolType string) ApiGetFinanceStakingDefiOrdersActiveV5Request {
	r.protocolType = &protocolType
	return r
}

// Investment currency, e.g. &#x60;BTC&#x60;
func (r ApiGetFinanceStakingDefiOrdersActiveV5Request) Ccy(ccy string) ApiGetFinanceStakingDefiOrdersActiveV5Request {
	r.ccy = &ccy
	return r
}

// Order state  &#x60;8&#x60;: Pending   &#x60;13&#x60;: Cancelling   &#x60;9&#x60;: Onchain   &#x60;1&#x60;: Earning   &#x60;2&#x60;: Redeeming
func (r ApiGetFinanceStakingDefiOrdersActiveV5Request) State(state string) ApiGetFinanceStakingDefiOrdersActiveV5Request {
	r.state = &state
	return r
}

func (r ApiGetFinanceStakingDefiOrdersActiveV5Request) Execute() (*GetFinanceStakingDefiOrdersActiveV5Resp, *http.Response, error) {
	return r.ApiService.GetFinanceStakingDefiOrdersActiveV5Execute(r)
}

/*
GetFinanceStakingDefiOrdersActiveV5 GET / Active orders

#### Rate Limit: 3 requests per second 

#### Rate limit rule: User ID 

#### Permission: Read 



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetFinanceStakingDefiOrdersActiveV5Request
*/
func (a *OnChainEarnAPIService) GetFinanceStakingDefiOrdersActiveV5(ctx context.Context) ApiGetFinanceStakingDefiOrdersActiveV5Request {
	return ApiGetFinanceStakingDefiOrdersActiveV5Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetFinanceStakingDefiOrdersActiveV5Resp
func (a *OnChainEarnAPIService) GetFinanceStakingDefiOrdersActiveV5Execute(r ApiGetFinanceStakingDefiOrdersActiveV5Request) (*GetFinanceStakingDefiOrdersActiveV5Resp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetFinanceStakingDefiOrdersActiveV5Resp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OnChainEarnAPIService.GetFinanceStakingDefiOrdersActiveV5")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v5/finance/staking-defi/orders-active"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.productId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "productId", r.productId, "form", "")
	} else {
		var defaultValue string = ""
		r.productId = &defaultValue
	}
	if r.protocolType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "protocolType", r.protocolType, "form", "")
	} else {
		var defaultValue string = ""
		r.protocolType = &defaultValue
	}
	if r.ccy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ccy", r.ccy, "form", "")
	} else {
		var defaultValue string = ""
		r.ccy = &defaultValue
	}
	if r.state != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "state", r.state, "form", "")
	} else {
		var defaultValue string = ""
		r.state = &defaultValue
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

type ApiGetFinanceStakingDefiOrdersHistoryV5Request struct {
	ctx context.Context
	ApiService *OnChainEarnAPIService
	productId *string
	protocolType *string
	ccy *string
	after *string
	before *string
	limit *string
}

// Product ID
func (r ApiGetFinanceStakingDefiOrdersHistoryV5Request) ProductId(productId string) ApiGetFinanceStakingDefiOrdersHistoryV5Request {
	r.productId = &productId
	return r
}

// Protocol type  &#x60;defi&#x60;: on-chain earn
func (r ApiGetFinanceStakingDefiOrdersHistoryV5Request) ProtocolType(protocolType string) ApiGetFinanceStakingDefiOrdersHistoryV5Request {
	r.protocolType = &protocolType
	return r
}

// Investment currency, e.g. &#x60;BTC&#x60;
func (r ApiGetFinanceStakingDefiOrdersHistoryV5Request) Ccy(ccy string) ApiGetFinanceStakingDefiOrdersHistoryV5Request {
	r.ccy = &ccy
	return r
}

// Pagination of data to return records earlier than the requested ID. The value passed is the corresponding &#x60;ordId&#x60;
func (r ApiGetFinanceStakingDefiOrdersHistoryV5Request) After(after string) ApiGetFinanceStakingDefiOrdersHistoryV5Request {
	r.after = &after
	return r
}

// Pagination of data to return records newer than the requested ID. The value passed is the corresponding &#x60;ordId&#x60;
func (r ApiGetFinanceStakingDefiOrdersHistoryV5Request) Before(before string) ApiGetFinanceStakingDefiOrdersHistoryV5Request {
	r.before = &before
	return r
}

// Number of results per request. The default is &#x60;100&#x60;. The maximum is &#x60;100&#x60;.
func (r ApiGetFinanceStakingDefiOrdersHistoryV5Request) Limit(limit string) ApiGetFinanceStakingDefiOrdersHistoryV5Request {
	r.limit = &limit
	return r
}

func (r ApiGetFinanceStakingDefiOrdersHistoryV5Request) Execute() (*GetFinanceStakingDefiOrdersHistoryV5Resp, *http.Response, error) {
	return r.ApiService.GetFinanceStakingDefiOrdersHistoryV5Execute(r)
}

/*
GetFinanceStakingDefiOrdersHistoryV5 GET / Order history

#### Rate Limit: 3 requests per second 

#### Rate limit rule: User ID 

#### Permission: Read 



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetFinanceStakingDefiOrdersHistoryV5Request
*/
func (a *OnChainEarnAPIService) GetFinanceStakingDefiOrdersHistoryV5(ctx context.Context) ApiGetFinanceStakingDefiOrdersHistoryV5Request {
	return ApiGetFinanceStakingDefiOrdersHistoryV5Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetFinanceStakingDefiOrdersHistoryV5Resp
func (a *OnChainEarnAPIService) GetFinanceStakingDefiOrdersHistoryV5Execute(r ApiGetFinanceStakingDefiOrdersHistoryV5Request) (*GetFinanceStakingDefiOrdersHistoryV5Resp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetFinanceStakingDefiOrdersHistoryV5Resp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OnChainEarnAPIService.GetFinanceStakingDefiOrdersHistoryV5")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v5/finance/staking-defi/orders-history"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.productId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "productId", r.productId, "form", "")
	} else {
		var defaultValue string = ""
		r.productId = &defaultValue
	}
	if r.protocolType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "protocolType", r.protocolType, "form", "")
	} else {
		var defaultValue string = ""
		r.protocolType = &defaultValue
	}
	if r.ccy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ccy", r.ccy, "form", "")
	} else {
		var defaultValue string = ""
		r.ccy = &defaultValue
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
