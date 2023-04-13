package cim

import (
	"encoding/xml"
)

const (
	AssociatedPowerManagementServiceSpace = "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_AssociatedPowerManagementService"
)

type AssociatedPowerManagementServiceServiceProvided CimReference
type AssociatedPowerManagementServicePowerState uint16
type AssociatedPowerManagementServiceOtherPowerState CimString
type AssociatedPowerManagementServiceRequestedPowerState uint16
type AssociatedPowerManagementServiceOtherRequestedPowerState CimString
type AssociatedPowerManagementServicePowerOnTime CimDateTime
type AssociatedPowerManagementServiceAvailableRequestedPowerStates uint16
type AssociatedPowerManagementServiceTransitioningToPowerState uint16
type AssociatedPowerManagementServiceUserOfService CimReference

type AssociatedPowerManagementServiceType struct {
	XMLName                       xml.Name                                                        `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_AssociatedPowerManagementService CIM_AssociatedPowerManagementService"`
	AvailableRequestedPowerStates []AssociatedPowerManagementServiceAvailableRequestedPowerStates `xml:"AvailableRequestedPowerStates,omitempty"`
	OtherPowerState               AssociatedPowerManagementServiceOtherPowerState                 `xml:"OtherPowerState,omitempty"`
	OtherRequestedPowerState      AssociatedPowerManagementServiceOtherRequestedPowerState        `xml:"OtherRequestedPowerState,omitempty"`
	PowerOnTime                   AssociatedPowerManagementServicePowerOnTime                     `xml:"PowerOnTime,omitempty"`
	PowerState                    AssociatedPowerManagementServicePowerState                      `xml:"PowerState,omitempty"`
	RequestedPowerState           AssociatedPowerManagementServiceRequestedPowerState             `xml:"RequestedPowerState,omitempty"`
	ServiceProvided               *AssociatedPowerManagementServiceServiceProvided                `xml:"ServiceProvided,omitempty"`
	TransitioningToPowerState     AssociatedPowerManagementServiceTransitioningToPowerState       `xml:"TransitioningToPowerState,omitempty"`
	UserOfService                 *AssociatedPowerManagementServiceUserOfService                  `xml:"UserOfService,omitempty"`
}

type AssociatedPowerManagementServiceAndEPRItem struct {
	ItemWithEndpointReference
	Service AssociatedPowerManagementServiceType `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_AssociatedPowerManagementService CIM_AssociatedPowerManagementService"`
}

type AssociatedPowerManagementService struct {
	client *CimClient
}

func (c *CimClient) AssociatedPowerManagementService() *AssociatedPowerManagementService {
	return &AssociatedPowerManagementService{client: c}
}

func (r *AssociatedPowerManagementService) EnumerateObjectAndEPR() ([]AssociatedPowerManagementServiceAndEPRItem, error) {
	var unmarshaledResponse []AssociatedPowerManagementServiceAndEPRItem
	if err := r.client.EnumerateObjectAndEPR(AssociatedPowerManagementServiceSpace, &unmarshaledResponse); err != nil {
		return nil, err
	}
	return unmarshaledResponse, nil
}

func (r *AssociatedPowerManagementService) Enumerate() ([]AssociatedPowerManagementServiceType, error) {
	var unmarshaledResponse []AssociatedPowerManagementServiceType
	if err := r.client.Enumerate(AssociatedPowerManagementServiceSpace, &unmarshaledResponse); err != nil {
		return nil, err
	}
	return unmarshaledResponse, nil
}

func (r *AssociatedPowerManagementService) EnumerateEPR() ([]EndpointReference, error) {
	var unmarshaledResponse []EndpointReference
	if err := r.client.EnumerateEPR(AssociatedPowerManagementServiceSpace, &unmarshaledResponse); err != nil {
		return nil, err
	}
	return unmarshaledResponse, nil
}
