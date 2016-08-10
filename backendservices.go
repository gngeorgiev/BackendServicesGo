package backendservices

import (
	"github.com/gngeorgiev/BackendServicesGo/data"
	"github.com/go-errors/errors"
)

var (
	ErrorInvalidAppId = errors.New("Invalid appId. The appId must be a string 16 characters long.")
)

type BackendServicesSdk struct {
	appId string
}

func New(appId string) (*BackendServicesSdk, error) {
	if len(appId) != 16 {
		return nil, ErrorInvalidAppId
	}

	return &BackendServicesSdk{
		appId: appId,
	}, nil
}

func (b *BackendServicesSdk) GetAppId() string {
	return b.appId
}

func (b *BackendServicesSdk) Data(contentType string) *data.Data {
	return data.New(contentType)
}
