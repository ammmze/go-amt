package amt

import (
	"fmt"

	"github.com/ammmze/go-amt/cim"
)

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

// IsPoweredOn checks current power state.
func (c *Client) IsPoweredOn() (bool, error) {
	return isPoweredOn(c)
}

func getPowerStatus(client *Client) (*cim.AssociatedPowerManagementServiceType, error) {
	// https://software.intel.com/sites/manageability/AMT_Implementation_and_Reference_Guide/default.htm?turl=WordDocuments%2Fgetsystempowerstate.htm
	items, err := client.AssociatedPowerManagementService().EnumerateObjectAndEPR()
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, fmt.Errorf("could not find the associated power management service")
	}
	return &items[0].Service, nil
}

func powerOn(client *Client) error {
	isOn, err := isPoweredOn(client)
	if err != nil {
		return err
	}
	if isOn {
		return nil
	}
	return requestPowerState(client, cim.RequestedPowerStateOn)
}

func powerOff(client *Client) error {
	status, err := getPowerStatus(client)
	if err != nil {
		return err
	}
	if isPoweredOnGivenStatus(status) {
		request := selectNextState(getPowerOffStates(), status.AvailableRequestedPowerStates)

		if request != cim.RequestedPowerStateUnknown {
			return requestPowerState(client, request)
		}
		return fmt.Errorf("there is no implemented transition state to power off the machine from the current machine state %d. available states are: %v", status.PowerState, status.AvailableRequestedPowerStates)
	}
	return nil
}

func powerCycle(client *Client) error {
	status, err := getPowerStatus(client)
	if err != nil {
		return err
	}

	if !isPoweredOnGivenStatus(status) {
		return powerOn(client)
	}

	request := selectNextState(getPowerCycleStates(), status.AvailableRequestedPowerStates)

	if request >= 0 {
		return requestPowerState(client, request)
	}

	return fmt.Errorf("there is no implemented transition state to power cycle the machine from the current machine state %d. available states are: %v", status.PowerState, status.AvailableRequestedPowerStates)
}

func isPoweredOn(client *Client) (bool, error) {
	status, err := getPowerStatus(client)
	if err != nil {
		return false, err
	}
	return isPoweredOnGivenStatus(status), nil
}

func isPoweredOnGivenStatus(status *cim.AssociatedPowerManagementServiceType) bool {
	// log.Printf("Current state: %s; Available states: %s\n", status.PowerState, status.AvailableRequestedPowerStates)
	return status.PowerState == cim.PowerStateOn
}

// https://software.intel.com/sites/manageability/AMT_Implementation_and_Reference_Guide/default.htm?turl=WordDocuments%2Fchangesystempowerstate.htm
func requestPowerState(client *Client, requestedPowerState cim.RequestedPowerState) error {
	status, err := getPowerStatus(client)
	if err != nil {
		return err
	}
	if !containsRequestedPowerState(status.AvailableRequestedPowerStates, requestedPowerState) {
		return fmt.Errorf("there is no implemented transition state to <%d> from the current machine state <%d>. available states are: %v", requestedPowerState, status.PowerState, status.AvailableRequestedPowerStates)
	}

	err = client.PowerManagementService().RequestPowerStateChange(requestedPowerState, *status.UserOfService)
	if err != nil {
		return err
	}
	return nil
}

func getPowerOffStates() []cim.RequestedPowerState {
	return []cim.RequestedPowerState{
		cim.RequestedPowerStateOffSoftGraceful,
		cim.RequestedPowerStateOffHardGraceful,
		cim.RequestedPowerStateOffSoft,
		cim.RequestedPowerStateOffHard,
	}
}

func getPowerCycleStates() []cim.RequestedPowerState {
	return []cim.RequestedPowerState{
		cim.RequestedPowerStateMasterBusResetGraceful,
		cim.RequestedPowerStatePowerCycleOffSoftGraceful,
		cim.RequestedPowerStatePowerCycleOffHardGraceful,
		cim.RequestedPowerStateMasterBusReset,
		cim.RequestedPowerStatePowerCycleOffSoft,
		cim.RequestedPowerStatePowerCycleOffHard,
	}
}

func selectNextState(requestedStates []cim.RequestedPowerState, availableStates []cim.RequestedPowerState) cim.RequestedPowerState {
	for _, a := range requestedStates {
		if containsRequestedPowerState(availableStates, a) {
			return a
		}
	}
	return cim.RequestedPowerStateUnknown
}

func containsRequestedPowerState(s []cim.RequestedPowerState, e cim.RequestedPowerState) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
