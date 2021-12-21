//go:generate stringer -type=bootConfigRole -trimprefix=bootConfigRole -linecomment

package amt

import (
	"fmt"
	"strconv"

	"github.com/VictorLowther/simplexml/dom"
	"github.com/VictorLowther/simplexml/search"
)

type bootConfigRole int

const (
	bootConfigRoleIsNext          bootConfigRole = 0
	bootConfigRoleIsNextSingleUse bootConfigRole = 1
	bootConfigRoleIsDefault       bootConfigRole = 2
	bootConfigRoleDMTFReserved    bootConfigRole = 3     // 3..32767
	bootConfigRoleVendorSpecified bootConfigRole = 32768 // 32768..65535
)

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

func setPXE(client *Client) error {
	// clear existing boot order per meshcommander's implementation...
	// "Set the boot order to null, this is needed for some AMT versions that don't clear this automatically."
	// err := changeBootOrder(client, []string{})
	// if err != nil {
	// 	return err
	// }

	err := setBootSettingData(client)
	if err != nil {
		return err
	}

	err = setBootConfigRole(client, bootConfigRoleIsNextSingleUse)
	if err != nil {
		return err
	}

	err = changeBootOrder(client, []string{"Intel(r) AMT: Force PXE Boot"})
	if err != nil {
		return err
	}

	return nil
}

func getBootConfigSettingRef(client *Client, name string) (*dom.Element, error) {
	return getEndpointReferenceByInstanceID(client, resourceCIMBootConfigSetting, name)
}

func getBootSourceRef(client *Client, name string) (*dom.Element, error) {
	return getEndpointReferenceByInstanceID(client, resourceCIMBootSourceSetting, name)
}
