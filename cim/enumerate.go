package cim

import (
	"encoding/xml"

	"github.com/VictorLowther/simplexml/dom"
	"github.com/ammmze/wsman"
)

type Filter struct {
	XMLName xml.Name `xml:"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd Filter"`
	Dialect string   `xml:"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd Dialect,attr,omitempty"`
}

type AssociationInstancesFilter struct {
	*Filter
	Value *AssociationInstances `xml:"http://schemas.dmtf.org/wbem/wsman/1/cimbinding.xsd AssociationInstances"`
}

type AssociationInstances struct {
	XMLName               xml.Name      `xml:"http://schemas.dmtf.org/wbem/wsman/1/cimbinding.xsd AssociationInstances"`
	Object                *CimReference `xml:"http://schemas.dmtf.org/wbem/wsman/1/cimbinding.xsd Object"`
	Role                  string        `xml:"http://schemas.dmtf.org/wbem/wsman/1/cimbinding.xsd Role,omitempty"`
	ResultClassName       string        `xml:"http://schemas.dmtf.org/wbem/wsman/1/cimbinding.xsd ResultClassName,omitempty"`
	ResultRole            string        `xml:"http://schemas.dmtf.org/wbem/wsman/1/cimbinding.xsd ResultRole,omitempty"`
	IncludeResultProperty []string      `xml:"http://schemas.dmtf.org/wbem/wsman/1/cimbinding.xsd IncludeResultProperty,omitempty"`
}

type Item struct {
	XMLName xml.Name `xml:"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd Item"`
}

type ItemWithEndpointReference struct {
	Item
	EndpointReference *CimReference `xml:"http://schemas.xmlsoap.org/ws/2004/08/addressing EndpointReference,omitempty"`
}

func (item *Item) DomElement() (*dom.Element, error) {
	elem, err := wsman.MarshalToElement(item)
	if err != nil {
		return nil, err
	}
	return elem, nil
}

func (item *Item) DomElementOrPanic() *dom.Element {
	elem, err := item.DomElement()
	if err != nil {
		panic(err)
	}
	return elem
}

func (c *CimClient) Enumerate(resource string, unmarshaledResponse interface{}) error {
	return sendAndUnmarshalEnumerateMessage(c.wsmanClient.Enumerate(resource), unmarshaledResponse)
}

func (c *CimClient) EnumerateEPR(resource string, unmarshaledResponse interface{}) error {
	return sendAndUnmarshalEnumerateMessage(c.wsmanClient.EnumerateEPR(resource), unmarshaledResponse)
}

func (c *CimClient) EnumerateObjectAndEPR(resource string, unmarshaledResponse interface{}) error {
	return sendAndUnmarshalEnumerateMessage(c.wsmanClient.EnumerateObjectAndEPR(resource), unmarshaledResponse)
}

func sendAndUnmarshalEnumerateMessage(msg *wsman.Message, unmarshaledResponse interface{}) error {
	response, err := msg.Send()
	if err != nil {
		return err
	}

	if err := response.UnmarshalEnumItems(unmarshaledResponse); err != nil {
		return err
	}

	return nil

	// enumResponse := response.GetBody(dom.Elem("EnumerateResponse", wsman.NS_WSMEN))
	// if enumResponse == nil {
	// 	return fmt.Errorf("no EnumerateResponse was present in the response. %v", response)
	// }

	// fmt.Println(enumResponse)

	// items := search.FirstTag("Items", wsman.NS_WSMAN, enumResponse.Children())
	// if items == nil {
	// 	return fmt.Errorf("no Items was present in the response. %v", response)
	// }

	// xml.Unmarshal(items.Bytes(), &unmarshaledResponse)

	// return nil
}
