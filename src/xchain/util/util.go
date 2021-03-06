package util

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"hash"
	"time"
)

func NewId() int64 {
	return time.Now().UnixNano()
}

// SHA256 hashes using sha256 algorithm
func SHA256(text string) string {
	algorithm := sha256.New()
	return stringHasher(algorithm, text)
}

func stringHasher(algorithm hash.Hash, text string) string {
	algorithm.Write([]byte(text))
	return hex.EncodeToString(algorithm.Sum(nil))
}

func ToString(value interface{}) string {
	if result, ok := value.(string); ok {
		return result
	}
	return ""
}

func ToBool(value interface{}) bool {
	if result, ok := value.(bool); ok {
		return result
	}
	return false
}

func ToFloat64(value interface{}) float64 {
	if result, ok := value.(float64); ok {
		return result
	}
	return 0
}

func ToSlice(value interface{}) []interface{} {
	if result, ok := value.([]interface{}); ok {
		return result
	}
	return nil
}

func ToJson(value interface{}) string {
	bytes, err := json.Marshal(value)
	if err == nil {
		return string(bytes)
	}
	return ""
}

func JsonToSlice(str string) []interface{} {
	var value []interface{}
	err := json.Unmarshal([]byte(str), &value)
	if err == nil {
		return value
	}
	return nil
}

func JsonToMap(str string) map[string]interface{} {
	var value map[string]interface{}
	err := json.Unmarshal([]byte(str), &value)
	if err == nil {
		return value
	}
	return nil
}
