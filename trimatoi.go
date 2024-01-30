package piscine

func TrimAtoi(s string) int {
	result := 0
	isNegative := false
	validDigit := false

	for _, char := range s {
		if char == '-' && result == 0 && !validDigit {
			isNegative = true
		} else if char >= '0' && char <= '9' {
			result = result*10 + int(char-'0')
			validDigit = true
		}
	}

	if isNegative {
		return -result
	}
	return result
}
