package main

func contains(arrayOfStrings []string, string_item string) bool {
	for _, val := range arrayOfStrings {
		if val == string_item {
			return true
		}
	}
	return false
}
