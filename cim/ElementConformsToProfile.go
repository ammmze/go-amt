package cim

import (
	"encoding/xml"
)

const (
	resourceUriElementConformsToProfile = "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ElementConformsToProfile"
)

// ElementConformsToProfile Client used to perform actions on the machine
type ElementConformsToProfile struct {
	client *CimClient
}

type ElementConformsToProfileItems struct {
	XMLName xml.Name                       `xml:"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd Items"`
	Items   []ElementConformsToProfileItem `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ElementConformsToProfile CIM_ElementConformsToProfile"`
}

type ElementConformsToProfileItem struct {
	XMLName            xml.Name           `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ElementConformsToProfile CIM_ElementConformsToProfile"`
	ConformantStandard ConformantStandard `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ElementConformsToProfile ConformantStandard"`
	ManagedElement     ManagedElement     `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ElementConformsToProfile ManagedElement"`
}

type ConformantStandard struct {
	XMLName             xml.Name            `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ElementConformsToProfile ConformantStandard"`
	Address             string              `xml:"http://schemas.xmlsoap.org/ws/2004/08/addressing Address"`
	ReferenceParameters ReferenceParameters `xml:"http://schemas.xmlsoap.org/ws/2004/08/addressing ReferenceParameters"`
}
type ManagedElement struct {
	XMLName             xml.Name            `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ElementConformsToProfile ManagedElement"`
	Address             string              `xml:"http://schemas.xmlsoap.org/ws/2004/08/addressing Address"`
	ReferenceParameters ReferenceParameters `xml:"http://schemas.xmlsoap.org/ws/2004/08/addressing ReferenceParameters"`
}

func (c *CimClient) ElementConformsToProfile() *ElementConformsToProfile {
	return &ElementConformsToProfile{client: c}
}

func (r *ElementConformsToProfile) Enumerate() ([]ElementConformsToProfileItem, error) {
	var unmarshaledResponse ElementConformsToProfileItems
	err := r.client.Enumerate(resourceUriElementConformsToProfile, &unmarshaledResponse)
	if err != nil {
		return nil, err
	}
	return unmarshaledResponse.Items, nil
}

func (r *ElementConformsToProfile) EnumerateEPR() ([]EndpointReference, error) {
	var unmarshaledResponse EndpointReferenceItems
	err := r.client.EnumerateEPR(resourceUriElementConformsToProfile, &unmarshaledResponse)
	if err != nil {
		return nil, err
	}
	return unmarshaledResponse.EndpointReference, nil
}
