/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.4.9
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package goesiv1

import (
	"golang.org/x/net/context"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type UserInterfaceApiService service

/* UserInterfaceApiService Set Autopilot Waypoint
Set a solar system as autopilot waypoint  ---  Alternate route: &#x60;/legacy/ui/autopilot/waypoint/&#x60;

* @param ctx context.Context Authentication Context
@param addToBeginning Whether this solar system should be added to the beginning of all waypoints
@param clearOtherWaypoints Whether clean other waypoints beforing adding this one
@param solarSystemId The solar system to set as autopilot waypoint
@param optional (nil or map[string]interface{}) with one or more of:
    @param "datasource" (string) The server name you would like data from
    @param "token" (string) Access token to use, if preferred over a header
    @param "userAgent" (string) Client identifier, takes precedence over headers
    @param "xUserAgent" (string) Client identifier, takes precedence over User-Agent
@return */
func (a *UserInterfaceApiService) PostUiAutopilotWaypoint(ctx context.Context, addToBeginning bool, clearOtherWaypoints bool, solarSystemId int32, localVarOptionals map[string]interface{}) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.basePath + "/ui/autopilot/waypoint/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["datasource"], "string", "datasource"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(localVarOptionals["token"], "string", "token"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(localVarOptionals["userAgent"], "string", "userAgent"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(localVarOptionals["xUserAgent"], "string", "xUserAgent"); err != nil {
		return nil, err
	}

	localVarQueryParams.Add("add_to_beginning", parameterToString(addToBeginning, ""))
	localVarQueryParams.Add("clear_other_waypoints", parameterToString(clearOtherWaypoints, ""))
	if localVarTempParam, localVarOk := localVarOptionals["datasource"].(string); localVarOk {
		localVarQueryParams.Add("datasource", parameterToString(localVarTempParam, ""))
	}
	localVarQueryParams.Add("solar_system_id", parameterToString(solarSystemId, ""))
	if localVarTempParam, localVarOk := localVarOptionals["token"].(string); localVarOk {
		localVarQueryParams.Add("token", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["userAgent"].(string); localVarOk {
		localVarQueryParams.Add("user_agent", parameterToString(localVarTempParam, ""))
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam, localVarOk := localVarOptionals["xUserAgent"].(string); localVarOk {
		localVarHeaderParams["X-User-Agent"] = parameterToString(localVarTempParam, "")
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		return localVarHttpResponse, reportError(localVarHttpResponse.Status)
	}

	return localVarHttpResponse, err
}

/* UserInterfaceApiService Open Contract Window
Open the contract window inside the client  ---  Alternate route: &#x60;/legacy/ui/openwindow/contract/&#x60;  Alternate route: &#x60;/latest/ui/openwindow/contract/&#x60;  Alternate route: &#x60;/dev/ui/openwindow/contract/&#x60;

* @param ctx context.Context Authentication Context
@param contractId The contract to open
@param optional (nil or map[string]interface{}) with one or more of:
    @param "datasource" (string) The server name you would like data from
    @param "token" (string) Access token to use, if preferred over a header
    @param "userAgent" (string) Client identifier, takes precedence over headers
    @param "xUserAgent" (string) Client identifier, takes precedence over User-Agent
@return */
func (a *UserInterfaceApiService) PostUiOpenwindowContract(ctx context.Context, contractId int32, localVarOptionals map[string]interface{}) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.basePath + "/ui/openwindow/contract/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["datasource"], "string", "datasource"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(localVarOptionals["token"], "string", "token"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(localVarOptionals["userAgent"], "string", "userAgent"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(localVarOptionals["xUserAgent"], "string", "xUserAgent"); err != nil {
		return nil, err
	}

	localVarQueryParams.Add("contract_id", parameterToString(contractId, ""))
	if localVarTempParam, localVarOk := localVarOptionals["datasource"].(string); localVarOk {
		localVarQueryParams.Add("datasource", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["token"].(string); localVarOk {
		localVarQueryParams.Add("token", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["userAgent"].(string); localVarOk {
		localVarQueryParams.Add("user_agent", parameterToString(localVarTempParam, ""))
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam, localVarOk := localVarOptionals["xUserAgent"].(string); localVarOk {
		localVarHeaderParams["X-User-Agent"] = parameterToString(localVarTempParam, "")
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		return localVarHttpResponse, reportError(localVarHttpResponse.Status)
	}

	return localVarHttpResponse, err
}

/* UserInterfaceApiService Open Information Window
Open the information window for a character, corporation or alliance inside the client  ---  Alternate route: &#x60;/legacy/ui/openwindow/information/&#x60;  Alternate route: &#x60;/latest/ui/openwindow/information/&#x60;  Alternate route: &#x60;/dev/ui/openwindow/information/&#x60;

* @param ctx context.Context Authentication Context
@param targetId The target to open
@param optional (nil or map[string]interface{}) with one or more of:
    @param "datasource" (string) The server name you would like data from
    @param "token" (string) Access token to use, if preferred over a header
    @param "userAgent" (string) Client identifier, takes precedence over headers
    @param "xUserAgent" (string) Client identifier, takes precedence over User-Agent
@return */
func (a *UserInterfaceApiService) PostUiOpenwindowInformation(ctx context.Context, targetId int32, localVarOptionals map[string]interface{}) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.basePath + "/ui/openwindow/information/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["datasource"], "string", "datasource"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(localVarOptionals["token"], "string", "token"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(localVarOptionals["userAgent"], "string", "userAgent"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(localVarOptionals["xUserAgent"], "string", "xUserAgent"); err != nil {
		return nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["datasource"].(string); localVarOk {
		localVarQueryParams.Add("datasource", parameterToString(localVarTempParam, ""))
	}
	localVarQueryParams.Add("target_id", parameterToString(targetId, ""))
	if localVarTempParam, localVarOk := localVarOptionals["token"].(string); localVarOk {
		localVarQueryParams.Add("token", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["userAgent"].(string); localVarOk {
		localVarQueryParams.Add("user_agent", parameterToString(localVarTempParam, ""))
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam, localVarOk := localVarOptionals["xUserAgent"].(string); localVarOk {
		localVarHeaderParams["X-User-Agent"] = parameterToString(localVarTempParam, "")
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		return localVarHttpResponse, reportError(localVarHttpResponse.Status)
	}

	return localVarHttpResponse, err
}

/* UserInterfaceApiService Open Market Details
Open the market details window for a specific typeID inside the client  ---  Alternate route: &#x60;/legacy/ui/openwindow/marketdetails/&#x60;  Alternate route: &#x60;/latest/ui/openwindow/marketdetails/&#x60;  Alternate route: &#x60;/dev/ui/openwindow/marketdetails/&#x60;

* @param ctx context.Context Authentication Context
@param typeId The item type to open in market window
@param optional (nil or map[string]interface{}) with one or more of:
    @param "datasource" (string) The server name you would like data from
    @param "token" (string) Access token to use, if preferred over a header
    @param "userAgent" (string) Client identifier, takes precedence over headers
    @param "xUserAgent" (string) Client identifier, takes precedence over User-Agent
@return */
func (a *UserInterfaceApiService) PostUiOpenwindowMarketdetails(ctx context.Context, typeId int32, localVarOptionals map[string]interface{}) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.basePath + "/ui/openwindow/marketdetails/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["datasource"], "string", "datasource"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(localVarOptionals["token"], "string", "token"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(localVarOptionals["userAgent"], "string", "userAgent"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(localVarOptionals["xUserAgent"], "string", "xUserAgent"); err != nil {
		return nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["datasource"].(string); localVarOk {
		localVarQueryParams.Add("datasource", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["token"].(string); localVarOk {
		localVarQueryParams.Add("token", parameterToString(localVarTempParam, ""))
	}
	localVarQueryParams.Add("type_id", parameterToString(typeId, ""))
	if localVarTempParam, localVarOk := localVarOptionals["userAgent"].(string); localVarOk {
		localVarQueryParams.Add("user_agent", parameterToString(localVarTempParam, ""))
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam, localVarOk := localVarOptionals["xUserAgent"].(string); localVarOk {
		localVarHeaderParams["X-User-Agent"] = parameterToString(localVarTempParam, "")
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		return localVarHttpResponse, reportError(localVarHttpResponse.Status)
	}

	return localVarHttpResponse, err
}

/* UserInterfaceApiService Open New Mail Window
Open the New Mail window, according to settings from the request if applicable  ---  Alternate route: &#x60;/legacy/ui/openwindow/newmail/&#x60;  Alternate route: &#x60;/latest/ui/openwindow/newmail/&#x60;  Alternate route: &#x60;/dev/ui/openwindow/newmail/&#x60;

* @param ctx context.Context Authentication Context
@param newMail The details of mail to create
@param optional (nil or map[string]interface{}) with one or more of:
    @param "datasource" (string) The server name you would like data from
    @param "token" (string) Access token to use, if preferred over a header
    @param "userAgent" (string) Client identifier, takes precedence over headers
    @param "xUserAgent" (string) Client identifier, takes precedence over User-Agent
@return */
func (a *UserInterfaceApiService) PostUiOpenwindowNewmail(ctx context.Context, newMail PostUiOpenwindowNewmailNewMail, localVarOptionals map[string]interface{}) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.basePath + "/ui/openwindow/newmail/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["datasource"], "string", "datasource"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(localVarOptionals["token"], "string", "token"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(localVarOptionals["userAgent"], "string", "userAgent"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(localVarOptionals["xUserAgent"], "string", "xUserAgent"); err != nil {
		return nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["datasource"].(string); localVarOk {
		localVarQueryParams.Add("datasource", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["token"].(string); localVarOk {
		localVarQueryParams.Add("token", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["userAgent"].(string); localVarOk {
		localVarQueryParams.Add("user_agent", parameterToString(localVarTempParam, ""))
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam, localVarOk := localVarOptionals["xUserAgent"].(string); localVarOk {
		localVarHeaderParams["X-User-Agent"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &newMail

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		return localVarHttpResponse, reportError(localVarHttpResponse.Status)
	}

	return localVarHttpResponse, err
}
