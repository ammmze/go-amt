package amt

import (
	"fmt"

	"github.com/ammmze/go-amt/cim"
	"github.com/ammmze/wsman"
)

// Client used to perform actions on the machine
type Client struct {
	*cimClient
	wsManClient *wsman.Client
}

type cimClient = cim.CimClient

// NewClient creates an amt client to use.
func NewClient(connection Connection) (*Client, error) {
	port := int(connection.Port)
	if port == 0 {
		port = 16992
	}
	path := connection.Path
	if path == "" {
		path = "/wsman"
	}
	target := fmt.Sprintf("http://%s:%d%s", connection.Host, port, path)
	wsmanClient, err := wsman.NewClient(target, connection.User, connection.Pass, true)
	if err != nil {
		return nil, err
	}
	wsmanClient.Debug = connection.Debug
	cimClient := cim.NewClient(wsmanClient)
	return &Client{
		cimClient:   cimClient,
		wsManClient: wsmanClient,
	}, nil
}

// Close the client.
func (c *Client) Close() error {
	return nil
}
