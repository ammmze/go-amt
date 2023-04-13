package cim

import (
	"encoding/xml"
	"fmt"
)

const (
	resourceUriComputerSystem = "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ComputerSystem"
)

// ComputerSystem Client used to perform actions on the machine
type ComputerSystem struct {
	client *CimClient
}

type ComputerSystemEnumerateResponse struct {
	XMLName xml.Name            `xml:"http://schemas.xmlsoap.org/ws/2004/09/enumeration EnumerateResponse"`
	Items   ComputerSystemItems `xml:"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd Items"`
}

type ComputerSystemItems struct {
	XMLName xml.Name             `xml:"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd Items"`
	Items   []ComputerSystemItem `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ComputerSystem CIM_ComputerSystem"`
}

type ComputerSystemItem struct {
	XMLName                 xml.Name `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ComputerSystem CIM_ComputerSystem"`
	CreationClassName       string   `xml:"CreationClassName,omitempty"`
	Dedicated               uint16   `xml:"Dedicated,omitempty"`
	ElementName             string   `xml:"ElementName,omitempty"`
	EnabledDefault          uint16   `xml:"EnabledDefault,omitempty"`
	EnabledState            uint16   `xml:"EnabledState,omitempty"`
	HealthState             uint16   `xml:"HealthState,omitempty"`
	IdentifyingDescriptions string   `xml:"IdentifyingDescriptions,omitempty"`
	Name                    string   `xml:"Name,omitempty"`
	NameFormat              string   `xml:"NameFormat,omitempty"`
	OperationalStatus       uint16   `xml:"OperationalStatus,omitempty"`
	OtherIdentifyingInfo    string   `xml:"OtherIdentifyingInfo,omitempty"`
	RequestedState          uint16   `xml:"RequestedState,omitempty"`
}

func (c *CimClient) ComputerSystem() *ComputerSystem {
	return &ComputerSystem{client: c}
}

func (r *ComputerSystem) Enumerate() ([]ComputerSystemItem, error) {
	var unmarshaledResponse ComputerSystemItems
	err := r.client.Enumerate(resourceUriComputerSystem, &unmarshaledResponse)
	if err != nil {
		return nil, err
	}
	return unmarshaledResponse.Items, nil
}

func (r *ComputerSystem) EnumerateEPR() {
	msg := r.client.wsmanClient.EnumerateEPR(resourceUriComputerSystem)
	response, err := msg.Send()
	fmt.Printf("Error: %v\n", err)
	fmt.Printf("Response: %v\n", response)
}
