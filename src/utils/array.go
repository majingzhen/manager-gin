package utils

// Contains 判断obj是否在array中
func Contains(array []string, str string) bool {
	// 判断array是否包含obj
	for _, v := range array {
		if v == str {
			return true
		}
	}
	return false
}
