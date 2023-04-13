package cim

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/ammmze/wsman"
	"github.com/stretchr/testify/require"
)

func Test_ComputerSystem_Enumerate(t *testing.T) {
	// amtInfo := amt.Connection{
	// 	Host: "192.168.32.6",
	// 	User: "demo",
	// 	Pass: "dem0Pa$$word",
	// 	// Debug: true,
	// }
	// client, err := amt.NewClient(amtInfo)
	wsman, err := wsman.NewClient("http://192.168.32.6:16992/wsman", "demo", "dem0Pa$$word", true)
	require.NoError(t, err)
	client := NewClient(wsman)

	items, err := client.ComputerSystem().Enumerate()
	require.NoError(t, err)
	out, _ := xml.MarshalIndent(items, " ", "  ")
	fmt.Println(string(out))
}
func Test_ComputerSystem_EnumerateEPR(t *testing.T) {
	// amtInfo := amt.Connection{
	// 	Host: "192.168.32.6",
	// 	User: "demo",
	// 	Pass: "dem0Pa$$word",
	// 	// Debug: true,
	// }
	// client, err := amt.NewClient(amtInfo)
	wsman, err := wsman.NewClient("http://192.168.32.6:16992/wsman", "demo", "dem0Pa$$word", true)
	require.NoError(t, err)
	client := NewClient(wsman)

	client.ComputerSystem().EnumerateEPR()
}
