package gemini

// A Request represents a Gemini Protocol request received by a server or to be sent by a client.
type Request struct {
	Target        string
	RemoteNetwork string
	RemoteAddress string
}
