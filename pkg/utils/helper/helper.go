package helper

import "os"
import "encoding/json"

// MustGetEnv get environment value
func MustGetEnv(key string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return ""
	}
	return value
}


func AdjustStructToStruct(a interface{},b interface{}) interface{} {
	JsonStruct, _ := json.Marshal(a)
	json.Unmarshal([]byte(JsonStruct), &b)
	return b
}