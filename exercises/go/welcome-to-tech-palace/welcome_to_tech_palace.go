package techpalace

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return fmt.Sprintf("Welcome to the Tech Palace, %s", strings.ToUpper(customer))
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	// create stars string literal
	stars := ""
	for i := 0; i < numStarsPerLine; i++ {
		stars = stars + "*"
	}

	// return string message
	return stars + "\n" + welcomeMsg + "\n" + stars
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	noStars := strings.ReplaceAll(oldMsg, "*", "")

	return strings.TrimSpace(noStars)
}
