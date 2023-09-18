package main

// Refer also to 8, which is a similar set of functions
import (
	"fmt"
	"regexp"
)

/*
The idea is to give the user two additional tries: if he fails at the first or second attempt then no answer or hint is
to be displayed. But, after the third and final attempt (the maximum allowed) we will pass a skipFlag of false so that
the answer and hint WILL be displayed after that third wrong guess -- when re-entering this first function:
that is, when it is finally re-called by the second func (the secondTry_ func, below).
.
When it IS re-called by the second func, skipFlag will be false. But, upon first-entry to this first func,
skipFlag will be true, so that the answer, and hint, will NOT be displayed after the first attempt. The answer and
hint can only be displayed by the first func (it is absent in the second func). The second func will pass-back to the
first func a skipFlag val of false, thereby causing the first to display the answer plus one-line of hints from
checkForHints(value string).
.
This, first of two functions, meatOfKataExerciseD is passed the first guess which was collected by the caller

If the user is found to have guessed incorrectly the user will be prompted for a second guess.

That guess is passed to the secondTry_ func below -- which may obtain a third guess
Finally, this first func is re-called by the second,: the first func is, then, passed the third and final guess
.
*/
func meatOfKataExerciseD(in string, skipFlag bool) {
	// Used for processing either a Romaji OR a Hiragana guess which was obtained prior to this func being called
	// ... either by TouchTypingExercise, or by: secondTry_ (secondTry_ actually solicits and obtains the THIRD guess)
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
	if isAlphanumeric == true && in == aCardD.KeyR { // After this line we know both that it was a Romaji, and that it
		// ... IS the correct Romaji
		fmt.Printf("%s", colorGreen)
		fmt.Printf("      　^^Right! ")
		fmt.Printf("%s", colorReset)
		//
		logHits_in_cyclicArrayHits("Right", aCardD.KeyK)
		logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCardD.KeyR)
		//
		fmt.Println() // This line will be the last-executed before returning to the caller: TouchTypingExercise
		//
		// The following 'else' will be executed, potentially only twice: 1: for the first, and 2: for the final guess
	} else if isAlphanumeric == true && in != aCardD.KeyR { // If user typed an alpha, but not the correct Romaji,
		// ... a second attempt will be solicited via: secondTry_meatOfKataExercise(in) -- UNLESS skipFlag is false
		fmt.Printf("%s", colorRed)
		fmt.Printf("      　^^Oops! ") // This ALWAYS prints unless ^^Right! is printed
		//
		logHits_in_cyclicArrayHits("Oops", aCardD.KeyK)
		logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCardD.KeyK +
			":it was:" + aCardD.KeyR + ":but you had guessed:" + in)
		logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCardD.KeyR)
		//
		// When this func is re-called by the second, with skipFlag set false, we skip this, and do the next
		if skipFlag == true {
			// Solicit the second guess ... and pass it to: secondTry_meatOfKataExercise(in)
			fmt.Println("Try again") // This ONLY prints on the first pass through this func
			fmt.Printf(colorReset)
			// Re-prompt, and will send that second guess to the secondTry_meatOfKataExercise func
			in = Prompt_Scan_4_Romaji_or_HiraResponse(aCardD.KeyK) // Obtain second guess, and pass it as 'in'
				branchOnUserSelectedDirectiveIfGiven(in, "Most_Difficult") // Perform the Directive
			secondTry_meatOfKataExerciseD(in)   // This instance of 'in' is the user's second guess.
		}
		// If user guesses incorrectly on his third-and-final try, then, and only then, execute the rest of this func
		if skipFlag == false { // skipFlag is true on first entry, & false when recalled after third failed attempt
			fmt.Printf("%s", colorReset)
			fmt.Printf("\n It was: ")
			fmt.Printf("%s", colorCyan)
			fmt.Printf("%s", aCardD.KeyR)
			fmt.Printf("%s", colorReset)
			// Display hints in white
			// Only the fields lacking the correct Romaji will be shown (only the last field: HintSansR)
			fmt.Printf("\n\n%s\n", aCardD.HintSansR)
			fmt.Println("")
		}
	} // If user was ^^Right!, then we return to TouchTypingExercise(selectedExercise) (directly from this very line)
	/*
	   .
	   .
	   .　An assumed Hiragana guess will be processed below ...
	*/
	if isAlphanumeric == false && in == aCardD.KeyH { // If user HAS typed the CORRECT Hiragana
		fmt.Printf("%s", colorGreen)
		fmt.Printf("      　^^Right! ")
		fmt.Printf("%s", colorReset)
		//
		logHits_in_cyclicArrayHits("Right", aCardD.KeyK)
		logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCardD.KeyR)
		//
		fmt.Println() // This line is the last-executed of this func, returns to TouchTypingExercise(SE)
		//
		//
	} else if isAlphanumeric == false && in != aCardD.KeyH { // <-- The user has typed the INCORRECT Hiragana ...
		fmt.Printf("%s", colorRed)
		fmt.Printf("      　^^Oops! ") // This ALWAYS prints unless ^^Right! is printed
		//
		logHits_in_cyclicArrayHits("Oops", aCardD.KeyK)
		logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCardD.KeyK +
			":it was:" + aCardD.KeyR + ":but you had guessed:" + in)
		logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCardD.KeyR)
		//
		// When this func is re-called, with flag set false, we skip this, and do the next
		if skipFlag == true {
			// So, we solicit another guess ... (user's second guess) and ...
			fmt.Println("Try again") // This only prints on first pass
			// Re-prompt, and send that second guess to the secondTry_meatOfKataExercise func
			in = Prompt_Scan_4_Romaji_or_HiraResponse(aCardD.KeyK)
				branchOnUserSelectedDirectiveIfGiven(in,"Most_Difficult") // <-- Perform the Directive
			secondTry_meatOfKataExerciseD(in)
		}
		// If user guesses incorrectly on his third-and-final try, then, and only then, execute the rest of this func
		if skipFlag == false { // skipFlag is true on first entry, & false when recalled after the third failed attempt
			fmt.Printf("%s", colorReset)
			fmt.Printf(" It was: ")
			fmt.Printf("%s", colorCyan)
			fmt.Printf("%s", aCardD.KeyH)
			fmt.Printf("%s", colorReset)
			// Display hints in white
			fmt.Printf("\n\n%s\n%s\n%s\n", aCardD.Hint1h, aCardD.Hint2k, aCardD.Hint3TT)
			fmt.Println("")
			// In this SECOND case: of having typed a Hiragana, all four lines of hints can be displayed from the card
			// ... though, NOT in the FIRST case: of having typed a Romaji (then only the last line will be shown)
			// .
		} // Due to the fact that the next } below is paired with an '} else if' ...
	} // ... if ^^Right!, then we return to TouchTypingExercise(selectedExercise) (directly from this line)
}

/*
.
*/
// Second-Try version of the above func (tests the second guess, and then obtains the third and final guess)
func secondTry_meatOfKataExerciseD(in string) { // <-- This second-instance of 'in' is the user's second guess
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
	if isAlphanumeric == true && in == aCardD.KeyR { // If user had typed an alpha, and it IS the correct Romaji
		fmt.Printf("%s", colorGreen)
		fmt.Printf("      　^^Right! ")
		fmt.Printf("%s", colorReset)
		//
		logHits_in_cyclicArrayHits("Right", aCardD.KeyK)
		logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCardD.KeyR)
		//
		fmt.Println() // Returns to caller: meatOfKataExercise(na)
	} else if isAlphanumeric == true && in != aCardD.KeyR { // If user had typed an alpha, but NOT the correct Romaji
		// Solicit the final guess ... will pass it back to: meatOfKataExercise(in, false)
		fmt.Printf("%s", colorRed)
		fmt.Println("       Try again, you have one last attempt ... ")
		fmt.Printf("%s", colorReset)
		// Re-prompt, will be sending that final guess back to: meatOfKataExercise(in, false)
		in = Prompt_Scan_4_Romaji_or_HiraResponse(aCardD.KeyK) // <-- Obtain the final guess, will pass it as 'in'
			branchOnUserSelectedDirectiveIfGiven(in, "Most_Difficult") // <-- Do dir
		fmt.Printf("%s", colorRed)
		fmt.Printf("      　^^Oops! ")
		//
		logHits_in_cyclicArrayHits("Oops", aCardD.KeyK)
		logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCardD.KeyK +
			":it was:" + aCardD.KeyR + ":but you had guessed:" + in)
		logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCardD.KeyR)
		//
		meatOfKataExerciseD(in, false) // Process the third try
		//
	} else if isAlphanumeric == false && in == aCardD.KeyH { // If the user HAS typed the CORRECT Hiragana ...
		fmt.Printf("%s", colorGreen)
		fmt.Printf("      　^^Right! ")
		fmt.Printf("%s", colorReset)
		//
		logHits_in_cyclicArrayHits("Right", aCardD.KeyK)
		logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCardD.KeyR)
		//
		fmt.Println() // Returns to caller from this line
	} else if isAlphanumeric == false && in != aCardD.KeyH { // User typed the INCORRECT Hiragana at Kata prompt
		// Solicit the third guess ... and pass it to: thirdTry_meatOfKataExercise(in, true)
		fmt.Printf("%s", colorRed)
		fmt.Println("       Try again, you have one last attempt ... ")
		fmt.Printf("%s", colorReset)
		// Re-prompt, will that third guess back to: meatOfKataExercise(in, false)
		in = Prompt_Scan_4_Romaji_or_HiraResponse(aCardD.KeyK)                   // Obtain the second guess, will pass it as 'in'
			branchOnUserSelectedDirectiveIfGiven(in, "Most_Difficult") // <-- Do dir
		meatOfKataExerciseD(in, false)                                           // Process the third and final try
	}
}
