package prettyfie

import "encoding/json"

func This(data interface{}) string {
	pretty, _ := json.MarshalIndent(data, " ", "  ")
	return string(pretty)
}
