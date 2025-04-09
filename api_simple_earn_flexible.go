/*
OKX v5 API

OpenAPI specification for Okx exchange - Rest API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rest

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// SimpleEarnFlexibleAPIService SimpleEarnFlexibleAPI service
type SimpleEarnFlexibleAPIService service

type ApiCreateFinanceSavingsPurchaseRedemptV5Request struct {
	ctx context.Context
	ApiService *SimpleEarnFlexibleAPIService
	createFinanceSavingsPurchaseRedemptV5Req *CreateFinanceSavingsPurchaseRedemptV5Req
}

// The request body for CreateFinanceSavingsPurchaseRedemptV5
func (r ApiCreateFinanceSavingsPurchaseRedemptV5Request) CreateFinanceSavingsPurchaseRedemptV5Req(createFinanceSavingsPurchaseRedemptV5Req CreateFinanceSavingsPurchaseRedemptV5Req) ApiCreateFinanceSavingsPurchaseRedemptV5Request {
	r.createFinanceSavingsPurchaseRedemptV5Req = &createFinanceSavingsPurchaseRedemptV5Req
	return r
}

func (r ApiCreateFinanceSavingsPurchaseRedemptV5Request) Execute() (*CreateFinanceSavingsPurchaseRedemptV5Resp, *http.Response, error) {
	return r.ApiService.CreateFinanceSavingsPurchaseRedemptV5Execute(r)
}

/*
CreateFinanceSavingsPurchaseRedemptV5 Only the assets in the funding account can be used for saving.  

Only the assets in the funding account can be used for saving.

#### Rate Limit: 6 requests per second 

#### Rate limit rule: User ID 

#### Permission: Trade 



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateFinanceSavingsPurchaseRedemptV5Request
*/
func (a *SimpleEarnFlexibleAPIService) CreateFinanceSavingsPurchaseRedemptV5(ctx context.Context) ApiCreateFinanceSavingsPurchaseRedemptV5Request {
	return ApiCreateFinanceSavingsPurchaseRedemptV5Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CreateFinanceSavingsPurchaseRedemptV5Resp
func (a *SimpleEarnFlexibleAPIService) CreateFinanceSavingsPurchaseRedemptV5Execute(r ApiCreateFinanceSavingsPurchaseRedemptV5Request) (*CreateFinanceSavingsPurchaseRedemptV5Resp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateFinanceSavingsPurchaseRedemptV5Resp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SimpleEarnFlexibleAPIService.CreateFinanceSavingsPurchaseRedemptV5")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v5/finance/savings/purchase-redempt"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createFinanceSavingsPurchaseRedemptV5Req == nil {
		return localVarReturnValue, nil, reportError("createFinanceSavingsPurchaseRedemptV5Req is required and must be specified")
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
	localVarPostBody = r.createFinanceSavingsPurchaseRedemptV5Req
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
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiCreateFinanceSavingsSetLendingRateV5Request struct {
	ctx context.Context
	ApiService *SimpleEarnFlexibleAPIService
	createFinanceSavingsSetLendingRateV5Req *CreateFinanceSavingsSetLendingRateV5Req
}

// The request body for CreateFinanceSavingsSetLendingRateV5
func (r ApiCreateFinanceSavingsSetLendingRateV5Request) CreateFinanceSavingsSetLendingRateV5Req(createFinanceSavingsSetLendingRateV5Req CreateFinanceSavingsSetLendingRateV5Req) ApiCreateFinanceSavingsSetLendingRateV5Request {
	r.createFinanceSavingsSetLendingRateV5Req = &createFinanceSavingsSetLendingRateV5Req
	return r
}

func (r ApiCreateFinanceSavingsSetLendingRateV5Request) Execute() (*CreateFinanceSavingsSetLendingRateV5Resp, *http.Response, error) {
	return r.ApiService.CreateFinanceSavingsSetLendingRateV5Execute(r)
}

/*
CreateFinanceSavingsSetLendingRateV5 Method for CreateFinanceSavingsSetLendingRateV5

#### Rate Limit: 6 requests per second 

#### Rate limit rule: User ID 

#### Permission: Trade 



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateFinanceSavingsSetLendingRateV5Request
*/
func (a *SimpleEarnFlexibleAPIService) CreateFinanceSavingsSetLendingRateV5(ctx context.Context) ApiCreateFinanceSavingsSetLendingRateV5Request {
	return ApiCreateFinanceSavingsSetLendingRateV5Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CreateFinanceSavingsSetLendingRateV5Resp
func (a *SimpleEarnFlexibleAPIService) CreateFinanceSavingsSetLendingRateV5Execute(r ApiCreateFinanceSavingsSetLendingRateV5Request) (*CreateFinanceSavingsSetLendingRateV5Resp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateFinanceSavingsSetLendingRateV5Resp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SimpleEarnFlexibleAPIService.CreateFinanceSavingsSetLendingRateV5")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v5/finance/savings/set-lending-rate"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createFinanceSavingsSetLendingRateV5Req == nil {
		return localVarReturnValue, nil, reportError("createFinanceSavingsSetLendingRateV5Req is required and must be specified")
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
	localVarPostBody = r.createFinanceSavingsSetLendingRateV5Req
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
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetFinanceSavingsBalanceV5Request struct {
	ctx context.Context
	ApiService *SimpleEarnFlexibleAPIService
	ccy *string
}

// Currency, e.g. &#x60;BTC&#x60;
func (r ApiGetFinanceSavingsBalanceV5Request) Ccy(ccy string) ApiGetFinanceSavingsBalanceV5Request {
	r.ccy = &ccy
	return r
}

func (r ApiGetFinanceSavingsBalanceV5Request) Execute() (*GetFinanceSavingsBalanceV5Resp, *http.Response, error) {
	return r.ApiService.GetFinanceSavingsBalanceV5Execute(r)
}

/*
GetFinanceSavingsBalanceV5 Method for GetFinanceSavingsBalanceV5

#### Rate Limit: 6 requests per second 

#### Rate limit rule: User ID 

#### Permission: Read 



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetFinanceSavingsBalanceV5Request
*/
func (a *SimpleEarnFlexibleAPIService) GetFinanceSavingsBalanceV5(ctx context.Context) ApiGetFinanceSavingsBalanceV5Request {
	return ApiGetFinanceSavingsBalanceV5Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetFinanceSavingsBalanceV5Resp
func (a *SimpleEarnFlexibleAPIService) GetFinanceSavingsBalanceV5Execute(r ApiGetFinanceSavingsBalanceV5Request) (*GetFinanceSavingsBalanceV5Resp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetFinanceSavingsBalanceV5Resp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SimpleEarnFlexibleAPIService.GetFinanceSavingsBalanceV5")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v5/finance/savings/balance"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetFinanceSavingsLendingHistoryV5Request struct {
	ctx context.Context
	ApiService *SimpleEarnFlexibleAPIService
	ccy *string
	after *string
	before *string
	limit *string
}

// Currency, e.g. &#x60;BTC&#x60;
func (r ApiGetFinanceSavingsLendingHistoryV5Request) Ccy(ccy string) ApiGetFinanceSavingsLendingHistoryV5Request {
	r.ccy = &ccy
	return r
}

// Pagination of data to return records earlier than the requested &#x60;ts&#x60;, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60;
func (r ApiGetFinanceSavingsLendingHistoryV5Request) After(after string) ApiGetFinanceSavingsLendingHistoryV5Request {
	r.after = &after
	return r
}

// Pagination of data to return records newer than the requested &#x60;ts&#x60;, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60;
func (r ApiGetFinanceSavingsLendingHistoryV5Request) Before(before string) ApiGetFinanceSavingsLendingHistoryV5Request {
	r.before = &before
	return r
}

// Number of results per request. The maximum is &#x60;100&#x60;. The default is &#x60;100&#x60;.
func (r ApiGetFinanceSavingsLendingHistoryV5Request) Limit(limit string) ApiGetFinanceSavingsLendingHistoryV5Request {
	r.limit = &limit
	return r
}

func (r ApiGetFinanceSavingsLendingHistoryV5Request) Execute() (*GetFinanceSavingsLendingHistoryV5Resp, *http.Response, error) {
	return r.ApiService.GetFinanceSavingsLendingHistoryV5Execute(r)
}

/*
GetFinanceSavingsLendingHistoryV5 Return data in the past month.  

Return data in the past month.

#### Rate Limit: 6 requests per second 

#### Rate limit rule: User ID 

#### Permission: Read 



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetFinanceSavingsLendingHistoryV5Request
*/
func (a *SimpleEarnFlexibleAPIService) GetFinanceSavingsLendingHistoryV5(ctx context.Context) ApiGetFinanceSavingsLendingHistoryV5Request {
	return ApiGetFinanceSavingsLendingHistoryV5Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetFinanceSavingsLendingHistoryV5Resp
func (a *SimpleEarnFlexibleAPIService) GetFinanceSavingsLendingHistoryV5Execute(r ApiGetFinanceSavingsLendingHistoryV5Request) (*GetFinanceSavingsLendingHistoryV5Resp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetFinanceSavingsLendingHistoryV5Resp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SimpleEarnFlexibleAPIService.GetFinanceSavingsLendingHistoryV5")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v5/finance/savings/lending-history"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetFinanceSavingsLendingRateHistoryV5Request struct {
	ctx context.Context
	ApiService *SimpleEarnFlexibleAPIService
	ccy *string
	after *string
	before *string
	limit *string
}

// Currency, e.g. &#x60;BTC&#x60;
func (r ApiGetFinanceSavingsLendingRateHistoryV5Request) Ccy(ccy string) ApiGetFinanceSavingsLendingRateHistoryV5Request {
	r.ccy = &ccy
	return r
}

// Pagination of data to return records earlier than the requested &#x60;ts&#x60;, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60;
func (r ApiGetFinanceSavingsLendingRateHistoryV5Request) After(after string) ApiGetFinanceSavingsLendingRateHistoryV5Request {
	r.after = &after
	return r
}

// Pagination of data to return records newer than the requested &#x60;ts&#x60;, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60;
func (r ApiGetFinanceSavingsLendingRateHistoryV5Request) Before(before string) ApiGetFinanceSavingsLendingRateHistoryV5Request {
	r.before = &before
	return r
}

// Number of results per request. The maximum is &#x60;100&#x60;. The default is &#x60;100&#x60;.  If &#x60;ccy&#x60; is not specified, all data under the same &#x60;ts&#x60; will be returned, not limited by &#x60;limit&#x60;
func (r ApiGetFinanceSavingsLendingRateHistoryV5Request) Limit(limit string) ApiGetFinanceSavingsLendingRateHistoryV5Request {
	r.limit = &limit
	return r
}

func (r ApiGetFinanceSavingsLendingRateHistoryV5Request) Execute() (*GetFinanceSavingsLendingRateHistoryV5Resp, *http.Response, error) {
	return r.ApiService.GetFinanceSavingsLendingRateHistoryV5Execute(r)
}

/*
GetFinanceSavingsLendingRateHistoryV5 Authentication is not required for this public endpoint.   Only returned records after December 14, 2021.  

Authentication is not required for this public endpoint.


Only returned records after December 14, 2021.

#### Rate Limit: 6 requests per second 

#### Rate limit rule: IP 



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetFinanceSavingsLendingRateHistoryV5Request
*/
func (a *SimpleEarnFlexibleAPIService) GetFinanceSavingsLendingRateHistoryV5(ctx context.Context) ApiGetFinanceSavingsLendingRateHistoryV5Request {
	return ApiGetFinanceSavingsLendingRateHistoryV5Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetFinanceSavingsLendingRateHistoryV5Resp
func (a *SimpleEarnFlexibleAPIService) GetFinanceSavingsLendingRateHistoryV5Execute(r ApiGetFinanceSavingsLendingRateHistoryV5Request) (*GetFinanceSavingsLendingRateHistoryV5Resp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetFinanceSavingsLendingRateHistoryV5Resp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SimpleEarnFlexibleAPIService.GetFinanceSavingsLendingRateHistoryV5")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v5/finance/savings/lending-rate-history"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetFinanceSavingsLendingRateSummaryV5Request struct {
	ctx context.Context
	ApiService *SimpleEarnFlexibleAPIService
	ccy *string
}

// Currency, e.g. &#x60;BTC&#x60;
func (r ApiGetFinanceSavingsLendingRateSummaryV5Request) Ccy(ccy string) ApiGetFinanceSavingsLendingRateSummaryV5Request {
	r.ccy = &ccy
	return r
}

func (r ApiGetFinanceSavingsLendingRateSummaryV5Request) Execute() (*GetFinanceSavingsLendingRateSummaryV5Resp, *http.Response, error) {
	return r.ApiService.GetFinanceSavingsLendingRateSummaryV5Execute(r)
}

/*
GetFinanceSavingsLendingRateSummaryV5 Authentication is not required for this public endpoint.  

Authentication is not required for this public endpoint.

#### Rate Limit: 6 requests per second 

#### Rate limit rule: IP 



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetFinanceSavingsLendingRateSummaryV5Request
*/
func (a *SimpleEarnFlexibleAPIService) GetFinanceSavingsLendingRateSummaryV5(ctx context.Context) ApiGetFinanceSavingsLendingRateSummaryV5Request {
	return ApiGetFinanceSavingsLendingRateSummaryV5Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetFinanceSavingsLendingRateSummaryV5Resp
func (a *SimpleEarnFlexibleAPIService) GetFinanceSavingsLendingRateSummaryV5Execute(r ApiGetFinanceSavingsLendingRateSummaryV5Request) (*GetFinanceSavingsLendingRateSummaryV5Resp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetFinanceSavingsLendingRateSummaryV5Resp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SimpleEarnFlexibleAPIService.GetFinanceSavingsLendingRateSummaryV5")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v5/finance/savings/lending-rate-summary"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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

	return localVarReturnValue, localVarHTTPResponse, nil
}
