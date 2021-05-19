package gstorage

import (
	"fmt"

	"github.com/G-Core/gcorelabs-storage-sdk-go/swagger/client/storage"
	"github.com/G-Core/gcorelabs-storage-sdk-go/swagger/models"
)

type sdkStorage struct {
	*apiCore
}

//StoragesList getter for g-core storage api
//same result like on UI here https://storage.gcorelabs.com/storage/list
func (sdk *sdkStorage) StoragesList(opts ...func(params *storage.StorageListHTTPV2Params)) ([]models.Storage, error) {
	params := &storage.StorageListHTTPV2Params{}
	for _, opt := range opts {
		opt(params)
	}
	res, err := sdk.client.Storage.StorageListHTTPV2(params, sdk.authWriter)
	if err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}
	list := make([]models.Storage, len(res.Payload.Data))
	for i := range res.Payload.Data {
		list[i] = *res.Payload.Data[i]
	}
	return list, nil
}

//CreateStorage writer for g-core storage api
func (sdk *sdkKey) CreateStorage(opts ...func(params *storage.StorageCreateHTTPParams)) (*models.Storage, error) {
	params := &storage.StorageCreateHTTPParams{}
	for _, opt := range opts {
		opt(params)
	}
	res, err := sdk.client.Storage.StorageCreateHTTP(params, sdk.authWriter)
	if err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}
	return res.Payload, nil
}

//ModifyStorage writer for g-core storage api
func (sdk *sdkKey) ModifyStorage(opts ...func(params *storage.StorageUpdateHTTPParams)) (*models.Storage, error) {
	params := &storage.StorageUpdateHTTPParams{}
	for _, opt := range opts {
		opt(params)
	}
	res, err := sdk.client.Storage.StorageUpdateHTTP(params, sdk.authWriter)
	if err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}
	return res.Payload, nil
}

//DeleteStorage writer for g-core storage api
//be noticed that delete action is async in g-core end service
func (sdk *sdkKey) DeleteStorage(opts ...func(params *storage.StorageDeleteHTTPParams)) error {
	params := &storage.StorageDeleteHTTPParams{}
	for _, opt := range opts {
		opt(params)
	}
	_, err := sdk.client.Storage.StorageDeleteHTTP(params, sdk.authWriter)
	if err != nil {
		return fmt.Errorf("request: %w", err)
	}
	return nil
}

//LinkKeyToStorage writer for g-core storage api
func (sdk *sdkKey) LinkKeyToStorage(opts ...func(params *storage.KeyLinkHTTPParams)) error {
	params := &storage.KeyLinkHTTPParams{}
	for _, opt := range opts {
		opt(params)
	}
	_, err := sdk.client.Storage.KeyLinkHTTP(params, sdk.authWriter)
	if err != nil {
		return fmt.Errorf("request: %w", err)
	}
	return nil
}
