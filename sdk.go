package gstorage

import (
	"github.com/G-Core/gcorelabs-storage-sdk-go/swagger/client"
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
	}
	apiCore struct {
		client     *client.GCDNStorageAPI
		authWriter runtime.ClientAuthInfoWriter
	}
)

//WithBearerAuth opt to setup SDK authWriter
func WithBearerAuth(tokenGetter func() string) SdkOpt {
	return func(sdk *apiCore) {
		sdk.authWriter = httptransport.BearerToken(tokenGetter())
	}
}

//NewSDK constructor of storage api swagger client wrapper
//see UI of your data here https://storage.gcorelabs.com/
func NewSDK(apiHost, apiBasePath string, opts ...SdkOpt) *SDK {
	transport := httptransport.New(apiHost, apiBasePath, nil)
	core := &apiCore{
		client: client.New(transport, strfmt.Default),
	}
	for _, opt := range opts {
		opt(core)
	}
	return &SDK{
		sdkKey:     &sdkKey{apiCore: core},
		sdkStorage: &sdkStorage{apiCore: core},
	}
}
