package arraysandstrings

// IsUnique determines if a string has all unique characters
func IsUnique(stringData string) bool {
	var hashMap = make(map[rune]bool)
	for _, c := range stringData {
		//Ignore space code point (rune)
		if hashMap[c] && c != 32 {
			return false
		} else {
			hashMap[c] = true
		}
	}
	return true
}

// CheckPermutation returns true if a string is permutation of the other
// func CheckPermutation(string1, string2 string) bool {
// 	if len(string1) != len(string2) {
// 		return false
// 	}
// 	var hashMap = make(map[rune]map[rune]bool)
// 	for _, c := range string1 {
// 		_, ok := hashMap[c]
// 		if ok {
// 			var testMap = make(map[rune]bool)
// 			testMap[c] = true
// 		}
// 	}
// 	for _, c := range string2 {
// 		_, ok := hashMap[c]
// 		if !ok {
// 			return false
// 		} else {
// 			if len(hashMap[c]) > 0 {
// 				delete(hashMap[c], c)
// 			}
// 			delete(hashMap, c)
// 		}
// 	}
// 	fmt.Println(len(hashMap))
// 	return true
// }
