package connection

var (
	magicPreamble     = []byte{0x60, 0x60, 0xb0, 0x17}
	supportedVersions = []byte{
		0x00, 0x00, 0x00, 0x04, // bolt v4
		0x00, 0x00, 0x00, 0x03, // bolt v3
		0x00, 0x00, 0x00, 0x02, // bolt v2
		0x00, 0x00, 0x00, 0x01, // bolt v1
	}
	handShake          = append(magicPreamble, supportedVersions...)
	noVersionSupported = []byte{0x00, 0x00, 0x00, 0x00}
	// Version is the current version of this driver
	Version = "0.1"
	// ClientID is the id of this client
	ClientID = "GoBolt/" + Version
)

type QueryParams = map[string]interface{}

type NeoRows [][]interface{}

func (n *NeoRows) Get2DSlice() [][]interface{} {
	return *n
}
