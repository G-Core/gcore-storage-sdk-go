package gstorage

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/G-Core/gcorelabs-storage-sdk-go/swagger/client/key"
	"github.com/G-Core/gcorelabs-storage-sdk-go/swagger/client/storage"
)

func TestNewSDK(t *testing.T) {
	apiUrl := strings.TrimSpace(os.Getenv("TESTS_API_URL"))
	apiPath := strings.TrimSpace(os.Getenv("TESTS_API_PATH"))
	apiToken := strings.TrimSpace(os.Getenv("TESTS_API_PERMANENT_TOKEN"))
	if apiUrl == "" || apiToken == "" {
		t.Skip("no defined TESTS_API_URL & TESTS_API_PERMANENT_TOKEN")
	}

	sdk := NewSDK(apiUrl, apiPath, WithPermanentTokenAuth(func() string { return apiToken }))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	sold := time.Now().Unix()

	// create key

	keyName := fmt.Sprintf("sdk-test-key-%d", sold)
	sshKey := `ssh-rsa AAAAB3NzaC1yc2EAAAABIwAAAQEAklOUpkDHrfHY17SbrmTIpNLTGK9Tjom/BWDSUGPl+nafzlHDTYW7hdI4yZ5ew18JH4JW9jbhUFrviQzM7xlELEVf4h9lFX5QVkbPppSwg0cda3Pbv7kOdJ/MTyBlWXFCR+HAo3FXRitBqxiX1nKhXpHAZsMciLq8V6RjsNAQwdsdMFvSlVK/7XAt3FaoJoAsncM1Q9x5+3V0Ww68/eIFmb1zuUFljQJKprrX88XypNDvjYNby6vw/Pb0rwert/EnmZ+AW4OZPnTPI89ZPmVMLuayrD2cE86Z/il8b+gw3r3+1nKatmIkjn2so1d01QraTlMqVSsbxNrRFi9wrf+M7Q== schacon@mylaptop.local`

	ckOpts := []func(params *key.KeyCreateHTTPParams){
		func(params *key.KeyCreateHTTPParams) { params.Context = ctx },
		func(params *key.KeyCreateHTTPParams) { params.Body.Name = keyName },
		func(params *key.KeyCreateHTTPParams) { params.Body.Key = sshKey },
	}
	ckResp, err := sdk.CreateKey(ckOpts...)
	if err != nil {
		t.Fatal("create key", err)
	}
	keyID := ckResp.ID

	// read key

	lkOpts := []func(p *key.KeyListHTTPV2Params){
		func(p *key.KeyListHTTPV2Params) { p.Context = ctx },
		func(p *key.KeyListHTTPV2Params) { id := fmt.Sprint(keyID); p.ID = &id },
	}
	lkResp, err := sdk.KeysList(lkOpts...)
	if err != nil {
		t.Fatal("read key", err)
	}
	if len(lkResp) == 0 {
		t.Fatal("read key empty response")
	}
	if lkResp[0].ID != keyID && lkResp[0].Name != keyName {
		t.Fatalf("read key want %d:%s got %d:%s", keyID, keyName, lkResp[0].ID, lkResp[0].Name)
	}

	// create storage

	stName := fmt.Sprintf("sdk-test-st-%d", sold)
	csOpts := []func(p *storage.StorageCreateHTTPParams){
		func(p *storage.StorageCreateHTTPParams) { p.Context = ctx },
		func(p *storage.StorageCreateHTTPParams) {
			p.Body.Type = "sftp"
			p.Body.Name = stName
			p.Body.Location = "lux"
		},
	}
	csResp, err := sdk.CreateStorage(csOpts...)
	if err != nil {
		t.Fatal("create storage", err)
	}
	stID := csResp.ID

	// modify storage

	msOpts := []func(p *storage.KeyLinkHTTPParams){
		func(p *storage.KeyLinkHTTPParams) {
			p.Context = ctx
			p.KeyID = keyID
			p.ID = stID
		},
	}
	err = sdk.LinkKeyToStorage(msOpts...)
	if err != nil {
		t.Fatal("link key to storage", err)
	}

	// read storage

	opts := make([]func(opt *storage.StorageListHTTPV2Params), 0)
	opts = append(opts, func(opt *storage.StorageListHTTPV2Params) {
		id := fmt.Sprint(stID)
		opt.Context = ctx
		opt.ID = &id
	})
	stResp, err := sdk.StoragesList(opts...)
	if err != nil {
		t.Fatal("read storage", err)
	}
	if len(stResp) == 0 {
		t.Fatal("read storage empty response")
	}
	if stResp[0].ID != stID && !strings.Contains(stResp[0].Name, stName) {
		t.Fatalf("read storage want %d:%s got %d:%s", stID, stName, stResp[0].ID, stResp[0].Name)
	}

	// delete key

	dkOpts := []func(p *key.KeyDeleteHTTPParams){
		func(p *key.KeyDeleteHTTPParams) {
			p.Context = ctx
			p.ID = keyID
		},
	}
	err = sdk.DeleteKey(dkOpts...)
	if err != nil {
		t.Fatal("delete key", err)
	}

	// delete storage

	dsOpts := []func(p *storage.StorageDeleteHTTPParams){
		func(p *storage.StorageDeleteHTTPParams) {
			p.Context = ctx
			p.ID = stID
		},
	}
	err = sdk.DeleteStorage(dsOpts...)
	if err != nil {
		t.Fatal("delete storage", err)
	}
}
