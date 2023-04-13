package cim

import "encoding/xml"

type ReferenceParameters struct {
	XMLName     xml.Name     `xml:"http://schemas.xmlsoap.org/ws/2004/08/addressing ReferenceParameters"`
	ResourceURI string       `xml:"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd ResourceURI,omitempty"`
	SelectorSet *SelectorSet `xml:"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd SelectorSet,omitempty"`
}

type SelectorSet struct {
	XMLName  xml.Name    `xml:"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd SelectorSet"`
	Selector []*Selector `xml:"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd Selector,omitempty"`
}

type Selector struct {
	XMLName           xml.Name           `xml:"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd Selector"`
	Name              string             `xml:"Name,attr,omitempty"`
	EndpointReference *EndpointReference `xml:"http://schemas.xmlsoap.org/ws/2004/08/addressing EndpointReference,omitempty"`
	Value             string             `xml:",chardata"`
}

type Reference struct {
	Address             string               `xml:"http://schemas.xmlsoap.org/ws/2004/08/addressing Address,omitempty"`
	ReferenceParameters *ReferenceParameters `xml:"http://schemas.xmlsoap.org/ws/2004/08/addressing ReferenceParameters,omitempty"`
}
