package steamAPI

import (
	"encoding/json"
	"net/http"
	"steamAPI/models/API"
)

func (client *client) GetSupportedAPIList() ([]API.Interface, error) {
	var bytes []byte
	var err error

	if bytes, err = client.processAPIRequest(http.MethodGet, "ISteamWebAPIUtil", "GetSupportedAPIList", 1, nil); err != nil {
		return nil, err
	}

	var apiResponse API.APIResponse
	if err = json.Unmarshal(bytes, &apiResponse); err != nil {
		return nil, err
	}

	return apiResponse.APIList.Interfaces, err
}

func (client *client) GetServerInfo() (*API.ServerInfo, error) {
	var bytes []byte
	var err error

	if bytes, err = client.processAPIRequest(http.MethodGet, "ISteamWebAPIUtil", "GetServerInfo", 1, nil); err != nil {
		return nil, err
	}

	var serverInfo API.ServerInfo
	if err = json.Unmarshal(bytes, &serverInfo); err != nil {
		return nil, err
	}

	return &serverInfo, err
}
