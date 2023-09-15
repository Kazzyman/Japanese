package main

// **do-this**
import (
	"fmt"
)

/*
.
Refer to 8, which is a similar function

All this means that:
The first func is passed the first guess, and obtains the second guess
The second func is passed the second guess, and obtains the third guess
Finally, the first func is re-called by the second,: the first func is, then, passed the third and final guess
.
*/
func meatOfRomajiExorcise(in string, skipFlag bool) { // NOTE: we have already been prompted with KeyR
	// ..^ ^ ^ called as|with: (usersGuessOrOptionDirective, (bool)'skipFlag')
	// so, 'in' == the-users-guess, which is ...
	// obtained prior to this func being called ...
	// ^^^^^^^^ either by TouchTypingExorcise, ...
	// or by: secondTry_ (secondTry_ actually solicits and obtains the THIRD guess)
	//
	var thisCaseOfAnInHasAlreadyBeenProcessedAbove bool
	thisCaseOfAnInHasAlreadyBeenProcessedAbove = false
	// *******************************************************************************************
	// Handle special prompt cases prior to doing the normal "if in == " processing, i.e.,
	// One 'very'-special key handler to emphasize the sameness of two variants of the zu sound
	// *******************************************************************************************

	if aCard.KeyR == "zu" {
		thisCaseOfAnInHasAlreadyBeenProcessedAbove = true
		if in == "づ" {
			fmt.Printf("%s", colorGreen)
			fmt.Printf("      ^^Right! ")
			fmt.Printf("%s", colorReset)
			fmt.Printf("It could have been either ず or づ as they are the same sound: zu\n")
			// log the hit:
			logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCard.KeyR)
			logHits_in_cyclicArrayHits("Right", aCard.KeyR)
		} else if in == "ず" {
			fmt.Printf("%s", colorGreen)
			fmt.Printf("      ^^Right! ")
			fmt.Printf("%s", colorReset)
			fmt.Printf("It could have been either ず or づ as they are the same sound: zu\n")
			// log the hit:
			logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCard.KeyR)
			logHits_in_cyclicArrayHits("Right", aCard.KeyR)
		} else {
			fmt.Printf("%s", colorRed)
			fmt.Printf("     ^^Oops! ")
			// log the miss:
			logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.KeyR)
			logHits_in_cyclicArrayHits("Oops", aCard.KeyR)
			logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCard.KeyR +
				":it was:" + aCard.KeyH + ":but you had guessed:" + in)

			// When this func is re-called by secondTry_, with flag set false, we skip this, and do the next
			if skipFlag == true {
				// Solicit the second guess ... and pass it to: secondTry_meatOfRomajiExorcise(in, true)
				fmt.Println("Try again") // This only prints on first pass
				fmt.Printf(colorReset)
				// Re-prompt, will send that second guess to the secondTry_meatOfRomajiExorcise func ...
				// Obtain the second guess, will pass it as 'in' to the secondTry_ func, below
				in = semi_Universal_Prompt_Scan_4_HiraResponse_NOT_a_KataPrompt(aCard.KeyR)
					branchOnUserSelectedDirectiveIfGiven(in, "Romaji_Prompt") // <-- Do the directive
				secondTry_meatOfRomajiExorcise(in) // This instance of 'in' is actually the third and final guess
			}
			// If the user guesses incorrectly on third-and-final try, then, and only then, execute the following
			if skipFlag == false { // skipFlag is true on first entry, & false when recalled after third failed attempt
				fmt.Printf("%s", colorReset)
				fmt.Printf(" It was: ")
				fmt.Printf("%s", colorCyan)
				fmt.Printf("%s", aCard.KeyH)
				fmt.Printf("%s", colorReset)
				// Display hints in white
				fmt.Printf("\n\n%s\n%s\n%s\n", aCard.Hint1h, aCard.Hint2k, aCard.Hint3TT)
				fmt.Println("")
			}
		} // If the user was ^^Right!, then we return to TouchTypingExorcise(SE) (directly from this line)
	}

	// The next two conditions are for all remaining normal (not special) prompt(key)/value events or conditions
	// ...

	if thisCaseOfAnInHasAlreadyBeenProcessedAbove != true {
		// ********************************************************
		// 'else', no special cases were found, so we process normal cases of "if 'in' == aCard.KeyR"
		// ********************************************************
		if in == aCard.KeyH { // if 'in' is the appropriate hira to match the Romaji prompt
			fmt.Printf("%s", colorGreen)
			fmt.Printf("     　^^Right! \n") // intentional '\n'
			fmt.Printf("%s", colorReset)
			// log the hit:
			logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCard.KeyR)
			logHits_in_cyclicArrayHits("Right", aCard.KeyR)
		} else {
			fmt.Printf("%s", colorRed)
			fmt.Printf("      　^^Oops! ")
			// log the miss:
			logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.KeyR)
			logHits_in_cyclicArrayHits("Oops", aCard.KeyR)
			logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCard.KeyR +
				":it was:" + aCard.KeyH + ":but you had guessed:" + in)
			//
			// When this func is re-called by secondTry_, with flag set false, we skip this, and do the next
			if skipFlag == true {
				// So, we solicit another guess ... (his second guess) and ...
				fmt.Println("Try again") // This only prints on first pass
				fmt.Printf(colorReset)
				// ... Re-prompt, and send that second guess to the secondTry_meatOfKataExorcise func
				in = semi_Universal_Prompt_Scan_4_HiraResponse_NOT_a_KataPrompt(aCard.KeyR)
					branchOnUserSelectedDirectiveIfGiven(in, "Romaji_Prompt") // <-- Do the directive
				secondTry_meatOfRomajiExorcise(in)
			}
			// If the user guesses incorrectly on third-and-final try, then, and only then, execute the following
			if skipFlag == false { // skipFlag is true on first entry, & false when recalled after third failed attempt
				fmt.Printf("%s", colorReset)
				fmt.Printf(" It was: ")
				fmt.Printf("%s", colorCyan)
				fmt.Printf("%s", aCard.KeyH)
				fmt.Printf("%s", colorReset)
				// Display hints in white
				fmt.Printf("\n\n%s\n%s\n%s\n", aCard.Hint1h, aCard.Hint2k, aCard.Hint3TT)
				fmt.Println("")
			}
		}
	}
}

/*
.
.
.
.
*/
// Second-Try version of the above func (tests the second guess, and then obtains the third and final one)
func secondTry_meatOfRomajiExorcise(in string) { // NOTE: we have already been prompted with KeyR
	var thisCaseOfAnInHasAlreadyBeenProcessedAbove bool
	thisCaseOfAnInHasAlreadyBeenProcessedAbove = false
	// *******************************************************************************************
	// Handel special prompt cases prior to doing the normal "if in == " processing
	// One 'very'-special key handler to emphasize the sameness of two variants of the zu sound
	// *******************************************************************************************
	if aCard.KeyR == "zu" {
		thisCaseOfAnInHasAlreadyBeenProcessedAbove = true
		if in == "づ" {
			fmt.Printf("%s", colorGreen)
			fmt.Printf("      ^^Right! ")
			fmt.Printf("%s", colorReset)
			fmt.Printf("It could have been either ず or づ as they are the same sound: zu\n")
			// log the hit:
			logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCard.KeyR)
			logHits_in_cyclicArrayHits("Right", aCard.KeyR)
		} else if in == "ず" {
			fmt.Printf("%s", colorGreen)
			fmt.Printf("      ^^Right! ")
			fmt.Printf("%s", colorReset)
			fmt.Printf("It could have been either ず or づ as they are the same sound: zu\n")
			// log the hit:
			logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCard.KeyR)
			logHits_in_cyclicArrayHits("Right", aCard.KeyR)
		} else {
			//
			fmt.Printf("%s", colorRed)
			fmt.Printf("     ^^Oops! ")
			// log the miss:
			logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.KeyR)
			logHits_in_cyclicArrayHits("Oops", aCard.KeyR)
			logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCard.KeyR +
				":it was:" + aCard.KeyH + ":but you had guessed:" + in)

			// Solicit the third and final guess ... and pass it to: secondTry_meatOfRomajiExorcise(in, false)
			fmt.Println("Try again") // This only prints on first pass
			fmt.Printf(colorReset)
			// Obtain the second guess, and pass it as 'in'
			// Re-prompt, and send that third and final guess to the secondTry_meatOfRomajiExorcise func
			in = semi_Universal_Prompt_Scan_4_HiraResponse_NOT_a_KataPrompt(aCard.KeyR)
				branchOnUserSelectedDirectiveIfGiven(in, "Romaji_Prompt") // <-- handle the directive
			fmt.Printf("%s", colorRed)
			fmt.Printf("      　^^Oops! ")

			logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.KeyR)
			logHits_in_cyclicArrayHits("Oops", aCard.KeyR)
			logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCard.KeyR +
				":it was:" + aCard.KeyH + ":but you had guessed:" + in)

			//
			// User failed third and final attempt, so do a "recursion", but with a skipFlag false
			meatOfRomajiExorcise(in, false)
		}
	}

	// The next two conditions are for all remaining normal (not special) prompt(key)/value events or conditions
	// ...

	if thisCaseOfAnInHasAlreadyBeenProcessedAbove != true {
		// ********************************************************
		// 'else', no special cases were found, so we process normal cases of "if in == aCard.KeyH"
		// ********************************************************
		if in == aCard.KeyH { // If 'in' is the appropriate hira (aCard.KeyH) to match the Romaji prompt
			fmt.Printf("%s", colorGreen)
			fmt.Printf("     　^^Right! \n") // Intentional '\n'
			fmt.Printf("%s", colorReset)
			// log the hit:
			logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCard.KeyR)
			logHits_in_cyclicArrayHits("Right", aCard.KeyR)
		} else {

			// Solicit the third and final guess ... and pass it to: meatOfRomajiExorcise(in, false)
			fmt.Printf("%s", colorRed)
			fmt.Println("       Try again, you have one last attempt ... ")
			fmt.Printf("%s", colorReset)
			// Obtain the last guess, and pass it as 'in'
			// Re-prompt, and send that third and last guess back to: meatOfRomajiExorcise(in, false)
			in = semi_Universal_Prompt_Scan_4_HiraResponse_NOT_a_KataPrompt(aCard.KeyR)
				branchOnUserSelectedDirectiveIfGiven(in, "Romaji_Prompt") // <-- handle the directive
			// User failed third and final attempt, so do a "recursion", but with a skipFlag false
			meatOfRomajiExorcise(in, false) // <-- Process the third try
		}
	}
}
