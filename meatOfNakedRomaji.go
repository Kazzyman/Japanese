package main

// **do-this**
import (
	"fmt"
)

// regardless, this func and activity seems to be perfect as of 08-19-2023
// **do-this** this func employs else-if's instead of a series of simple if's as is done in meatOfRomajiKata
func meatOfRomajiNakedExorcise(in, key, value string) {
	/*  **do-this**
	??? Retain the first sections with the "^^Right!" & "^^Oops!" messages and actions, and maybe copy the rest to G-A-H-H ???
		This func provides some special "^^Right!" or "^^Oops!" messages and actions in a "'fall'-through" manner ...
		... and, THAT is ALL that it does! (SOME hints ARE also appended in the case of either a "^^Right!" or an "^^Oops!")
		... It only calls one func: checkForHints(value)
		-----------------------------------------------------------------
	*/
	var thisCaseOfAnInHasAlreadyBeenProcessedAbove bool
	thisCaseOfAnInHasAlreadyBeenProcessedAbove = false
	// ****************************************************************************
	// handel special prompt cases prior to doing the normal "if in == " processing
	// ****************************************************************************
	// Proper use of the ji sound from the "ta" group : ta->da,ji,zu , another special 'key' handler
	// た　ち   つ　て　と
	// ta chi tsu te to
	// da ji  zu  de do゛
	// key is the second parameter passed to this func, and is now KeyR
	if key == "ji" { // there is no other Romaji (or sound) for じ or ぢ
		// じ or ... not really ぢ
		if in == "じ" { // nearly-always it is じ
			thisCaseOfAnInHasAlreadyBeenProcessedAbove = true
			fmt.Printf("%s", colorGreen)
			fmt.Printf("     　　^^Right! ")
			//
			logHitsRomaji("Right", aCard.KeyR)
			//
			fmt.Printf("%s", colorReset)
			fmt.Printf("ji:じ　while ヂ:ぢ is only-SOMETIMES from chi ち ,both are the sound ji, NEVER gi (that being ぎ:gi ギ)\n")
		} else {
			thisCaseOfAnInHasAlreadyBeenProcessedAbove = true
			fmt.Printf("%s", colorRed)
			fmt.Printf("     　　^^Oops! ")
			//
			logHitsRomaji("Oops", aCard.KeyR)
			logRomajiGottenWrong(key + ":it was:" + aCard.KeyH + ":but you had guessed:" + in)
			//
			fmt.Printf("%s", colorReset)
			fmt.Printf("It was: ")
			fmt.Printf("%s", colorCyan)
			fmt.Printf("%s ", value) // shows the Hiragana じ
			fmt.Printf("%s", colorReset)
			fmt.Printf("ta->da,ji,zu, the ji sound じヂ is usually from shi:し, Rarely chi:ち")
		}

		// Proper use of the "gi" sound from the ka group : ka->ga , another special 'key' handler
		// か　き　く　け　こ
		// ka ki ku ke ko
		// ga gi gu ge go゛
	} else if key == "gi" {
		if in == "ぎ" {
			thisCaseOfAnInHasAlreadyBeenProcessedAbove = true
			fmt.Printf("%s", colorGreen)
			fmt.Printf("     ^^Right! \n")
			//
			logHitsRomaji("Right", aCard.KeyR)
			//
			fmt.Printf("%s", colorReset)
			fmt.Printf("It is always from ki き and NEVER from shi し ,it is the sound gi, NEVER ji (that being じジ)\n")
		} else {
			fmt.Printf("%s", colorRed)
			fmt.Printf("     ^^Oops! ")
			//
			logHitsRomaji("Oops", aCard.KeyR)
			logRomajiGottenWrong(key + ":it was:" + aCard.KeyH + ":but you had guessed:" + in)
			//
			fmt.Printf("%s", colorReset)
			fmt.Printf("It was: ")
			fmt.Printf("%s", colorCyan)
			fmt.Printf("%s ", value) // displays ぎ
			fmt.Printf("%s", colorReset)
			fmt.Printf("ka->ga, so the sound gi is always from ki き ,and NEVER from shi し or chi ち\n")
		}
	} else

	// Properly-constructed forms of ji used as ja, ju, jo which are nearly always from shi し and really never from chi ち
	// , three more special 'key' handlers
	if key == "ja" { // 1 of 3
		thisCaseOfAnInHasAlreadyBeenProcessedAbove = true
		if in == "じゃ" {
			fmt.Printf("%s", colorGreen)
			fmt.Printf("     ^^Right! ")
			//
			logHitsRomaji("Right", aCard.KeyR)
			//
			fmt.Printf("%s", colorReset)
			fmt.Printf("It is nearly-always from shi し and very-rarely from chi ち\n")
		} else if in != "?" {
			fmt.Printf("%s", colorRed)
			fmt.Printf("     ^^Oops! ")
			//
			logHitsRomaji("Oops", aCard.KeyR)
			logRomajiGottenWrong(key + ":it was:" + aCard.KeyH + ":but you had guessed:" + in)
			//
			fmt.Printf("%s", colorReset)
			fmt.Printf("It was: ")
			fmt.Printf("%s", colorCyan)
			fmt.Printf("%s ", value)
			fmt.Printf("%s", colorReset)
			fmt.Printf("ta->da,ji,zu: it is nearly-never from chi ち but is nearly-always from shi し\n")
		} // 1 of 3

	} else if key == "ju" { // 2 of 3
		thisCaseOfAnInHasAlreadyBeenProcessedAbove = true
		if in == "じゅ" {
			fmt.Printf("%s", colorGreen)
			fmt.Printf("     ^^Right! \n")
			//
			logHitsRomaji("Right", aCard.KeyR)
			//
			fmt.Printf("%s", colorReset)
			fmt.Printf("It is nearly always from shi し and rarely from chi ち ")
		} else if in != "?" {
			fmt.Printf("%s", colorRed)
			fmt.Printf("     ^^Oops! ")
			//
			logHitsRomaji("Oops", aCard.KeyR)
			logRomajiGottenWrong(key + ":it was:" + aCard.KeyH + ":but you had guessed:" + in)
			//
			fmt.Printf("%s", colorReset)
			fmt.Printf("It was: ")
			fmt.Printf("%s", colorCyan)
			fmt.Printf("%s ", value)
			fmt.Printf("%s", colorReset)
			fmt.Printf("ta->da,ji,zu: it is nearly-never from chi ち but is nearly-always from shi し ")
		} // 2 of 3
	} else if key == "jo" { // 3 of 3
		thisCaseOfAnInHasAlreadyBeenProcessedAbove = true
		if in == "じょ" {
			fmt.Printf("%s", colorGreen)
			fmt.Printf("     ^^Right! \n")
			//
			logHitsRomaji("Right", aCard.KeyR)
			//
			fmt.Printf("%s", colorReset)
			fmt.Printf("It is nearly always from shi し and rarely from chi ち ")
		} else if in != "?" {
			fmt.Printf("%s", colorRed)
			fmt.Printf("     ^^Oops! ")
			//
			logHitsRomaji("Oops", aCard.KeyR)
			logRomajiGottenWrong(key + ":it was:" + aCard.KeyH + ":but you had guessed:" + in)
			//
			fmt.Printf("%s", colorReset)
			fmt.Printf("It was: ")
			fmt.Printf("%s", colorCyan)
			fmt.Printf("%s ", value)
			fmt.Printf("%s", colorReset)
			fmt.Printf("ta->da,ji,zu: it is nearly-never from chi:ち but is nearly-always from shi:し ")
		}
	} else // end 3 of 3

	// 3 /- obsolete -/ forms of ja, ju and jo , special 'key' handlers for obsolete rarely-used forms
	if key == "ja obs" { // 1 of 3
		thisCaseOfAnInHasAlreadyBeenProcessedAbove = true
		if in == "ぢゃ" { // asked for the wrong (obs) way, and user has given it
			fmt.Printf("%s", colorGreen)
			fmt.Printf("     ^^Right! (as in: \"you gave the wrong way:-)\" ")
			//
			logHitsRomaji("Right", aCard.KeyR)
			//
			fmt.Printf("%s", colorReset)
			fmt.Printf("The ji sound is nearly-always from shi:し and really never from chi:ち\n")
		} else if in != "?" {
			fmt.Printf("%s", colorRed)
			fmt.Printf("     ^^Oops! ")
			//
			logHitsRomaji("Oops", aCard.KeyR)
			logRomajiGottenWrong(key + ":it was:" + aCard.KeyH + ":but you had guessed:" + in)
			//
			fmt.Printf("%s", colorReset)
			fmt.Printf("It was: ")
			fmt.Printf("%s", colorCyan)
			fmt.Printf("%s ", value)
			fmt.Printf("%s", colorReset)
			fmt.Printf("ta->da,ji,zu: ji is nearly-never from chi:ち but is nearly-always from shi:し ")
		} // 1 of 3
	} else if key == "ju obs" { // 2 of 3
		thisCaseOfAnInHasAlreadyBeenProcessedAbove = true
		if in == "ぢゅ" {
			fmt.Printf("%s", colorGreen)
			fmt.Printf("     ^^Right! (as in: \"you gave the wrong way:-)\" ")
			//
			logHitsRomaji("Right", aCard.KeyR)
			//
			fmt.Printf("%s", colorReset)
			fmt.Printf("The ji sound is nearly-always from shi し and really never from chi ち ")
		} else if in != "?" {
			fmt.Printf("%s", colorRed)
			fmt.Printf("     ^^Oops! ")
			//
			logHitsRomaji("Oops", aCard.KeyR)
			logRomajiGottenWrong(key + ":it was:" + aCard.KeyH + ":but you had guessed:" + in)
			//
			fmt.Printf("%s", colorReset)
			fmt.Printf("It was: ")
			fmt.Printf("%s", colorCyan)
			fmt.Printf("%s ", value)
			fmt.Printf("%s", colorReset)
			fmt.Printf("ta->da,ji,zu: ji is nearly-never from chi:ち but is nearly-always from shi:し ")
		} // 2 of 3
	} else if key == "jo obs" { // 3 of 3
		thisCaseOfAnInHasAlreadyBeenProcessedAbove = true
		if in == "ぢょ" {
			fmt.Printf("%s", colorGreen)
			fmt.Printf("     ^^Right! (as in: \"you gave the wrong way:-)\" ")
			//
			logHitsRomaji("Right", aCard.KeyR)
			//
			fmt.Printf("%s", colorReset)
			fmt.Printf("It is nearly always from shi し and really never from chi ち  ")
		} else {
			fmt.Printf("%s", colorRed)
			fmt.Printf("     ^^Oops! ")
			//
			logHitsRomaji("Oops", aCard.KeyR)
			logRomajiGottenWrong(key + ":it was:" + aCard.KeyH + ":but you had guessed:" + in)
			//
			fmt.Printf("%s", colorReset)
			fmt.Printf("It was: ")
			fmt.Printf("%s", colorCyan)
			fmt.Printf("%s ", value)
			fmt.Printf("%s", colorReset)
			fmt.Printf("ta->da,ji,zu: ji is nearly-never from chi:ち but is nearly-always from shi:し ")
		}
	} else // end 3 of 3  ---  'obsolete' :: ya, yu, yo forms of ja, ju, jo "obsolete"
	// */  -- 'fall' through --- "fall-through" --- >>>

	// Next, we check for some additional 'very-special' prompt(key)/value events or conditions ...
	// ... one 'very'-special key handler to emphasize the sameness of the two variants of zu sound
	if key == "zu" { // || key == "zu: from つ or す" // does not seem to work ????????? **do-this**
		//     ^^     ^^      ^^   ^^^^^^^^^ are both in main!!!!
		// zu has two KeyR values **** when prompting with naked Romaji ****** "zu", and "zu: from つ or す"
		thisCaseOfAnInHasAlreadyBeenProcessedAbove = true
		if in == "づ" {
			fmt.Printf("%s", colorGreen)
			fmt.Printf("     ^^Right! ")
			logHitsRomaji("Right", aCard.KeyR)
			fmt.Printf("%s", colorReset)
			fmt.Printf("It could have been either ず or づ as they are the same sound ")
		} else if in == "ず" {
			fmt.Printf("%s", colorGreen)
			fmt.Printf("     ^^Right! ")
			//
			logHitsRomaji("Right", aCard.KeyR)
			//
			fmt.Printf("%s", colorReset)
			fmt.Printf("It could have been either ず or づ as they are the same sound ")
		} else {
			fmt.Printf("%s", colorRed)
			fmt.Printf("     ^^Oops! ")
			//
			logHitsRomaji("Oops", aCard.KeyR)
			logRomajiGottenWrong(key + ":it was:" + aCard.KeyH + ":but you had guessed:" + in)
			//
			fmt.Printf("%s", colorReset)
			fmt.Printf("It was: ")
			fmt.Printf("%s", colorCyan)
			fmt.Printf("%s \n", value)   // v v v v v right up here v v v v v v
			fmt.Printf("%s", colorReset) // ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^
		} // That last Printf directly-above needs that \n as in: ("It was: %s \n", value)
		// ... but maybe it don't need no '        //wasDoneAbove = true' as all the others above ???
	} // last else if, below should be the final else
	// else {
	// This 'else' (immediately above) will cover a lot of if's, else if's: Dozens of them!
	// ...
	// the next two conditions are for all remaining normal (not special) prompt(key)/value events or conditions
	// ...
	//} else {

	if thisCaseOfAnInHasAlreadyBeenProcessedAbove != true {
		// ********************************************************
		// 'else', no special cases were found, so we process normal cases of "if in == value"
		// ********************************************************
		if in == value { // value is the Hiragana
			// v v v v v v v v v v v v v v : unspecified '^^Right!' v v v
			fmt.Printf("%s", colorGreen)
			fmt.Printf("      　^^Right! \n")
			fmt.Printf("%s", colorReset)
			//
			logHitsRomaji("Right", aCard.KeyR)
			//
		} else { // wasDoneAbove == false && dontSayOops == false && // if  in != "?" && in != "set-key"
			// v v v v v v v v v v v v v v : unspecified '^^Oops!' v v v
			fmt.Printf("%s", colorRed)
			fmt.Printf("      　^^Oops! ")
			//
			logHitsRomaji("Oops", aCard.KeyR)
			logRomajiGottenWrong(key + ":it was:" + aCard.KeyH + ":but you had guessed:" + in)
			//
			fmt.Printf("%s", colorReset)
			fmt.Printf(" It was: ")
			fmt.Printf("%s", colorCyan)
			fmt.Printf("%s", value)
			fmt.Printf("%s", colorReset)
			// added new func here v v v  **do-this**
			checkForHints(value) // this new func is found in meatOfNakedKata.go : we only give hints for non-composites **do-this**
		}
	}
}
