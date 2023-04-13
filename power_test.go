package amt

import (
	"testing"

	"github.com/ammmze/go-amt/cim"
	"github.com/stretchr/testify/assert"
)

func TestSelectNextState_When_NoneOfTheRequestedStatesAreAvailable_Expect_Unknown(t *testing.T) {
	availableStates := []cim.RequestedPowerState{cim.RequestedPowerStateOn}
	nextState := selectNextState(getPowerOffStates(), availableStates)
	assert.Equal(t, cim.RequestedPowerStateUnknown, nextState)
}

func TestSelectNextState_When_OneOfTheRequestedStatesAreAvailable_Expect_RequestedState(t *testing.T) {
	requestedStates := getPowerOffStates()
	availableStates := []cim.RequestedPowerState{requestedStates[0]}
	nextState := selectNextState(requestedStates, availableStates)
	assert.Equal(t, requestedStates[0], nextState)
}
func TestSelectNextState_When_MultipleOfTheRequestedStatesAreAvailable_Expect_FirstAvailableRequestedState(t *testing.T) {
	requestedStates := getPowerOffStates()
	availableStates := []cim.RequestedPowerState{requestedStates[1], requestedStates[2]}
	nextState := selectNextState(requestedStates, availableStates)
	assert.Equal(t, requestedStates[1], nextState)
}

func TestIsPoweredOnGivenStatus_When_powerStateOn_Expect_True(t *testing.T) {
	status := &cim.AssociatedPowerManagementServiceType{PowerState: cim.PowerStateOn}
	actual := isPoweredOnGivenStatus(status)
	assert.Equal(t, true, actual)
}
func TestIsPoweredOnGivenStatus_When_powerStateOffSoft_Expect_False(t *testing.T) {
	status := &cim.AssociatedPowerManagementServiceType{PowerState: cim.PowerStateOffSoft}
	actual := isPoweredOnGivenStatus(status)
	assert.Equal(t, false, actual)
}
