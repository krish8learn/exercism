package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	//panic("Please implement the ParseCard function")
    var cardValues = map[string]int{
        "ace":11,
        "two":2,
        "three":3,
        "four":4,
        "five":5,
        "six":6,
        "seven":7,
        "eight":8,
        "nine":9,
        "ten":10,
        "jack":10,
        "queen":10,
        "king":10,
        "other":0,
    }
	return cardValues[card]
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	//panic("Please implement the FirstTurn function")
    cardValueTotal := ParseCard(card1) + ParseCard(card2)
    dealerCardValue := ParseCard(dealerCard)
    result := ""
    switch {
        case card1 == "ace" && card2 == "ace":
		result = "P"
	case cardValueTotal == 21 && (dealerCardValue != 11 && dealerCardValue != 10): 
		result = "W"
	case cardValueTotal == 21 && (dealerCardValue == 11 || dealerCardValue == 10):
		result = "S"
	case cardValueTotal >= 17 && cardValueTotal <= 20:
		result = "S"
	case cardValueTotal >= 12 && cardValueTotal <= 16 && dealerCardValue >= 7:
		result = "H"
	case cardValueTotal >= 12 && cardValueTotal <= 16:
		result = "S"	
	case cardValueTotal <= 11:
		result = "H"
    }
	return result
}
