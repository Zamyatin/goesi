/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 1.8.2
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

package esi

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"fmt"

	"github.com/antihax/goesi/optional"
)

// Linger please
var (
	_ context.Context
)

type LocationApiService service

/*
LocationApiService Get character location
Information about the characters current location. Returns the current solar system id, and also the current station or structure ID if applicable  ---  This route is cached for up to 5 seconds
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param characterId An EVE character ID
 * @param optional nil or *GetCharactersCharacterIdLocationOpts - Optional Parameters:
     * @param "Datasource" (optional.String) -  The server name you would like data from
     * @param "IfNoneMatch" (optional.String) -  ETag from a previous request. A 304 will be returned if this matches the current ETag
     * @param "Token" (optional.String) -  Access token to use if unable to set a header

@return GetCharactersCharacterIdLocationOk
*/

type GetCharactersCharacterIdLocationOpts struct {
	Datasource  optional.String
	IfNoneMatch optional.String
	Token       optional.String
}

func (a *LocationApiService) GetCharactersCharacterIdLocation(ctx context.Context, characterId int32, localVarOptionals *GetCharactersCharacterIdLocationOpts) (GetCharactersCharacterIdLocationOk, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue GetCharactersCharacterIdLocationOk
	)

	// create path and map variables
	localVarPath := a.client.basePath + "/v1/characters/{character_id}/location/"
	localVarPath = strings.Replace(localVarPath, "{"+"character_id"+"}", fmt.Sprintf("%v", characterId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if characterId < 1 {
		return localVarReturnValue, nil, reportError("characterId must be greater than 1")
	}

	if localVarOptionals != nil && localVarOptionals.Datasource.IsSet() {
		localVarQueryParams.Add("datasource", parameterToString(localVarOptionals.Datasource.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Token.IsSet() {
		localVarQueryParams.Add("token", parameterToString(localVarOptionals.Token.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.IfNoneMatch.IsSet() {
		localVarHeaderParams["If-None-Match"] = parameterToString(localVarOptionals.IfNoneMatch.Value(), "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 400 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("content-type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 400 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v GetCharactersCharacterIdLocationOk
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("content-type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 400 {
			var v BadRequest
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("content-type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 401 {
			var v Unauthorized
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("content-type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 403 {
			var v Forbidden
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("content-type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 420 {
			var v ErrorLimited
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("content-type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 500 {
			var v InternalServerError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("content-type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 503 {
			var v ServiceUnavailable
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("content-type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 504 {
			var v GatewayTimeout
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("content-type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
LocationApiService Get character online
Checks if the character is currently online  ---  This route is cached for up to 60 seconds
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param characterId An EVE character ID
 * @param optional nil or *GetCharactersCharacterIdOnlineOpts - Optional Parameters:
     * @param "Datasource" (optional.String) -  The server name you would like data from
     * @param "IfNoneMatch" (optional.String) -  ETag from a previous request. A 304 will be returned if this matches the current ETag
     * @param "Token" (optional.String) -  Access token to use if unable to set a header

@return GetCharactersCharacterIdOnlineOk
*/

type GetCharactersCharacterIdOnlineOpts struct {
	Datasource  optional.String
	IfNoneMatch optional.String
	Token       optional.String
}

func (a *LocationApiService) GetCharactersCharacterIdOnline(ctx context.Context, characterId int32, localVarOptionals *GetCharactersCharacterIdOnlineOpts) (GetCharactersCharacterIdOnlineOk, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue GetCharactersCharacterIdOnlineOk
	)

	// create path and map variables
	localVarPath := a.client.basePath + "/v2/characters/{character_id}/online/"
	localVarPath = strings.Replace(localVarPath, "{"+"character_id"+"}", fmt.Sprintf("%v", characterId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if characterId < 1 {
		return localVarReturnValue, nil, reportError("characterId must be greater than 1")
	}

	if localVarOptionals != nil && localVarOptionals.Datasource.IsSet() {
		localVarQueryParams.Add("datasource", parameterToString(localVarOptionals.Datasource.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Token.IsSet() {
		localVarQueryParams.Add("token", parameterToString(localVarOptionals.Token.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.IfNoneMatch.IsSet() {
		localVarHeaderParams["If-None-Match"] = parameterToString(localVarOptionals.IfNoneMatch.Value(), "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 400 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("content-type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 400 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v GetCharactersCharacterIdOnlineOk
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("content-type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 400 {
			var v BadRequest
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("content-type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 401 {
			var v Unauthorized
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("content-type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 403 {
			var v Forbidden
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("content-type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 420 {
			var v ErrorLimited
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("content-type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 500 {
			var v InternalServerError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("content-type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 503 {
			var v ServiceUnavailable
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("content-type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 504 {
			var v GatewayTimeout
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("content-type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
LocationApiService Get current ship
Get the current ship type, name and id  ---  This route is cached for up to 5 seconds
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param characterId An EVE character ID
 * @param optional nil or *GetCharactersCharacterIdShipOpts - Optional Parameters:
     * @param "Datasource" (optional.String) -  The server name you would like data from
     * @param "IfNoneMatch" (optional.String) -  ETag from a previous request. A 304 will be returned if this matches the current ETag
     * @param "Token" (optional.String) -  Access token to use if unable to set a header

@return GetCharactersCharacterIdShipOk
*/

type GetCharactersCharacterIdShipOpts struct {
	Datasource  optional.String
	IfNoneMatch optional.String
	Token       optional.String
}

func (a *LocationApiService) GetCharactersCharacterIdShip(ctx context.Context, characterId int32, localVarOptionals *GetCharactersCharacterIdShipOpts) (GetCharactersCharacterIdShipOk, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue GetCharactersCharacterIdShipOk
	)

	// create path and map variables
	localVarPath := a.client.basePath + "/v1/characters/{character_id}/ship/"
	localVarPath = strings.Replace(localVarPath, "{"+"character_id"+"}", fmt.Sprintf("%v", characterId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if characterId < 1 {
		return localVarReturnValue, nil, reportError("characterId must be greater than 1")
	}

	if localVarOptionals != nil && localVarOptionals.Datasource.IsSet() {
		localVarQueryParams.Add("datasource", parameterToString(localVarOptionals.Datasource.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Token.IsSet() {
		localVarQueryParams.Add("token", parameterToString(localVarOptionals.Token.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.IfNoneMatch.IsSet() {
		localVarHeaderParams["If-None-Match"] = parameterToString(localVarOptionals.IfNoneMatch.Value(), "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 400 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("content-type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 400 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v GetCharactersCharacterIdShipOk
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("content-type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 400 {
			var v BadRequest
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("content-type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 401 {
			var v Unauthorized
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("content-type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 403 {
			var v Forbidden
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("content-type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 420 {
			var v ErrorLimited
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("content-type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 500 {
			var v InternalServerError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("content-type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 503 {
			var v ServiceUnavailable
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("content-type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 504 {
			var v GatewayTimeout
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("content-type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
