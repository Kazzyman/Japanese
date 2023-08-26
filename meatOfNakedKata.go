package main

// **do-this**
import (
	"fmt"
	"regexp"
)

func meatOfNakedKataExorcise(in string) {
	// Used for processing either a Romaji or a Hiragana guess
	//
	// if 'in' is alpha
	var isAlphanumeric bool
	findAlphasIn := regexp.MustCompile(`[a-zA-Z]`)
	switch true {
	case findAlphasIn.MatchString(in): // <-- 'in'
		isAlphanumeric = true
	default:
		isAlphanumeric = false
	}

	// User has typed a Romaji -- aCard.KeyR would match the correct Romaji
	if isAlphanumeric == true && in == aCard.KeyR { // if user has typed an alpha, and it IS the correct Romaji
		fmt.Printf("%s", colorGreen)
		fmt.Printf("      　^^Right! ")
		fmt.Printf("%s", colorReset)
		//
		logHitsKata("Right", aCard.KeyK)
		logSkipThisPrompt(aCard.KeyR)
		//
		fmt.Println()
	} else if isAlphanumeric == true && in != aCard.KeyR { // if user typed an alpha, but NOT the correct Romaji
		fmt.Printf("%s", colorRed)
		fmt.Printf("      　^^Oops! ")
		//
		logHitsKata("Oops", aCard.KeyK)
		logKataGottenWrong(aCard.KeyK + ":it was:" + aCard.KeyR + ":but you had guessed:" + in)
		//
		fmt.Printf("%s", colorReset)
		fmt.Printf(" It was: ")
		fmt.Printf("%s", colorCyan)
		fmt.Printf("%s", aCard.KeyR)
		fmt.Printf("%s", colorReset)
		// check for hints ... the old way, via hira (so sad!, so fix this, by using the NEW way)
		checkForHints(aCard.KeyH) // we only give hints for non-composites **do-this**
		// In this FIRST case: hint card should be found via KeyK and only the KeyH field should be shown, and maybe some other info sans Romaji
	}

	if isAlphanumeric == false && in == aCard.KeyH { // if user HAS typed the CORRECT Hiragana corresponding to the Kata prompt
		fmt.Printf("%s", colorGreen)
		fmt.Printf("      　^^Right! ")
		fmt.Printf("%s", colorReset)
		//
		logHitsKata("Right", aCard.KeyK)
		logSkipThisPrompt(aCard.KeyK)
		//
		fmt.Println()
	} else if isAlphanumeric == false && in != aCard.KeyH { // user has typed the INCORRECT Hiragana corresponding to the Kata prompt
		fmt.Printf("%s", colorRed)
		fmt.Printf("      　^^Oops! ")
		//
		logHitsKata("Oops", aCard.KeyK)
		logKataGottenWrong(aCard.KeyK + ":it was:" + aCard.KeyH + ":but you had guessed:" + in)
		//
		fmt.Printf("%s", colorReset)
		fmt.Printf(" It was: ")
		fmt.Printf("%s", colorCyan)
		fmt.Printf("%s", aCard.KeyH)
		fmt.Printf("%s", colorReset)
		checkForHints(aCard.KeyH) // we only give hints for non-composites **do-this**
		// In this SECOND case: of having typed a Hiragana, all four lines of hints can be displayed from the card
		// ... though, not in the first case: of having typed a Romaji (then only the last line can be shown ??)
	}
}
