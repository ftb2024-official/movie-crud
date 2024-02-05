package utils

import "regexp"

func IsISBN(isbn string) bool {
	regex := `^(?:(978|979)-\d{1,5}-\d{1,7}-\d{1,7}-\d|\d{1,9}-\d{1,5}-\d{1,7}-\d{1,7}-\d)(?:-\d|)$`
	match, _ := regexp.MatchString(regex, isbn)
	return match
}
