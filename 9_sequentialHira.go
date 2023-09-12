package main

import (
	"fmt"
)

// Finished September 12 of 2023 at 0132 hours
/*
This is essentially a clone of meatOfSequentialKata(in string, skipFlag bool), sans the Alpha testing,
refer to that func for documentation
*/
func meatOfSequentialHira(in string, skipFlag bool) {

	if in == aCardS.KeyR {
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
	} else if in != aCardS.KeyR {
		// ... a second attempt will be solicited via: secondTry_  -- UNLESS skipFlag is false
		fmt.Printf("%s", colorRed)
		fmt.Printf("      　^^Oops! ") // This ALWAYS prints unless ^^Right! is printed
		//
		logHits_in_cyclicArrayHits("Oops", aCardS.KeyH)
		logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCardS.KeyH +
			":it was:" + aCardS.KeyR + ":but you had guessed:" + in)
		logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCardS.KeyR)
		//
		// When this func is re-called by the second, with skipFlag set false, we skip this, and do the next
		if skipFlag == true {
			// Solicit the second guess
			fmt.Println("Try again") // This ONLY prints on the first pass through this func
			fmt.Printf(colorReset)
			// Re-prompt
			in = Prompt_Scan_4_Romaji_or_HiraResponse(aCardS.KeyH) // Obtain second guess, and pass it as 'in'
			if in == "set" ||
				in == "?" ||
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
					"Kata_Prompt-Respond-w-Hira|Romaji") // <-- Perform the Directive
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
			//
			checkForHints(aCardS.KeyH)
			fmt.Println("")
		}
	} // If user was ^^Right!, then we return to the caller (directly from this very line)
}

/*
.
*/
// Second-Try version of the above func (tests the second guess, and then obtains the third and final guess)
func secondTry_meatOfSequentialHira(in string) { // <-- This second-instance of 'in' is the user's second guess

	if in == aCardS.KeyR {
		fmt.Printf("%s", colorGreen)
		fmt.Printf("      　^^Right! ")
		fmt.Printf("%s", colorReset)
		//
		logHits_in_cyclicArrayHits("Right", aCardS.KeyH)
		logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCardS.KeyR)
		//
		fmt.Println() // Returns to caller: meatOfKataExorcise(na)
	} else if in != aCardS.KeyR { // If user had typed an alpha, but NOT the correct Romaji
		// Solicit the final guess ... will pass it back to: meatOfKataExorcise(in, false)
		fmt.Printf("%s", colorRed)
		fmt.Println("       Try again, you have one last attempt ... ")
		fmt.Printf("%s", colorReset)
		// Re-prompt, will be sending that final guess back to: meatOfKataExorcise(in, false)
		in = Prompt_Scan_4_Romaji_or_HiraResponse(aCardS.KeyH) // <-- Obtain the final guess, will pass it as 'in'
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
			branchOnUserSelectedDirectiveIfGiven(in, "Respond-w-Hira|Romaji") // <-- Do dir
		}
		fmt.Printf("%s", colorRed)
		fmt.Printf("      　^^Oops! ")
		//
		logHits_in_cyclicArrayHits("Oops", aCardS.KeyH)
		logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCardS.KeyH +
			":it was:" + aCardS.KeyR + ":but you had guessed:" + in)
		logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCardS.KeyR)
		//
		meatOfSequentialHira(in, false) // Process the third and final try
	}
}
