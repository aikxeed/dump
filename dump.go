package dump

import (
	"encoding/json"
	"fmt"
)

// Dump dumps the given variables in a pretty format
func Dump(variables ...interface{}) {
	var a []interface{}

	for _, v := range variables {
		m := map[string]interface{}{
			"type":  fmt.Sprintf("%T", v),
			"value": fmt.Sprintf("%+v", v),
		}
		a = append(a, m)
	}

	json, _ := json.MarshalIndent(a, "", "  ")
	fmt.Println(string(json))
}
