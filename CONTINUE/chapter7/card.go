package chapter7

import "fmt"

func UnitTest() {
	// UNIT TEST
	fmt.Println("[1] UNIT TEST")

	println()
}

func DrawCard(cardLeft *int) string {
	*cardLeft = *cardLeft - 1

	return "Monster Card"
}

func PlayCard(handCards *int) string {
	*handCards = *handCards - 1

	return "Spell Card"
}

func PrintCard(cardName string) string {
	return cardName
}
