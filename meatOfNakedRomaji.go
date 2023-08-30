package main

// **do-this** 'these'
import (
	"fmt"
)

func meatOfRomajiNakedExorcise(in string, skipFlag bool) { // NOTE: we have already been prompted with KeyR
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
			logSkipThisPrompt(aCard.KeyR)
			logHitsRomaji("Right", aCard.KeyR)
		} else if in == "ず" {
			fmt.Printf("%s", colorGreen)
			fmt.Printf("      ^^Right! ")
			fmt.Printf("%s", colorReset)
			fmt.Printf("It could have been either ず or づ as they are the same sound: zu\n")
			// log the hit:
			logSkipThisPrompt(aCard.KeyR)
			logHitsRomaji("Right", aCard.KeyR)
		} else {
			fmt.Printf("%s", colorRed)
			fmt.Printf("     ^^Oops! ")
			// log the miss:
			logReinforceThisPrompt(aCard.KeyR)
			logHitsRomaji("Oops", aCard.KeyR)
			logRomajiGottenWrong(aCard.KeyR + ":it was:" + aCard.KeyH + ":but you had guessed:" + in)
			// Solicit another guess ...
			fmt.Println("Try again")
			// re-prompt
			usersGuessOrOptionDirective := semi_Universal_Prompt_Scan_4_HiraResponse_NOT_a_KataPrompt(aCard.KeyR)
			meatOfRomajiNakedExorcise_secondTry(usersGuessOrOptionDirective)
			// if the user guesses correctly on this ^ second try, skip the rest of this func
			if skipFlag == false {
				fmt.Printf("%s", colorReset)
				fmt.Printf("It was: ")
				fmt.Printf("%s", colorCyan)
				fmt.Printf("%s \n", aCard.KeyH) // needs the '\n'
				fmt.Printf("%s", colorReset)
				checkForHints(aCard.KeyH) // Note: we only give hints for non-conjunctives
			}
		}
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
			logSkipThisPrompt(aCard.KeyR)
			logHitsRomaji("Right", aCard.KeyR)
		} else {
			fmt.Printf("%s", colorRed)
			fmt.Printf("      　^^Oops! ")
			// log the miss:
			logReinforceThisPrompt(aCard.KeyR)
			logHitsRomaji("Oops", aCard.KeyR)
			logRomajiGottenWrong(aCard.KeyR + ":it was:" + aCard.KeyH + ":but you had guessed:" + in)
			// Solicit another guess ...
			fmt.Println("Try again")
			// re-prompt
			usersGuessOrOptionDirective := semi_Universal_Prompt_Scan_4_HiraResponse_NOT_a_KataPrompt(aCard.KeyR)
			meatOfRomajiNakedExorcise_secondTry(usersGuessOrOptionDirective)
			// if the user guesses correctly on this ^ second try, skip the rest of this func
			if skipFlag == false {
				fmt.Printf("%s", colorReset)
				fmt.Printf(" It was: ")
				fmt.Printf("%s", colorCyan)
				fmt.Printf("%s", aCard.KeyH)
				fmt.Printf("%s", colorReset)
				checkForHints(aCard.KeyH) // Note: we only give hints for non-conjunctives
			}
		}
	}
}

/*
.
*/
func meatOfRomajiNakedExorcise_secondTry(in string) { // NOTE: we have already been prompted with KeyR
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
			logSkipThisPrompt(aCard.KeyR)
			logHitsRomaji("Right", aCard.KeyR)
		} else if in == "ず" {
			fmt.Printf("%s", colorGreen)
			fmt.Printf("      ^^Right! ")
			fmt.Printf("%s", colorReset)
			fmt.Printf("It could have been either ず or づ as they are the same sound: zu\n")
			// log the hit:
			logSkipThisPrompt(aCard.KeyR)
			logHitsRomaji("Right", aCard.KeyR)
		} else {
			fmt.Printf("%s", colorRed)
			fmt.Printf("     ^^Oops! ")
			// log the miss:
			logReinforceThisPrompt(aCard.KeyR)
			logHitsRomaji("Oops", aCard.KeyR)
			logRomajiGottenWrong(aCard.KeyR + ":it was:" + aCard.KeyH + ":but you had guessed:" + in)
			// User failed second attempt, so do a "recursion", but with a skipFlag
			meatOfRomajiNakedExorcise(in, true)
		}
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
			logSkipThisPrompt(aCard.KeyR)
			logHitsRomaji("Right", aCard.KeyR)
		} else {
			fmt.Printf("%s", colorRed)
			fmt.Printf("      　^^Oops! ")
			// log the miss:
			logReinforceThisPrompt(aCard.KeyR)
			logHitsRomaji("Oops", aCard.KeyR)
			logRomajiGottenWrong(aCard.KeyR + ":it was:" + aCard.KeyH + ":but you had guessed:" + in)
			// User failed second attempt, so do a "recursion", but with a skipFlag
			meatOfRomajiNakedExorcise(in, true)
		}
	}
}
