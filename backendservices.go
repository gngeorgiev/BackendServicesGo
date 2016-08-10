package backendservices

import (
	"github.com/gngeorgiev/BackendServicesGo/data"
	"github.com/go-errors/errors"
)

var (
	ErrorInvalidAppId = errors.New("Invalid appId. The appId must be a string 16 characters long.")
)

type Settings struct {
	appId, scheme, url, apiVersion string
}

func (s Settings) GetAppId() string {
	return s.appId
}

type BackendServicesSdk struct {
	settings Settings
}

func New(settings Settings) (*BackendServicesSdk, error) {
	if len(settings.appId) != 16 {
		return nil, ErrorInvalidAppId
	}

	sdkSettings := &Settings{}
	sdkSettings.appId = settings.appId
	if settings.scheme == "" {
		sdkSettings.scheme = "http://"
	} else {
		sdkSettings.scheme = settings.scheme
	}

	if settings.apiVersion == "" {
		sdkSettings.apiVersion = "v1"
	} else {
		sdkSettings.apiVersion = settings.apiVersion
	}

	if settings.url == "" {
		sdkSettings.url = "api.everlive.com"
	} else {
		sdkSettings.url = settings.url
	}

	return &BackendServicesSdk{*sdkSettings}, nil
}

func (b *BackendServicesSdk) GetSettings() Settings {
	return b.settings
}

func (b *BackendServicesSdk) Data(contentType string) *data.Data {
	return data.New(contentType)
}
