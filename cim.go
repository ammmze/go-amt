package amt

import (
	"fmt"
	"strconv"

	"github.com/VictorLowther/simplexml/dom"
	"github.com/VictorLowther/simplexml/search"
	"github.com/ammmze/wsman"
)

const (
	resourceCIMAssociatedPowerManagementService = "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_AssociatedPowerManagementService"
	resourceCIMBootConfigSetting                = "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootConfigSetting"
	resourceCIMBootService                      = "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootService"
	resourceCIMBootSourceSetting                = "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting"
	resourceCIMPowerManagementService           = "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_PowerManagementService"
	resourceCIMComputerSystem                   = "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ComputerSystem"
)

func getEndpointReferenceBySelector(client *Client, namespace string, selectorName string, selectorValue string) (*dom.Element, error) {
	message := client.wsManClient.EnumerateEPR(namespace)

	response, err := message.Send()
	if err != nil {
		return nil, err
	}
	items, err := response.EnumItems()
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		selector := search.First(search.Attr("Name", "*", selectorName), item.Descendants())
		if selector == nil || string(selector.Content) != selectorValue {
			continue
		}
		return item, nil
	}
	return nil, fmt.Errorf("could not find endpoint reference with selector %s=%s", selectorName, selectorValue)
}

func getEndpointReferenceByInstanceID(client *Client, namespace string, instanceID string) (*dom.Element, error) {
	return getEndpointReferenceBySelector(client, namespace, "InstanceID", instanceID)
}

func getComputerSystemRef(client *Client, name string) (*dom.Element, error) {
	return getEndpointReferenceBySelector(client, resourceCIMComputerSystem, "Name", name)
}

func getReturnValueInt(response *wsman.Message) (int, error) {
	returnElement := search.FirstTag("ReturnValue", "*", response.AllBodyElements())
	if returnElement == nil {
		return -1, fmt.Errorf("no ReturnValue found in the response")
	}
	return strconv.Atoi(string(returnElement.Content))
}

func sendMessageForReturnValueInt(message *wsman.Message) (int, error) {
	response, err := message.Send()
	if err != nil {
		return -1, err
	}
	returnValue, err := getReturnValueInt(response)

	if err != nil {
		return -1, err
	}

	if returnValue == 0 {
		return returnValue, nil
	}

	return returnValue, fmt.Errorf("received invalid return value %d", returnValue)
}
