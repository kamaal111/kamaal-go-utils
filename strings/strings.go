package strings

func Contains(items []string, searchValue string) bool {
	for _, item := range items {
		if item == searchValue {
			return true
		}
	}
	return false
}
