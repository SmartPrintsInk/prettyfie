package prettyfie

import "encoding/json"

func Pretty(data any) string {
	pretty, _ := json.MarshalIndent(data, " ", "  ")
	return string(pretty)
}

func DisBytes(data []byte) string {
	var aux any

	json.Unmarshal(data, &aux)
	pretty, _ := json.MarshalIndent(aux, " ", "  ")
	return string(pretty)
}
