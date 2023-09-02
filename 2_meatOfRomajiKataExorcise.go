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
			// Solicit another guess ...
			fmt.Println("Try again")
			// re-prompt
			usersGuessOrOptionDirective := semi_Universal_Prompt_Scan_4_HiraResponse_NOT_a_KataPrompt(aCard.KeyRK)
			secondTry_meatOfRomajiKataExorcise(usersGuessOrOptionDirective)
			// if the user guesses correctly on this ^ second try, skip the rest of this func
			if skipFlag == false {
				fmt.Printf("%s", colorReset)
				fmt.Printf("It was: ")
				fmt.Printf("%s", colorCyan)
				fmt.Printf("%s ", aCard.KeyH) // v v v v v right up here v v v v v v
				fmt.Printf("%s", colorReset)  // ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^
				fmt.Printf("It could have been either ず or づ as they are the same sound\n")
				checkForHints(aCard.KeyH) // Note: we only give hints for non-conjunctives
			}
		}
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
				logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.KeyR)
				logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCard.KeyRK + ":it was:" + aCard.KeyH + ":but you had guessed:" + in)
				//
				// Solicit another guess ...
				fmt.Println("Try again")
				// re-prompt
				usersGuessOrOptionDirective := semi_Universal_Prompt_Scan_4_HiraResponse_NOT_a_KataPrompt(aCard.KeyRK)
				secondTry_meatOfRomajiKataExorcise(usersGuessOrOptionDirective)
				// if the user guesses correctly on this ^ second try, skip the rest of this func
				if skipFlag == false {
					fmt.Printf("%s", colorReset)
					fmt.Printf(" It was: ")
					fmt.Printf("%s", colorCyan)
					fmt.Printf("%s", aCard.KeyH)
					fmt.Printf("%s", colorReset)
					// added new func here v v v  **do-this**
					checkForHints(aCard.KeyH) // Note: we only give hints for non-conjunctives
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
				branchOnUserSelectedDirectiveIfGiven(in, "Kata_Prompt-Respond-w-Hira|Romaji") // <-- handle the directive
			}
			fmt.Printf("%s", colorRed)
			fmt.Printf("        ^^Oops! ")
			//
			logHits_in_cyclicArrayHits("Oops", aCard.KeyRK)
			logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.KeyR)
			logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCard.KeyRK + ":it was:" + aCard.KeyH + ":but you had guessed:" + in)
			//
			// User failed second attempt, so do a "recursion", but with a skipFlag
			meatOfRomajiKataExorcise(in, true)
		}
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
				fmt.Printf("      　 ^^Oops! ")
				//
				logHits_in_cyclicArrayHits("Oops", aCard.KeyRK)
				logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.KeyR)
				logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCard.KeyRK + ":it was:" + aCard.KeyH + ":but you had guessed:" + in)
				//
				// User failed second attempt, so do a "recursion", but with a skipFlag
				meatOfRomajiKataExorcise(in, true)
			}
		}
	}
}
