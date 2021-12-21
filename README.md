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