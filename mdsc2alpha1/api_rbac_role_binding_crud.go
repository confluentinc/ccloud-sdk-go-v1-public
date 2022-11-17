package mdsv2alpha1

import (
	_bytes "bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

type RBACRoleBindingCRUDApi interface {
	/*
	 * AddRoleForPrincipal Binds the principal to a role for a specific cluster or in the given scope. Callable by Admins.
	 *
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param principal Fully-qualified KafkaPrincipal string for a user.
	 * @param roleName The name of the role to bind the user to.
	 * @param scope
	 */
	AddRoleForPrincipal(ctx _context.Context, principal string, roleName string, scope Scope) (*_nethttp.Response, error)

	/*
	 * AddRoleResourcesForPrincipal Incrementally grant the resources to the principal at the given scope/cluster using the given role.
	 *
	 * Callable by Admins+ResourceOwners.
	 *
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param principal Fully-qualified KafkaPrincipal string for a user.
	 * @param roleName The name of the role.
	 * @param resourcesRequest
	 */
	AddRoleResourcesForPrincipal(ctx _context.Context, principal string, roleName string, resourcesRequest ResourcesRequest) (*_nethttp.Response, error)

	/*
	 * DeleteRoleForPrincipal Remove the role from the principal at the given scope/cluster. No-op if the user doesn't have the role.  Callable by Admins.
	 *
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param principal Fully-qualified KafkaPrincipal string for a user.
	 * @param roleName The name of the role.
	 * @param scope
	 */
	DeleteRoleForPrincipal(ctx _context.Context, principal string, roleName string, scope Scope) (*_nethttp.Response, error)

	/*
	 * RemoveRoleResourcesForPrincipal Incrementally remove the resources from the principal at the given scope/cluster using the given role.
	 *
	 * Callable by Admins+ResourceOwners.
	 *
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param principal Fully-qualified KafkaPrincipal string for a user.
	 * @param roleName The name of the role.
	 * @param resourcesRequest
	 */
	RemoveRoleResourcesForPrincipal(ctx _context.Context, principal string, roleName string, resourcesRequest ResourcesRequest) (*_nethttp.Response, error)
}

// RBACRoleBindingCRUDApiService RBACRoleBindingCRUDApi service
type RBACRoleBindingCRUDApiService service

/*
 * AddRoleForPrincipal Binds the principal to a role for a specific cluster or in the given scope. Callable by Admins.
 *
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param principal Fully-qualified KafkaPrincipal string for a user.
 * @param roleName The name of the role to bind the user to.
 * @param scope
 */
func (a *RBACRoleBindingCRUDApiService) AddRoleForPrincipal(ctx _context.Context, principal string, roleName string, scope Scope) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/principals/{principal}/roles/{roleName}"
	localVarPath = strings.Replace(localVarPath, "{"+"principal"+"}", _neturl.PathEscape(parameterToString(principal, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"roleName"+"}", _neturl.PathEscape(parameterToString(roleName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	localVarPostBody = &scope
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(_bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ErrorResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

/*
 * AddRoleResourcesForPrincipal Incrementally grant the resources to the principal at the given scope/cluster using the given role.
 *
 * Callable by Admins+ResourceOwners.
 *
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param principal Fully-qualified KafkaPrincipal string for a user.
 * @param roleName The name of the role.
 * @param resourcesRequest
 */
func (a *RBACRoleBindingCRUDApiService) AddRoleResourcesForPrincipal(ctx _context.Context, principal string, roleName string, resourcesRequest ResourcesRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/principals/{principal}/roles/{roleName}/bindings"
	localVarPath = strings.Replace(localVarPath, "{"+"principal"+"}", _neturl.PathEscape(parameterToString(principal, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"roleName"+"}", _neturl.PathEscape(parameterToString(roleName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	localVarPostBody = &resourcesRequest
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(_bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ErrorResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

/*
 * DeleteRoleForPrincipal Remove the role from the principal at the given scope/cluster. No-op if the user doesn't have the role.  Callable by Admins.
 *
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param principal Fully-qualified KafkaPrincipal string for a user.
 * @param roleName The name of the role.
 * @param scope
 */
func (a *RBACRoleBindingCRUDApiService) DeleteRoleForPrincipal(ctx _context.Context, principal string, roleName string, scope Scope) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/principals/{principal}/roles/{roleName}"
	localVarPath = strings.Replace(localVarPath, "{"+"principal"+"}", _neturl.PathEscape(parameterToString(principal, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"roleName"+"}", _neturl.PathEscape(parameterToString(roleName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	localVarPostBody = &scope
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(_bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ErrorResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

/*
 * RemoveRoleResourcesForPrincipal Incrementally remove the resources from the principal at the given scope/cluster using the given role.
 *
 * Callable by Admins+ResourceOwners.
 *
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param principal Fully-qualified KafkaPrincipal string for a user.
 * @param roleName The name of the role.
 * @param resourcesRequest
 */
func (a *RBACRoleBindingCRUDApiService) RemoveRoleResourcesForPrincipal(ctx _context.Context, principal string, roleName string, resourcesRequest ResourcesRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/principals/{principal}/roles/{roleName}/bindings"
	localVarPath = strings.Replace(localVarPath, "{"+"principal"+"}", _neturl.PathEscape(parameterToString(principal, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"roleName"+"}", _neturl.PathEscape(parameterToString(roleName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	localVarPostBody = &resourcesRequest
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(_bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ErrorResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
