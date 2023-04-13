package cim

import (
	"encoding/xml"
)

type AnyXMLTag struct {
	XMLName xml.Name
	Content string `xml:",innerxml"`
}

type CimDateTime string
type CimString string

type CimReference struct {
	Address             string                  `xml:"http://schemas.xmlsoap.org/ws/2004/08/addressing Address"`
	ReferenceParameters *CimReferenceParameters `xml:"http://schemas.xmlsoap.org/ws/2004/08/addressing ReferenceParameters"`
	Values              []AnyXMLTag             `xml:",any"`
}

type CimReferenceParameters struct {
	XMLName     xml.Name     `xml:"http://schemas.xmlsoap.org/ws/2004/08/addressing ReferenceParameters"`
	ResourceURI string       `xml:"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd ResourceURI"`
	SelectorSet *SelectorSet `xml:"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd SelectorSet"`
	Values      []*AnyXMLTag `xml:",any"`
}
