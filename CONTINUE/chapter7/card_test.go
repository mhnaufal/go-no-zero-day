package chapter7

import (
	"fmt"
	"testing"
)

func TestDrawCard(t *testing.T) {
	deckCards := 30

	drawed := DrawCard(&deckCards)

	if deckCards != 29 || drawed != "Monster Card" {
		t.Error("ERROR: Failed to draw card!")
	}

	fmt.Println("✔️ TestDrawCard Done")
}
