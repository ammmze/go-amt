package cim

import (
	"encoding/xml"
)

type EndpointReferenceItems struct {
	XMLName           xml.Name            `xml:"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd Items"`
	EndpointReference []EndpointReference `xml:"http://schemas.xmlsoap.org/ws/2004/08/addressing EndpointReference"`
}

type EndpointReference struct {
	CimReference
	XMLName             xml.Name                    `xml:"http://schemas.xmlsoap.org/ws/2004/08/addressing EndpointReference"`
	ReferenceParameters EndpointReferenceParameters `xml:"http://schemas.xmlsoap.org/ws/2004/08/addressing ReferenceParameters"`
}

type EndpointReferenceParameters struct {
	XMLName     xml.Name    `xml:"http://schemas.xmlsoap.org/ws/2004/08/addressing ReferenceParameters"`
	ResourceURI string      `xml:"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd ResourceURI"`
	SelectorSet SelectorSet `xml:"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd SelectorSet"`
}
