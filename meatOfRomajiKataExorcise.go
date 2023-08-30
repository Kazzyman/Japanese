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
			logHitsRomajiKata("Right")
			logSkipThisPrompt(aCard.KeyR)
			//
			fmt.Printf("%s", colorReset)
			fmt.Printf("It could have been either ず or づ as they are the same sound\n")
		} else {
			fmt.Printf("%s", colorRed)
			fmt.Printf("        ^^Oops! ")
			//
			logHitsRomajiKata("Oops")
			logReinforceThisPrompt(aCard.KeyR)
			logRomajiKataGottenWrong(aCard.KeyRK + ":it was:" + aCard.KeyH + ":but you had guessed:" + in)
			//
			// Solicit another guess ...
			fmt.Println("Try again")
			// re-prompt
			usersGuessOrOptionDirective := semi_Universal_Prompt_Scan_4_HiraResponse_NOT_a_KataPrompt(aCard.KeyRK)
			meatOfRomajiKataExorcise_secondTry(usersGuessOrOptionDirective)
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
				logHitsRomajiKata("Right")
				logSkipThisPrompt(aCard.KeyR)
				//
			} else {
				fmt.Printf("%s", colorRed)
				fmt.Printf("      　 ^^Oops! ")
				//
				logHitsRomajiKata("Oops")
				logReinforceThisPrompt(aCard.KeyR)
				logRomajiKataGottenWrong(aCard.KeyRK + ":it was:" + aCard.KeyH + ":but you had guessed:" + in)
				//
				// Solicit another guess ...
				fmt.Println("Try again")
				// re-prompt
				usersGuessOrOptionDirective := semi_Universal_Prompt_Scan_4_HiraResponse_NOT_a_KataPrompt(aCard.KeyRK)
				meatOfRomajiKataExorcise_secondTry(usersGuessOrOptionDirective)
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
func meatOfRomajiKataExorcise_secondTry(in string) {
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
			logHitsRomajiKata("Right")
			logSkipThisPrompt(aCard.KeyR)
			//
			fmt.Printf("%s", colorReset)
			fmt.Printf("It could have been either ず or づ as they are the same sound\n")
		} else {
			fmt.Printf("%s", colorRed)
			fmt.Printf("        ^^Oops! ")
			//
			logHitsRomajiKata("Oops")
			logReinforceThisPrompt(aCard.KeyR)
			logRomajiKataGottenWrong(aCard.KeyRK + ":it was:" + aCard.KeyH + ":but you had guessed:" + in)
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
				logHitsRomajiKata("Right")
				logSkipThisPrompt(aCard.KeyR)
				//
			} else {
				fmt.Printf("%s", colorRed)
				fmt.Printf("      　 ^^Oops! ")
				//
				logHitsRomajiKata("Oops")
				logReinforceThisPrompt(aCard.KeyR)
				logRomajiKataGottenWrong(aCard.KeyRK + ":it was:" + aCard.KeyH + ":but you had guessed:" + in)
				//
				// User failed second attempt, so do a "recursion", but with a skipFlag
				meatOfRomajiKataExorcise(in, true)
			}
		}
	}
}
