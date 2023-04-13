# go-amt
Client library to interact with Intel AMT api (via wsman)


## Usage

```golang
connection := Connection{
    Host: "192.168.32.6",
    User: "admin",
    Pass: "yourreallyawesomepassword",
}
client, err := NewClient(&connection)
assert.NoError(t, err)

err = client.PowerOn()
assert.NoError(t, err)
```


```
root@docker-desktop:/tmp# dashcli -h 192.168.32.6 -p 16992 -u demo -P 'dem0Pa$$word' -C -v 2 -t bootconfig[0] changebootorder 1 0



--------------------------------------------
<?xml version="1.0"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" xmlns:wsa="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:wsman="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns:wsen="http://schemas.xmlsoap.org/ws/2004/09/enumeration">
  <s:Header>
    <wsa:Action s:mustUnderstand="true">http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate</wsa:Action>
    <wsa:To s:mustUnderstand="true">http://192.168.32.6:16992/wsman</wsa:To>
    <wsman:ResourceURI s:mustUnderstand="true">http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ElementConformsToProfile</wsman:ResourceURI>
    <wsa:MessageID s:mustUnderstand="true">uuid:be4d223b-d3ad-13ad-8002-010000005002</wsa:MessageID>
    <wsa:ReplyTo>
      <wsa:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</wsa:Address>
    </wsa:ReplyTo>
  </s:Header>
  <s:Body>
    <wsen:Enumerate/>
  </s:Body>
</s:Envelope>

<?xml version="1.0" encoding="UTF-8"?>
<a:Envelope xmlns:a="http://www.w3.org/2003/05/soap-envelope" xmlns:b="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:c="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns:d="http://schemas.xmlsoap.org/ws/2005/02/trust" xmlns:e="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" xmlns:f="http://schemas.dmtf.org/wbem/wsman/1/cimbinding.xsd" xmlns:g="http://schemas.xmlsoap.org/ws/2004/09/enumeration" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <a:Header>
    <b:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:To>
    <b:RelatesTo>uuid:be4d223b-d3ad-13ad-8002-010000005002</b:RelatesTo>
    <b:Action a:mustUnderstand="true">http://schemas.xmlsoap.org/ws/2004/09/enumeration/EnumerateResponse</b:Action>
    <b:MessageID>uuid:00000000-8086-8086-8086-000000010DA0</b:MessageID>
    <c:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ElementConformsToProfile</c:ResourceURI>
  </a:Header>
  <a:Body>
    <g:EnumerateResponse>
      <g:EnumerationContext>B41C0000-0000-0000-0000-000000000000</g:EnumerationContext>
    </g:EnumerateResponse>
  </a:Body>
</a:Envelope>




--------------------------------------------
<?xml version="1.0"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" xmlns:wsa="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:wsman="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns:wsen="http://schemas.xmlsoap.org/ws/2004/09/enumeration">
  <s:Header>
    <wsa:Action s:mustUnderstand="true">http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull</wsa:Action>
    <wsa:To s:mustUnderstand="true">http://192.168.32.6:16992/wsman</wsa:To>
    <wsman:ResourceURI s:mustUnderstand="true">http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ElementConformsToProfile</wsman:ResourceURI>
    <wsa:MessageID s:mustUnderstand="true">uuid:be4e6396-d3ad-13ad-8003-010000005002</wsa:MessageID>
    <wsa:ReplyTo>
      <wsa:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</wsa:Address>
    </wsa:ReplyTo>
  </s:Header>
  <s:Body>
    <wsen:Pull>
      <wsen:EnumerationContext>B41C0000-0000-0000-0000-000000000000</wsen:EnumerationContext>
    </wsen:Pull>
  </s:Body>
</s:Envelope>

<?xml version="1.0" encoding="UTF-8"?>
<a:Envelope xmlns:a="http://www.w3.org/2003/05/soap-envelope" xmlns:b="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:c="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns:d="http://schemas.xmlsoap.org/ws/2005/02/trust" xmlns:e="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" xmlns:f="http://schemas.dmtf.org/wbem/wsman/1/cimbinding.xsd" xmlns:g="http://schemas.xmlsoap.org/ws/2004/09/enumeration" xmlns:h="http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ElementConformsToProfile" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <a:Header>
    <b:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:To>
    <b:RelatesTo>uuid:be4e6396-d3ad-13ad-8003-010000005002</b:RelatesTo>
    <b:Action a:mustUnderstand="true">http://schemas.xmlsoap.org/ws/2004/09/enumeration/PullResponse</b:Action>
    <b:MessageID>uuid:00000000-8086-8086-8086-000000010DA1</b:MessageID>
    <c:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ElementConformsToProfile</c:ResourceURI>
  </a:Header>
  <a:Body>
    <g:PullResponse>
      <g:Items>
        <h:CIM_ElementConformsToProfile>
          <h:ConformantStandard>
            <b:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:Address>
            <b:ReferenceParameters>
              <c:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_RegisteredProfile</c:ResourceURI>
              <c:SelectorSet>
                <c:Selector Name="InstanceID">Intel(r) ME:Base Desktop and Mobile</c:Selector>
              </c:SelectorSet>
            </b:ReferenceParameters>
          </h:ConformantStandard>
          <h:ManagedElement>
            <b:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:Address>
            <b:ReferenceParameters>
              <c:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ComputerSystem</c:ResourceURI>
              <c:SelectorSet>
                <c:Selector Name="CreationClassName">CIM_ComputerSystem</c:Selector>
                <c:Selector Name="Name">ManagedSystem</c:Selector>
              </c:SelectorSet>
            </b:ReferenceParameters>
          </h:ManagedElement>
        </h:CIM_ElementConformsToProfile>
      </g:Items>
      <g:EndOfSequence/>
    </g:PullResponse>
  </a:Body>
</a:Envelope>




--------------------------------------------
<?xml version="1.0"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" xmlns:wsa="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:wsman="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns:wsen="http://schemas.xmlsoap.org/ws/2004/09/enumeration">
  <s:Header>
    <wsa:Action s:mustUnderstand="true">http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate</wsa:Action>
    <wsa:To s:mustUnderstand="true">http://192.168.32.6:16992/wsman</wsa:To>
    <wsman:ResourceURI s:mustUnderstand="true">http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootConfigSetting</wsman:ResourceURI>
    <wsa:MessageID s:mustUnderstand="true">uuid:be4ed46d-d3ad-13ad-8004-010000005002</wsa:MessageID>
    <wsa:ReplyTo>
      <wsa:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</wsa:Address>
    </wsa:ReplyTo>
  </s:Header>
  <s:Body>
    <wsen:Enumerate>
      <wsman:OptimizeEnumeration/>
      <wsman:EnumerationMode>EnumerateObjectAndEPR</wsman:EnumerationMode>
      <wsman:MaxElements>250</wsman:MaxElements>
    </wsen:Enumerate>
  </s:Body>
</s:Envelope>

<?xml version="1.0" encoding="UTF-8"?>
<a:Envelope xmlns:a="http://www.w3.org/2003/05/soap-envelope" xmlns:b="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:c="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns:d="http://schemas.xmlsoap.org/ws/2005/02/trust" xmlns:e="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" xmlns:f="http://schemas.dmtf.org/wbem/wsman/1/cimbinding.xsd" xmlns:g="http://schemas.xmlsoap.org/ws/2004/09/enumeration" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <a:Header>
    <b:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:To>
    <b:RelatesTo>uuid:be4ed46d-d3ad-13ad-8004-010000005002</b:RelatesTo>
    <b:Action a:mustUnderstand="true">http://schemas.xmlsoap.org/ws/2004/09/enumeration/EnumerateResponse</b:Action>
    <b:MessageID>uuid:00000000-8086-8086-8086-000000010DA2</b:MessageID>
    <c:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootConfigSetting</c:ResourceURI>
  </a:Header>
  <a:Body>
    <g:EnumerateResponse>
      <g:EnumerationContext>B51C0000-0000-0000-0000-000000000000</g:EnumerationContext>
    </g:EnumerateResponse>
  </a:Body>
</a:Envelope>




--------------------------------------------
<?xml version="1.0"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" xmlns:wsa="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:wsman="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns:wsen="http://schemas.xmlsoap.org/ws/2004/09/enumeration">
  <s:Header>
    <wsa:Action s:mustUnderstand="true">http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull</wsa:Action>
    <wsa:To s:mustUnderstand="true">http://192.168.32.6:16992/wsman</wsa:To>
    <wsman:ResourceURI s:mustUnderstand="true">http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootConfigSetting</wsman:ResourceURI>
    <wsa:MessageID s:mustUnderstand="true">uuid:be4fb19f-d3ad-13ad-8005-010000005002</wsa:MessageID>
    <wsa:ReplyTo>
      <wsa:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</wsa:Address>
    </wsa:ReplyTo>
  </s:Header>
  <s:Body>
    <wsen:Pull>
      <wsen:EnumerationContext>B51C0000-0000-0000-0000-000000000000</wsen:EnumerationContext>
      <wsen:MaxElements>250</wsen:MaxElements>
    </wsen:Pull>
  </s:Body>
</s:Envelope>

<?xml version="1.0" encoding="UTF-8"?>
<a:Envelope xmlns:a="http://www.w3.org/2003/05/soap-envelope" xmlns:b="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:c="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns:d="http://schemas.xmlsoap.org/ws/2005/02/trust" xmlns:e="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" xmlns:f="http://schemas.dmtf.org/wbem/wsman/1/cimbinding.xsd" xmlns:g="http://schemas.xmlsoap.org/ws/2004/09/enumeration" xmlns:h="http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootConfigSetting" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <a:Header>
    <b:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:To>
    <b:RelatesTo>uuid:be4fb19f-d3ad-13ad-8005-010000005002</b:RelatesTo>
    <b:Action a:mustUnderstand="true">http://schemas.xmlsoap.org/ws/2004/09/enumeration/PullResponse</b:Action>
    <b:MessageID>uuid:00000000-8086-8086-8086-000000010DA3</b:MessageID>
    <c:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootConfigSetting</c:ResourceURI>
  </a:Header>
  <a:Body>
    <g:PullResponse>
      <g:Items>
        <c:Item>
          <h:CIM_BootConfigSetting>
            <h:ElementName>Intel(r) AMT: Boot Configuration</h:ElementName>
            <h:InstanceID>Intel(r) AMT: Boot Configuration 0</h:InstanceID>
          </h:CIM_BootConfigSetting>
          <b:EndpointReference>
            <b:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:Address>
            <b:ReferenceParameters>
              <c:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootConfigSetting</c:ResourceURI>
              <c:SelectorSet>
                <c:Selector Name="InstanceID">Intel(r) AMT: Boot Configuration 0</c:Selector>
              </c:SelectorSet>
            </b:ReferenceParameters>
          </b:EndpointReference>
        </c:Item>
      </g:Items>
      <g:EndOfSequence/>
    </g:PullResponse>
  </a:Body>
</a:Envelope>




--------------------------------------------
<?xml version="1.0"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" xmlns:wsa="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:wsman="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns:wsen="http://schemas.xmlsoap.org/ws/2004/09/enumeration" xmlns:wsmb="http://schemas.dmtf.org/wbem/wsman/1/cimbinding.xsd">
  <s:Header>
    <wsa:Action s:mustUnderstand="true">http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate</wsa:Action>
    <wsa:To s:mustUnderstand="true">http://192.168.32.6:16992/wsman</wsa:To>
    <wsman:ResourceURI s:mustUnderstand="true">http://schemas.dmtf.org/wbem/wscim/1/*</wsman:ResourceURI>
    <wsa:MessageID s:mustUnderstand="true">uuid:be500563-d3ad-13ad-8006-010000005002</wsa:MessageID>
    <wsa:ReplyTo>
      <wsa:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</wsa:Address>
    </wsa:ReplyTo>
  </s:Header>
  <s:Body>
    <wsen:Enumerate>
      <wsman:OptimizeEnumeration/>
      <wsman:EnumerationMode>EnumerateObjectAndEPR</wsman:EnumerationMode>
      <wsman:Filter Dialect="http://schemas.dmtf.org/wbem/wsman/1/cimbinding/associationFilter">
        <wsmb:AssociationInstances>
          <wsmb:Object>
            <wsa:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</wsa:Address>
            <wsa:ReferenceParameters>
              <wsman:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootConfigSetting</wsman:ResourceURI>
              <wsman:SelectorSet>
                <wsman:Selector Name="InstanceID">Intel(r) AMT: Boot Configuration 0</wsman:Selector>
              </wsman:SelectorSet>
            </wsa:ReferenceParameters>
          </wsmb:Object>
          <wsmb:Role>GroupComponent</wsmb:Role>
          <wsmb:ResultClassName>CIM_OrderedComponent</wsmb:ResultClassName>
        </wsmb:AssociationInstances>
      </wsman:Filter>
      <wsman:MaxElements>250</wsman:MaxElements>
    </wsen:Enumerate>
  </s:Body>
</s:Envelope>

<?xml version="1.0" encoding="UTF-8"?>
<a:Envelope xmlns:a="http://www.w3.org/2003/05/soap-envelope" xmlns:b="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:c="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns:d="http://schemas.xmlsoap.org/ws/2005/02/trust" xmlns:e="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" xmlns:f="http://schemas.dmtf.org/wbem/wsman/1/cimbinding.xsd" xmlns:g="http://schemas.xmlsoap.org/ws/2004/09/enumeration" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <a:Header>
    <b:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:To>
    <b:RelatesTo>uuid:be500563-d3ad-13ad-8006-010000005002</b:RelatesTo>
    <b:Action a:mustUnderstand="true">http://schemas.xmlsoap.org/ws/2004/09/enumeration/EnumerateResponse</b:Action>
    <b:MessageID>uuid:00000000-8086-8086-8086-000000010DA4</b:MessageID>
    <c:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/*</c:ResourceURI>
  </a:Header>
  <a:Body>
    <g:EnumerateResponse>
      <g:EnumerationContext>B61C0000-0000-0000-0000-000000000000</g:EnumerationContext>
    </g:EnumerateResponse>
  </a:Body>
</a:Envelope>




--------------------------------------------
<?xml version="1.0"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" xmlns:wsa="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:wsman="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns:wsen="http://schemas.xmlsoap.org/ws/2004/09/enumeration">
  <s:Header>
    <wsa:Action s:mustUnderstand="true">http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull</wsa:Action>
    <wsa:To s:mustUnderstand="true">http://192.168.32.6:16992/wsman</wsa:To>
    <wsman:ResourceURI s:mustUnderstand="true">http://schemas.dmtf.org/wbem/wscim/1/*</wsman:ResourceURI>
    <wsa:MessageID s:mustUnderstand="true">uuid:be53266e-d3ad-13ad-8007-010000005002</wsa:MessageID>
    <wsa:ReplyTo>
      <wsa:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</wsa:Address>
    </wsa:ReplyTo>
  </s:Header>
  <s:Body>
    <wsen:Pull>
      <wsen:EnumerationContext>B61C0000-0000-0000-0000-000000000000</wsen:EnumerationContext>
      <wsen:MaxElements>250</wsen:MaxElements>
    </wsen:Pull>
  </s:Body>
</s:Envelope>

<?xml version="1.0" encoding="UTF-8"?>
<a:Envelope xmlns:a="http://www.w3.org/2003/05/soap-envelope" xmlns:b="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:c="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns:d="http://schemas.xmlsoap.org/ws/2005/02/trust" xmlns:e="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" xmlns:f="http://schemas.dmtf.org/wbem/wsman/1/cimbinding.xsd" xmlns:g="http://schemas.xmlsoap.org/ws/2004/09/enumeration" xmlns:h="http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_OrderedComponent" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <a:Header>
    <b:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:To>
    <b:RelatesTo>uuid:be53266e-d3ad-13ad-8007-010000005002</b:RelatesTo>
    <b:Action a:mustUnderstand="true">http://schemas.xmlsoap.org/ws/2004/09/enumeration/PullResponse</b:Action>
    <b:MessageID>uuid:00000000-8086-8086-8086-000000010DA5</b:MessageID>
    <c:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/*</c:ResourceURI>
  </a:Header>
  <a:Body>
    <g:PullResponse>
      <g:Items>
        <c:Item>
          <h:CIM_OrderedComponent>
            <h:AssignedSequence>1</h:AssignedSequence>
            <h:GroupComponent>
              <b:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:Address>
              <b:ReferenceParameters>
                <c:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootConfigSetting</c:ResourceURI>
                <c:SelectorSet>
                  <c:Selector Name="InstanceID">Intel(r) AMT: Boot Configuration 0</c:Selector>
                </c:SelectorSet>
              </b:ReferenceParameters>
            </h:GroupComponent>
            <h:PartComponent>
              <b:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:Address>
              <b:ReferenceParameters>
                <c:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting</c:ResourceURI>
                <c:SelectorSet>
                  <c:Selector Name="InstanceID">Intel(r) AMT: Force Hard-drive Boot</c:Selector>
                </c:SelectorSet>
              </b:ReferenceParameters>
            </h:PartComponent>
          </h:CIM_OrderedComponent>
          <b:EndpointReference>
            <b:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:Address>
            <b:ReferenceParameters>
              <c:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_OrderedComponent</c:ResourceURI>
              <c:SelectorSet>
                <c:Selector Name="GroupComponent">
                  <b:EndpointReference>
                    <b:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:Address>
                    <b:ReferenceParameters>
                      <c:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootConfigSetting</c:ResourceURI>
                      <c:SelectorSet>
                        <c:Selector Name="InstanceID">Intel(r) AMT: Boot Configuration 0</c:Selector>
                      </c:SelectorSet>
                    </b:ReferenceParameters>
                  </b:EndpointReference>
                </c:Selector>
                <c:Selector Name="PartComponent">
                  <b:EndpointReference>
                    <b:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:Address>
                    <b:ReferenceParameters>
                      <c:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting</c:ResourceURI>
                      <c:SelectorSet>
                        <c:Selector Name="InstanceID">Intel(r) AMT: Force Hard-drive Boot</c:Selector>
                      </c:SelectorSet>
                    </b:ReferenceParameters>
                  </b:EndpointReference>
                </c:Selector>
              </c:SelectorSet>
            </b:ReferenceParameters>
          </b:EndpointReference>
        </c:Item>
        <c:Item>
          <h:CIM_OrderedComponent>
            <h:AssignedSequence>2</h:AssignedSequence>
            <h:GroupComponent>
              <b:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:Address>
              <b:ReferenceParameters>
                <c:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootConfigSetting</c:ResourceURI>
                <c:SelectorSet>
                  <c:Selector Name="InstanceID">Intel(r) AMT: Boot Configuration 0</c:Selector>
                </c:SelectorSet>
              </b:ReferenceParameters>
            </h:GroupComponent>
            <h:PartComponent>
              <b:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:Address>
              <b:ReferenceParameters>
                <c:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting</c:ResourceURI>
                <c:SelectorSet>
                  <c:Selector Name="InstanceID">Intel(r) AMT: Force PXE Boot</c:Selector>
                </c:SelectorSet>
              </b:ReferenceParameters>
            </h:PartComponent>
          </h:CIM_OrderedComponent>
          <b:EndpointReference>
            <b:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:Address>
            <b:ReferenceParameters>
              <c:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_OrderedComponent</c:ResourceURI>
              <c:SelectorSet>
                <c:Selector Name="GroupComponent">
                  <b:EndpointReference>
                    <b:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:Address>
                    <b:ReferenceParameters>
                      <c:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootConfigSetting</c:ResourceURI>
                      <c:SelectorSet>
                        <c:Selector Name="InstanceID">Intel(r) AMT: Boot Configuration 0</c:Selector>
                      </c:SelectorSet>
                    </b:ReferenceParameters>
                  </b:EndpointReference>
                </c:Selector>
                <c:Selector Name="PartComponent">
                  <b:EndpointReference>
                    <b:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:Address>
                    <b:ReferenceParameters>
                      <c:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting</c:ResourceURI>
                      <c:SelectorSet>
                        <c:Selector Name="InstanceID">Intel(r) AMT: Force PXE Boot</c:Selector>
                      </c:SelectorSet>
                    </b:ReferenceParameters>
                  </b:EndpointReference>
                </c:Selector>
              </c:SelectorSet>
            </b:ReferenceParameters>
          </b:EndpointReference>
        </c:Item>
      </g:Items>
      <g:EndOfSequence/>
    </g:PullResponse>
  </a:Body>
</a:Envelope>




--------------------------------------------
<?xml version="1.0"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" xmlns:wsa="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:wsman="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns:wsen="http://schemas.xmlsoap.org/ws/2004/09/enumeration" xmlns:wsmb="http://schemas.dmtf.org/wbem/wsman/1/cimbinding.xsd">
  <s:Header>
    <wsa:Action s:mustUnderstand="true">http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate</wsa:Action>
    <wsa:To s:mustUnderstand="true">http://192.168.32.6:16992/wsman</wsa:To>
    <wsman:ResourceURI s:mustUnderstand="true">http://schemas.dmtf.org/wbem/wscim/1/*</wsman:ResourceURI>
    <wsa:MessageID s:mustUnderstand="true">uuid:be5427cc-d3ad-13ad-8008-010000005002</wsa:MessageID>
    <wsa:ReplyTo>
      <wsa:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</wsa:Address>
    </wsa:ReplyTo>
  </s:Header>
  <s:Body>
    <wsen:Enumerate>
      <wsman:OptimizeEnumeration/>
      <wsman:EnumerationMode>EnumerateObjectAndEPR</wsman:EnumerationMode>
      <wsman:Filter Dialect="http://schemas.dmtf.org/wbem/wsman/1/cimbinding/associationFilter">
        <wsmb:AssociationInstances>
          <wsmb:Object>
            <wsa:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</wsa:Address>
            <wsa:ReferenceParameters>
              <wsman:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootConfigSetting</wsman:ResourceURI>
              <wsman:SelectorSet>
                <wsman:Selector Name="InstanceID">Intel(r) AMT: Boot Configuration 0</wsman:Selector>
              </wsman:SelectorSet>
            </wsa:ReferenceParameters>
          </wsmb:Object>
          <wsmb:Role>GroupComponent</wsmb:Role>
          <wsmb:ResultClassName>CIM_OrderedComponent</wsmb:ResultClassName>
        </wsmb:AssociationInstances>
      </wsman:Filter>
      <wsman:MaxElements>250</wsman:MaxElements>
    </wsen:Enumerate>
  </s:Body>
</s:Envelope>

<?xml version="1.0" encoding="UTF-8"?>
<a:Envelope xmlns:a="http://www.w3.org/2003/05/soap-envelope" xmlns:b="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:c="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns:d="http://schemas.xmlsoap.org/ws/2005/02/trust" xmlns:e="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" xmlns:f="http://schemas.dmtf.org/wbem/wsman/1/cimbinding.xsd" xmlns:g="http://schemas.xmlsoap.org/ws/2004/09/enumeration" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <a:Header>
    <b:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:To>
    <b:RelatesTo>uuid:be5427cc-d3ad-13ad-8008-010000005002</b:RelatesTo>
    <b:Action a:mustUnderstand="true">http://schemas.xmlsoap.org/ws/2004/09/enumeration/EnumerateResponse</b:Action>
    <b:MessageID>uuid:00000000-8086-8086-8086-000000010DA6</b:MessageID>
    <c:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/*</c:ResourceURI>
  </a:Header>
  <a:Body>
    <g:EnumerateResponse>
      <g:EnumerationContext>B71C0000-0000-0000-0000-000000000000</g:EnumerationContext>
    </g:EnumerateResponse>
  </a:Body>
</a:Envelope>




--------------------------------------------
<?xml version="1.0"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" xmlns:wsa="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:wsman="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns:wsen="http://schemas.xmlsoap.org/ws/2004/09/enumeration">
  <s:Header>
    <wsa:Action s:mustUnderstand="true">http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull</wsa:Action>
    <wsa:To s:mustUnderstand="true">http://192.168.32.6:16992/wsman</wsa:To>
    <wsman:ResourceURI s:mustUnderstand="true">http://schemas.dmtf.org/wbem/wscim/1/*</wsman:ResourceURI>
    <wsa:MessageID s:mustUnderstand="true">uuid:be5722af-d3ad-13ad-8009-010000005002</wsa:MessageID>
    <wsa:ReplyTo>
      <wsa:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</wsa:Address>
    </wsa:ReplyTo>
  </s:Header>
  <s:Body>
    <wsen:Pull>
      <wsen:EnumerationContext>B71C0000-0000-0000-0000-000000000000</wsen:EnumerationContext>
      <wsen:MaxElements>250</wsen:MaxElements>
    </wsen:Pull>
  </s:Body>
</s:Envelope>

<?xml version="1.0" encoding="UTF-8"?>
<a:Envelope xmlns:a="http://www.w3.org/2003/05/soap-envelope" xmlns:b="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:c="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns:d="http://schemas.xmlsoap.org/ws/2005/02/trust" xmlns:e="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" xmlns:f="http://schemas.dmtf.org/wbem/wsman/1/cimbinding.xsd" xmlns:g="http://schemas.xmlsoap.org/ws/2004/09/enumeration" xmlns:h="http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_OrderedComponent" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <a:Header>
    <b:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:To>
    <b:RelatesTo>uuid:be5722af-d3ad-13ad-8009-010000005002</b:RelatesTo>
    <b:Action a:mustUnderstand="true">http://schemas.xmlsoap.org/ws/2004/09/enumeration/PullResponse</b:Action>
    <b:MessageID>uuid:00000000-8086-8086-8086-000000010DA7</b:MessageID>
    <c:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/*</c:ResourceURI>
  </a:Header>
  <a:Body>
    <g:PullResponse>
      <g:Items>
        <c:Item>
          <h:CIM_OrderedComponent>
            <h:AssignedSequence>1</h:AssignedSequence>
            <h:GroupComponent>
              <b:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:Address>
              <b:ReferenceParameters>
                <c:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootConfigSetting</c:ResourceURI>
                <c:SelectorSet>
                  <c:Selector Name="InstanceID">Intel(r) AMT: Boot Configuration 0</c:Selector>
                </c:SelectorSet>
              </b:ReferenceParameters>
            </h:GroupComponent>
            <h:PartComponent>
              <b:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:Address>
              <b:ReferenceParameters>
                <c:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting</c:ResourceURI>
                <c:SelectorSet>
                  <c:Selector Name="InstanceID">Intel(r) AMT: Force Hard-drive Boot</c:Selector>
                </c:SelectorSet>
              </b:ReferenceParameters>
            </h:PartComponent>
          </h:CIM_OrderedComponent>
          <b:EndpointReference>
            <b:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:Address>
            <b:ReferenceParameters>
              <c:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_OrderedComponent</c:ResourceURI>
              <c:SelectorSet>
                <c:Selector Name="GroupComponent">
                  <b:EndpointReference>
                    <b:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:Address>
                    <b:ReferenceParameters>
                      <c:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootConfigSetting</c:ResourceURI>
                      <c:SelectorSet>
                        <c:Selector Name="InstanceID">Intel(r) AMT: Boot Configuration 0</c:Selector>
                      </c:SelectorSet>
                    </b:ReferenceParameters>
                  </b:EndpointReference>
                </c:Selector>
                <c:Selector Name="PartComponent">
                  <b:EndpointReference>
                    <b:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:Address>
                    <b:ReferenceParameters>
                      <c:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting</c:ResourceURI>
                      <c:SelectorSet>
                        <c:Selector Name="InstanceID">Intel(r) AMT: Force Hard-drive Boot</c:Selector>
                      </c:SelectorSet>
                    </b:ReferenceParameters>
                  </b:EndpointReference>
                </c:Selector>
              </c:SelectorSet>
            </b:ReferenceParameters>
          </b:EndpointReference>
        </c:Item>
        <c:Item>
          <h:CIM_OrderedComponent>
            <h:AssignedSequence>2</h:AssignedSequence>
            <h:GroupComponent>
              <b:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:Address>
              <b:ReferenceParameters>
                <c:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootConfigSetting</c:ResourceURI>
                <c:SelectorSet>
                  <c:Selector Name="InstanceID">Intel(r) AMT: Boot Configuration 0</c:Selector>
                </c:SelectorSet>
              </b:ReferenceParameters>
            </h:GroupComponent>
            <h:PartComponent>
              <b:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:Address>
              <b:ReferenceParameters>
                <c:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting</c:ResourceURI>
                <c:SelectorSet>
                  <c:Selector Name="InstanceID">Intel(r) AMT: Force PXE Boot</c:Selector>
                </c:SelectorSet>
              </b:ReferenceParameters>
            </h:PartComponent>
          </h:CIM_OrderedComponent>
          <b:EndpointReference>
            <b:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:Address>
            <b:ReferenceParameters>
              <c:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_OrderedComponent</c:ResourceURI>
              <c:SelectorSet>
                <c:Selector Name="GroupComponent">
                  <b:EndpointReference>
                    <b:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:Address>
                    <b:ReferenceParameters>
                      <c:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootConfigSetting</c:ResourceURI>
                      <c:SelectorSet>
                        <c:Selector Name="InstanceID">Intel(r) AMT: Boot Configuration 0</c:Selector>
                      </c:SelectorSet>
                    </b:ReferenceParameters>
                  </b:EndpointReference>
                </c:Selector>
                <c:Selector Name="PartComponent">
                  <b:EndpointReference>
                    <b:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:Address>
                    <b:ReferenceParameters>
                      <c:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting</c:ResourceURI>
                      <c:SelectorSet>
                        <c:Selector Name="InstanceID">Intel(r) AMT: Force PXE Boot</c:Selector>
                      </c:SelectorSet>
                    </b:ReferenceParameters>
                  </b:EndpointReference>
                </c:Selector>
              </c:SelectorSet>
            </b:ReferenceParameters>
          </b:EndpointReference>
        </c:Item>
      </g:Items>
      <g:EndOfSequence/>
    </g:PullResponse>
  </a:Body>
</a:Envelope>




--------------------------------------------
<?xml version="1.0"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" xmlns:wsa="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:wsman="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd">
  <s:Header>
    <wsa:Action s:mustUnderstand="true">http://schemas.xmlsoap.org/ws/2004/09/transfer/Get</wsa:Action>
    <wsa:To s:mustUnderstand="true">http://192.168.32.6:16992/wsman</wsa:To>
    <wsman:ResourceURI s:mustUnderstand="true">http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting</wsman:ResourceURI>
    <wsa:MessageID s:mustUnderstand="true">uuid:be581c2d-d3ad-13ad-800a-010000005002</wsa:MessageID>
    <wsa:ReplyTo>
      <wsa:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</wsa:Address>
    </wsa:ReplyTo>
    <wsman:SelectorSet>
      <wsman:Selector Name="InstanceID">Intel(r) AMT: Force PXE Boot</wsman:Selector>
    </wsman:SelectorSet>
  </s:Header>
  <s:Body/>
</s:Envelope>

<?xml version="1.0" encoding="UTF-8"?>
<a:Envelope xmlns:a="http://www.w3.org/2003/05/soap-envelope" xmlns:b="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:c="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns:d="http://schemas.xmlsoap.org/ws/2005/02/trust" xmlns:e="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" xmlns:f="http://schemas.dmtf.org/wbem/wsman/1/cimbinding.xsd" xmlns:g="http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <a:Header>
    <b:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:To>
    <b:RelatesTo>uuid:be581c2d-d3ad-13ad-800a-010000005002</b:RelatesTo>
    <b:Action a:mustUnderstand="true">http://schemas.xmlsoap.org/ws/2004/09/transfer/GetResponse</b:Action>
    <b:MessageID>uuid:00000000-8086-8086-8086-000000010DA8</b:MessageID>
    <c:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting</c:ResourceURI>
  </a:Header>
  <a:Body>
    <g:CIM_BootSourceSetting>
      <g:ElementName>Intel(r) AMT: Boot Source</g:ElementName>
      <g:FailThroughSupported>2</g:FailThroughSupported>
      <g:InstanceID>Intel(r) AMT: Force PXE Boot</g:InstanceID>
      <g:StructuredBootString>CIM:Network:1</g:StructuredBootString>
    </g:CIM_BootSourceSetting>
  </a:Body>
</a:Envelope>




--------------------------------------------
<?xml version="1.0"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" xmlns:wsa="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:wsman="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd">
  <s:Header>
    <wsa:Action s:mustUnderstand="true">http://schemas.xmlsoap.org/ws/2004/09/transfer/Get</wsa:Action>
    <wsa:To s:mustUnderstand="true">http://192.168.32.6:16992/wsman</wsa:To>
    <wsman:ResourceURI s:mustUnderstand="true">http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting</wsman:ResourceURI>
    <wsa:MessageID s:mustUnderstand="true">uuid:be585506-d3ad-13ad-800b-010000005002</wsa:MessageID>
    <wsa:ReplyTo>
      <wsa:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</wsa:Address>
    </wsa:ReplyTo>
    <wsman:SelectorSet>
      <wsman:Selector Name="InstanceID">Intel(r) AMT: Force Hard-drive Boot</wsman:Selector>
    </wsman:SelectorSet>
  </s:Header>
  <s:Body/>
</s:Envelope>

<?xml version="1.0" encoding="UTF-8"?>
<a:Envelope xmlns:a="http://www.w3.org/2003/05/soap-envelope" xmlns:b="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:c="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns:d="http://schemas.xmlsoap.org/ws/2005/02/trust" xmlns:e="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" xmlns:f="http://schemas.dmtf.org/wbem/wsman/1/cimbinding.xsd" xmlns:g="http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <a:Header>
    <b:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:To>
    <b:RelatesTo>uuid:be585506-d3ad-13ad-800b-010000005002</b:RelatesTo>
    <b:Action a:mustUnderstand="true">http://schemas.xmlsoap.org/ws/2004/09/transfer/GetResponse</b:Action>
    <b:MessageID>uuid:00000000-8086-8086-8086-000000010DA9</b:MessageID>
    <c:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting</c:ResourceURI>
  </a:Header>
  <a:Body>
    <g:CIM_BootSourceSetting>
      <g:ElementName>Intel(r) AMT: Boot Source</g:ElementName>
      <g:FailThroughSupported>2</g:FailThroughSupported>
      <g:InstanceID>Intel(r) AMT: Force Hard-drive Boot</g:InstanceID>
      <g:StructuredBootString>CIM:Hard-Disk:1</g:StructuredBootString>
    </g:CIM_BootSourceSetting>
  </a:Body>
</a:Envelope>




--------------------------------------------
<?xml version="1.0"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" xmlns:wsa="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:wsman="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd">
  <s:Header>
    <wsa:Action s:mustUnderstand="true">http://schemas.xmlsoap.org/ws/2004/09/transfer/Get</wsa:Action>
    <wsa:To s:mustUnderstand="true">http://192.168.32.6:16992/wsman</wsa:To>
    <wsman:ResourceURI s:mustUnderstand="true">http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting</wsman:ResourceURI>
    <wsa:MessageID s:mustUnderstand="true">uuid:be589972-d3ad-13ad-800c-010000005002</wsa:MessageID>
    <wsa:ReplyTo>
      <wsa:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</wsa:Address>
    </wsa:ReplyTo>
    <wsman:SelectorSet>
      <wsman:Selector Name="InstanceID">Intel(r) AMT: Force PXE Boot</wsman:Selector>
    </wsman:SelectorSet>
  </s:Header>
  <s:Body/>
</s:Envelope>

<?xml version="1.0" encoding="UTF-8"?>
<a:Envelope xmlns:a="http://www.w3.org/2003/05/soap-envelope" xmlns:b="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:c="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns:d="http://schemas.xmlsoap.org/ws/2005/02/trust" xmlns:e="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" xmlns:f="http://schemas.dmtf.org/wbem/wsman/1/cimbinding.xsd" xmlns:g="http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <a:Header>
    <b:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:To>
    <b:RelatesTo>uuid:be589972-d3ad-13ad-800c-010000005002</b:RelatesTo>
    <b:Action a:mustUnderstand="true">http://schemas.xmlsoap.org/ws/2004/09/transfer/GetResponse</b:Action>
    <b:MessageID>uuid:00000000-8086-8086-8086-000000010DAA</b:MessageID>
    <c:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting</c:ResourceURI>
  </a:Header>
  <a:Body>
    <g:CIM_BootSourceSetting>
      <g:ElementName>Intel(r) AMT: Boot Source</g:ElementName>
      <g:FailThroughSupported>2</g:FailThroughSupported>
      <g:InstanceID>Intel(r) AMT: Force PXE Boot</g:InstanceID>
      <g:StructuredBootString>CIM:Network:1</g:StructuredBootString>
    </g:CIM_BootSourceSetting>
  </a:Body>
</a:Envelope>




--------------------------------------------
<?xml version="1.0"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" xmlns:wsa="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:wsman="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd">
  <s:Header>
    <wsa:Action s:mustUnderstand="true">http://schemas.xmlsoap.org/ws/2004/09/transfer/Get</wsa:Action>
    <wsa:To s:mustUnderstand="true">http://192.168.32.6:16992/wsman</wsa:To>
    <wsman:ResourceURI s:mustUnderstand="true">http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting</wsman:ResourceURI>
    <wsa:MessageID s:mustUnderstand="true">uuid:be58e17c-d3ad-13ad-800d-010000005002</wsa:MessageID>
    <wsa:ReplyTo>
      <wsa:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</wsa:Address>
    </wsa:ReplyTo>
    <wsman:SelectorSet>
      <wsman:Selector Name="InstanceID">Intel(r) AMT: Force Hard-drive Boot</wsman:Selector>
    </wsman:SelectorSet>
  </s:Header>
  <s:Body/>
</s:Envelope>

<?xml version="1.0" encoding="UTF-8"?>
<a:Envelope xmlns:a="http://www.w3.org/2003/05/soap-envelope" xmlns:b="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:c="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns:d="http://schemas.xmlsoap.org/ws/2005/02/trust" xmlns:e="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" xmlns:f="http://schemas.dmtf.org/wbem/wsman/1/cimbinding.xsd" xmlns:g="http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <a:Header>
    <b:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:To>
    <b:RelatesTo>uuid:be58e17c-d3ad-13ad-800d-010000005002</b:RelatesTo>
    <b:Action a:mustUnderstand="true">http://schemas.xmlsoap.org/ws/2004/09/transfer/GetResponse</b:Action>
    <b:MessageID>uuid:00000000-8086-8086-8086-000000010DAB</b:MessageID>
    <c:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting</c:ResourceURI>
  </a:Header>
  <a:Body>
    <g:CIM_BootSourceSetting>
      <g:ElementName>Intel(r) AMT: Boot Source</g:ElementName>
      <g:FailThroughSupported>2</g:FailThroughSupported>
      <g:InstanceID>Intel(r) AMT: Force Hard-drive Boot</g:InstanceID>
      <g:StructuredBootString>CIM:Hard-Disk:1</g:StructuredBootString>
    </g:CIM_BootSourceSetting>
  </a:Body>
</a:Envelope>




--------------------------------------------
<?xml version="1.0"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" xmlns:wsa="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:wsman="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns:n1="http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootConfigSetting">
  <s:Header>
    <wsa:Action s:mustUnderstand="true">http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootConfigSetting/ChangeBootOrder</wsa:Action>
    <wsa:To s:mustUnderstand="true">http://192.168.32.6:16992/wsman</wsa:To>
    <wsman:ResourceURI s:mustUnderstand="true">http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootConfigSetting</wsman:ResourceURI>
    <wsa:MessageID s:mustUnderstand="true">uuid:be5927a4-d3ad-13ad-800e-010000005002</wsa:MessageID>
    <wsa:ReplyTo>
      <wsa:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</wsa:Address>
    </wsa:ReplyTo>
    <wsman:SelectorSet>
      <wsman:Selector Name="InstanceID">Intel(r) AMT: Boot Configuration 0</wsman:Selector>
    </wsman:SelectorSet>
  </s:Header>
  <s:Body>
    <n1:ChangeBootOrder_INPUT>
      <n1:Source>
        <wsa:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</wsa:Address>
        <wsa:ReferenceParameters>
          <wsman:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting</wsman:ResourceURI>
          <wsman:SelectorSet>
            <wsman:Selector Name="InstanceID">Intel(r) AMT: Force PXE Boot</wsman:Selector>
          </wsman:SelectorSet>
        </wsa:ReferenceParameters>
      </n1:Source>
      <n1:Source>
        <wsa:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</wsa:Address>
        <wsa:ReferenceParameters>
          <wsman:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting</wsman:ResourceURI>
          <wsman:SelectorSet>
            <wsman:Selector Name="InstanceID">Intel(r) AMT: Force Hard-drive Boot</wsman:Selector>
          </wsman:SelectorSet>
        </wsa:ReferenceParameters>
      </n1:Source>
    </n1:ChangeBootOrder_INPUT>
  </s:Body>
</s:Envelope>

<?xml version="1.0" encoding="UTF-8"?>
<a:Envelope xmlns:a="http://www.w3.org/2003/05/soap-envelope" xmlns:b="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:c="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns:d="http://schemas.xmlsoap.org/ws/2005/02/trust" xmlns:e="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" xmlns:f="http://schemas.dmtf.org/wbem/wsman/1/cimbinding.xsd" xmlns:g="http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootConfigSetting" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <a:Header>
    <b:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:To>
    <b:RelatesTo>uuid:be5927a4-d3ad-13ad-800e-010000005002</b:RelatesTo>
    <b:Action a:mustUnderstand="true">http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootConfigSetting/ChangeBootOrderResponse</b:Action>
    <b:MessageID>uuid:00000000-8086-8086-8086-000000010DAC</b:MessageID>
    <c:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootConfigSetting</c:ResourceURI>
  </a:Header>
  <a:Body>
    <g:ChangeBootOrder_OUTPUT>
      <g:ReturnValue>0</g:ReturnValue>
    </g:ChangeBootOrder_OUTPUT>
  </a:Body>
</a:Envelope>


Boot Order Changed Successfully
```