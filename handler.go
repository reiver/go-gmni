package gemini

// A Handler responds to a Gemini Protocol (or Naked) request.
type Handler interface {
	ServeGemini(ResponseWriter, Request)
}
