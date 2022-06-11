package steamAPI

import (
	"net/http"
	"steamAPI/models/API"
	"time"
)

type Client interface {
	GetSupportedAPIList() ([]API.Interface, error)
	GetServerInfo() (*API.ServerInfo, error)
}

type client struct {
	httpClient *http.Client
	APIKey     string
}

var instance Client

func GetClient() Client {
	if instance == nil {
		instance = NewClient()
	}
	return instance
}

func NewClient() Client {
	return &client{
		httpClient: &http.Client{
			Transport:     nil,
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       5 * time.Second,
		},
		APIKey: "53C1D792C8C437BA54C3ABE1DF0ED5B5",
	}
}
