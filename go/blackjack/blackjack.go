package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
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
	case "jack", "queen", "king", "ten":
		return 10
	case "ace":
		return 11
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	// - Stand (S)
	// - Hit (H)
	// - Split (P)
	// - Automatically win (W)

	// - if you have a Blackjack (two cards that sum up to a value of 21), and the dealer does not have an ace,
	// a figure or a ten then you automatically win. If the dealer does have any of those cards then you'll have to stand and wait for the reveal of the other card.
	// - if your cards sum up to a value within the range [17, 20] you should always stand.
	// - if your cards sum up to a value within the range [12, 16] you should always stand unless the dealer has a 7 or higher, in which case you should always hit.
	// - if your cards sum up to 11 or lower you should always hit.
	sumCard := ParseCard(card1) + ParseCard(card2)
	isBlackjack := sumCard == 21
	// - If you have a pair of aces you must always split them.
	if card1 == "ace" && card2 == "ace" {
		return "P"
	}
	if isBlackjack && !(dealerCard == "jack" || dealerCard == "queen" || dealerCard == "king" || dealerCard == "ten" || dealerCard == "ace") {
		return "W"
	}
	if sumCard <= 11 || (12 <= sumCard && sumCard <= 16 && ParseCard(dealerCard) >= 7) {
		return "H"
	}
	return "S"
}
