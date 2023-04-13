package cim

import "github.com/ammmze/wsman"

type CimClient struct {
	wsmanClient *wsman.Client
}

func NewClient(wsmanClient *wsman.Client) *CimClient {
	client := &CimClient{wsmanClient: wsmanClient}
	return client
}
