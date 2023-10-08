package luhn

// CheckLuhn проверяет номер карты на соответствие алгоритму Луна
func CheckLuhn(cardNumber string) bool {
	var sum int
	var alt bool
	for i := len(cardNumber) - 1; i >= 0; i-- {
		if !alt {
			sum += int(cardNumber[i] - '0')
		} else {
			val := 2 * int(cardNumber[i]-'0')
			if val > 9 {
				val -= 9
			}
			sum += val
		}
		alt = !alt
	}
	return sum%10 == 0
}
