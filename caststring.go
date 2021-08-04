package gemini

import (
	"fmt"
)

func caststring(value interface{}) string {
	switch casted := value.(type) {
	case interface{String()(string,error)}:
		s, err := casted.String()
		if nil != err {
			return fmt.Sprintf("%v", value)
		}
		return s
	case interface{String()(string,bool)}:
		s, ok := casted.String()
		if !ok {
			return fmt.Sprintf("%v", value)
		}
		return s
	case fmt.Stringer:
		return casted.String()
	case []byte:
		return string(casted)
	case rune:
		return string(casted)
	case []rune:
		return string(casted)
	case string:
		return casted
	default:
		return fmt.Sprintf("%v", value)
	}
}
