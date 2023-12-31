package main

// **do-this** : is the tag I have used to denote sections that I have yet to complete

import "fmt"

/*
.
Refer to 8, which is a similar function

All this means that:
The first of these two functions is passed the user's first guess, and, optionally, obtains a second guess
Optionally, the second func is passed the second guess, and obtains a third guess
Finally, the first func is re-called by the second,: the first func is, then, passed the third and final guess
.
*/
func meatOfRomajiKataExercise(in, selectedExercise string, skipFlag bool) {
	// ..^ ^ ^ called as|with: (usersGuessOrOptionDirective, (bool)'skipFlag')
	// so, 'in' == the-users-guess, which is ...
	// obtained prior to this func being called ...
	// ^^^^^^^^ either by TouchTypingExercises12346_10, or by: secondTry_ (secondTry_ actually solicits and obtains the THIRD guess)
	//
	var thisCaseOfAnInHasAlreadyBeenProcessedAbove bool
	thisCaseOfAnInHasAlreadyBeenProcessedAbove = false

	// *******************************************************************************************
	// Handel special prompt cases prior to doing the normal "if in == " processing
	// One 'very'-special key handler to emphasize the sameness of two variants of the zu sound
	// *******************************************************************************************
	//
	if aCard.KeyRK == "zuズ" || aCard.KeyRK == "zuヅ" {
		if in == "づ" || in == "ず" {
			thisCaseOfAnInHasAlreadyBeenProcessedAbove = true
			fmt.Printf("%s", colorGreen)
			fmt.Printf("        ^^Right! ")
			//
			logHits_in_cyclicArrayHits("Right", aCard.KeyRK)
			logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCard.KeyR)
			//
			fmt.Printf("%s", colorReset)
			fmt.Printf("It could have been either ず or づ as they are the same sound\n")
		} else {
			fmt.Printf("%s", colorRed)
			fmt.Printf("        ^^Oops! ")
			//
			logHits_in_cyclicArrayHits("Oops", aCard.KeyRK)
			// logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.KeyRK) // Because, it was probably just a typo
			logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCard.KeyRK + ":it was:" + aCard.KeyH + ":but you had guessed:" + in)
			//
			// When this func is re-called by secondTry_, with flag set false, we skip this, and do the next
			if skipFlag == true {
				// Solicit the second guess ... and pass it to: secondTry_meatOfRomajiKataExercise(in, true)

				fmt.Println("Try again") // This only prints on first pass
				fmt.Printf(colorReset)
				// Re-prompt, will send that second guess to the secondTry_meatOfRomajiKataExercise func ...
				// Obtain the second guess, will pass it as 'in' to the secondTry_ func, below
				in = semi_Universal_Prompt_Scan_4_HiraResponse_to_RomajiPrompt(aCard.KeyRK, selectedExercise)
					branchOnUserSelectedDirectiveIfGiven(in, "Romaji_w_Kata_Prompt") // <-- Do directive
				secondTry_meatOfRomajiKataExercise(in, selectedExercise)  // This instance of 'in' is actually the third and final guess
			}
			// If the user guesses incorrectly on third-and-final try, then, and only then, execute the following
			if skipFlag == false { // skipFlag is true on first entry, & false when recalled after third failed attempt
				fmt.Printf("%s", colorReset)
				fmt.Printf(" It was: ")
				fmt.Printf("%s", colorCyan)
				fmt.Printf("%s", aCard.KeyH) // done
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
		// 'else', no special cases were found, so we process normal cases of "if 'in' == aCard.KeyH"
		// ********************************************************
		if in == aCard.KeyH { // if 'in' is the appropriate hira to match the Romaji+Kata prompt
			fmt.Printf("%s", colorGreen)
			fmt.Printf("        ^^Right! \n")
			fmt.Printf("%s", colorReset)
			//
			logHits_in_cyclicArrayHits("Right", aCard.KeyRK)
			logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCard.KeyR)
			//
		} else {
			fmt.Printf("%s", colorRed)
			fmt.Printf("      　 ^^Oops! ")
			//
			logHits_in_cyclicArrayHits("Oops", aCard.KeyRK)
			// logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.KeyRK) // Because, it was probably just a typo
			logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCard.KeyRK +
				":it was:" + aCard.KeyH + ":but you had guessed:" + in)
			//
			//
			// When this func is re-called by secondTry_, with flag set false, we skip this, and do the next
			if skipFlag == true {
				// So, we solicit another guess ... (his second guess) and ...
				fmt.Println("Try again") // This only prints on first pass
				fmt.Printf(colorReset)
				// ... Re-prompt, and send that second guess to the secondTry_meatOfRomajiKataExercise func
				in = semi_Universal_Prompt_Scan_4_HiraResponse_to_RomajiPrompt(aCard.KeyRK, selectedExercise) // done RK ?
					branchOnUserSelectedDirectiveIfGiven(in, "Romaji_w_Kata_Prompt")        // <-- Do the directive
				secondTry_meatOfRomajiKataExercise(in, selectedExercise)                                      // done RK
			}
			// If the user guesses incorrectly on third-and-final try, then, and only then, execute the following
			if skipFlag == false { // skipFlag is true on first entry, & false when recalled after third failed attempt
				fmt.Printf("%s", colorReset)
				fmt.Printf(" It was: ")
				fmt.Printf("%s", colorCyan)
				fmt.Printf("%s", aCard.KeyH) // done RK
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
*/
// Second-Try version of the above func (tests the second guess, and then obtains the third and final one)
func secondTry_meatOfRomajiKataExercise(in, selectedExercise string) {
	var thisCaseOfAnInHasAlreadyBeenProcessedAbove bool
	thisCaseOfAnInHasAlreadyBeenProcessedAbove = false

	// *******************************************************************************************
	// Handel special prompt cases prior to doing the normal "if in == " processing
	// One 'very'-special key handler to emphasize the sameness of two variants of the zu sound
	// *******************************************************************************************
	//
	if aCard.KeyRK == "zuズ" || aCard.KeyRK == "zuヅ" {
		if in == "づ" || in == "ず" {
			thisCaseOfAnInHasAlreadyBeenProcessedAbove = true
			fmt.Printf("%s", colorGreen)
			fmt.Printf("        ^^Right! ")
			//
			logHits_in_cyclicArrayHits("Right", aCard.KeyRK)
			logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCard.KeyR)
			//
			fmt.Printf("%s", colorReset)
			fmt.Printf("It could have been either ず or づ as they are the same sound\n")
		} else {
				branchOnUserSelectedDirectiveIfGiven(in, "Romaji_w_Kata_Prompt") // <-- Do the directive
			fmt.Printf("%s", colorRed)
			fmt.Printf("        ^^Oops! ")
			//
			logHits_in_cyclicArrayHits("Oops", aCard.KeyRK)
			logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.KeyR) // Don't even think about "fixing" this
			logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCard.KeyRK +
				":it was:" + aCard.KeyH + ":but you had guessed:" + in)
			//
			// User failed third and final attempt, so do a "recursion", but with a skipFlag false
			meatOfRomajiKataExercise(in, selectedExercise,false)
		}
	} // end of zu if 
	// ...
	// The next two conditions are for all remaining normal (not special) prompt(key)/value events or conditions
	// ...

	if thisCaseOfAnInHasAlreadyBeenProcessedAbove != true {
		// ********************************************************
		// 'else', no special cases were found, so we process normal cases of "if in == aCard.KeyH"
		// ********************************************************
		if in == aCard.KeyH {
			fmt.Printf("%s", colorGreen)
			fmt.Printf("        ^^Right! \n")
			fmt.Printf("%s", colorReset)
			//
			logHits_in_cyclicArrayHits("Right", aCard.KeyRK)
			logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCard.KeyR)
			//
		} else {

			// Solicit the third and final guess ... and pass it to: meatOfRomajiExercise(in, false)
			fmt.Printf("%s", colorRed)
			fmt.Println("       Try again, you have one last attempt ... ")
			fmt.Printf("%s", colorReset)
			// Re-prompt, and send that third and final guess back to: meatOfRomajiKataExercise(in, false)
			in = semi_Universal_Prompt_Scan_4_HiraResponse_to_RomajiPrompt(aCard.KeyRK, selectedExercise) // Obtain the last guess, and pass it as 'in' // done RK
				branchOnUserSelectedDirectiveIfGiven(in, "Romaji_w_Kata_Prompt")        // <-- handle the directive // done
			// User failed third and final attempt, so do a "recursion", but with a skipFlag false
			meatOfRomajiKataExercise(in, selectedExercise,false) // Process the third try // done
		}
	}
}
