package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten":
		return 10
	case "jack":
		return 10
	case "queen":
		return 10
	case "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	cards := ParseCard(card1) + ParseCard(card2)
	dealer := ParseCard(dealerCard)

	switch {
	case cards == 22:
		return "P"
	case cards == 21:
		return HighRange(dealer)
	case cards >= 17 && cards <= 20:
		return "S"
	case cards >= 12 && cards <= 16:
		return LowRange(dealer)
	default:
		return "H"
	}
}

func HighRange(dealer int) string {
	if dealer == 10 || dealer == 11 {
		return "S"
	} else {
		return "W"
	}
}

func LowRange(dealer int) string {
	if dealer < 7 {
		return "S"
	} else {
		return "H"
	}
}
