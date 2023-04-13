package cim

// PowerState represents the possible values for CIM_AssociatedPowerManagementService.PowerState
type PowerState int

//go:generate stringer -type=PowerState -trimprefix=PowerState -linecomment
const (
	PowerStateUnknown          PowerState = 0
	PowerStateOn               PowerState = 2
	PowerStateSleepLight       PowerState = 3
	PowerStateSleepDeep        PowerState = 4
	PowerStateOffHard          PowerState = 6
	PowerStateHibernateOffSoft PowerState = 7
	PowerStateOffSoft          PowerState = 8
)

// RequestedPowerState represents the possible values power states that can be requested.
// CIM_AssociatedPowerManagementService.AvailableRequestedPowerStates and CIM_AssociatedPowerManagementService.RequestedPowerState
type RequestedPowerState int

//go:generate stringer -type=RequestedPowerState -trimprefix=RequestedPowerState -linecomment
const (
	RequestedPowerStateUnknown                   RequestedPowerState = 0
	RequestedPowerStateOn                        RequestedPowerState = 2
	RequestedPowerStateSleepLight                RequestedPowerState = 3
	RequestedPowerStateSleepDeep                 RequestedPowerState = 4
	RequestedPowerStatePowerCycleOffSoft         RequestedPowerState = 5
	RequestedPowerStateOffHard                   RequestedPowerState = 6
	RequestedPowerStateHibernateOffSoft          RequestedPowerState = 7
	RequestedPowerStateOffSoft                   RequestedPowerState = 8
	RequestedPowerStatePowerCycleOffHard         RequestedPowerState = 9
	RequestedPowerStateMasterBusReset            RequestedPowerState = 10
	RequestedPowerStateDiagnosticInterruptNMI    RequestedPowerState = 11
	RequestedPowerStateOffSoftGraceful           RequestedPowerState = 12
	RequestedPowerStateOffHardGraceful           RequestedPowerState = 13
	RequestedPowerStateMasterBusResetGraceful    RequestedPowerState = 14
	RequestedPowerStatePowerCycleOffSoftGraceful RequestedPowerState = 15
	RequestedPowerStatePowerCycleOffHardGraceful RequestedPowerState = 16
)
