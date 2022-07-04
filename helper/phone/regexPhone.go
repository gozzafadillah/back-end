package regexPhone

import "regexp"

func GenerateNewPhone(phone string) string {
	var regex, _ = regexp.Compile(`[a-z]+`)
	regex.FindStringIndex(phone)

	lenght := len(phone)
	var regexPhone = phone[1:lenght]
	var newPhone = "+62" + regexPhone

	return newPhone
}
