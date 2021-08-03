package gemini

// A Handler responds to a Gemini (or Naked) request.
type Handler interface {
	ServeGemini(ResponseWriter, string)
}
