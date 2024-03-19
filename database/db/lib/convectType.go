package lib

import (
	"fmt"
	"time"
)

func ConvertValue(targetValue interface{}, dataType string, length int) (interface{}, error) {
	switch dataType {
	case "int":
		intValue, ok := targetValue.(int)
		if !ok {
			return nil, fmt.Errorf("cannot convert value to int")
		}
		return intValue, nil
	case "string":
		stringValue, ok := targetValue.(string)
		if !ok {
			return nil, fmt.Errorf("cannot convert value to string")
		}
		return stringValue, nil
	case "float":
		floatValue, ok := targetValue.(float64)
		if !ok {
			return nil, fmt.Errorf("cannot convert value to float")
		}
		return floatValue, nil
	case "varchar":
		if length <= 0 {
			return nil, fmt.Errorf("invalid length for varchar type")
		}
		stringValue, ok := targetValue.(string)
		if !ok {
			return nil, fmt.Errorf("cannot convert value to varchar")
		}
		if len(stringValue) > length {
			return nil, fmt.Errorf("length of varchar is too large")
		}
		return stringValue, nil
	case "byte":
		if length <= 0 {
			return nil, fmt.Errorf("invalid length for byte type")
		}
		byteValue, ok := targetValue.([]byte)
		if !ok {
			return nil, fmt.Errorf("cannot convert value to byte")
		}
		if len(byteValue) > length {
			return nil, fmt.Errorf("length of byte is too large")
		}
		return byteValue, nil
	case "date":
		dateValue, ok := targetValue.(time.Time)
		if !ok {
			return nil, fmt.Errorf("cannot convert value to date")
		}
		return dateValue, nil
	case "bool":
		boolValue, ok := targetValue.(bool)
		if !ok {
			return nil, fmt.Errorf("cannot convert value to bool")
		}
		return boolValue, nil
	case "intArray":
		intArray, ok := targetValue.([]int)
		if !ok {
			return nil, fmt.Errorf("cannot convert value to int array")
		}
		return intArray, nil
	case "stringArray":
		stringArray, ok := targetValue.([]string)
		if !ok {
			return nil, fmt.Errorf("cannot convert value to string array")
		}
		return stringArray, nil
	case "floatArray":
		floatArray, ok := targetValue.([]float64)
		if !ok {
			return nil, fmt.Errorf("cannot convert value to float array")
		}
		return floatArray, nil
	case "varcharArray":
		varcharArray, ok := targetValue.([]string)
		if !ok {
			return nil, fmt.Errorf("cannot convert value to varchar array")
		}
		for _, value := range varcharArray {
			if len(value) > length {
				return nil, fmt.Errorf("varchar array contains values longer than maximum length")
			}
		}
		return varcharArray, nil
	case "boolArray":
		boolArray, ok := targetValue.([]bool)
		if !ok {
			return nil, fmt.Errorf("cannot convert value to bool array")
		}
		return boolArray, nil
	default:
		return nil, fmt.Errorf("unsupported data type: %s", dataType)
	}
}
