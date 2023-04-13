package cim

import (
	"fmt"

	"github.com/ammmze/wsman"
)

const (
	PowerManagementServiceSpace = "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_PowerManagementService"
)

type PowerManagementService struct {
	client *CimClient
}

func (c *CimClient) PowerManagementService() *PowerManagementService {
	return &PowerManagementService{client: c}
}

func (r *PowerManagementService) RequestPowerStateChange(requestedPowerState RequestedPowerState, managedElementRef Reference) error {
	msg := r.client.wsmanClient.Invoke(PowerManagementServiceSpace, "RequestPowerStateChange")
	msg.Parameters("PowerState", fmt.Sprint(int(requestedPowerState)))
	managedElement := msg.MakeParameter("ManagedElement")
	ref, err := wsman.MarshalToElement(managedElementRef)
	if err != nil {
		return err
	}
	managedElement.AddChildren(ref.Children()...)
	msg.AddParameter(managedElement)

	response, err := msg.Send()
	if err != nil {
		return err
	}

	body, returnValue, err := response.InvokeResponse()
	if err != nil {
		return err
	}
	if returnValue != "0" {
		return fmt.Errorf("received non-zero return value requesting power state change to <%s>. response was %s", requestedPowerState, body)
	}
	return nil
}
