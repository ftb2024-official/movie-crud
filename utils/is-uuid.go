package utils

import "regexp"

func IsUUID(id string) bool {
	uuidPattern := "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[1-5][0-9a-fA-F]{3}-[89aAbB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$"
	regex := regexp.MustCompile(uuidPattern)

	return regex.MatchString(id)
}
