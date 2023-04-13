package cim

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/ammmze/wsman"
	"github.com/stretchr/testify/require"
)

func Test_ElementConformsToProfile_Enumerate(t *testing.T) {
	wsman, err := wsman.NewClient("http://192.168.32.6:16992/wsman", "demo", "dem0Pa$$word", true)
	require.NoError(t, err)
	client := NewClient(wsman)

	items, err := client.ElementConformsToProfile().Enumerate()
	require.NoError(t, err)
	out, _ := xml.MarshalIndent(items, " ", "  ")
	fmt.Println("result:")
	fmt.Println(string(out))

	// for _, item := range items {
	// 	var anyParams map[string]string
	// 	client.wsmanClient.Enumerate(item.ConformantStandard.ReferenceParameters.ResourceURI, &anyParams)
	// 	fmt.Println("result:")
	// 	fmt.Println(anyParams)
	// }
}

func Test_ElementConformsToProfile_EnumerateEPR(t *testing.T) {
	wsman, err := wsman.NewClient("http://192.168.32.6:16992/wsman", "demo", "dem0Pa$$word", true)
	require.NoError(t, err)
	client := NewClient(wsman)

	items, err := client.ElementConformsToProfile().EnumerateEPR()
	require.NoError(t, err)
	out, _ := xml.MarshalIndent(items, " ", "  ")
	fmt.Println("result:")
	fmt.Println(string(out))
}
