package main

// **do-this**
import (
	"fmt"
	"regexp"
)

/*
.
The idea is to give the user two additional tries: if he fails at the first or second attempt then no answer or hint is
to be displayed. But, after the third and final attempt (the maximum allowed) we will pass a skipFlag of false so that
the answer and hint WILL be displayed after that third wrong guess -- when re-entering this first-of-three functions:
that is, when it is finally re-called by the second func.
.
When it IS re-called by the second func, skipFlag will be false. But, upon first-entry to this first func,
skipFlag will be true, so that the answer, and hint, will NOT be displayed after the first attempt. The answer and
hint can only be displayed by the first func (it is absent in the second func). The second func will pass-back to the
first func a skipFlag val of false, thereby causing the first to display the answer and one-line of hints from
checkForHints(value string).
.
meatOfKataExorcise(in string, skipFlag bool) is passed the first guess which was collected by the caller:
... TouchTypingExorcise(selectedExorcise)

meatOfKataExorcise(in string, skipFlag bool) then determines what should be done in response.

If the user is found to have guessed incorrectly the user will be prompted for a second guess.

That guess is passed to the second func guess, and obtains the third guess
Finally, the first func is re-called by the second,: the first func is, then, passed the third and final guess
.
*/
func meatOfSequentialHira(in string, skipFlag bool) {
	// Used for processing either a Romaji OR a Hiragana guess which was obtained prior to this func being called
	// ... either by TouchTypingExorcise, or by: secondTry_ (secondTry_ actually solicits and obtains the THIRD guess)
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
	if isAlphanumeric == true && in == aCardS.KeyR { // After this line we know both that it was a Romaji, and that it
		// ... IS the correct Romaji
		fmt.Printf("%s", colorGreen)
		fmt.Printf("      　^^Right! ")
		fmt.Printf("%s", colorReset)
		//
		logHits_in_cyclicArrayHits("Right", aCardS.KeyH)
		logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCardS.KeyR)
		//
		fmt.Println() // This line will be the last-executed before returning to the caller: TouchTypingExorcise
		//
		// The following 'else' will be executed, potentially only twice: 1: for the first, and 2: for the final guess
	} else if isAlphanumeric == true && in != aCardS.KeyR { // If user typed an alpha, but not the correct Romaji,
		// ... a second attempt will be solicited via: secondTry_meatOfKataExorcise(in) -- UNLESS skipFlag is false
		fmt.Printf("%s", colorRed)
		fmt.Printf("      　^^Oops! ") // This ALWAYS prints unless ^^Right! is printed
		logHits_in_cyclicArrayHits("Oops", aCardS.KeyH)
		logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCardS.KeyH +
			":it was:" + aCardS.KeyR + ":but you had guessed:" + in)
		logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCardS.KeyR)
		// When this func is re-called by the second, with skipFlag set false, we skip this, and do the next
		if skipFlag == true {
			// Solicit the second guess ... and pass it to: secondTry_meatOfKataExorcise(in)
			fmt.Println("Try again") // This ONLY prints on the first pass through this func
			fmt.Printf(colorReset)
			// Re-prompt, and will send that second guess to the secondTry_meatOfKataExorcise func
			in = Kata_Prompt_Scan_4_Romaji_or_HiraResponse(aCardS.KeyH) // Obtain second guess, and pass it as 'in'
			if in == "set" ||
				in == "?" || // <-- If it IS a directive
				in == "??" ||
				in == "menu" ||
				in == "reset" ||
				in == "stat" ||
				in == "dir" ||
				in == "notes" ||
				in == "quit" ||
				in == "exit" ||
				in == "stats" ||
				in == "rm" ||
				in == "stack" {
				branchOnUserSelectedDirectiveIfGiven(in,
					"Kata_Prompt-Respond-w-Hira|Romaji") // <-- Perform the directive
			}
			secondTry_meatOfSequentialHira(in) // This instance of 'in' is the user's second guess.
		}
		// If user guesses incorrectly on his third-and-final try, then, and only then, execute the rest of this func
		if skipFlag == false { // skipFlag is true on first entry, & false when recalled after third failed attempt
			fmt.Printf("%s", colorReset)
			fmt.Printf("\n It was: ")
			fmt.Printf("%s", colorCyan)
			fmt.Printf("%s", aCardS.KeyR)
			fmt.Printf("%s", colorReset)
			// Only the fields lacking the correct Romaji will be shown (only the last field: HintSansR)
			checkForHints(aCardS.KeyH) // Note: we only give hints for non-conjunctives
			fmt.Println("")
		}
	} // If user was ^^Right!, then we return to TouchTypingExorcise(selectedExorcise) (directly from this very line)
	/*
	   .
	   .
	   .　An assumed Hiragana guess will be processed below ...
	*/
	if isAlphanumeric == false && in == aCardS.KeyH { // If user HAS typed the CORRECT Hiragana
		fmt.Printf("%s", colorGreen)
		fmt.Printf("      　^^Right! ")
		fmt.Printf("%s", colorReset)
		//
		logHits_in_cyclicArrayHits("Right", aCardS.KeyH)
		logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCardS.KeyR)
		//
		fmt.Println() // This line is the last-executed of this func, returns to TouchTypingExorcise(SE)
		//
		//
	} else if isAlphanumeric == false && in != aCardS.KeyH { // <-- The user has typed the INCORRECT Hiragana ...
		fmt.Printf("%s", colorRed)
		fmt.Printf("      　^^Oops! ") // This ALWAYS prints unless ^^Right! is printed
		logHits_in_cyclicArrayHits("Oops", aCardS.KeyH)
		logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCardS.KeyH +
			":it was:" + aCardS.KeyH + ":but you had guessed:" + in)
		logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCardS.KeyR)
		//
		// When this func is re-called, with flag set false, we skip this, and do the next
		if skipFlag == true {
			// So, we solicit another guess ... (user's second guess) and ...
			fmt.Println("Try again") // This only prints on first pass
			// Re-prompt, and send that second guess to the secondTry_meatOfKataExorcise func
			in = Kata_Prompt_Scan_4_Romaji_or_HiraResponse(aCardS.KeyH)
			if in == "set" ||
				in == "?" || // <-- if it IS a directive
				in == "??" ||
				in == "menu" ||
				in == "reset" ||
				in == "stat" ||
				in == "dir" ||
				in == "notes" ||
				in == "quit" ||
				in == "exit" ||
				in == "stats" ||
				in == "rm" ||
				in == "stack" {
				branchOnUserSelectedDirectiveIfGiven(in,
					"Kata_Prompt-Respond-w-Hira|Romaji") // <-- Perform the directive
			}
			secondTry_meatOfSequentialHira(in)
		}
		// If user guesses incorrectly on his third-and-final try, then, and only then, execute the rest of this func
		if skipFlag == false { // skipFlag is true on first entry, & false when recalled after the third failed attempt
			fmt.Printf("%s", colorReset)
			fmt.Printf("\n It was: ")
			fmt.Printf("%s", colorCyan)
			fmt.Printf("%s", aCardS.KeyH)
			fmt.Printf("%s", colorReset)
			checkForHints(aCardS.KeyH) // Note: we only give hints for non-conjunctives
			fmt.Println("")
			// In this SECOND case: of having typed a Hiragana, all four lines of hints can be displayed from the card
			// ... though, NOT in the FIRST case: of having typed a Romaji (then only the last line will be shown)
			// .
		} // Due to the fact that the next } below is paired with an '} else if' ...
	} // ... if ^^Right!, then we return to TouchTypingExorcise(selectedExorcise) (directly from this line)
}

/*
.
*/
// Second-Try version of the above func (tests the second guess, and then obtains the third and final guess)
func secondTry_meatOfSequentialHira(in string) { // <-- This second-instance of 'in' is the user's second guess
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
		logHits_in_cyclicArrayHits("Right", aCardS.KeyH)
		logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCardS.KeyR)
		//
		fmt.Println() // Returns to caller: meatOfKataExorcise(na)
	} else if isAlphanumeric == true && in != aCardS.KeyR { // If user had typed an alpha, but NOT the correct Romaji
		// Solicit the final guess ... will pass it back to: meatOfKataExorcise(in, false)
		fmt.Printf("%s", colorRed)
		fmt.Println("       Try again, you have one last attempt ... ")
		fmt.Printf("%s", colorReset)
		// Re-prompt, will be sending that final guess back to: meatOfKataExorcise(in, false)
		in = Kata_Prompt_Scan_4_Romaji_or_HiraResponse(aCardS.KeyH) // <-- Obtain the final guess, will pass it as 'in'
		if in == "set" ||
			in == "?" || // <-- if it IS a directive
			in == "??" ||
			in == "menu" ||
			in == "reset" ||
			in == "stat" ||
			in == "dir" ||
			in == "notes" ||
			in == "quit" ||
			in == "exit" ||
			in == "stats" ||
			in == "rm" ||
			in == "stack" {
			branchOnUserSelectedDirectiveIfGiven(in, "Kata_Prompt-Respond-w-Hira|Romaji") // <-- Do dir
		}
		fmt.Printf("%s", colorRed)
		fmt.Printf("      　^^Oops! ")
		logHits_in_cyclicArrayHits("Oops", aCardS.KeyH)
		logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCardS.KeyH +
			":it was:" + aCardS.KeyR + ":but you had guessed:" + in)
		logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCardS.KeyR)
		meatOfSequentialHira(in, false) // Process the third try
		//
	} else if isAlphanumeric == false && in == aCardS.KeyH { // If the user HAS typed the CORRECT Hiragana ...
		fmt.Printf("%s", colorGreen)
		fmt.Printf("      　^^Right! ")
		fmt.Printf("%s", colorReset)
		//
		logHits_in_cyclicArrayHits("Right", aCardS.KeyH)
		logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCardS.KeyR)
		//
		fmt.Println() // Returns to caller from this line
	} else if isAlphanumeric == false && in != aCardS.KeyH { // User typed the INCORRECT Hiragana at Kata prompt
		// Solicit the third guess ... and pass it to: thirdTry_meatOfKataExorcise(in, true)
		fmt.Printf("%s", colorRed)
		fmt.Println("       Try again, you have one last attempt ... ")
		fmt.Printf("%s", colorReset)
		// Re-prompt, will that third guess back to: meatOfKataExorcise(in, false)
		in = Kata_Prompt_Scan_4_Romaji_or_HiraResponse(aCardS.KeyH) // Obtain the second guess, will pass it as 'in'
		if in == "set" ||
			in == "?" || // <-- If it IS a directive
			in == "??" ||
			in == "menu" ||
			in == "reset" ||
			in == "stat" ||
			in == "dir" ||
			in == "notes" ||
			in == "quit" ||
			in == "exit" ||
			in == "stats" ||
			in == "rm" ||
			in == "stack" {
			branchOnUserSelectedDirectiveIfGiven(in, "Kata_Prompt-Respond-w-Hira|Romaji") // <-- Do dir
		}
		meatOfSequentialHira(in, false) // Process the third and final try
	}
}
