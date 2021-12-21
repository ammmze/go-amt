//go:generate stringer -type=powerState -trimprefix=powerState -linecomment

package amt

import (
	"log"
	"fmt"
	"strconv"

	"github.com/VictorLowther/simplexml/dom"
	"github.com/ammmze/wsman"
)

type powerState int

// https://software.intel.com/sites/manageability/AMT_Implementation_and_Reference_Guide/HTMLDocuments/WS-Management_Class_Reference/CIM_AssociatedPowerManagementService.htm#powerState
// https://software.intel.com/sites/manageability/AMT_Implementation_and_Reference_Guide/default.htm?turl=WordDocuments%2Fgetsystempowerstate.htm
const (
	powerStateUnknown                   powerState = 0
	powerStateOther                     powerState = 1
	powerStateOn                        powerState = 2
	powerStateSleepLight                powerState = 3
	powerStateSleepDeep                 powerState = 4
	powerStatePowerCycleOffSoft         powerState = 5
	powerStateOffHard                   powerState = 6
	powerStateHibernateOffSoft          powerState = 7
	powerStateOffSoft                   powerState = 8
	powerStatePowerCycleOffHard         powerState = 9
	powerStateMasterBusReset            powerState = 10
	powerStateDiagnosticInterruptNMI    powerState = 11
	powerStateOffSoftGraceful           powerState = 12
	powerStateOffHardGraceful           powerState = 13
	powerStateMasterBusResetGraceful    powerState = 14
	powerStatePowerCycleOffSoftGraceful powerState = 15
	powerStatePowerCycleOffHardGraceful powerState = 16
	powerStateDiagnosticInterruptInit   powerState = 17
	// DMTF Reserverd = ..
	// Vendor Specific = 0x7FFF..0xFFFF
)

type powerStatus struct {
	AvailableRequestedpowerStates []powerState
	powerState                    powerState
	RequestedpowerState           powerState
}

func getPowerStatus(client *Client) (*powerStatus, error) {
	// https://software.intel.com/sites/manageability/AMT_Implementation_and_Reference_Guide/default.htm?turl=WordDocuments%2Fgetsystempowerstate.htm
	message := client.wsManClient.Enumerate(resourceCIMAssociatedPowerManagementService)

	response, err := message.Send()
	if err != nil {
		return nil, err
	}
	pmElms, err := getPowerManagementElements(response)
	if err != nil {
		return nil, err
	}

	status := &powerStatus{
		AvailableRequestedpowerStates: []powerState{},
	}
	for _, e := range pmElms {
		switch e.Name.Local {
		case "PowerState":
			val, err := strconv.Atoi(string(e.Content))
			if err != nil {
				return nil, err
			}
			status.powerState = powerState(val)
		case "RequestedPowerState":
			val, err := strconv.Atoi(string(e.Content))
			if err != nil {
				return nil, err
			}
			status.RequestedpowerState = powerState(val)
		case "AvailableRequestedPowerStates":
			val, err := strconv.Atoi(string(e.Content))
			if err != nil {
				return nil, err
			}
			status.AvailableRequestedpowerStates = append(status.AvailableRequestedpowerStates, powerState(val))
		}
	}

	return status, nil
}

func powerOn(client *Client) error {
	isOn, err := isPoweredOn(client)
	if err != nil {
		return err
	}
	if isOn {
		return nil
	}
	_, err = requestpowerState(client, powerStateOn)
	return err
}

func powerOff(client *Client) error {
	status, err := getPowerStatus(client)
	if err != nil {
		return err
	}
	if isPoweredOnGivenStatus(status) {
		request := selectNextState(getPowerOffStates(), status.AvailableRequestedpowerStates)

		if request != powerStateUnknown {
			_, err := requestpowerState(client, request)
			return err
		}
		return fmt.Errorf("there is no implemented transition state to power off the machine from the current machine state %d. available states are: %v", status.powerState, status.AvailableRequestedpowerStates)
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

	request := selectNextState(getPowerCycleStates(), status.AvailableRequestedpowerStates)

	if request >= 0 {
		_, err := requestpowerState(client, request)
		return err
	}

	return fmt.Errorf("there is no implemented transition state to power cycle the machine from the current machine state %d. available states are: %v", status.powerState, status.AvailableRequestedpowerStates)
}

func isPoweredOn(client *Client) (bool, error) {
	status, err := getPowerStatus(client)
	if err != nil {
		return false, err
	}
	return isPoweredOnGivenStatus(status), nil
}

func isPoweredOnGivenStatus(status *powerStatus) bool {
	log.Printf("Current state: %s; Available states: %s\n", status.powerState, status.AvailableRequestedpowerStates)
	switch status.powerState {
	case powerStateOn:
		return true
	default:
		return false
	}
}

// https://software.intel.com/sites/manageability/AMT_Implementation_and_Reference_Guide/default.htm?turl=WordDocuments%2Fchangesystempowerstate.htm
func requestpowerState(client *Client, requestedpowerState powerState) (int, error) {
	status, err := getPowerStatus(client)
	if err != nil {
		return -1, err
	}
	if !containspowerState(status.AvailableRequestedpowerStates, requestedpowerState) {
		// fmt.Printf("there is no implemented transition state to <%d> from the current machine state <%d>. available states are: %v\n", requestedpowerState, status.powerState, status.AvailableRequestedpowerStates)
		return -1, fmt.Errorf("there is no implemented transition state to <%d> from the current machine state <%d>. available states are: %v", requestedpowerState, status.powerState, status.AvailableRequestedpowerStates)
	}
	fmt.Printf("sending request to machine: <%s>\n", requestedpowerState)
	message := client.wsManClient.Invoke(resourceCIMPowerManagementService, "RequestPowerStateChange")
	message.Parameters("PowerState", fmt.Sprint(int(requestedpowerState)))
	managedElement, err := makeManagedElement(client, message)
	if err != nil {
		return -1, err
	}
	message.AddParameter(managedElement)

	response, err := message.Send()
	if err != nil {
		return -1, err
	}

	body := response.GetBody(dom.Elem("RequestPowerStateChange_OUTPUT", resourceCIMPowerManagementService))
	if body == nil || len(body.Children()) != 1 {
		return -1, fmt.Errorf("received unknown response requesting power state change: %v", response)
	}
	val, err := strconv.Atoi(string(body.Children()[0].Content))
	if err != nil {
		return -1, err
	}
	fmt.Printf("RequestPowerState response <%d>\n", val)
	return val, nil
}

func getPowerManagementElements(response *wsman.Message) ([]*dom.Element, error) {
	items, err := response.EnumItems()

	if err != nil {
		return nil, err
	}

	for _, e := range items {
		if e.Name.Local == "CIM_AssociatedPowerManagementService" && e.Name.Space == resourceCIMAssociatedPowerManagementService {
			return e.Children(), nil
		}
	}
	return nil, fmt.Errorf("did not receive %s enumeration item", "CIM_AssociatedPowerManagementService")
}

func makeManagedElement(client *Client, message *wsman.Message) (*dom.Element, error) {
	managedSystemRef, err := getComputerSystemRef(client, "ManagedSystem")
	if err != nil {
		return nil, err
	}
	if managedSystemRef == nil {
		return nil, fmt.Errorf("could not retrieve the managed system endpoint reference")
	}
	managedElement := message.MakeParameter("ManagedElement")
	managedElement.AddChildren(managedSystemRef.Children()...)
	return managedElement, nil
}

func getPowerOffStates() []powerState {
	return []powerState{
		powerStateOffSoftGraceful,
		powerStateOffSoft,
		powerStateOffHardGraceful,
		powerStateOffHard,
	}
}

func getPowerCycleStates() []powerState {
	return []powerState{
		powerStatePowerCycleOffSoftGraceful,
		powerStatePowerCycleOffSoft,
		powerStateMasterBusResetGraceful,
		powerStatePowerCycleOffHardGraceful,
		powerStatePowerCycleOffHard,
		powerStateMasterBusReset,
	}
}

func selectNextState(requestedStates []powerState, availableStates []powerState) powerState {
	for _, a := range requestedStates {
		if containspowerState(availableStates, a) {
			return a
		}
	}
	return powerStateUnknown
}

func containspowerState(s []powerState, e powerState) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
