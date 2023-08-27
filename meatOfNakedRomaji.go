package main

// **do-this** 'these'
import (
	"fmt"
)

func meatOfRomajiNakedExorcise(in string) { // NOTE: we have already been prompted with KeyR
	// ..^ ^ ^ called as|with: (usersGuessOrOptionDirective, aCard.KeyH) that KeyH being the particular Hira in question
	// so, in=the-users-guess, keyH=aCard.KeyH
	/*  **do-this**
	??? Retain the first sections with the "^^Right!" & "^^Oops!" messages and actions, and maybe copy the rest to G-A-H-H ???
		This func provides some special "^^Right!" or "^^Oops!" messages and actions in a "'fall'-through" manner ...
		... and, THAT is ALL that it does! (SOME hints ARE also appended in the case of either a "^^Right!" or an "^^Oops!")
		... It only calls one func: checkForHints(value)
		-----------------------------------------------------------------
	*/
	var thisCaseOfAnInHasAlreadyBeenProcessedAbove bool
	thisCaseOfAnInHasAlreadyBeenProcessedAbove = false
	// *******************************************************************************************
	// Handel special prompt cases prior to doing the normal "if in == " processing
	// 'gi' and 'ji' are done first; followed by Properly-Constructed forms of 'ji' used as: ja, ju, jo
	// And, finally: one 'very'-special key handler to emphasize the sameness of two variants of the zu sound
	// *******************************************************************************************
	// Proper use of the ji sound from the "ta" group : ta->da,ji,zu , a special 'key' handler
	// た　ち   つ　て　と
	// ta chi tsu te to
	// da ji  zu  de do゛
	if aCard.KeyR == "ji" { // There is no other Romaji (or sound) for じ, or for the seldom-seen: ぢ
		// じ or ... not really: ぢ
		if in == "じ" { // Nearly-always it is じ instead of the seldom-seen: ぢ
			thisCaseOfAnInHasAlreadyBeenProcessedAbove = true
			fmt.Printf("%s", colorGreen)
			fmt.Printf("     　^^Right! ") // Missing an: '/n' because it is followed by a comment with its own
			fmt.Printf("%s", colorReset)
			fmt.Printf("ji:じジ　RARELY: ぢ:chi:ヂ ,both are the sound ji, but NEVER gi (that being ぎ:gi:ギ)\n")
			// Log the hit:
			logSkipThisPrompt(aCard.KeyR)
			logHitsRomaji("Right", aCard.KeyR)
		} else {
			thisCaseOfAnInHasAlreadyBeenProcessedAbove = true
			fmt.Printf("%s", colorRed)
			fmt.Printf("     　　^^Oops! ")
			fmt.Printf("%s", colorReset)
			fmt.Printf("It was: ")
			fmt.Printf("%s", colorCyan)
			fmt.Printf("%s ", aCard.KeyH) // shows the Hiragana じ
			fmt.Printf("%s", colorReset)
			fmt.Printf("ta->da,ji,zu, the ji sound じジ is usually from shi:し, Very-Rarely as chi:ぢヂ\n")
			// log the miss:
			logReinforceThisPrompt(aCard.KeyR)
			logHitsRomaji("Oops", aCard.KeyR)
			logRomajiGottenWrong(aCard.KeyR + ":it was:" + aCard.KeyH + ":but you had guessed:" + in)
			checkForHints(aCard.KeyH) // Note: we only give hints for non-conjunctive-composites
		}

		// Proper use of the "gi" sound from the ka group : ka->ga , another special 'key' handler
		// か　き　く　け　こ
		// ka ki ku ke ko
		// ga gi gu ge go゛
	}
	if aCard.KeyR == "gi" {
		if in == "ぎ" {
			thisCaseOfAnInHasAlreadyBeenProcessedAbove = true
			fmt.Printf("%s", colorGreen)
			fmt.Printf("      ^^Right! ")
			fmt.Printf("%s", colorReset)
			fmt.Printf("It is always from ki:き and NEVER from shi:し ,it is the sound gi, NEVER ji (that being じジ)\n")
			// log the hit:
			logSkipThisPrompt(aCard.KeyR)
			logHitsRomaji("Right", aCard.KeyR)
		} else {
			fmt.Printf("%s", colorRed)
			fmt.Printf("     ^^Oops! ")
			fmt.Printf("%s", colorReset)
			fmt.Printf("It was: ")
			fmt.Printf("%s", colorCyan)
			fmt.Printf("%s ", aCard.KeyH) // displays ぎ
			fmt.Printf("%s", colorReset)
			fmt.Printf("ka->ga, so the sound gi is always from ki:き , and NEVER from shi:し or as chi:ぢヂ\n")
			// log the miss:
			logReinforceThisPrompt(aCard.KeyR)
			logHitsRomaji("Oops", aCard.KeyR)
			logRomajiGottenWrong(aCard.KeyR + ":it was:" + aCard.KeyH + ":but you had guessed:" + in)
			checkForHints(aCard.KeyH) // Note: we only give hints for non-conjunctive-composites
		}
	}

	// Properly-constructed forms of ji used as ja, ju, jo which are nearly always from shi し and really never from chi ち
	// , three more special 'key' handlers
	if aCard.KeyR == "ja" { // 1 of 3
		thisCaseOfAnInHasAlreadyBeenProcessedAbove = true
		if in == "じゃ" {
			fmt.Printf("%s", colorGreen)
			fmt.Printf("      ^^Right! ")
			fmt.Printf("%s", colorReset)
			fmt.Printf("It is nearly-always from shi:し , and NEVER from shi:し or as chi:ぢヂ\n")
			// log the hit:
			logSkipThisPrompt(aCard.KeyR)
			logHitsRomaji("Right", aCard.KeyR)
		} else if in != "?" {
			fmt.Printf("%s", colorRed)
			fmt.Printf("     ^^Oops! ")
			fmt.Printf("%s", colorReset)
			fmt.Printf("It was: ")
			fmt.Printf("%s", colorCyan)
			fmt.Printf("%s ", aCard.KeyH)
			fmt.Printf("%s", colorReset)
			fmt.Printf("ta->da,ji,zu: it is nearly-never from chi:ち but is nearly-always from shi:し\n")
			// log the miss:
			logReinforceThisPrompt(aCard.KeyR)
			logHitsRomaji("Oops", aCard.KeyR)
			logRomajiGottenWrong(aCard.KeyR + ":it was:" + aCard.KeyH + ":but you had guessed:" + in)
			checkForHints(aCard.KeyH) // Note: we only give hints for non-conjunctive-composites
		} // 1 of 3

	}
	if aCard.KeyR == "ju" { // 2 of 3
		thisCaseOfAnInHasAlreadyBeenProcessedAbove = true
		if in == "じゅ" {
			fmt.Printf("%s", colorGreen)
			fmt.Printf("      ^^Right! ")
			fmt.Printf("%s", colorReset)
			fmt.Printf("It is nearly-always from shi:し -- rarely-if-ever from chi:ち\n")
			// log the hit:
			logSkipThisPrompt(aCard.KeyR)
			logHitsRomaji("Right", aCard.KeyR)
		} else if in != "?" {
			fmt.Printf("%s", colorRed)
			fmt.Printf("     ^^Oops! ")
			fmt.Printf("%s", colorReset)
			fmt.Printf("It was: ")
			fmt.Printf("%s", colorCyan)
			fmt.Printf("%s ", aCard.KeyH)
			fmt.Printf("%s", colorReset)
			fmt.Printf("ta->da,ji,zu: it is nearly-never from chi:ち and is nearly-always from shi:し\n")
			// log the miss:
			logReinforceThisPrompt(aCard.KeyR)
			logHitsRomaji("Oops", aCard.KeyR)
			logRomajiGottenWrong(aCard.KeyR + ":it was:" + aCard.KeyH + ":but you had guessed:" + in)
			checkForHints(aCard.KeyH) // Note: we only give hints for non-conjunctive-composites
		} // 2 of 3
	}
	if aCard.KeyR == "jo" { // 3 of 3
		thisCaseOfAnInHasAlreadyBeenProcessedAbove = true
		if in == "じょ" {
			fmt.Printf("%s", colorGreen)
			fmt.Printf("      ^^Right! ")
			fmt.Printf("%s", colorReset)
			fmt.Printf("It is nearly-always from shi:し and very-rarely from chi:ち\n")
			// log the hit:
			logSkipThisPrompt(aCard.KeyR)
			logHitsRomaji("Right", aCard.KeyR)
		} else if in != "?" {
			fmt.Printf("%s", colorRed)
			fmt.Printf("     ^^Oops! ")
			fmt.Printf("%s", colorReset)
			fmt.Printf("It was: ")
			fmt.Printf("%s", colorCyan)
			fmt.Printf("%s ", aCard.KeyH)
			fmt.Printf("%s", colorReset)
			fmt.Printf("ta->da,ji,zu: it is nearly-never from chi:ち and is nearly-always from shi:し\n")
			// log the miss:
			logReinforceThisPrompt(aCard.KeyR)
			logHitsRomaji("Oops", aCard.KeyR)
			logRomajiGottenWrong(aCard.KeyR + ":it was:" + aCard.KeyH + ":but you had guessed:" + in)
			checkForHints(aCard.KeyH) // Note: we only give hints for non-conjunctive-composites
		}
	} // end 3 of 3

	// Next, we check for some additional 'very-special' prompt events or conditions ...
	// and, one 'very'-special handler to emphasize the sameness of the two variants of the zu sound
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
			fmt.Printf("%s", colorReset)
			fmt.Printf("It was: ")
			fmt.Printf("%s", colorCyan)
			fmt.Printf("%s \n", aCard.KeyH) // needs the '\n'
			fmt.Printf("%s", colorReset)
			// log the miss:
			logReinforceThisPrompt(aCard.KeyR)
			logHitsRomaji("Oops", aCard.KeyR)
			logRomajiGottenWrong(aCard.KeyR + ":it was:" + aCard.KeyH + ":but you had guessed:" + in)
			checkForHints(aCard.KeyH) // Note: we only give hints for non-conjunctive-composites
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
			fmt.Printf("%s", colorReset)
			fmt.Printf(" It was: ")
			fmt.Printf("%s", colorCyan)
			fmt.Printf("%s", aCard.KeyH)
			fmt.Printf("%s", colorReset)
			// log the miss:
			logReinforceThisPrompt(aCard.KeyR)
			logHitsRomaji("Oops", aCard.KeyR)
			logRomajiGottenWrong(aCard.KeyR + ":it was:" + aCard.KeyH + ":but you had guessed:" + in)
			checkForHints(aCard.KeyH) // Note: we only give hints for non-conjunctive-composites
		}
	}
}
