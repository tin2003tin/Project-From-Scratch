package lib

func CompareInt(value int, operator string, target interface{}) bool {
	targetValue, ok := target.(int)
	if !ok {
		return false
	}
	switch operator {
	case "=":
		return value == targetValue
	case "<":
		return value < targetValue
	case "<=":
		return value <= targetValue
	case ">":
		return value > targetValue
	case ">=":
		return value >= targetValue
	default:
		return false
	}
}

// compareFloat64 performs comparison for float64 values
func CompareFloat64(value float64, operator string, target interface{}) bool {
	targetValue, ok := target.(float64)
	if !ok {
		return false
	}
	switch operator {
	case "=":
		return value == targetValue
	case "<":
		return value < targetValue
	case "<=":
		return value <= targetValue
	case ">":
		return value > targetValue
	case ">=":
		return value >= targetValue
	default:
		return false
	}
}

// compareString performs comparison for string values
func CompareString(value string, operator string, target interface{}) bool {
	targetValue, ok := target.(string)
	if !ok {
		return false
	}
	switch operator {
	case "=":
		return value == targetValue
	default:
		return false
	}
}