package chapter7

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// BEFORE
	fmt.Println("[] run before test")

	m.Run()

	// AFTER
	fmt.Println("[] run after test")
}

func TestDrawCard(t *testing.T) {
	deckCards := 30

	drawed := DrawCard(&deckCards)

	if deckCards != 29 || drawed != "Monster Card" {
		t.Error("ERROR: Failed to draw card!")
	}

	fmt.Println("✔️ TestDrawCard Done")
}

func TestPlayCard(t *testing.T) {
	handCards := 5

	played := PlayCard(&handCards)

	assert.Equal(t, 4, handCards, "Play 1 card. 4 left")
	assert.Equal(t, "Spell Card", played, "Play 'Spell Card'")

	t.Run("Sub test play another card", func(t *testing.T) {
		assert.NotEqual(t, "Spell Cards", played, "Play another card")
		fmt.Println("✔️ Sub test play another card Done")
	})

	fmt.Println("✔️ TestPlayCard Done")
}

// Table Test
func TestPrintCard(t *testing.T) {
	testCases := []struct {
		name     string
		reqest   string
		expected string
	}{
		{
			name:     "Jinx",
			reqest:   "Jinx",
			expected: "Jinx",
		},
		{
			name:     "Nautilus",
			reqest:   "Nautilus",
			expected: "Nautilus",
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result := PrintCard(test.reqest)
			assert.Equal(t, test.expected, result)
		})
	}
}

func BenchmarkPlayCard(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		PrintCard("Dark Magician")
	}
}
