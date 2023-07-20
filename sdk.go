package gstorage

import (
	"net/http"
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
		agent      string
	}
)

// WithBearerAuth opt to setup SDK authWriter
func WithBearerAuth(tokenGetter func() string) SdkOpt {
	return func(sdk *apiCore) {
		if tokenGetter != nil && tokenGetter() != "" {
			sdk.authWriter = httptransport.BearerToken(tokenGetter())
		}
	}
}

// WithUserAgent opt
func WithUserAgent(agent string) SdkOpt {
	return func(sdk *apiCore) {
		sdk.agent = agent
	}
}

// WithPermanentTokenAuth opt to setup SDK authWriter
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

// NewSDK constructor of storage api swagger client wrapper
// see UI of your data here https://api.gcore.com/storage
// apiHost = https://api.gcore.com
// apiBasePath = /storage
func NewSDK(apiHost, apiBasePath string, opts ...SdkOpt) *SDK {
	schema := strings.Split(apiHost, "://")
	if len(schema) > 1 {
		apiHost = schema[1]
		schema = schema[0:1]
	} else {
		schema = nil
	}
	transport := httptransport.New(apiHost, apiBasePath, schema)
	core := &apiCore{}
	for _, opt := range opts {
		opt(core)
	}
	if core.agent != "" {
		transport.Transport = customUserAgent{
			root:  transport.Transport,
			value: core.agent,
		}
	}
	core.client = client.New(transport, strfmt.Default)

	return &SDK{
		sdkKey:     &sdkKey{apiCore: core},
		sdkStorage: &sdkStorage{apiCore: core},
		sdkBucket:  &sdkBucket{apiCore: core},
	}
}

type customUserAgent struct {
	root  http.RoundTripper
	value string
}

// RoundTrip implemented
func (c customUserAgent) RoundTrip(request *http.Request) (*http.Response, error) {
	request.Header.Set("User-Agent", c.value)
	return c.root.RoundTrip(request)
}
