package modtest

import (
	"errors"
	"fmt"
)

// Hi returns a friendly greeting in language lang
func Hi(name, lastName, lang string) (string, error) {
	switch lang {
	case "en":
		return fmt.Sprintf("Hi, %s! Or should I say: %v?", name, lastName), nil
	case "pt":
		return fmt.Sprintf("Oi, %s! Or should I say: %v?", name, lastName), nil
	case "gn":
		return fmt.Sprintf("Sup, %s! Or should I say: %v?", name, lastName), nil
	case "es":
		return fmt.Sprintf("¡Hola, %s! Or should I say: %v?", name, lastName), nil
	case "fr":
		return fmt.Sprintf("Bonjour, %s! Or should I say: %v?", name, lastName), nil
	default:
		return "", errors.New("unknown language")
	}
}
