// Code generated by "stringer -type=RequestedPowerState -trimprefix=RequestedPowerState -linecomment"; DO NOT EDIT.

package cim

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[RequestedPowerStateOn-2]
	_ = x[RequestedPowerStateSleepLight-3]
	_ = x[RequestedPowerStateSleepDeep-4]
	_ = x[RequestedPowerStatePowerCycleOffSoft-5]
	_ = x[RequestedPowerStateOffHard-6]
	_ = x[RequestedPowerStateHibernateOffSoft-7]
	_ = x[RequestedPowerStateOffSoft-8]
	_ = x[RequestedPowerStatePowerCycleOffHard-9]
	_ = x[RequestedPowerStateMasterBusReset-10]
	_ = x[RequestedPowerStateDiagnosticInterruptNMI-11]
	_ = x[RequestedPowerStateOffSoftGraceful-12]
	_ = x[RequestedPowerStateOffHardGraceful-13]
	_ = x[RequestedPowerStateMasterBusResetGraceful-14]
	_ = x[RequestedPowerStatePowerCycleOffSoftGraceful-15]
	_ = x[RequestedPowerStatePowerCycleOffHardGraceful-16]
}

const _RequestedPowerState_name = "OnSleepLightSleepDeepPowerCycleOffSoftOffHardHibernateOffSoftOffSoftPowerCycleOffHardMasterBusResetDiagnosticInterruptNMIOffSoftGracefulOffHardGracefulMasterBusResetGracefulPowerCycleOffSoftGracefulPowerCycleOffHardGraceful"

var _RequestedPowerState_index = [...]uint8{0, 2, 12, 21, 38, 45, 61, 68, 85, 99, 121, 136, 151, 173, 198, 223}

func (i RequestedPowerState) String() string {
	i -= 2
	if i < 0 || i >= RequestedPowerState(len(_RequestedPowerState_index)-1) {
		return "RequestedPowerState(" + strconv.FormatInt(int64(i+2), 10) + ")"
	}
	return _RequestedPowerState_name[_RequestedPowerState_index[i]:_RequestedPowerState_index[i+1]]
}