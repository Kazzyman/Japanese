package main

// **do-this** : is the tag I have used to denote sections that I have yet to complete

import "fmt"

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
func meatOfRomajiKataExorcise(in string, skipFlag bool) {
	// ..^ ^ ^ called as|with: (usersGuessOrOptionDirective, (bool)'skipFlag')
	// so, 'in' == the-users-guess, which is ...
	// obtained prior to this func being called ...
	// ^^^^^^^^ either by TouchTypingExorcise, or by: secondTry_ (secondTry_ actually solicits and obtains the THIRD guess)
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
			logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.KeyR)
			logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCard.KeyRK + ":it was:" + aCard.KeyH + ":but you had guessed:" + in)
			//
			// When this func is re-called by secondTry_, with flag set false, we skip this, and do the next
			if skipFlag == true {
				// Solicit the second guess ... and pass it to: secondTry_meatOfRomajiKataExorcise(in, true)

				fmt.Println("Try again") // This only prints on first pass
				fmt.Printf(colorReset)
				// Re-prompt, will send that second guess to the secondTry_meatOfRomajiKataExorcise func ...
				// Obtain the second guess, will pass it as 'in' to the secondTry_ func, below
				in = semi_Universal_Prompt_Scan_4_HiraResponse_NOT_a_KataPrompt(aCard.KeyRK)
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
					branchOnUserSelectedDirectiveIfGiven(in, "Romaji+Kata_Prompt") // <-- Do directive
				}
				secondTry_meatOfRomajiKataExorcise(in) // This instance of 'in' is actually the third and final guess
			}
			// If the user guesses incorrectly on third-and-final try, then, and only then, execute the following
			if skipFlag == false { // skipFlag is true on first entry, & false when recalled after third failed attempt
				fmt.Printf("%s", colorReset)
				fmt.Printf("\n It was: ")
				fmt.Printf("%s", colorCyan)
				fmt.Printf("%s", aCard.KeyH) // done
				fmt.Printf("%s", colorReset)
				checkForHints(aCard.KeyH) // Note: we only give hints for non-conjunctives
				fmt.Println("")
			}
		} // If the user was ^^Right!, then we return to TouchTypingExorcise(SE) (directly from this line)

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
			logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.KeyR) // This is correct
			logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCard.KeyRK +
				":it was:" + aCard.KeyH + ":but you had guessed:" + in)
			//
			//
			// When this func is re-called by secondTry_, with flag set false, we skip this, and do the next
			if skipFlag == true {
				// So, we solicit another guess ... (his second guess) and ...
				fmt.Println("Try again") // This only prints on first pass
				fmt.Printf(colorReset)
				// ... Re-prompt, and send that second guess to the secondTry_meatOfRomajiKataExorcise func
				in = semi_Universal_Prompt_Scan_4_HiraResponse_NOT_a_KataPrompt(aCard.KeyRK) // done RK ?
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
					branchOnUserSelectedDirectiveIfGiven(in, "Romaji+Kata_Prompt") // <-- Do the directive
				}
				secondTry_meatOfRomajiKataExorcise(in) // done RK
			}
			// If the user guesses incorrectly on third-and-final try, then, and only then, execute the following
			if skipFlag == false { // skipFlag is true on first entry, & false when recalled after third failed attempt
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

/*
.
.
*/
// Second-Try version of the above func (tests the second guess, and then obtains the third and final one)
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
				branchOnUserSelectedDirectiveIfGiven(in, "Romaji+Kata_Prompt") // <-- Do the directive
			}
			fmt.Printf("%s", colorRed)
			fmt.Printf("        ^^Oops! ")
			//
			logHits_in_cyclicArrayHits("Oops", aCard.KeyRK)
			logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.KeyR) // Don't even think about "fixing" this
			logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCard.KeyRK +
				":it was:" + aCard.KeyH + ":but you had guessed:" + in)
			//
			// User failed third and final attempt, so do a "recursion", but with a skipFlag false
			meatOfRomajiKataExorcise(in, false)
		}
	} //else {
	// This 'else' (immediately above) covers a lot of if, else ifs: Dozens of them!
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
			// User failed third and final attempt, so do a "recursion", but with a skipFlag false
			meatOfRomajiKataExorcise(in, false) // Process the third try // done
		}
	}
}