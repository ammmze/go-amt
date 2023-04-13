package cim

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/ammmze/wsman"
	"github.com/stretchr/testify/require"
)

func Test_AssociatedPowerManagementService_EnumerateObjectAndEPR(t *testing.T) {
	wsman, err := wsman.NewClient("http://192.168.32.6:16992/wsman", "demo", "dem0Pa$$word", true)
	require.NoError(t, err)
	client := NewClient(wsman)

	{
		items, err := client.AssociatedPowerManagementService().EnumerateObjectAndEPR()
		require.NoError(t, err)
		out, _ := xml.MarshalIndent(items, " ", "  ")
		fmt.Printf("Result: %s\n", string(out))
	}

	{
		items, err := client.AssociatedPowerManagementService().Enumerate()
		require.NoError(t, err)
		out, _ := xml.MarshalIndent(items, " ", "  ")
		fmt.Printf("Result: %s\n", string(out))
	}

	{
		items, err := client.AssociatedPowerManagementService().EnumerateEPR()
		require.NoError(t, err)
		out, _ := xml.MarshalIndent(items, " ", "  ")
		fmt.Printf("Result: %s\n", string(out))
	}
	// ps := PowerState(0)
	// fmt.Println(ps)
	// fmt.Println(items[0].Service.PowerState)

	// out, err = xml.Marshal(items[0].Service.UserOfService)
	// assert.NoError(t, err)
	// doc, err := dom.Parse(strings.NewReader(string(out)))
	// fmt.Printf("Result: %s\n", doc)

	// msg := client.Invoke("http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_PowerManagementService", "RequestPowerStateChange")
	// msg.Parameters("PowerState", fmt.Sprint(2))
	// managedElement := msg.MakeParameter("ManagedElement")
	// managedElement.AddChildren(doc.Root().Children()...)
	// msg.AddParameter(managedElement)
	// response, err := msg.Send()
	// assert.NoError(t, err)
	// fmt.Printf("Response: %v\n", response)

}
