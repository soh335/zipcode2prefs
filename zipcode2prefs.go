package zipcode2prefs

import "strings"

// Get return prefecture list via yubin no
// yubin no format expect numbers like that "1234567"
// some yubin no return multiple prefectures.
func Get(no string) []string {
	s, ok := prefs[no]
	if !ok {
		return nil
	}
	return strings.Split(s, ",")
}
