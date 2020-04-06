package util

import (
	"encoding/json"
	"strconv"
)

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

func ToInt(value interface{}) int {
	if result, ok := value.(int); ok {
		return result
	}
	return -1
}

func ToFloat64(value interface{}) float64 {
	if result, ok := value.(float64); ok {
		return result
	}
	return -1
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

func StrToFloat64(value string) float64 {
	if value == "" {
		return -1
	}
	float, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return -1
	}
	return float
}

func Float64ToStr(value float64) string {
	return strconv.FormatFloat(value, 'f', -1, 64)
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
