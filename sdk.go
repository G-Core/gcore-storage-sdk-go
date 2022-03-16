package gstorage

import (
	"strings"

	"github.com/G-Core/gcore-storage-sdk-go/swagger/client"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

type (
	//SdkOpt is optional arg for setup apiCore
	SdkOpt func(sdk *apiCore)
	//SDK for g-core storage api
	SDK struct {
		*sdkKey
		*sdkStorage
		*sdkBucket
	}
	apiCore struct {
		client     *client.GCDNStorageAPI
		authWriter runtime.ClientAuthInfoWriter
	}
)

//WithBearerAuth opt to setup SDK authWriter
func WithBearerAuth(tokenGetter func() string) SdkOpt {
	return func(sdk *apiCore) {
		if tokenGetter != nil && tokenGetter() != "" {
			sdk.authWriter = httptransport.BearerToken(tokenGetter())
		}
	}
}

//WithPermanentTokenAuth opt to setup SDK authWriter
func WithPermanentTokenAuth(tokenGetter func() string) SdkOpt {
	return func(sdk *apiCore) {
		if tokenGetter != nil && tokenGetter() != "" {
			sdk.authWriter = runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
				// nolint: wrapcheck
				return r.SetHeaderParam("Authorization", "APIKey "+tokenGetter())
			})
		}
	}
}

//NewSDK constructor of storage api swagger client wrapper
//see UI of your data here https://api.gcorelabs.com/storage
//apiHost = https://api.gcorelabs.com
//apiBasePath = /storage
func NewSDK(apiHost, apiBasePath string, opts ...SdkOpt) *SDK {
	schema := strings.Split(apiHost, "://")
	if len(schema) > 1 {
		apiHost = schema[1]
		schema = schema[0:1]
	} else {
		schema = nil
	}
	transport := httptransport.New(apiHost, apiBasePath, schema)
	core := &apiCore{
		client: client.New(transport, strfmt.Default),
	}
	for _, opt := range opts {
		opt(core)
	}
	return &SDK{
		sdkKey:     &sdkKey{apiCore: core},
		sdkStorage: &sdkStorage{apiCore: core},
		sdkBucket:  &sdkBucket{apiCore: core},
	}
}
