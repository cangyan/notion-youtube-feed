package utils

func StringInArr(s string, arr []string) bool {
	for _, v := range arr {
		if s == v {
			return true
		}
	}

	return false
}
