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
that is, when it is finally re-called by the third func.
.
When it IS re-called by the third and final func, skipFlag will be false. But, upon first-entry to this first func,
skipFlag will be true, so that the answer, and hint, will NOT be displayed after only the first attempt. And, since the
answer and hint can only be displayed by this first func, it will not be displayed by the second, and will only be
INDIRECTLY displayed by the third func: the third will pass-back to this func a skipFlag val which is finally false.

All this means that only functions 1 & 2 will be soliciting and obtaining new guesses
The first func is passed the first guess, and obtains the second guess
The second func is passed the second guess, and obtains the third guess
Finally, the first func is re-called by the second,: the first func is, then, passed the third and final guess
.
*/
// 1
func meatOfKataExorcise(in string, skipFlag bool) {
	// Used for processing either a Romaji OR a Hiragana guess which was obtained prior to this func being called
	// ... either by TouchTypingExorcise, or by: secondTry_ (secondTry_ actually solicits and obtains the THIRD guess)
	// If 'in' (the user's guess) is alpha v v v v v      (as first-instance, 'in' is the user's first guess)
	var isAlphanumeric bool
	findAlphasIn := regexp.MustCompile(`[a-zA-Z]`)
	switch true {
	case findAlphasIn.MatchString(in): // <-- 'in'
		isAlphanumeric = true
	default:
		isAlphanumeric = false
	}

	// First, we will deal with the case that ... the
	// user, typed an alpha (as first, though not necessarily his first, guess) -- user, presumably, typed a Romaji
	/*
		.
		.　The assumed Romaji guess, though not necessarily his first guess, will be processed below ...
	*/
	// The following 'if' will be executed, potentially only twice: for the first, and for the final, guess
	if isAlphanumeric == true && in == aCard.KeyR { // After this line we know both that it was a Romaji, and that it
		// IS the correct Romaji
		fmt.Printf("%s", colorGreen)
		fmt.Printf("      　^^Right! ")
		fmt.Printf("%s", colorReset)
		//
		logHits_in_cyclicArrayHits("Right", aCard.KeyK)
		logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCard.KeyR)
		//
		fmt.Println() // This line will be the last-executed before returning to the caller: TouchTypingExorcise
		//
		// The following 'else' will be executed, potentially only twice: for the first, and for the final, guess
	} else if isAlphanumeric == true && in != aCard.KeyR { // If user typed an alpha, but not the correct Romaji,
		// ... a second attempt will be solicited via: secondTry_meatOfKataExorcise(in) -- UNLESS flag is false
		fmt.Printf("%s", colorRed)
		fmt.Printf("      　^^Oops! ") // This always prints unless ^^Right! is printed
		logHits_in_cyclicArrayHits("Oops", aCard.KeyK)
		logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCard.KeyK + ":it was:" + aCard.KeyR + ":but you had guessed:" + in)
		logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.KeyR)
		//
		// When this func is re-called by the second, with flag set false, we skip this, and do the next
		if skipFlag == true {
			// Solicit the second guess ... and pass it to: secondTry_meatOfKataExorcise(in, true)
			fmt.Println("Try again") // This only prints on first pass
			fmt.Printf(colorReset)
			// Re-prompt, and send that second guess to the secondTry_meatOfKataExorcise func
			in = Kata_Prompt_Scan_4_Romaji_or_HiraResponse(aCard.KeyK) // Obtain the second guess, and pass it as 'in'
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
				branchOnUserSelectedDirectiveIfGiven(in, "Kata_Prompt-Respond-w-Hira|Romaji") // <-- handle the directive
			}
			secondTry_meatOfKataExorcise(in) // This instance of 'in' is actually the third and final guess.
		}
		// If the user guesses incorrectly on his third-and-final try, then, and only then, execute the rest of this func
		if skipFlag == false { // skipFlag will be true on first entry, but false upon being recalled after the third failed attempt
			fmt.Printf("%s", colorReset)
			fmt.Printf("\n It was: ")
			fmt.Printf("%s", colorCyan)
			fmt.Printf("%s", aCard.KeyR)
			fmt.Printf("%s", colorReset)
			checkForHints(aCard.KeyH) // Note: we only give hints for non-conjunctives
			fmt.Println("")
			// v v v v v NOT TRUE v v v v v But no need to 'fix' it since the answer was given above
			// **do-this** ??? ^ v
			// ... only the fields lacking the correct Romaji should be shown, and maybe some other info sans-Romaji
		}
	} // If the user was ^^Right!, then we return to TouchTypingExorcise(selectedExorcise) (directly from this very line)
	/*
	   .
	   .
	   .　An assumed Hiragana guess will be processed below ...
	*/
	if isAlphanumeric == false && in == aCard.KeyH { // if user HAS typed the CORRECT Hiragana
		fmt.Printf("%s", colorGreen)
		fmt.Printf("      　^^Right! ")
		fmt.Printf("%s", colorReset)
		//
		logHits_in_cyclicArrayHits("Right", aCard.KeyK)
		logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCard.KeyR)
		//
		fmt.Println() // Then, this line will be the last-executed of this func
		//
		//
	} else if isAlphanumeric == false && in != aCard.KeyH { // Here then, the user has typed the INCORRECT Hiragana ...
		fmt.Printf("%s", colorRed)
		fmt.Printf("      　^^Oops! ") // This always prints unless ^^Right! is printed
		logHits_in_cyclicArrayHits("Oops", aCard.KeyK)
		logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCard.KeyK + ":it was:" + aCard.KeyH + ":but you had guessed:" + in)
		logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.KeyR)
		//
		// When this func is re-called by the third, with flag set false, we skip this, and do the next
		if skipFlag == true {
			// So, we solicit another guess ... (his second guess) and ...
			fmt.Println("Try again") // This only prints on first pass
			// Re-prompt, and send that second guess to the secondTry_meatOfKataExorcise func
			in = Kata_Prompt_Scan_4_Romaji_or_HiraResponse(aCard.KeyK)
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
				branchOnUserSelectedDirectiveIfGiven(in, "Kata_Prompt-Respond-w-Hira|Romaji") // <-- handle the directive
			}
			secondTry_meatOfKataExorcise(in)
		}
		// If the user guesses incorrectly on his third-and-final try, then, and only then, execute the rest of this func
		if skipFlag == false { // skipFlag will be true on first entry, but false upon being recalled after the third failed attempt
			fmt.Printf("%s", colorReset)
			fmt.Printf("\n It was: ")
			fmt.Printf("%s", colorCyan)
			fmt.Printf("%s", aCard.KeyH)
			fmt.Printf("%s", colorReset)
			checkForHints(aCard.KeyH) // Note: we only give hints for non-conjunctives
			fmt.Println("")
			// In this SECOND case: of having typed a Hiragana, all four lines of hints can be displayed from the card
			// **do-this** ^ v
			// ... though, not in the first case: of having typed a Romaji (then only the last line can be shown ??)
		}
	} // If the user was ^^Right!, then we return to TouchTypingExorcise(selectedExorcise) (directly from this very line)
}

/*
.
*/
// Second-Try version of the above func (tests the second guess, and then obtains the third and final one)
func secondTry_meatOfKataExorcise(in string) { // <-- This second-instance of 'in' is the user's second guess
	// Used for processing either a Romaji or a Hiragana guess, at a Katakana prompt
	//
	// If 'in' (the user's guess) is alpha  v v v v v
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
	if isAlphanumeric == true && in == aCard.KeyR { // If user previously typed an alpha, and it IS the correct Romaji
		fmt.Printf("%s", colorGreen)
		fmt.Printf("      　^^Right! ")
		fmt.Printf("%s", colorReset)
		//
		logHits_in_cyclicArrayHits("Right", aCard.KeyK)
		logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCard.KeyR)
		//
		fmt.Println() // Returns to caller: meatOfKataExorcise(na)
	} else if isAlphanumeric == true && in != aCard.KeyR { // If user previously typed an alpha, but NOT the correct Romaji
		// Solicit the final guess ... will pass it back to: meatOfKataExorcise(in, false)
		fmt.Printf("%s", colorRed)
		fmt.Println("       Try again, you have one last attempt ... ")
		fmt.Printf("%s", colorReset)
		// Re-prompt, will be sending that final guess back to: meatOfKataExorcise(in, false)
		in = Kata_Prompt_Scan_4_Romaji_or_HiraResponse(aCard.KeyK) // <-- Obtain the final guess, will pass it as 'in'
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
			branchOnUserSelectedDirectiveIfGiven(in, "Kata_Prompt-Respond-w-Hira|Romaji") // <-- handle the directive
		}
		fmt.Printf("%s", colorRed)
		fmt.Printf("      　^^Oops! ")
		logHits_in_cyclicArrayHits("Oops", aCard.KeyK)
		logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCard.KeyK + ":it was:" + aCard.KeyR + ":but you had guessed:" + in)
		logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.KeyR)
		meatOfKataExorcise(in, false) // Process the third try
		//
	} else if isAlphanumeric == false && in == aCard.KeyH { // if user HAS typed the CORRECT Hiragana
		fmt.Printf("%s", colorGreen)
		fmt.Printf("      　^^Right! ")
		fmt.Printf("%s", colorReset)
		//
		logHits_in_cyclicArrayHits("Right", aCard.KeyK)
		logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCard.KeyR)
		//
		fmt.Println() // And, return to caller
	} else if isAlphanumeric == false && in != aCard.KeyH { // user has typed the INCORRECT Hiragana corresponding to the Kata prompt
		// Solicit the third guess ... and pass it to: thirdTry_meatOfKataExorcise(in, true)
		fmt.Printf("%s", colorRed)
		fmt.Println("       Try again, you have one last attempt ... ")
		fmt.Printf("%s", colorReset)
		// Re-prompt, and send that third guess back to: meatOfKataExorcise(in, false)
		in = Kata_Prompt_Scan_4_Romaji_or_HiraResponse(aCard.KeyK) // Obtain the second guess, and pass it as 'in'
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
			branchOnUserSelectedDirectiveIfGiven(in, "Kata_Prompt-Respond-w-Hira|Romaji") // <-- handle the directive
		}
		meatOfKataExorcise(in, false) // Process the third try
	}
}
