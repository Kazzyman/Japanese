package main

// **do-this** : is the tag I have used to denote sections that I have yet to complete

import "fmt"

func meatOfRomajiKataExorcise(in string, skipFlag bool) {
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
			logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.KeyR)
			logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCard.KeyRK + ":it was:" + aCard.KeyH + ":but you had guessed:" + in)
			//
			// When this func is re-called by the second, with flag set false, we skip this, and do the next

			if skipFlag == true {
				// Solicit the second guess ... and pass it to: secondTry_meatOfRomajiKataExorcise(in, true)
				fmt.Println("Try again") // This only prints on first pass
				fmt.Printf(colorReset)
				// Re-prompt, and send that second guess to the secondTry_meatOfRomajiKataExorcise func
				in = semi_Universal_Prompt_Scan_4_HiraResponse_NOT_a_KataPrompt(aCard.KeyRK) // Obtain the second guess, and pass it as 'in' // done
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
					branchOnUserSelectedDirectiveIfGiven(in, "Romaji+Kata_Prompt") // <-- handle the directive // done
				}
				secondTry_meatOfRomajiKataExorcise(in) // This instance of 'in' is actually the third and final guess // done
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

	} else {
		// This 'else' (immediately above) covers a lot of if, else ifs: Dozens of them!
		// ...
		// the next two conditions are for all remaining normal (not special) prompt(key)/value events or conditions
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
				fmt.Printf("%s", colorRed)
				fmt.Printf("      　 ^^Oops! ")
				//
				logHits_in_cyclicArrayHits("Oops", aCard.KeyRK)
				logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.KeyR) // This is correct
				logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCard.KeyRK + ":it was:" + aCard.KeyH + ":but you had guessed:" + in)
				//
				//
				if skipFlag == true {
					// So, we solicit another guess ... (his second guess) and ...
					fmt.Println("Try again") // This only prints on first pass
					fmt.Printf(colorReset)
					// Re-prompt, and send that second guess to the secondTry_meatOfRomajiKataExorcise func
					in = semi_Universal_Prompt_Scan_4_HiraResponse_NOT_a_KataPrompt(aCard.KeyRK) // done RK ?
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
						branchOnUserSelectedDirectiveIfGiven(in, "Romaji+Kata_Prompt") // <-- handle the directive
					}
					secondTry_meatOfRomajiKataExorcise(in) // done RK
				}
				// If the user guesses incorrectly on his third-and-final try, then, and only then, execute the rest of this func
				if skipFlag == false { // skipFlag will be true on first entry, but false upon being recalled after the third failed attempt
					fmt.Printf("%s", colorReset)
					fmt.Printf("\n It was: ")
					fmt.Printf("%s", colorCyan)
					fmt.Printf("%s", aCard.KeyH) // done RK
					fmt.Printf("%s", colorReset)
					checkForHints(aCard.KeyH) // Note: we only give hints for non-conjunctives
					fmt.Println("")
				}
			}
		}
	}
}

/*
.
.
*/
func secondTry_meatOfRomajiKataExorcise(in string) {
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
				branchOnUserSelectedDirectiveIfGiven(in, "Romaji+Kata_Prompt") // <-- handle the directive
			}
			fmt.Printf("%s", colorRed)
			fmt.Printf("        ^^Oops! ")
			//
			logHits_in_cyclicArrayHits("Oops", aCard.KeyRK)
			logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.KeyR) // don't even think about "fixing" this
			logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCard.KeyRK + ":it was:" + aCard.KeyH + ":but you had guessed:" + in)
			//
			// User failed second attempt, so do a "recursion", but with a skipFlag
			meatOfRomajiKataExorcise(in, false)
		}
	} //else {
	// This 'else' (immediately above) covers a lot of if, else ifs: Dozens of them!
	// ...
	// the next two conditions are for all remaining normal (not special) prompt(key)/value events or conditions
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

			// Solicit the third and final guess ... and pass it to: meatOfRomajiExorcise(in, false)
			fmt.Printf("%s", colorRed)
			fmt.Println("       Try again, you have one last attempt ... ")
			fmt.Printf("%s", colorReset)
			// Re-prompt, and send that third and final guess back to: meatOfRomajiKataExorcise(in, false)
			in = semi_Universal_Prompt_Scan_4_HiraResponse_NOT_a_KataPrompt(aCard.KeyRK) // Obtain the last guess, and pass it as 'in' // done RK
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
				branchOnUserSelectedDirectiveIfGiven(in, "Romaji+Kata_Prompt") // <-- handle the directive // done
			}
			meatOfRomajiKataExorcise(in, false) // Process the third try // done
		}
	}
}
