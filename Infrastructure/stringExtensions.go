package Infrastructure

import "unicode"

func StringCoalesce(value *string, defaultValue string) string {
	if IsNullOrWhiteSpace(value) {
		return defaultValue
	}
	return *value
}

func StringPtr(value string) *string {
	if IsEmptyOrWhiteSpace(value) {
		return nil
	}
	return &value
}

func StringArrPtr(value []byte) *string {
	if len(value) == 0 {
		return nil
	}
	str := string(value)
	if IsEmptyOrWhiteSpace(str) {
		return nil
	}
	return &str
}

func IsNullOrWhiteSpace(value *string) bool {
	if value == nil {
		return true
	}
	val := *value
	if IsEmptyOrWhiteSpace(val) {
		return true
	}
	return false
}

func IsEmptyOrWhiteSpace(value string) bool {
	return IsEmpty(value) || IsWhiteSpace(value)
}

func IsEmpty(value string) bool {
	return len(value) == 0
}

func IsWhiteSpace(value string) bool {
	for _, c := range value {
		if !unicode.IsSpace(c) {
			return false
		}
	}
	return true
}
