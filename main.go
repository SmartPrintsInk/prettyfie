package prettyfie

import "encoding/json"

func Pretty(data interface{}) string {
	pretty, _ := json.MarshalIndent(data, " ", "  ")
	return string(pretty)
}
