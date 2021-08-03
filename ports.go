package gemini

// The default TCP port used by a Naked server.
//
// A Naked server is basically a Gemini Protocol server without the TLS encryption.
// The reason one might want to create a Naked server is when the TLS encryption is handled elsewhere.
//
// The Gemini Protocol is meant to be encrypted by default.
// Unless you know what you are doing, don't create a Naked server.
const DefaultNakedServerPort = 1961

// The default TCP port used by a Gemini Protocol server.
const DefaultServerPort = 1965
