package main

// **do-this**
import (
	"fmt"
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
func meatOf_Mixed_HiraKataExercise(in string, skipFlag bool) { // NOTE: we have already been prompted with KeyR
	// ..^ ^ ^ called as|with: (usersGuessOrOptionDirective, (bool)'skipFlag')
	// so, 'in' == the-users-guess, which is ...
	// obtained prior to this func being called ...
	// ^^^^^^^^ either by TouchTypingExercises12346_10, ...
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
		if in == "zu" {
			fmt.Printf("%s", colorGreen)
			fmt.Printf("      ^^Right! ")
			fmt.Printf("%s", colorReset)
			fmt.Printf("It could have been either ず or づ as they are the same sound: zu\n")
			// log the hit:
			if Mixed_prompt_is == "hira" {
				logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCard.KeyH)
				logHits_in_cyclicArrayHits("Right", aCard.KeyH)
			} else if Mixed_prompt_is == "kata" {
				logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCard.KeyK)
				logHits_in_cyclicArrayHits("Right", aCard.KeyK)
			}
		} else if in == "zu" {
			fmt.Printf("%s", colorGreen)
			fmt.Printf("      ^^Right! ")
			fmt.Printf("%s", colorReset)
			fmt.Printf("It could have been either ず or づ as they are the same sound: zu\n")
			// log the hit:
			if Mixed_prompt_is == "hira" {
				logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCard.KeyH)
				logHits_in_cyclicArrayHits("Right", aCard.KeyH)
			} else if Mixed_prompt_is == "kata" {
				logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCard.KeyK)
				logHits_in_cyclicArrayHits("Right", aCard.KeyK)
			}
		} else {
			fmt.Printf("%s", colorRed)
			fmt.Printf("     ^^Oops! ")
			// log the miss:
			if Mixed_prompt_is == "hira" {
				// logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.KeyH) // Because, it was prob a typo
				logHits_in_cyclicArrayHits("Oops", aCard.KeyH)
				logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCard.KeyH +
					":it was:" + aCard.KeyR + ":but you had guessed:" + in)
			} else if Mixed_prompt_is == "kata" {
				// logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.KeyK) // Because, it was prob a typo
				logHits_in_cyclicArrayHits("Oops", aCard.KeyK)
				logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCard.KeyK +
					":it was:" + aCard.KeyR + ":but you had guessed:" + in)
			}

			// When this func is re-called by secondTry_, with flag set false, we skip this, and do the next
			if skipFlag == true {
				// Solicit the second guess ... and pass it to: secondTry_meatOfRomajiExercise(in, true)
				fmt.Println("Try again") // This only prints on first pass
				fmt.Printf(colorReset)
				// Re-prompt, will send that second guess to the secondTry_meatOfRomajiExercise func ...
				// Obtain the second guess, will pass it as 'in' to the secondTry_ func, below
				in = prompt_and_Scan_4_RomajiResponse_to_Any_Prompt(Mixed_prompts_KeyX)
					branchOnUserSelectedDirectiveIfGiven(in, "Mixed_prompts") // <-- Do the directive
				secondTry_meatOf_mixed_HiraKataExercise(in)                   // This instance of 'in' is actually the third and final guess
			}
			// If the user guesses incorrectly on third-and-final try, then, and only then, execute the following
			if skipFlag == false { // skipFlag is true on first entry, & false when recalled after third failed attempt
				fmt.Printf("%s", colorReset)
				fmt.Printf(" It was: ")
				fmt.Printf("%s", colorCyan)
				fmt.Printf("%s", aCard.KeyR)
				fmt.Printf("%s", colorReset)
				// Display hints in white
				fmt.Printf("\n\n%s\n%s\n%s\n", aCard.Hint1h, aCard.Hint2k, aCard.Hint3TT)
				fmt.Println("")
			}
		} // If the user was ^^Right!, then we return to TouchTypingExercises12346_10(SE) (directly from this line)
	}

	// The next two conditions are for all remaining normal (not special) prompt(key)/value events or conditions
	// ...

	if thisCaseOfAnInHasAlreadyBeenProcessedAbove != true {
		// ********************************************************
		// 'else', no special cases were found, so we process normal cases of "if 'in' == aCard.KeyR"
		// ********************************************************
		if in == aCard.KeyR { // if 'in' is the appropriate Romaji to match the hira or kata prompt
			fmt.Printf("%s", colorGreen)
			fmt.Printf("     　^^Right! \n") // intentional '\n'
			fmt.Printf("%s", colorReset)
			// log the hit:
			if Mixed_prompt_is == "hira" {
				logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCard.KeyH)
				logHits_in_cyclicArrayHits("Right", aCard.KeyH)
			} else if Mixed_prompt_is == "kata" {
				logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCard.KeyK)
				logHits_in_cyclicArrayHits("Right", aCard.KeyK)
			}
		} else {
			fmt.Printf("%s", colorRed)
			fmt.Printf("      　^^Oops! ")
			// log the miss:
			if Mixed_prompt_is == "hira" {
				// logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.KeyH) // Because, it was prob a typo
				logHits_in_cyclicArrayHits("Oops", aCard.KeyH)
				logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCard.KeyH +
					":it was:" + aCard.KeyR + ":but you had guessed:" + in)
			} else if Mixed_prompt_is == "kata" {
				// logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.KeyK) // Because, it was prob a typo
				logHits_in_cyclicArrayHits("Oops", aCard.KeyK)
				logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCard.KeyK +
					":it was:" + aCard.KeyR + ":but you had guessed:" + in)
			}
			//
			// When this func is re-called by secondTry_, with flag set false, we skip this, and do the next
			if skipFlag == true {
				// So, we solicit another guess ... (his second guess) and ...
				fmt.Println("Try again") // This only prints on first pass
				fmt.Printf(colorReset)
				// ... Re-prompt, and send that second guess to the secondTry_meatOfKataExercise func
				in = prompt_and_Scan_4_RomajiResponse_to_Any_Prompt(Mixed_prompts_KeyX)
				// in = semi_Universal_Prompt_Scan_4_HiraResponse_to_RomajiPrompt(aCard.KeyH)
					branchOnUserSelectedDirectiveIfGiven(in, "Mixed_prompts") // <-- Do the directive
				secondTry_meatOf_mixed_HiraKataExercise(in)
			}
			// If the user guesses incorrectly on third-and-final try, then, and only then, execute the following
			if skipFlag == false { // skipFlag is true on first entry, & false when recalled after third failed attempt
				fmt.Printf("%s", colorReset)
				fmt.Printf(" It was: ")
				fmt.Printf("%s", colorCyan)
				fmt.Printf("%s", aCard.KeyR)
				fmt.Printf("%s", colorReset)
				// Display hints in white
				fmt.Printf("\n\n%s\n%s\n%s\n", aCard.Hint1h, aCard.Hint2k, aCard.Hint3TT)
				fmt.Println()

// These probably belong here ?????????????????????????????????????????????, rather than below 
				if Mixed_prompt_is == "hira" {
					logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.KeyH)
					logHits_in_cyclicArrayHits("Oops", aCard.KeyH)
					logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCard.KeyH +
						":it was:" + aCard.KeyR + ":but you had guessed:" + in)
				} else if Mixed_prompt_is == "kata" {
					logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.KeyK)
					logHits_in_cyclicArrayHits("Oops", aCard.KeyK)
					logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCard.KeyK +
						":it was:" + aCard.KeyR + ":but you had guessed:" + in)
				}
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
func secondTry_meatOf_mixed_HiraKataExercise(in string) { // NOTE: we have already been prompted with KeyR
	var thisCaseOfAnInHasAlreadyBeenProcessedAbove bool
	thisCaseOfAnInHasAlreadyBeenProcessedAbove = false
	// *******************************************************************************************
	// Handel special prompt cases prior to doing the normal "if in == " processing
	// One 'very'-special key handler to emphasize the sameness of two variants of the zu sound
	// *******************************************************************************************
	if aCard.KeyR == "zu" {
		thisCaseOfAnInHasAlreadyBeenProcessedAbove = true
		if in == "zu" {
			fmt.Printf("%s", colorGreen)
			fmt.Printf("      ^^Right! ")
			fmt.Printf("%s", colorReset)
			fmt.Printf("It could have been either ず or づ as they are the same sound: zu\n")
			// log the hit:
			if Mixed_prompt_is == "hira" {
				logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCard.KeyH)
				logHits_in_cyclicArrayHits("Right", aCard.KeyH)
			} else if Mixed_prompt_is == "kata" {
				logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCard.KeyK)
				logHits_in_cyclicArrayHits("Right", aCard.KeyK)
			}
		} else if in == "zu" { // this duplicate does not seem to make sense ?????????????????????????????????????????????
			fmt.Printf("%s", colorGreen)
			fmt.Printf("      ^^Right! ")
			fmt.Printf("%s", colorReset)
			fmt.Printf("It could have been either ず or づ as they are the same sound: zu\n")
			// log the hit:
			if Mixed_prompt_is == "hira" {
				logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCard.KeyH)
				logHits_in_cyclicArrayHits("Right", aCard.KeyH)
			} else if Mixed_prompt_is == "kata" {
				logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCard.KeyK)
				logHits_in_cyclicArrayHits("Right", aCard.KeyK)
			}
		} else { // if aCard.KeyR == "zu" and in != "zu" // but it is a strange and confusing construct !!!!!!!!!!!!!!!!!!
			//
			fmt.Printf("%s", colorRed)
			fmt.Printf("     ^^Oops! ")
			// log the miss:
			if Mixed_prompt_is == "hira" {
				logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.KeyH)
				logHits_in_cyclicArrayHits("Oops", aCard.KeyH)
				logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCard.KeyH +
					":it was:" + aCard.KeyR + ":but you had guessed:" + in)
			} else if Mixed_prompt_is == "kata" {
				logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.KeyK)
				logHits_in_cyclicArrayHits("Oops", aCard.KeyK)
				logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCard.KeyK +
					":it was:" + aCard.KeyR + ":but you had guessed:" + in)
			}

			// Solicit the third and final guess ... and pass it to: secondTry_meatOfRomajiExercise(in, false)
			fmt.Println("Try again") // This only prints on first pass
			fmt.Printf(colorReset)
			// Obtain the second guess, and pass it as 'in'
			// Re-prompt, and send that third and final guess to the secondTry_meatOfRomajiExercise func
			in = prompt_and_Scan_4_RomajiResponse_to_Any_Prompt(Mixed_prompts_KeyX)
				branchOnUserSelectedDirectiveIfGiven(in, "Mixed_prompts") // <-- handle the directive
				
				// ???????????????????????????????????????????????????
				// like 1, 10, and 2 -- 6 works rightly (because this silly oops is only in the zu if)
			fmt.Printf("%s", colorRed)
			fmt.Printf("      　^^Oops! ")

			if Mixed_prompt_is == "hira" {
				logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.KeyH)
				logHits_in_cyclicArrayHits("Oops", aCard.KeyH)
				logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCard.KeyH +
					":it was:" + aCard.KeyR + ":but you had guessed:" + in)
			} else if Mixed_prompt_is == "kata" {
				logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.KeyK)
				logHits_in_cyclicArrayHits("Oops", aCard.KeyK)
				logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCard.KeyK +
					":it was:" + aCard.KeyR + ":but you had guessed:" + in)
			}

			//
			// User failed third and final attempt, so do a "recursion", but with a skipFlag false
				
				 
			meatOf_Mixed_HiraKataExercise(in, false)
		}
	} // end of the if that only handles cases of zu 

	// The next two conditions are for all remaining normal (not special) prompt(key)/value events or conditions
	// ...

	if thisCaseOfAnInHasAlreadyBeenProcessedAbove != true {
		// ********************************************************
		// 'else', no special cases were found, so we process normal cases of "if in == aCard.KeyH"
		// ********************************************************
		if in == aCard.KeyR { // if 'in' is the appropriate Romaji to match the hira or kata prompt
			fmt.Printf("%s", colorGreen)
			fmt.Printf("     　^^Right! \n") // Intentional '\n'
			fmt.Printf("%s", colorReset)
			// log the hit:
			if Mixed_prompt_is == "hira" {
				logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCard.KeyH)
				logHits_in_cyclicArrayHits("Right", aCard.KeyH)
			} else if Mixed_prompt_is == "kata" {
				logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCard.KeyK)
				logHits_in_cyclicArrayHits("Right", aCard.KeyK)
			}
		} else {

			// Solicit the third and final guess ... and pass it to: meatOfRomajiExercise(in, false)
			fmt.Printf("%s", colorRed)
			fmt.Println("       Try again, you have one last attempt ... ")
			fmt.Printf("%s", colorReset)
			// Obtain the last guess, and pass it as 'in'
			// Re-prompt, and send that third and last guess back to: meatOfRomajiExercise(in, false)
			in = prompt_and_Scan_4_RomajiResponse_to_Any_Prompt(Mixed_prompts_KeyX)
				branchOnUserSelectedDirectiveIfGiven(in, "Mixed_prompts") // <-- handle the directive
			// User failed third and final attempt, so do a "recursion", but with a skipFlag false
			meatOf_Mixed_HiraKataExercise(in, false) // <-- Process the third try
		}
	}
}
