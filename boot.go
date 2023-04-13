//go:generate stringer -type=bootConfigRole -trimprefix=bootConfigRole -linecomment

package amt

import (
	"fmt"
	"strconv"

	"github.com/VictorLowther/simplexml/dom"
	"github.com/VictorLowther/simplexml/search"
	"github.com/ammmze/go-amt/cim"
	"github.com/ammmze/wsman"
)

type bootConfigRole int

const (
	bootConfigRoleIsNext          bootConfigRole = 0
	bootConfigRoleIsNextSingleUse bootConfigRole = 1
	bootConfigRoleIsDefault       bootConfigRole = 2
	bootConfigRoleDMTFReserved    bootConfigRole = 3     // 3..32767
	bootConfigRoleVendorSpecified bootConfigRole = 32768 // 32768..65535
)

// SetPXE makes sure the node will pxe boot next time.
func (c *Client) SetPXE() error {
	// clear existing boot order per meshcommander's implementation...
	// "Set the boot order to null, this is needed for some AMT versions that don't clear this automatically."
	// err := changeBootOrder(client, []string{})
	// if err != nil {
	// 	return err
	// }

	if err := setBootSettingData(c); err != nil {
		return err
	}

	if err := setBootConfigRole(c, bootConfigRoleIsNextSingleUse); err != nil {
		return err
	}

	if err := changeBootOrder(c, []string{"Intel(r) AMT: Force PXE Boot"}); err != nil {
		return err
	}

	return nil
}

func setBootConfigRole(client *Client, role bootConfigRole) error {
	bootConfigRef, err := getBootConfigSettingRef(client, "Intel(r) AMT: Boot Configuration 0")
	if err != nil {
		return err
	}

	message := client.wsManClient.Invoke(resourceCIMBootService, "SetBootConfigRole")
	bootConfigSetting := message.MakeParameter("BootConfigSetting")
	bootConfigSetting.AddChildren(bootConfigRef.Children()...)
	message.AddParameter(bootConfigSetting)
	message.Parameters("Role", strconv.Itoa(int(role)))

	_, err = sendMessageForReturnValueInt(message)
	if err != nil {
		return err
	}
	return nil
}

func changeBootOrder(client *Client, items []string) error {
	message := client.wsManClient.Invoke(resourceCIMBootConfigSetting, "ChangeBootOrder")

	if len(items) > 0 {
		// TODO: multiple?
		pxeEndpointRef, err := getBootSourceRef(client, "Intel(r) AMT: Force PXE Boot")
		if err != nil {
			return err
		}
		sourceParam := message.MakeParameter("Source")
		sourceParam.AddChildren(pxeEndpointRef.Children()...)
		message.AddParameter(sourceParam)
	}

	_, err := sendMessageForReturnValueInt(message)
	if err != nil {
		return err
	}
	return nil
}

func getBootSettingData(client *Client) ([]*dom.Element, error) {
	msg := client.wsManClient.Get(resourceAMTBootSettingData)
	response, err := msg.Send()
	if err != nil {
		return nil, err
	}
	data := search.FirstTag("AMT_BootSettingData", resourceAMTBootSettingData, response.Body())
	if data == nil {
		return nil, fmt.Errorf("response was missing the AMT_BootSettingData")
	}
	return data.Children(), nil
}

func setBootSettingData(client *Client) error {
	bootSettings, err := getBootSettingData(client)
	if err != nil {
		return err
	}

	settingsToKeep := []*dom.Element{}
	for _, setting := range bootSettings {
		switch setting.Name.Local {
		// omit these ones ... they are read-only parameters (per meshcommand implementation)
		case "WinREBootEnabled",
			"UEFILocalPBABootEnabled",
			"UEFIHTTPSBootEnabled",
			"SecureBootControlEnabled",
			"BootguardStatus",
			"OptionsCleared",
			"BIOSLastStatus",
			"UefiBootParametersArray":
			continue
		// gonna make sure these are set to "false"
		case "BIOSPause", "BIOSSetup":
			setting.Content = []byte("false")
			settingsToKeep = append(settingsToKeep, setting)
		// gonna make sure these are set to "0"
		case "BootMediaIndex":
			setting.Content = []byte("0")
			settingsToKeep = append(settingsToKeep, setting)
		default:
			settingsToKeep = append(settingsToKeep, setting)
		}
	}

	msg := client.wsManClient.Put(resourceAMTBootSettingData)
	data := dom.Elem("AMT_BootSettingData", resourceAMTBootSettingData)
	data.AddChildren(settingsToKeep...)
	msg.SetBody(data)
	_, err = msg.Send()

	return err
}

func getBootConfigSettingRef(client *Client, name string) (*dom.Element, error) {
	return getEndpointReferenceByInstanceID(client, resourceCIMBootConfigSetting, name)
}

func getBootSourceRef(client *Client, name string) (*dom.Element, error) {
	return getEndpointReferenceByInstanceID(client, resourceCIMBootSourceSetting, name)
}

func getStructuredBootOrder(client *Client) ([]string, error) {
	msg, err := client.wsManClient.EnumerateObjectAndEPR("http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootConfigSetting").Send()
	if err != nil {
		return nil, err
	}
	bootConfigs := []cim.BootConfigSettingTypeAndEPRItem{}
	err = msg.UnmarshalEnumItems(&bootConfigs)
	if err != nil {
		return nil, err
	}
	if len(bootConfigs) == 0 {
		return nil, fmt.Errorf("no boot config was found. cannot get boot order")
	}
	msg, err = client.wsManClient.EnumerateObjectAndEPR("http://schemas.dmtf.org/wbem/wscim/1/*").
		AddFilterWithDialect("http://schemas.dmtf.org/wbem/wsman/1/cimbinding/associationFilter", wsman.MarshalToElementOrPanic(
			cim.AssociationInstances{
				Object:          bootConfigs[0].EndpointReference,
				Role:            "GroupComponent",
				ResultClassName: "CIM_OrderedComponent",
			})).
		Send()
	if err != nil {
		return nil, err
	}
	bootSourceSettings := []cim.OrderedComponentTypeAndEPRItem{}
	msg.UnmarshalEnumItems(&bootSourceSettings)
	bootSources := []string{}
	for _, src := range bootSourceSettings {
		msg, err := client.wsManClient.Get(src.OrderedComponent.PartComponent.ReferenceParameters.ResourceURI).
			AddSelector(wsman.MarshalToElementOrPanic(src.OrderedComponent.PartComponent.ReferenceParameters.SelectorSet).Children()...).
			Send()
		if err != nil {
			return nil, err
		}
		settings := &cim.BootSourceSettingType{}
		err = msg.UnmarshalBody(settings)
		if err != nil {
			return nil, err
		}
		bootSources = append(bootSources, settings.StructuredBootString)
	}

	return bootSources, nil
}
