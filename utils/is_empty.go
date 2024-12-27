package utils

func IsEmpty(value any) bool {
	switch v := value.(type) {
	case string:
		return v == ""
	case int, int8, int16, int32, int64:
		return v == 0
	case float32, float64:
		return v == 0.0
	case bool:
		return !v
	case nil:
		return true
	case []any:
		return len(v) == 0
	default:
		return value == nil
	}
}
