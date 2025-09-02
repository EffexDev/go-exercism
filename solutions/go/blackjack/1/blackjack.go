package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	cardValues := map[string]int{
		"ace":   11,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"ten":   10,
		"jack":  10,
		"queen": 10,
		"king":  10,
	}
	if val, ok := cardValues[card]; ok {
		return val
	}
	return 0
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	firstCard := ParseCard(card1)
	secondCard := ParseCard(card2)
	dealer := ParseCard(dealerCard)
	sum := firstCard + secondCard

	switch {
	case firstCard == secondCard:
		switch {
		case sum >= 20 && sum != 22:
			return "S"
		case sum <= 10:
			return "H"
		default:
			return "P"
		}

	case sum == 21:
		if dealer == 11 || dealer == 10 {
			return "S"
		}
		return "W"

	case sum <= 17:
		if sum == 17 {
			return "S"
		}
		if dealer == 11 || (dealer < 11 && dealer > 6) || (sum == 10 && dealer == 2) {
			return "H"
		}
		return "S"

	case sum >= 18:
		return "S"
	}

	return ""
}
