package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"net/url"

	"github.com/fatih/structs"
)

func FixData(data string) []byte {
	if data == "" {
		data = "{}"
	}
	return []byte(data)
}

func ExtractMap(params map[string]interface{}, extract []string) map[string]interface{} {
	payload := make(map[string]interface{})

	for _, item := range extract {
		data, ok := params[item]
		if ok {
			payload[item] = data
		}
	}

	return payload
}

func EncryptPassword(data []byte) string {
	hash := sha256.New()
	hash.Write(data)
	md := hash.Sum(nil)
	mdStr := hex.EncodeToString(md)
	return mdStr
}

func ParseParams(url url.Values) map[string]interface{} {
	data := make(map[string]interface{})
	for item, _ := range url {
		data[item] = url.Get(item)
	}
	return data
}

// GetUpdateMap converts an structure to a map of string keys, to not nil values on the structure.
// This function is used to update the structure values because gorm also updates the default values
// errasing previous ones if passed in an structure.
func GetUpdateMap(structure interface{}) map[string]interface{} {
	var theMap = make(map[string]interface{})
	for _, field := range structs.Fields(structure) {
		if _, ok := field.Value().(bool); ok || !field.IsZero() {
			theMap[field.Name()] = field.Value()
		}
	}
	return theMap
}
