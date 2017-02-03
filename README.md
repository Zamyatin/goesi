# Go API client for esi

An OpenAPI for EVE Online ESI API

## Overview
A module to allow access to CCP's EVE Online ESI API.
This module offers:
* Versioned Endpoints
* Ability to pass OAuth2 Tokens (See [this module](https://github.com/antihax/eveapi) for authentication)
* Complete API coverage.

## Installation
```
    go get github.com/antihax/goesi
```

## New Client
```
  client, err := goesi.NewClient(*http.Client, userAgent string)
```

One client should be created that will serve as an agent for all requests. This allows http2 multiplexing and keep-alive be used to optimize connections.
It is also good manners to provide a user-agent describing the point of use of the API, allowing CCP to contact you in case of emergencies.

Example
```
  client, err := goesi.NewClient(nil, "my esi client http://mysite.com contact <SomeDude> ingame")
  result, response, err := client.V#.Endpoint.Operation(requiredParam, map[string]interface{} { 
                                                                        "optionalParam1": "stringParam",
                                                                        "optionalParam2": 1234.56
                                                                    })
```

## Passing Tokens
OAuth2 tokens are passed to endpoings via contexts. Example:
```
	ctx := context.WithValue(context.TODO(), goesi.ContextOAuth2, ESIPublicToken)
	struc, response, err := client.V1.UniverseApi.GetUniverseStructuresStructureId(ctx, structureID, nil)
```

## Testing
If you would rather not rely on public ESI for testing, a mock ESI server is available for local and CI use.
Information here: https://github.com/antihax/mock-esi

## What about the other stuff?
Legacy, Dev, and Latest endpoints are not covered by this module as they are not for production use.
If you are using any of these endpoints, they will break! Use only versioned endpoints.
If you need bleeding edge access, add the endpoint to the generator and rebuild this module. 
Generator is here: https://github.com/antihax/swagger-esi-goclient

## Documentation for API Endpoints

[V1 Endpoints](./v1/README.md)

[V2 Endpoints](./v2/README.md)

[V3 Endpoints](./v3/README.md)

[V4 Endpoints](./v4/README.md)

## Author
  antihax on #devfleet slack


## Credits
https://github.com/go-resty/resty (MIT license) Copyright Â© 2015-2016 Jeevanandam M (jeeva@myjeeva.com)
 - Uses modified setBody and detectContentType

https://github.com/gregjones/httpcache (MIT license) Copyright Â© 2012 Greg Jones (greg.jones@gmail.com)
  - Uses parseCacheControl and CacheExpires as a helper function


