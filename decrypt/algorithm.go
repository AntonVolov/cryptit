package decrypt

func Bimbus(input string) string {
	decryptedStr := ""
	for _, letter := range input {
		decryptedLetter := string(int(letter)-3)
		decryptedStr += decryptedLetter
	}
	return decryptedStr
}