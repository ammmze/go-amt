package amt

import (
	"context"
	"fmt"

	"github.com/ammmze/wsman"
	"github.com/go-logr/logr"
)

// Client used to perform actions on the machine
type Client struct {
	logger      logr.Logger
	wsManClient *wsman.Client
}

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
	if connection.Logger.GetSink() == nil {
		connection.Logger = logr.Discard()
	}
	return &Client{
		logger:      connection.Logger,
		wsManClient: wsmanClient,
	}, nil
}

// Close the client.
func (c *Client) Close() error {
	return nil
}

// PowerOn will power on a given machine.
func (c *Client) PowerOn(ctx context.Context) error {
	return powerOn(ctx, c)
}

// PowerOff will power off a given machine.
func (c *Client) PowerOff(ctx context.Context) error {
	return powerOff(ctx, c)
}

// PowerCycle will power cycle a given machine.
func (c *Client) PowerCycle(ctx context.Context) error {
	return powerCycle(ctx, c)
}

// SetPXE makes sure the node will pxe boot next time.
func (c *Client) SetPXE(ctx context.Context) error {
	return setPXE(ctx, c)
}

// IsPoweredOn checks current power state.
func (c *Client) IsPoweredOn(ctx context.Context) (bool, error) {
	return isPoweredOn(ctx, c)
}
