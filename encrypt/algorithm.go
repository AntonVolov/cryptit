package encrypt 

func Nimbus(str string) string {
	encryptedString := ""
	for _, c := range str {
		asciiCode := int(c)
		charecter := string(asciiCode + 3)
		encryptedString += charecter
	}
	return encryptedString
}