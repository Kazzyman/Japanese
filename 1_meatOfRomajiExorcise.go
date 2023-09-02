package main

// **do-this** 'these'
import (
	"fmt"
)

func meatOfRomajiExorcise(in string, skipFlag bool) { // NOTE: we have already been prompted with KeyR
	// ..^ ^ ^ called as|with: (usersGuessOrOptionDirective, (bool)'skipFlag')
	// so, 'in' == the-users-guess

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
			logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCard.KeyR + ":it was:" + aCard.KeyH + ":but you had guessed:" + in)

			// When this func is re-called by the second, with flag set false, we skip this, and do the next
			if skipFlag == true {
				// Solicit the second guess ... and pass it to: secondTry_meatOfRomajiExorcise(in, true)
				fmt.Println("Try again") // This only prints on first pass
				fmt.Printf(colorReset)
				// Re-prompt, and send that second guess to the secondTry_meatOfRomajiExorcise func
				in = semi_Universal_Prompt_Scan_4_HiraResponse_NOT_a_KataPrompt(aCard.KeyR) // Obtain the second guess, and pass it as 'in' // done
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
					branchOnUserSelectedDirectiveIfGiven(in, "Romaji_Prompt") // <-- handle the directive // done
				}
				secondTry_meatOfRomajiExorcise(in) // This instance of 'in' is actually the third and final guess // done
			}
			// If the user guesses incorrectly on his third-and-final try, then, and only then, execute the rest of this func
			if skipFlag == false { // skipFlag will be true on first entry, but false upon being recalled after the third failed attempt
				fmt.Printf("%s", colorReset)
				fmt.Printf("\n It was: ")
				fmt.Printf("%s", colorCyan)
				fmt.Printf("%s", aCard.KeyH) // done
				fmt.Printf("%s", colorReset)
				checkForHints(aCard.KeyH) // Note: we only give hints for non-conjunctives
				fmt.Println("")
			}
		} // If the user was ^^Right!, then we return to TouchTypingExorcise(selectedExorcise) (directly from this very line)
	}

	// the next two conditions are for all remaining normal (not special) prompt:as:KeyR events or conditions
	// ...

	if thisCaseOfAnInHasAlreadyBeenProcessedAbove != true {
		// ********************************************************
		// 'else', no special cases were found, so we process normal cases of "if in == aCard.KeyR"
		// ********************************************************
		if in == aCard.KeyH { // if in is the appropriate hira to match the Romaji prompt
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
			logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCard.KeyR + ":it was:" + aCard.KeyH + ":but you had guessed:" + in)
			//
			// When this func is re-called by the third, with flag set false, we skip this, and do the next
			if skipFlag == true {
				// So, we solicit another guess ... (his second guess) and ...
				fmt.Println("Try again") // This only prints on first pass
				fmt.Printf(colorReset)
				// Re-prompt, and send that second guess to the secondTry_meatOfKataExorcise func
				in = semi_Universal_Prompt_Scan_4_HiraResponse_NOT_a_KataPrompt(aCard.KeyR) // done
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
					branchOnUserSelectedDirectiveIfGiven(in, "Romaji_Prompt") // <-- handle the directive
				}
				secondTry_meatOfRomajiExorcise(in) // done
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
func secondTry_meatOfRomajiExorcise(in string) { // NOTE: we have already been prompted with KeyR
	// ..^ ^ ^ called as|with: (usersGuessOrOptionDirective, aCard.KeyH) that KeyH being the particular Hira in question
	// so, in=the-users-guess, keyH=aCard.KeyH

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

			fmt.Printf("%s", colorRed)
			fmt.Printf("     ^^Oops! ")
			// log the miss:
			logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.KeyR)
			logHits_in_cyclicArrayHits("Oops", aCard.KeyR)
			logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCard.KeyR + ":it was:" + aCard.KeyH + ":but you had guessed:" + in)

			// Solicit the second guess ... and pass it to: secondTry_meatOfRomajiExorcise(in, true)
			fmt.Println("Try again") // This only prints on first pass
			fmt.Printf(colorReset)
			// Re-prompt, and send that second guess to the secondTry_meatOfRomajiExorcise func
			in = semi_Universal_Prompt_Scan_4_HiraResponse_NOT_a_KataPrompt(aCard.KeyR) // Obtain the second guess, and pass it as 'in' // done
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
				branchOnUserSelectedDirectiveIfGiven(in, "Romaji_Prompt") // <-- handle the directive // done
			}
			fmt.Printf("%s", colorRed)
			fmt.Printf("      　^^Oops! ")

			logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.KeyR)
			logHits_in_cyclicArrayHits("Oops", aCard.KeyR)
			logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCard.KeyR + ":it was:" + aCard.KeyH + ":but you had guessed:" + in)

			//
			// User failed second attempt, so do a "recursion", but with a skipFlag
			meatOfRomajiExorcise(in, false) // done
		}
	}

	// the next two conditions are for all remaining normal (not special) prompt:as:KeyR events or conditions
	// ...

	if thisCaseOfAnInHasAlreadyBeenProcessedAbove != true {
		// ********************************************************
		// 'else', no special cases were found, so we process normal cases of "if in == aCard.KeyR"
		// ********************************************************
		if in == aCard.KeyH { // if 'in' is the appropriate hira (aCard.KeyH) to match the Romaji prompt
			fmt.Printf("%s", colorGreen)
			fmt.Printf("     　^^Right! \n") // intentional '\n'
			fmt.Printf("%s", colorReset)
			// log the hit:
			logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCard.KeyR)
			logHits_in_cyclicArrayHits("Right", aCard.KeyR)
		} else {

			// Solicit the third and final guess ... and pass it to: meatOfRomajiExorcise(in, false)
			fmt.Printf("%s", colorRed)
			fmt.Println("       Try again, you have one last attempt ... ")
			fmt.Printf("%s", colorReset)
			// Re-prompt, and send that third and last guess back to: meatOfRomajiExorcise(in, false)
			in = semi_Universal_Prompt_Scan_4_HiraResponse_NOT_a_KataPrompt(aCard.KeyR) // Obtain the last guess, and pass it as 'in' // done
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
				branchOnUserSelectedDirectiveIfGiven(in, "Romaji_Prompt") // <-- handle the directive // done
			}
			meatOfRomajiExorcise(in, false) // Process the third try // done
		}
	}
}
