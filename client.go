package amt

import (
	"fmt"

	"github.com/ammmze/wsman"
)

// Client used to perform actions on the machine
type Client struct {
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
	return &Client{
		wsManClient: wsmanClient,
	}, nil
}

// Close the client.
func (c *Client) Close() error {
	return nil
}

// PowerOn will power on a given machine.
func (c *Client) PowerOn() error {
	return powerOn(c)
}

// PowerOff will power off a given machine.
func (c *Client) PowerOff() error {
	return powerOff(c)
}

// PowerCycle will power cycle a given machine.
func (c *Client) PowerCycle() error {
	return powerCycle(c)
}

// SetPXE makes sure the node will pxe boot next time.
func (c *Client) SetPXE() error {
	return setPXE(c)
}

// IsPoweredOn checks current power state.
func (c *Client) IsPoweredOn() (bool, error) {
	return isPoweredOn(c)
}
