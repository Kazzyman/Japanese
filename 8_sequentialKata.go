package main

// Finished September 12 of 2023 at 0112 0729 hours
import (
	"fmt"
	"regexp"
)

/*
Basically: The first function, below, is passed the first guess, and it then obtains the second guess
The second func is passed the second guess, and obtains the third guess
Finally, the first func is re-called by the second,: the first func is, then, passed the third and final guess

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
meatOfSequentialKata(in string, skipFlag bool) is passed the first guess which was collected by the caller:
... TouchTypingExerciseSequential()

meatOfSequentialKata(in string, skipFlag bool) then determines what should be done in response.

If the user is found to have guessed incorrectly the user will be prompted for a second guess.

That guess is passed to secondTry_meatOfSequentialKata(in string), which may obtain a third guess
Finally, the first func is re-called by the second,: the first func is, then, passed the third and final guess
.
*/
func meatOfSequentialKata(in string, skipFlag bool) {
	// Used for processing either a Romaji OR a Hiragana guess which was obtained prior to this func being called
	// ... either by TouchTypingExerciseSequential, or by: secondTry_meatOfSequentialKata
	// ... (it is secondTry_meatOfSequentialKata which actually solicits and obtains the THIRD guess)
	// .
	// Since this func must deal with either Hiragana or Romaji inputs, we must first determine which kind we have
	// If 'in' (the user's guess) is alpha v v v v v      (a first-instance of 'in' is the user's first guess)
	var isAlphanumeric bool
	findAlphasIn := regexp.MustCompile(`[a-zA-Z]`)
	switch true { // Determine if the input was an Alpha, and therefor a Romaji as opposed to a Hiragana
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
	// The following 'if' will be executed, potentially only twice: 1: for the first guess, and 2: for the final guess
	if isAlphanumeric == true && in == aCardS.KeyR { // After this line we know both that it was a Romaji, and that it
		// ... IS the correct Romaji
		fmt.Printf("%s", colorGreen)
		fmt.Printf("      　^^Right! ")
		fmt.Printf("%s", colorReset)
		//
		logHits_in_cyclicArrayHits("Right", aCardS.KeyK)
		logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCardS.KeyR)
		//
		fmt.Println() // This line will be the last-executed before returning to the caller
		//
		// The following 'else' will be executed, potentially only twice: 1: for the first, and 2: for the final guess
	} else if isAlphanumeric == true && in != aCardS.KeyR { // If user typed an alpha, but not the correct Romaji,
		// ... a second attempt will be solicited via the secondTry_ func below -- UNLESS skipFlag is false
		fmt.Printf("%s", colorRed)
		fmt.Printf("      　^^Oops! ") // This ALWAYS prints unless ^^Right! is printed
		//
		logHits_in_cyclicArrayHits("Oops", aCardS.KeyK)
		logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCardS.KeyK +
			":it was:" + aCardS.KeyR + ":but you had guessed:" + in)
		logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCardS.KeyR)
		//
		// When this func is re-called by the second, with skipFlag set false, we skip this, and do the next
		if skipFlag == true {
			// Solicit the second guess ... and pass it to the secondTry_ func, below
			fmt.Println("Try again") // This ONLY prints on the first pass through this func
			fmt.Printf(colorReset)
			// Re-prompt, and will send that second guess to the secondTry_ func, below
			in = Prompt_Scan_4_Romaji_or_HiraResponse(aCardS.KeyK) // Obtain second guess, and pass it as 'in'
				branchOnUserSelectedDirectiveIfGiven(in,
					"Respond_w_Hira_or_Romaji") // <-- Perform the Directive
			secondTry_meatOfSequentialKata(in) // This instance of 'in' is the user's second guess.
		}
		// If user guesses incorrectly on his third-and-final try, then, and only then, execute the rest of this func
		if skipFlag == false { // skipFlag is true on first entry, & false when recalled after third failed attempt
			fmt.Printf("%s", colorReset)
			fmt.Printf(" It was: ")
			fmt.Printf("%s", colorCyan)
			fmt.Printf("%s", aCardS.KeyR)
			fmt.Printf("%s", colorReset)
			// Display hints in white
			fmt.Printf("\n\n%s\n%s\n%s\n", aCardS.Hint1h, aCardS.Hint2k, aCardS.Hint3TT)
			fmt.Println("")
		}
	} // If user was ^^Right!, then we return to the caller (directly from this very line)
	/*
	   .
	   .
	   .　An assumed Hiragana guess (a non-Alpha) will be processed below ...
	*/
	if isAlphanumeric == false && in == aCardS.KeyH { // If user HAS typed the CORRECT Hiragana
		fmt.Printf("%s", colorGreen)
		fmt.Printf("      　^^Right! ")
		fmt.Printf("%s", colorReset)
		//
		logHits_in_cyclicArrayHits("Right", aCardS.KeyK)
		logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCardS.KeyR)
		//
		fmt.Println() // This line is the last-executed of this func, returns to caller
		//
		//
	} else if isAlphanumeric == false && in != aCardS.KeyH { // <-- The user has typed an INCORRECT Hiragana ...
		fmt.Printf("%s", colorRed)
		fmt.Printf("      　^^Oops! ") // This ALWAYS prints unless ^^Right! is printed
		//
		logHits_in_cyclicArrayHits("Oops", aCardS.KeyK)
		logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCardS.KeyK +
			":it was:" + aCardS.KeyH + ":but you had guessed:" + in)
		logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCardS.KeyR)
		//
		// When this func is re-called, with skipFlag set false, we skip this, and do the next
		if skipFlag == true {
			// So, we solicit another guess ... (user's second guess) and ...
			fmt.Println("Try again") // This only prints on first pass
			// Re-prompt, and send that second guess to the secondTry_ func, below
			in = Prompt_Scan_4_Romaji_or_HiraResponse(aCardS.KeyK)
			// Prompt_Scan_4_Romaji_or_HiraResponse
				branchOnUserSelectedDirectiveIfGiven(in,
					"Respond_w_Hira_or_Romaji") // <-- Perform the Directive
			secondTry_meatOfSequentialKata(in)
		}
		// If user guesses incorrectly on his third-and-final try, then, and only then, execute the rest of this func
		if skipFlag == false { // skipFlag is true on first entry, & false when recalled after the third failed attempt
			fmt.Printf("%s", colorReset)
			fmt.Printf(" It was: ")
			fmt.Printf("%s", colorCyan)
			fmt.Printf("%s", aCardS.KeyH)
			fmt.Printf("%s", colorReset)
			// Display hints in white
			fmt.Printf("\n\n%s\n%s\n%s\n", aCardS.Hint1h, aCardS.Hint2k, aCardS.Hint3TT)
			fmt.Println("")
		} // Due to the fact that the next } below is paired with an '} else if' ...
	} // ... if ^^Right!, then we return to the caller (directly from this line)
}

/*
.
*/
// Second-Try version of the above func (tests the second guess, and then obtains the third and final guess)
// (as a partial clone of the above func, additional documentation can be found in that 'parent' func)
func secondTry_meatOfSequentialKata(in string) { // <-- This second-instance of 'in' is the user's second guess
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
	if isAlphanumeric == true && in == aCardS.KeyR { // If user had typed an alpha, and it IS the correct Romaji
		fmt.Printf("%s", colorGreen)
		fmt.Printf("      　^^Right! ")
		fmt.Printf("%s", colorReset)
		//
		logHits_in_cyclicArrayHits("Right", aCardS.KeyK)
		logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCardS.KeyR)
		//
		fmt.Println() // Returns to caller
	} else if isAlphanumeric == true && in != aCardS.KeyR { // If user had typed an alpha, but NOT the correct Romaji
		// Solicit the final guess ... will pass it back to caller
		fmt.Printf("%s", colorRed)
		fmt.Println("       Try again, you have one last attempt ... ")
		fmt.Printf("%s", colorReset)
		// Re-prompt, will be sending that final guess back to caller
		in = Prompt_Scan_4_Romaji_or_HiraResponse(aCardS.KeyK) // <-- Obtain the final guess, will pass it as 'in'
			branchOnUserSelectedDirectiveIfGiven(in, "Respond_w_Hira_or_Romaji") // <-- Do dir
		fmt.Printf("%s", colorRed)
		fmt.Printf("      　^^Oops! ")
		logHits_in_cyclicArrayHits("Oops", aCardS.KeyK)
		logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCardS.KeyK +
			":it was:" + aCardS.KeyR + ":but you had guessed:" + in)
		logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCardS.KeyR)
		meatOfSequentialKata(in, false) // Process the third try
		//
	} else if isAlphanumeric == false && in == aCardS.KeyH { // If the user HAS typed the CORRECT Hiragana ...
		fmt.Printf("%s", colorGreen)
		fmt.Printf("      　^^Right! ")
		fmt.Printf("%s", colorReset)
		//
		logHits_in_cyclicArrayHits("Right", aCardS.KeyK)
		logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCardS.KeyR)
		//
		fmt.Println() // Returns to caller from this line
	} else if isAlphanumeric == false && in != aCardS.KeyH { // User typed the INCORRECT Hiragana at Kata prompt
		// Solicit the third guess ... and pass it to: thirdTry_meatOfKataExercise(in, true)
		fmt.Printf("%s", colorRed)
		fmt.Println("       Try again, you have one last attempt ... ")
		fmt.Printf("%s", colorReset)
		// Re-prompt, will that third guess back to: meatOfKataExercise(in, false)
		in = Prompt_Scan_4_Romaji_or_HiraResponse(aCardS.KeyK) // Obtain the second guess, will pass it as 'in'
			branchOnUserSelectedDirectiveIfGiven(in, "Respond_w_Hira_or_Romaji") // <-- Do dir
		meatOfSequentialKata(in, false) // Process the third and final try
	}
}
