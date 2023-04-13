package cim

import (
	"encoding/xml"
)

const (
	resourceURIBootConfigSetting = "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootConfigSetting"
)

type BootConfigSetting struct {
	client *CimClient
}

type BootConfigSettingType struct {
	XMLName           xml.Name    `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootConfigSetting CIM_BootConfigSetting"`
	Caption           string      `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootConfigSetting Caption,omitempty"`
	ChangeableType    uint16      `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootConfigSetting ChangeableType,omitempty"`
	ComponentSetting  []AnyXMLTag `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootConfigSetting ComponentSetting"`
	ConfigurationName string      `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootConfigSetting ConfigurationName,omitempty"`
	Description       string      `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootConfigSetting Description,omitempty"`
	ElementName       string      `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootConfigSetting ElementName"`
	Generation        uint64      `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootConfigSetting Generation,omitempty"`
	InstanceID        string      `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootConfigSetting InstanceID"`
	SoID              string      `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootConfigSetting SoID,omitempty"`
	SoOrgID           string      `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootConfigSetting SoOrgID,omitempty"`
	Values            []AnyXMLTag `xml:",any"`
}

type BootConfigSettingTypeAndEPRItem struct {
	ItemWithEndpointReference `xml:"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd Item"`
	BootConfigSetting         BootConfigSettingType `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootConfigSetting CIM_BootConfigSetting"`
}

type BootConfigSettingTypeAndEPRItems struct {
	XMLName xml.Name                          `xml:"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd Items"`
	Items   []BootConfigSettingTypeAndEPRItem `xml:"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd Item"`
}

func (c *CimClient) BootConfigSetting() *BootConfigSetting {
	return &BootConfigSetting{client: c}
}

func (r *BootConfigSetting) EnumerateObjectAndEPR() ([]BootConfigSettingTypeAndEPRItem, error) {
	var unmarshaledResponse BootConfigSettingTypeAndEPRItems
	err := r.client.EnumerateObjectAndEPR(resourceURIBootConfigSetting, &unmarshaledResponse)
	if err != nil {
		return nil, err
	}
	return unmarshaledResponse.Items, nil
}
