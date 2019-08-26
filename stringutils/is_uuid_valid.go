package stringutils

import "regexp"

func IsUuidValid(uuid string) bool {
	// sample UUID: 2ab23561-58f5-4aa4-b58c-edd3a3592e40
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid)
}
