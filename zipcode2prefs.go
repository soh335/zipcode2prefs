package zipcode2prefs

import "strings"

// Get return prefecture list via zipcode
// zipcode format expect numbers like that "1234567"
// some zipcode return multiple prefectures. Because
// there are duplicated.
func Get(no string) []string {
	s, ok := prefs[no]
	if !ok {
		return nil
	}
	return strings.Split(s, ",")
}
