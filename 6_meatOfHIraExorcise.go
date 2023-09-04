package main

// **do-this**
import (
	"fmt"
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
func meatOfHiraExorcise(in string, skipFlag bool) { // NOTE: we have already been prompted with KeyR
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
		if in == "zu" {
			fmt.Printf("%s", colorGreen)
			fmt.Printf("      ^^Right! ")
			fmt.Printf("%s", colorReset)
			fmt.Printf("It could have been either ず or づ as they are the same sound: zu\n")
			// log the hit:
			if Hira_prompt_is == "hira" {
				logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCard.KeyH)
				logHits_in_cyclicArrayHits("Right", aCard.KeyH)
			} else if Hira_prompt_is == "kata" {
				logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCard.KeyK)
				logHits_in_cyclicArrayHits("Right", aCard.KeyK)
			}
		} else if in == "zu" {
			fmt.Printf("%s", colorGreen)
			fmt.Printf("      ^^Right! ")
			fmt.Printf("%s", colorReset)
			fmt.Printf("It could have been either ず or づ as they are the same sound: zu\n")
			// log the hit:
			if Hira_prompt_is == "hira" {
				logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCard.KeyH)
				logHits_in_cyclicArrayHits("Right", aCard.KeyH)
			} else if Hira_prompt_is == "kata" {
				logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCard.KeyK)
				logHits_in_cyclicArrayHits("Right", aCard.KeyK)
			}
		} else {
			fmt.Printf("%s", colorRed)
			fmt.Printf("     ^^Oops! ")
			// log the miss:
			if Hira_prompt_is == "hira" {
				logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.KeyH)
				logHits_in_cyclicArrayHits("Oops", aCard.KeyH)
				logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCard.KeyH +
					":it was:" + aCard.KeyR + ":but you had guessed:" + in)
			} else if Hira_prompt_is == "kata" {
				logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.KeyK)
				logHits_in_cyclicArrayHits("Oops", aCard.KeyK)
				logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCard.KeyK +
					":it was:" + aCard.KeyR + ":but you had guessed:" + in)
			}

			// When this func is re-called by secondTry_, with flag set false, we skip this, and do the next
			if skipFlag == true {
				// Solicit the second guess ... and pass it to: secondTry_meatOfRomajiExorcise(in, true)
				fmt.Println("Try again") // This only prints on first pass
				fmt.Printf(colorReset)
				// Re-prompt, will send that second guess to the secondTry_meatOfRomajiExorcise func ...
				// Obtain the second guess, will pass it as 'in' to the secondTry_ func, below
				in = prompt_and_Scan_4_RomajiResponse_to_HiraPrompt(Hira_prompt_KeyX)
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
					branchOnUserSelectedDirectiveIfGiven(in, "Hira_Prompt") // <-- Do the directive
				}
				secondTry_meatOfHiraExorcise(in) // This instance of 'in' is actually the third and final guess
			}
			// If the user guesses incorrectly on third-and-final try, then, and only then, execute the following
			if skipFlag == false { // skipFlag is true on first entry, & false when recalled after third failed attempt
				fmt.Printf("%s", colorReset)
				fmt.Printf("\n It was: ")
				fmt.Printf("%s", colorCyan)
				fmt.Printf("%s", aCard.KeyR)
				fmt.Printf("%s", colorReset)
				//checkForHints(aCard.KeyH) // Note: we only give hints for non-conjunctives
				giveHintInResponseToSingleQuestionMarkContextSensitiveDir(aCard.KeyH, aCard.Hint1h, aCard.Hint2k, aCard.Hint3TT) // dole-out a hint
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
		if in == aCard.KeyR { // if 'in' is the appropriate Romaji to match the hira or kata prompt
			fmt.Printf("%s", colorGreen)
			fmt.Printf("     　^^Right! \n") // intentional '\n'
			fmt.Printf("%s", colorReset)
			// log the hit:
			if Hira_prompt_is == "hira" {
				logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCard.KeyH)
				logHits_in_cyclicArrayHits("Right", aCard.KeyH)
			} else if Hira_prompt_is == "kata" {
				logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCard.KeyK)
				logHits_in_cyclicArrayHits("Right", aCard.KeyK)
			}
		} else {
			fmt.Printf("%s", colorRed)
			fmt.Printf("      　^^Oops! ")
			// log the miss:
			if Hira_prompt_is == "hira" {
				logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.KeyH)
				logHits_in_cyclicArrayHits("Oops", aCard.KeyH)
				logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCard.KeyH +
					":it was:" + aCard.KeyR + ":but you had guessed:" + in)
			} else if Hira_prompt_is == "kata" {
				logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.KeyK)
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
				// ... Re-prompt, and send that second guess to the secondTry_meatOfKataExorcise func
				in = prompt_and_Scan_4_RomajiResponse_to_HiraPrompt(Hira_prompt_KeyX)
				//in = semi_Universal_Prompt_Scan_4_HiraResponse_NOT_a_KataPrompt(aCard.KeyH)
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
					branchOnUserSelectedDirectiveIfGiven(in, "Hira_Prompt") // <-- Do the directive
				}
				secondTry_meatOfHiraExorcise(in)
			}
			// If the user guesses incorrectly on third-and-final try, then, and only then, execute the following
			if skipFlag == false { // skipFlag is true on first entry, & false when recalled after third failed attempt
				fmt.Printf("%s", colorReset)
				fmt.Printf("\n It was: ")
				fmt.Printf("%s", colorCyan)
				fmt.Printf("%s", aCard.KeyR)
				fmt.Printf("%s", colorReset)
				//checkForHints(aCard.KeyH) // Note: we only give hints for non-conjunctives
				giveHintInResponseToSingleQuestionMarkContextSensitiveDir(aCard.KeyH, aCard.Hint1h, aCard.Hint2k, aCard.Hint3TT) // dole-out a hint
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
func secondTry_meatOfHiraExorcise(in string) { // NOTE: we have already been prompted with KeyR
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
			if Hira_prompt_is == "hira" {
				logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCard.KeyH)
				logHits_in_cyclicArrayHits("Right", aCard.KeyH)
			} else if Hira_prompt_is == "kata" {
				logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCard.KeyK)
				logHits_in_cyclicArrayHits("Right", aCard.KeyK)
			}
		} else if in == "zu" {
			fmt.Printf("%s", colorGreen)
			fmt.Printf("      ^^Right! ")
			fmt.Printf("%s", colorReset)
			fmt.Printf("It could have been either ず or づ as they are the same sound: zu\n")
			// log the hit:
			if Hira_prompt_is == "hira" {
				logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCard.KeyH)
				logHits_in_cyclicArrayHits("Right", aCard.KeyH)
			} else if Hira_prompt_is == "kata" {
				logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCard.KeyK)
				logHits_in_cyclicArrayHits("Right", aCard.KeyK)
			}
		} else {
			//
			fmt.Printf("%s", colorRed)
			fmt.Printf("     ^^Oops! ")
			// log the miss:
			if Hira_prompt_is == "hira" {
				logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.KeyH)
				logHits_in_cyclicArrayHits("Oops", aCard.KeyH)
				logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCard.KeyH +
					":it was:" + aCard.KeyR + ":but you had guessed:" + in)
			} else if Hira_prompt_is == "kata" {
				logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.KeyK)
				logHits_in_cyclicArrayHits("Oops", aCard.KeyK)
				logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCard.KeyK +
					":it was:" + aCard.KeyR + ":but you had guessed:" + in)
			}

			// Solicit the third and final guess ... and pass it to: secondTry_meatOfRomajiExorcise(in, false)
			fmt.Println("Try again") // This only prints on first pass
			fmt.Printf(colorReset)
			// Obtain the second guess, and pass it as 'in'
			// Re-prompt, and send that third and final guess to the secondTry_meatOfRomajiExorcise func
			in = prompt_and_Scan_4_RomajiResponse_to_HiraPrompt(Hira_prompt_KeyX)
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
				branchOnUserSelectedDirectiveIfGiven(in, "Romaji_Prompt") // <-- handle the directive
			}
			fmt.Printf("%s", colorRed)
			fmt.Printf("      　^^Oops! ")

			if Hira_prompt_is == "hira" {
				logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.KeyH)
				logHits_in_cyclicArrayHits("Oops", aCard.KeyH)
				logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCard.KeyH +
					":it was:" + aCard.KeyR + ":but you had guessed:" + in)
			} else if Hira_prompt_is == "kata" {
				logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.KeyK)
				logHits_in_cyclicArrayHits("Oops", aCard.KeyK)
				logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCard.KeyK +
					":it was:" + aCard.KeyR + ":but you had guessed:" + in)
			}

			//
			// User failed third and final attempt, so do a "recursion", but with a skipFlag false
			meatOfHiraExorcise(in, false)
		}
	}

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
			if Hira_prompt_is == "hira" {
				logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCard.KeyH)
				logHits_in_cyclicArrayHits("Right", aCard.KeyH)
			} else if Hira_prompt_is == "kata" {
				logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCard.KeyK)
				logHits_in_cyclicArrayHits("Right", aCard.KeyK)
			}
		} else {

			// Solicit the third and final guess ... and pass it to: meatOfRomajiExorcise(in, false)
			fmt.Printf("%s", colorRed)
			fmt.Println("       Try again, you have one last attempt ... ")
			fmt.Printf("%s", colorReset)
			// Obtain the last guess, and pass it as 'in'
			// Re-prompt, and send that third and last guess back to: meatOfRomajiExorcise(in, false)
			in = prompt_and_Scan_4_RomajiResponse_to_HiraPrompt(Hira_prompt_KeyX)
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
				branchOnUserSelectedDirectiveIfGiven(in, "Romaji_Prompt") // <-- handle the directive
			}
			// User failed third and final attempt, so do a "recursion", but with a skipFlag false
			meatOfHiraExorcise(in, false) // <-- Process the third try
		}
	}
}
