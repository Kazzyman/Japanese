package main

// **do-this**
import (
	"fmt"
	"regexp"
)

/*
.
Refer to 8, which is a similar function

All this means that:
The first of these two functions is passed the user's first guess, and, optionally, obtains a second guess
Optionally, the second func is passed the second guess, and obtains a third guess
Finally, the first func is re-called by the second,: the first func is, then, passed the third and final guess
.
*/
func meatOfKataExercise(in string, skipFlag bool) {
	// Used for processing either a Romaji OR a Hiragana guess which was obtained prior to this func being called
	// ... either by TouchTypingExercises12346, or by: secondTry_ (secondTry_ actually solicits and obtains the THIRD guess)
	// If 'in' (the user's guess) is alpha v v v v v      (a first-instance of 'in' is the user's first guess)
	var isAlphanumeric bool
	findAlphasIn := regexp.MustCompile(`[a-zA-Z]`)
	switch true {
	case findAlphasIn.MatchString(in): // <-- 'in'
		isAlphanumeric = true
	default:
		isAlphanumeric = false
	}

	// First, we will deal with the case that ... the
	// user, typed an alpha (not necessarily his first guess) -- user, presumably, typed a Romaji
	/*
		.
		.　The assumed Romaji guess, though not necessarily his first guess, will be processed below ...
	*/
	// The following 'if' will be executed, potentially only twice: 1: for the first, and 2: for the final guess
	if isAlphanumeric == true && in == aCard.KeyR { // After this line we know both that it was a Romaji, and that it
		// ... IS the correct Romaji
		fmt.Printf("%s", colorGreen)
		fmt.Printf("      　^^Right! ")
		fmt.Printf("%s", colorReset)
		//
		logHits_in_cyclicArrayHits("Right", aCard.KeyK)
		logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCard.KeyR)
		//
		fmt.Println() // This line will be the last-executed before returning to the caller: TouchTypingExercises12346
		//
		// The following 'else' will be executed, potentially only twice: 1: for the first, and 2: for the final guess
	} else if isAlphanumeric == true && in != aCard.KeyR { // If user typed an alpha, but not the correct Romaji,
		// ... a second attempt will be solicited via: secondTry_meatOfKataExercise(in) -- UNLESS skipFlag is false
		fmt.Printf("%s", colorRed)
		fmt.Printf("      　^^Oops! ") // This ALWAYS prints unless ^^Right! is printed
		logHits_in_cyclicArrayHits("Oops", aCard.KeyK)
		logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCard.KeyK +
			":it was:" + aCard.KeyR + ":but you had guessed:" + in)
		logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.KeyR)
		// When this func is re-called by the second, with skipFlag set false, we skip this, and do the next
		if skipFlag == true {
			// Solicit the second guess ... and pass it to: secondTry_meatOfKataExercise(in)
			fmt.Println("Try again") // This ONLY prints on the first pass through this func
			fmt.Printf(colorReset)
			// Re-prompt, and will send that second guess to the secondTry_meatOfKataExercise func
			in = Prompt_Scan_4_Romaji_or_HiraResponse(aCard.KeyK) // Obtain second guess, and pass it as 'in'
				branchOnUserSelectedDirectiveIfGiven(in,
					"Respond_w_Hira_or_Romaji_to_kataPrompt_3") // <-- Perform the directive
			secondTry_meatOfKataExercise(in)    // This instance of 'in' is the user's second guess.
		}
		// If user guesses incorrectly on his third-and-final try, then, and only then, execute the rest of this func
		if skipFlag == false { // skipFlag is true on first entry, & false when recalled after third failed attempt
			fmt.Printf("%s", colorReset)
			fmt.Printf(" It was: ")
			fmt.Printf("%s", colorCyan)
			fmt.Printf("%s", aCard.KeyR)
			fmt.Printf("%s", colorReset)
			// Only the field lacking the correct Romaji will be shown (only the last field: HintSansR)
			// fmt.Printf("\n\n%s\n", aCard.HintSansR) // Optionally limiting hint to being sans all mention of Romaji 
			// Display hints in white
			fmt.Printf("\n\n%s\n%s\n%s\n", aCard.Hint1h, aCard.Hint2k, aCard.Hint3TT)
			fmt.Println("")
		}
	} // If user was ^^Right!, then we return to TouchTypingExercises12346(selectedExercise) (directly from this very line)
	/*
	   .
	   .
	   .　An assumed Hiragana guess will be processed below ...
	*/
	if isAlphanumeric == false && in == aCard.KeyH { // If user HAS typed the CORRECT Hiragana
		fmt.Printf("%s", colorGreen)
		fmt.Printf("      　^^Right! ")
		fmt.Printf("%s", colorReset)
		//
		logHits_in_cyclicArrayHits("Right", aCard.KeyK)
		logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCard.KeyR)
		//
		fmt.Println() // This line is the last-executed of this func, returns to TouchTypingExercises12346(SE)
		//
		//
	} else if isAlphanumeric == false && in != aCard.KeyH { // <-- The user has typed the INCORRECT Hiragana ...
		fmt.Printf("%s", colorRed)
		fmt.Printf("      　^^Oops! ") // This ALWAYS prints unless ^^Right! is printed
		logHits_in_cyclicArrayHits("Oops", aCard.KeyK)
		logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCard.KeyK +
			":it was:" + aCard.KeyH + ":but you had guessed:" + in)
		logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.KeyR)
		//
		// When this func is re-called, with flag set false, we skip this, and do the next
		if skipFlag == true {
			// So, we solicit another guess ... (user's second guess) and ...
			fmt.Println("Try again") // This only prints on first pass
			// Re-prompt, and send that second guess to the secondTry_meatOfKataExercise func
			in = Prompt_Scan_4_Romaji_or_HiraResponse(aCard.KeyK)
				branchOnUserSelectedDirectiveIfGiven(in,
					"Respond_w_Hira_or_Romaji_to_kataPrompt_3") // <-- Perform the directive
			secondTry_meatOfKataExercise(in)
		}
		// If user guesses incorrectly on his third-and-final try, then, and only then, execute the rest of this func
		if skipFlag == false { // skipFlag is true on first entry, & false when recalled after the third failed attempt
			fmt.Printf("%s", colorReset)
			fmt.Printf(" It was: ")
			fmt.Printf("%s", colorCyan)
			fmt.Printf("%s", aCard.KeyH)
			fmt.Printf("%s", colorReset)
			// Display hints in white
			fmt.Printf("\n\n%s\n%s\n%s\n", aCard.Hint1h, aCard.Hint2k, aCard.Hint3TT)
			
			fmt.Println("")
			// In this SECOND case: of having typed a Hiragana, the first 3 lines of hints can be displayed from the card
			// ... though, NOT in the FIRST case: of having typed a Romaji (then only the last line will be shown)
			// .
		} // Due to the fact that the next } below is paired with an '} else if' ...
	} // ... if ^^Right!, then we return to TouchTypingExercises12346(selectedExercise) (directly from this line)
}

/*
.
*/
// Second-Try version of the above func (tests the second guess, and then obtains the third and final guess)
func secondTry_meatOfKataExercise(in string) { // <-- This second-instance of 'in' is the user's second guess
	// Used for processing either a Romaji or a Hiragana guess, at a Katakana prompt
	//
	// If 'in' (the user's guess) is an alpha  v v v v v
	var isAlphanumeric bool
	findAlphasIn := regexp.MustCompile(`[a-zA-Z]`)
	switch true {
	case findAlphasIn.MatchString(in): // <-- 'in'
		isAlphanumeric = true
	default:
		isAlphanumeric = false
	}

	// User previously typed an alpha (as his second guess) -- And, the user has, presumably, typed a Romaji
	/*
		.
		.　The assumed Romaji guess will be processed below ...
	*/
	if isAlphanumeric == true && in == aCard.KeyR { // If user had typed an alpha, and it IS the correct Romaji
		fmt.Printf("%s", colorGreen)
		fmt.Printf("      　^^Right! ")
		fmt.Printf("%s", colorReset)
		//
		logHits_in_cyclicArrayHits("Right", aCard.KeyK)
		logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCard.KeyR)
		//
		fmt.Println() // Returns to caller: meatOfKataExercise(na)
	} else if isAlphanumeric == true && in != aCard.KeyR { // If user had typed an alpha, but NOT the correct Romaji
		// Solicit the final guess ... will pass it back to: meatOfKataExercise(in, false)
		fmt.Printf("%s", colorRed)
		fmt.Println("       Try again, you have one last attempt ... ")
		fmt.Printf("%s", colorReset)
		// Re-prompt, will be sending that final guess back to: meatOfKataExercise(in, false)
		in = Prompt_Scan_4_Romaji_or_HiraResponse(aCard.KeyK) // <-- Obtain the final guess, will pass it as 'in'
			branchOnUserSelectedDirectiveIfGiven(in, "Respond_w_Hira_or_Romaji_to_kataPrompt_3") // <-- Do dir
		fmt.Printf("%s", colorRed)
		fmt.Printf("      　^^Oops! ")
		logHits_in_cyclicArrayHits("Oops", aCard.KeyK)
		logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCard.KeyK +
			":it was:" + aCard.KeyR + ":but you had guessed:" + in)
		logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.KeyR)
		meatOfKataExercise(in, false) // Process the third try
		//
	} else if isAlphanumeric == false && in == aCard.KeyH { // If the user HAS typed the CORRECT Hiragana ...
		fmt.Printf("%s", colorGreen)
		fmt.Printf("      　^^Right! ")
		fmt.Printf("%s", colorReset)
		//
		logHits_in_cyclicArrayHits("Right", aCard.KeyK)
		logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCard.KeyR)
		//
		fmt.Println() // Returns to caller from this line
	} else if isAlphanumeric == false && in != aCard.KeyH { // User typed the INCORRECT Hiragana at Kata prompt
		// Solicit the third guess ... and pass it to: thirdTry_meatOfKataExercise(in, true)
		fmt.Printf("%s", colorRed)
		fmt.Println("       Try again, you have one last attempt ... ")
		fmt.Printf("%s", colorReset)
		// Re-prompt, will that third guess back to: meatOfKataExercise(in, false)
		in = Prompt_Scan_4_Romaji_or_HiraResponse(aCard.KeyK)                    // Obtain the second guess, will pass it as 'in'
			branchOnUserSelectedDirectiveIfGiven(in, "Respond_w_Hira_or_Romaji_to_kataPrompt_3") // <-- Do dir
		meatOfKataExercise(in, false)                                            // Process the third and final try
	}
}
