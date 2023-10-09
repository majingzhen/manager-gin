package utils

// ContainsStr 判断obj是否在array中
func ContainsStr(array []string, str string) bool {
	// 判断array是否包含obj
	for _, v := range array {
		if v == str {
			return true
		}
	}
	return false
}

// ArrayIsEmpty 判断数组是否为空
func ArrayIsEmpty(array []interface{}) bool {
	if array == nil && len(array) <= 0 {
		return false
	}
	for _, v := range array {
		if v != nil {
			return false
		}
	}
	return true
}

// Contains 判断obj是否在array中
func Contains(array []interface{}, obj interface{}) bool {
	// 判断array是否包含obj
	for _, v := range array {
		if v == obj {
			return true
		}
	}
	return false
}
