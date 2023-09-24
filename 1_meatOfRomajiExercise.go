package main

import (
	"fmt"
)
/*
The first of these two functions is passed the user's first guess, and, potentially, it obtains a second guess
.. the second func is passed that second guess, and, potentially, it obtains the third and final guess
Finally, the first func is re-called by the second: the first func is, then, passed the third and final guess ...
... with skipFlag false signifying that the offer to try again is to be replaced with a messages containing the correct
answer, and also a set of hints relating to how one could have known the correct answer.

Since this Exercise (1) expects only Hira guesses, we test for and handle Directives with a simple Alpha test in
... semi_Universal_Prompt_Scan_4_HiraResponse_to_RomajiPrompt()
*/
func meatOfRomajiExercise(in, selectedExercise string, skipFlag bool) { //  - -
	if aCard.KeyR == "zu" {
		if in == "づ" || in == "ず" {
			log_right(aCard.KeyR)
			fmt.Printf("%s", colorGreen)
			fmt.Printf("        ^^Right! ")
			fmt.Printf("%s", colorReset)
			fmt.Printf("It could have been either ず or づ as they are the same sound: zu\n")
		} else {
			log_oops(aCard.KeyR, aCard.KeyH, in)
			fmt.Printf("%s", colorRed)
			fmt.Printf("        ^^Oops! ")
			if skipFlag == true { // skipFlag is true on first entry, & false when recalled after third failed attempt
				// When current func is re-called by secondTry_, with skipFlag false, skip this, and do next ...
				fmt.Println("Try again") // ... in other words, this only prints on first pass
				fmt.Printf(colorReset)
				// Re-prompt, second guess
				in = semi_Universal_Prompt_Scan_4_HiraResponse_to_RomajiPrompt(aCard.KeyR, selectedExercise)
				secondTry_meatOfRomajiExercise(in, selectedExercise)
			} else
			if skipFlag == false { // skipFlag is true on first entry, & false when recalled after third failed attempt
				// If the user guesses incorrectly on third-and-final try, then, and only then, execute the following
				fmt.Printf("%s", colorReset)
				fmt.Printf(" It was: ")
				fmt.Printf("%s", colorCyan)
				fmt.Printf("%s", aCard.KeyH)
				fmt.Printf("%s", colorReset)
				// Display hints in white
				fmt.Printf("\n\n%s\n%s\n%s\n", aCard.Hint1h, aCard.Hint2k, aCard.Hint3TT)
				fmt.Println("")
			}
		} // If the user was ^^Right!, then we return to TouchTypingExercises12346_10(SE) (essentially from this line)

	} else if
	aCard.KeyR == "i" {
		if in == "い" {
			log_right(aCard.KeyR)
			fmt.Printf("%s", colorGreen)
			fmt.Printf("        ^^Right! ")
			fmt.Printf("%s", colorReset)
			fmt.Printf(" \"i\" is longer on the left\n")
		} else {
			log_oops(aCard.KeyR, aCard.KeyH, in)
			fmt.Printf("%s", colorRed)
			fmt.Printf("        ^^Oops! ")
			if skipFlag == true { // skipFlag is true on first entry, & false when recalled after third failed attempt
				// When current func is re-called by secondTry_, with skipFlag false, skip this, and do next ...
				fmt.Println("Try again") // ... in other words, this only prints on first pass
				fmt.Printf(colorReset)
				// Re-prompt, second guess
				in = semi_Universal_Prompt_Scan_4_HiraResponse_to_RomajiPrompt(aCard.KeyR, selectedExercise)
				secondTry_meatOfRomajiExerciseD(in, selectedExercise)
			} else
			if skipFlag == false { // skipFlag is true on first entry, & false when recalled after third failed attempt
				// If the user guesses incorrectly on third-and-final try, then, and only then, execute the following
				fmt.Printf("%s", colorReset)
				fmt.Printf(" It was: ")
				fmt.Printf("%s", colorCyan)
				fmt.Printf("%s", aCard.KeyH)
				fmt.Printf("%s", colorReset)
				// Display hints in white
				fmt.Printf("\n\n%s\n%s\n%s\n", aCard.Hint1h, aCard.Hint2k, aCard.Hint3TT)
				fmt.Println("")
			}
		} // If the user was ^^Right!, then we return to TouchTypingExercises12346_10(SE) (essentially from this line)

	} else if
	aCard.KeyR == "ri" {
		if in == "り" {
			log_right(aCard.KeyR)
			fmt.Printf("%s", colorGreen)
			fmt.Printf("        ^^Right! ")
			fmt.Printf("%s", colorReset)
			fmt.Printf(" \"ri\" is longer on the right\n")
		} else {
			log_oops(aCard.KeyR, aCard.KeyH, in)
			fmt.Printf("%s", colorRed)
			fmt.Printf("        ^^Oops! ")
			if skipFlag == true { // skipFlag is true on first entry, & false when recalled after third failed attempt
				// When current func is re-called by secondTry_, with skipFlag false, skip this, and do next ...
				fmt.Println("Try again") // ... in other words, this only prints on first pass
				fmt.Printf(colorReset)
				// Re-prompt, second guess
				in = semi_Universal_Prompt_Scan_4_HiraResponse_to_RomajiPrompt(aCard.KeyR, selectedExercise)
				secondTry_meatOfRomajiExerciseD(in, selectedExercise)
			} else
			if skipFlag == false { // skipFlag is true on first entry, & false when recalled after third failed attempt
				// If the user guesses incorrectly on third-and-final try, then, and only then, execute the following
				fmt.Printf("%s", colorReset)
				fmt.Printf(" It was: ")
				fmt.Printf("%s", colorCyan)
				fmt.Printf("%s", aCard.KeyH)
				fmt.Printf("%s", colorReset)
				// Display hints in white
				fmt.Printf("\n\n%s\n%s\n%s\n", aCard.Hint1h, aCard.Hint2k, aCard.Hint3TT)
				fmt.Println("")
			}
		} // If the user was ^^Right!, then we return to TouchTypingExercises12346_10(SE) (essentially from this line)
	} else 
		
	
	// Not "zu" or "ri" or "i" [except for two lines, i.e.: if in == "づ" || in == "ず" {  ~~>  if in == aCardD.KeyH {
	// , and the one commented-out line, the remainder of this func is similar to the above] plus one \n
	if in == aCard.KeyH {
		log_right(aCard.KeyR)
		fmt.Printf("%s", colorGreen)
		fmt.Printf("        ^^Right! \n")
		fmt.Printf("%s", colorReset)
		// commented-out line
	} else {
		log_oops(aCard.KeyR, aCard.KeyH, in)
		fmt.Printf("%s", colorRed)
		fmt.Printf("        ^^Oops! ")
		if skipFlag == true { // skipFlag is true on first entry, & false when recalled after third failed attempt
			// When current func is re-called by secondTry_, with skipFlag false, skip this, and do next ...
			fmt.Println("Try again") // ... in other words, this only prints on first pass
			fmt.Printf(colorReset)
			// Re-prompt, second guess
			in = semi_Universal_Prompt_Scan_4_HiraResponse_to_RomajiPrompt(aCard.KeyR, selectedExercise)
			secondTry_meatOfRomajiExercise(in, selectedExercise)
		} else
		if skipFlag == false { // skipFlag is true on first entry, & false when recalled after third failed attempt
			// If the user guesses incorrectly on third-and-final try, then, and only then, execute the following
			fmt.Printf("%s", colorReset)
			fmt.Printf(" It was: ")
			fmt.Printf("%s", colorCyan)
			fmt.Printf("%s", aCard.KeyH)
			fmt.Printf("%s", colorReset)
			// Display hints in white
			fmt.Printf("\n\n%s\n%s\n%s\n", aCard.Hint1h, aCard.Hint2k, aCard.Hint3TT)
			fmt.Println("")
		}
	}
}
/*
.
*/
// Second-Try version of the above func (tests the second guess, and then obtains the third and final one)
func secondTry_meatOfRomajiExercise(in, selectedExercise string) { //  - -
	if aCard.KeyR == "zu" {
		if in == "づ" || in == "ず" {
			log_right(aCard.KeyR)
			fmt.Printf("%s", colorGreen)
			fmt.Printf("        ^^Right! ")
			fmt.Printf("%s", colorReset)
			fmt.Printf("It could have been either ず or づ as they are the same sound: zu\n")
		} else {
			log_oops(aCard.KeyR, aCard.KeyH, in)
			fmt.Printf("%s", colorRed)
			fmt.Printf("        ^^Oops! Try again\n")
			fmt.Printf(colorReset)
			// Re-prompt, final guess
			in = semi_Universal_Prompt_Scan_4_HiraResponse_to_RomajiPrompt(aCard.KeyR, selectedExercise)
			meatOfRomajiExercise(in, selectedExercise, false)// Process the third and final guess
		}
	} else if
	aCard.KeyR == "ri" {
		if in == "り" {
			log_right(aCard.KeyR)
			fmt.Printf("%s", colorGreen)
			fmt.Printf("        ^^Right! ")
			fmt.Printf("%s", colorReset)
			fmt.Printf(" \"ri\" is longer on the right\n")
		} else {
			log_oops(aCard.KeyR, aCard.KeyH, in)
			fmt.Printf("%s", colorRed)
			fmt.Printf("        ^^Oops! Try again\n")
			fmt.Printf(colorReset)
			// Re-prompt, final guess
			in = semi_Universal_Prompt_Scan_4_HiraResponse_to_RomajiPrompt(aCard.KeyR, selectedExercise)
			meatOfRomajiExerciseD(in, selectedExercise, false) // Process the third and final guess
		}
	} else if
	aCard.KeyR == "i" {
		if in == "い" {
			log_right(aCard.KeyR)
			fmt.Printf("%s", colorGreen)
			fmt.Printf("        ^^Right! ")
			fmt.Printf("%s", colorReset)
			fmt.Printf(" \"i\" is longer on the left\n")
		} else {
			log_oops(aCard.KeyR, aCard.KeyH, in)
			fmt.Printf("%s", colorRed)
			fmt.Printf("        ^^Oops! Try again\n")
			fmt.Printf(colorReset)
			// Re-prompt, final guess
			in = semi_Universal_Prompt_Scan_4_HiraResponse_to_RomajiPrompt(aCard.KeyR, selectedExercise)
			meatOfRomajiExerciseD(in, selectedExercise, false)// Process the third and final guess
		}
	} else
	// Not "zu" or "ri" or "i" [except for two lines, i.e.: if in == "づ" || in == "ず" {  ~~>  if in == aCardD.KeyH {
	// , and the one commented-out line, the remainder of this func is similar to the above] plus one \n
	if in == aCard.KeyH {
		log_right(aCard.KeyR)
		fmt.Printf("%s", colorGreen)
		fmt.Printf("        ^^Right! \n")
		fmt.Printf("%s", colorReset)
		// commented-out line
	} else {
		log_oops(aCard.KeyR, aCard.KeyH, in)
		fmt.Printf("%s", colorRed)
		fmt.Printf("        ^^Oops! Try again\n")
		fmt.Printf(colorReset)
		// Re-prompt, final guess
		in = semi_Universal_Prompt_Scan_4_HiraResponse_to_RomajiPrompt(aCard.KeyR, selectedExercise)
		meatOfRomajiExercise(in, selectedExercise, false)// Process the third and final guess
	}
}
