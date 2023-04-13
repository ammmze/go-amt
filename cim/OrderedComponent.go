package cim

import (
	"encoding/xml"
)

const (
	resourceURIOrderedComponent = "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_OrderedComponent"
)

type OrderedComponent struct {
	client *CimClient
}

type OrderedComponentType struct {
	XMLName          xml.Name     `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_OrderedComponent CIM_OrderedComponent"`
	AssignedSequence uint64       `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_OrderedComponent AssignedSequence,omitempty"`
	GroupComponent   CimReference `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_OrderedComponent GroupComponent"`
	PartComponent    CimReference `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_OrderedComponent PartComponent"`
	Values           []AnyXMLTag  `xml:",any"`
}

type OrderedComponentTypeAndEPRItem struct {
	XMLName           xml.Name             `xml:"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd Item"`
	OrderedComponent  OrderedComponentType `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_OrderedComponent CIM_OrderedComponent"`
	EndpointReference EndpointReference    `xml:"http://schemas.xmlsoap.org/ws/2004/08/addressing EndpointReference"`
}

type OrderedComponentTypeAndEPRItems struct {
	XMLName xml.Name                         `xml:"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd Items"`
	Items   []OrderedComponentTypeAndEPRItem `xml:"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd Item"`
}

func (c *CimClient) OrderedComponent() *OrderedComponent {
	return &OrderedComponent{client: c}
}

func (r *OrderedComponent) EnumerateObjectAndEPR() ([]OrderedComponentTypeAndEPRItem, error) {
	var unmarshaledResponse OrderedComponentTypeAndEPRItems
	err := r.client.EnumerateObjectAndEPR(resourceURIOrderedComponent, &unmarshaledResponse)
	if err != nil {
		return nil, err
	}
	return unmarshaledResponse.Items, nil
}
