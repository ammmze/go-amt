package cim

import "encoding/xml"

const (
	resourceURIBootSourceSetting = "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting"
)

type BootSourceSettingType struct {
	XMLName              xml.Name    `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting CIM_BootSourceSetting"`
	BIOSBootString       string      `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting BIOSBootString"`
	BootString           string      `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting BootString"`
	Caption              string      `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting Caption"`        // max length 64
	ChangeableType       uint16      `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting ChangeableType"` // 0, 1, 2, 3
	ComponentSetting     []AnyXMLTag `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting ComponentSetting"`
	ConfigurationName    string      `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting ConfigurationName"`
	Description          string      `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting Description"`
	ElementName          string      `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting ElementName"`
	FailThroughSupported uint16      `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting FailThroughSupported"` // 0, 1, 2
	Generation           uint32      `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting Generation"`
	InstanceID           string      `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting InstanceID"`
	SoID                 string      `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting SoID"`
	StructuredBootString string      `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting StructuredBootString"`
	Values               []AnyXMLTag `xml:",any"`
}
